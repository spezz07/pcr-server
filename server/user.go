package server

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"pcrweb/config"
	"pcrweb/impl"
	"pcrweb/logger"
	"pcrweb/model"
	"pcrweb/response"
	"pcrweb/utils"
	"strings"
	"time"
)

type RMethods struct{}

var RedisMethods impl.RedisMethods = &RMethods{}
var ctx = context.Background()

func (R RMethods) RedisValueGet(key string) (result string, err error) {
	result, err = config.Redis.Get(ctx, key).Result()
	return result, err
}

func (R RMethods) RedisValueSetExp(key string, val string, expTime time.Duration) (err error) {
	_, err = config.Redis.Set(ctx, key, val, expTime).Result()
	return err
}

type userInfoResp struct {
	model.UserInfo
	Token string `json:"token"`
}

func Login(u model.UserModel) (info userInfoResp, err error) {
	var user model.UserModel
	var userInfo userInfoResp
	result, err := config.DB.Table("pcr_user").Where("account = ?", u.Account).Get(&user)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	if !result {
		return userInfo, errors.New("用户不存在！")
	}
	if user.Password != utils.StrToMd5(u.Password) {
		return userInfo, errors.New("密码错误！")
	}
	token, err := utils.GenerateToken(user.UserId, user.Account)
	if err != nil {
		logger.Log.Warn(err)
		return userInfo, errors.New("内部错误！")
	}
	userInfo.Account = user.Account
	userInfo.UserName = user.Account
	userInfo.UserId = user.UserId
	userInfo.Avatar = user.Avatar
	userInfo.Token = token
	redisErr := RedisMethods.RedisValueSetExp(token, user.UserId, time.Hour*24*7)
	if redisErr != nil {
		fmt.Errorf(redisErr.Error())
	}
	return userInfo, nil
}

func UserWxLogin(u model.UserWxModel) (token string, err error) {
	var uInfo model.UserWxModel
	type wxR struct {
		Session_key string `json:"session_key"`
		Openid      string `json:"openid"`
	}
	var wxAppResp wxR
	appid := config.WxAppCfg.Appid
	sercet := config.WxAppCfg.Secret
	code := u.Code
	grantType := "authorization_code"
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + appid + "&secret=" + sercet + "&js_code=" + code + "&grant_type=" + grantType
	resp, err := http.Get(url)
	if err != nil {
		logger.Log.Warn(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		err = jsoniter.Unmarshal(body, &wxAppResp)
		if err != nil {
			logger.Log.Warn(err)
			return "", errors.New("内部错误！")
		}
		total, err := config.DB.Table("pcr_user").Where("open_id = ?", wxAppResp.Openid).Count()
		if err != nil {
			logger.Log.Warn(err)
			return "", errors.New("内部错误！")
		}
		if total != 0 {
			_, err := config.DB.Table("pcr_user").Where("open_id = ?", wxAppResp.Openid).Get(&uInfo)
			token, err := utils.GenerateToken(uInfo.UserId, uInfo.UserName)
			if err != nil {
				logger.Log.Warn(err)
				return "", errors.New("内部错误！")
			}
			redisErr := RedisMethods.RedisValueSetExp(token, uInfo.UserId, time.Hour*24*7)
			if redisErr != nil {
				fmt.Errorf(redisErr.Error())
				logger.Log.Warn(redisErr)
			}
			return token, nil
		} else {
			uInfo.OpenId = wxAppResp.Openid
			uInfo.SessionKey = wxAppResp.Session_key
			uInfo.Avatar = u.Avatar
			uInfo.Gender = u.Gender
			uInfo.Province = u.Province
			uInfo.Password = utils.StrToMd5("PCR111")
			uInfo.UserId = strings.Replace(uuid.New().String(), "-", "", -1)[0:8]
			uInfo.UserName = u.UserName
			//uInfo.UserId = u.UserId
			if u.Account == "" {
				uInfo.Account = u.UserName
			}
			//u.RoleId = 1
			uInfo.UserType = 2
			uInfo.UserStatus = 1
			_, err = config.DB.Table("pcr_user").Insert(uInfo)
			if err != nil {
				logger.Log.Error("insertError:" + err.Error())
				return "", errors.New("内部错误！")
			}
			token, err := utils.GenerateToken(uInfo.UserId, uInfo.UserName)
			if err != nil {
				logger.Log.Warn(err)
				return "", errors.New("内部错误！")
			}
			err = RedisMethods.RedisValueSetExp(token, uInfo.UserId, time.Hour*24*7)
			if err != nil {
				logger.Log.Warn(err)
				return "", errors.New("内部错误！")
			}
			return token, nil
		}
	} else {
		return "", errors.New("内部错误！")
	}

}

func UserSignUp(u model.UserModel) (err error) {
	//var user model.UserModel
	result, err := config.DB.SqlMapClient("selUserAccount", u.Account).Query().Count()
	if result != 0 {
		logger.Log.Error(result)
		return errors.New("用户名已注册")
	} else {
		uInfo := &u
		uInfo.Password = utils.StrToMd5(uInfo.Password)
		uInfo.UserId = strings.Replace(uuid.New().String(), "-", "", -1)
		uInfo.UserId = u.UserId[0:8]
		if u.Account == "" {
			uInfo.Account = u.UserName
		}
		//u.RoleId = 1
		uInfo.UserType = 1
		uInfo.UserStatus = 1
		_, eInfo := config.DB.Insert(uInfo)
		err = eInfo

	}
	return err
}

func GetUserList(u model.UserList) (page response.Pagination, list interface{}, err error) {
	var uList []model.UserInfo
	total, _ := config.DB.Table("pcr_user").Count()
	limit, offset, page := response.PaginationInfo(u, int(total))
	err = config.DB.Table("pcr_user").Limit(limit, offset).Find(&uList)
	//result, err := config.DB.SQL("SELECT * FROM log LIMIT ? OFFSET ?", limit, offset).Query().List()
	return page, uList, err
}
