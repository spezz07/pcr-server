package server

import (
	"errors"
	"fmt"
	"pcrweb/config"
	"pcrweb/impl"
	"pcrweb/model"
	"pcrweb/response"
	"pcrweb/utils"
	"sort"
	"strconv"
	"strings"
	"time"
)

type fightServer struct {
	impl.BaseRequestsImpl
}

var FightServer fightServer = fightServer{}

func (s fightServer) Add(v interface{}) (data interface{}, err error) {
	session := config.DB.NewSession()
	defer session.Close()
	errSession := session.Begin()
	var f = v.(model.FightResultInfo)
	var fightInfo model.FightResult
	if len(f.Defensive) == 0 {
		return nil, errors.New("防守阵容不能为空！")
	}
	if len(f.Offensive) == 0 {
		return nil, errors.New("进攻阵容不能为空！")
	}
	sortD := make([]int, len(f.Defensive))
	sortO := make([]int, len(f.Offensive))
	copy(sortD, f.Defensive)
	copy(sortO, f.Offensive)
	sort.Ints(sortD)
	sort.Ints(sortO)
	fightInfo.Offensive = utils.IntArrayToString(f.Offensive, ",")
	fightInfo.Defensive = utils.IntArrayToString(f.Defensive, ",")
	//fightInfo.MappingName = f.MappingName

	fightInfo.SortDefensive = utils.IntArrayToString(sortD, ",")
	fightInfo.SortOffensive = utils.IntArrayToString(sortO, ",")
	fightInfo.MappingString = f.MappingString
	rid := strconv.Itoa(int(time.Now().UnixNano()))
	rid = rid[8 : len(rid)-4]
	fightInfo.IsDel = 2
	fightInfo.CreatorId = f.UserId
	fightInfo.AreaType = f.AreaType

	fightInfo.Title = f.Title
	fightInfo.RId, _ = strconv.Atoi(rid)
	_, errSession = config.DB.Table("pcr_arena_result").Insert(&fightInfo)
	save := model.FightApprove{
		ResultRelationId: fightInfo.RId,
	}
	if errSession != nil {
		_ = session.Rollback()
		return nil, errSession
	}
	_, errSession = config.DB.Table("pcr_arena_result_like").Insert(&save)
	if errSession != nil {
		_ = session.Rollback()
		return nil, errSession
	}
	errSession = session.Commit()
	return nil, errSession
}

type fightResultJoinList struct {
	CreatorName      string
	CreatorId        string
	ApproveNum       int
	DisapproveNum    int
	Offensive        string
	Defensive        string
	IsDel            int
	RId              int
	Created          impl.TimeFormat
	ApproveUserId    string
	DisapproveUserId string
	AreaType         int
	Description      string
}

func (s fightServer) GetPageList(v model.FightResultPageInfo) (page response.Pagination, data interface{}, err error) {
	var pageList []model.FightResultPageList
	var pageJoinList []fightResultJoinList
	params := map[string]interface{}{
		"Limit":     v.Rows,
		"Offset":    v.Rows * (v.PageNo - 1),
		"Defensive": utils.IntArrayToString(v.Defensive, ","),
	}
	if len(v.Defensive) == 0 {
		return page, pageList, errors.New("请选择阵容！")
	}
	sql := "SELECT l.*,r.approve_user_id,disapprove_user_id FROM pcr_arena_result l JOIN pcr_arena_result_like r ON l.r_id = r.result_relation_id WHERE  is_del!=1 AND sort_defensive LIKE ? "
	sqlL := "ORDER BY l.created DESC  Limit ? offset ? "
	err = config.DB.SQL(sql+sqlL, "%"+params["Defensive"].(string)+"%", params["Limit"], params["Offset"]).Find(&pageJoinList)
	count, _ := config.DB.SQL(sql, "%"+params["Defensive"].(string)+"%").Query().Count()
	for _, val := range pageJoinList {
		pItem := model.FightResultPageList{}
		approveUserId := strings.Split(val.ApproveUserId, ",")
		disapproveUserId := strings.Split(val.DisapproveUserId, ",")
		approveStatus := 0
		t, _ := utils.IsExistInStringArray(v.UserId, approveUserId)
		if t {
			approveStatus = 1
		} else {
			t, _ = utils.IsExistInStringArray(v.UserId, disapproveUserId)
			if t {
				approveStatus = 2
			}
		}
		pItem.Defensive = utils.StringArrayToIntArray(strings.Split(val.Defensive, ","))
		pItem.Offensive = utils.StringArrayToIntArray(strings.Split(val.Offensive, ","))
		pItem.AreaType = val.AreaType
		pItem.Created = val.Created
		pItem.CreatorName = val.CreatorName
		pItem.CreatorId = val.CreatorId
		pItem.RId = val.RId
		pItem.Status = val.IsDel
		pItem.ApproveStatus = approveStatus
		pItem.Description = val.Description
		pItem.ApproveNum = len(approveUserId)
		pItem.DisapproveNum = len(disapproveUserId)
		if approveUserId[0] == "" {
			pItem.ApproveNum = 0
		}
		if disapproveUserId[0] == "" {
			pItem.DisapproveNum = 0
		}
		pageList = append(pageList, pItem)
	}
	if count == 0 {
		pageList = []model.FightResultPageList{}
	}
	_, _, page = response.PaginationInfo(v, count)
	return page, pageList, err
}

func (s fightServer) ApproveHandle(v model.FightLike) (err error) {
	var f []model.FightApprove
	isE, err := config.DB.Table("pcr_arena_result").Where("r_id = ?", v.FightResultId).Exist()
	if !isE {
		_ = fmt.Errorf("%v", err.Error())
		return errors.New("竞技场结果Id错误！")
	}
	err = config.DB.Table("pcr_arena_result_like").Where("result_relation_id = ?", v.FightResultId).Find(&f)
	if err != nil {
		_ = fmt.Errorf("%v", err.Error())
		return errors.New("查询Id出错！")
	}
	approveStatus := v.Status
	//if f == nil || f[0].ResultRelationId == 0 {
	//	if approveStatus == 1 {
	//		save := model.FightApprove{
	//			ResultRelationId: v.FightResultId,
	//			ApproveUserId:    v.UserId,
	//		}
	//		_, err = config.DB.Table("pcr_arena_result_like").Insert(save)
	//
	//	} else {
	//		save := model.FightApprove{
	//			ResultRelationId: v.FightResultId,
	//			DisapproveUserId: v.UserId,
	//		}
	//		_, err = config.DB.Table("pcr_arena_result_like").Insert(save)
	//	}
	//} else {
	var appRList []string
	var disAppRList []string
	disAppRList = strings.Split(f[0].DisapproveUserId, ",")
	appRList = strings.Split(f[0].ApproveUserId, ",")
	isOk, okIndex := utils.IsExistInStringArray(v.UserId, appRList)
	isDiss, dissIndex := utils.IsExistInStringArray(v.UserId, disAppRList)
	if approveStatus == 1 {
		if isOk { // 已有，在点赞就取消
			utils.RemoveItemArray(&appRList, okIndex)
			upVal := model.FightApprove{
				ApproveUserId: trimComma(strings.Join(appRList, ",")),
			}
			_, err = config.DB.Table("pcr_arena_result_like").Where("result_relation_id = ?", v.FightResultId).Cols("approve_user_id").Update(&upVal)
			return err
		}
		if isDiss { // 已有，在反对就转为点赞
			utils.RemoveItemArray(&disAppRList, dissIndex)
			appRList = append(appRList, v.UserId)
			upVal := model.FightApprove{
				ApproveUserId:    trimComma(strings.Join(appRList, ",")),
				DisapproveUserId: trimComma(strings.Join(disAppRList, ",")),
			}
			_, err = config.DB.Table("pcr_arena_result_like").Where("result_relation_id = ?", v.FightResultId).Cols("approve_user_id").Cols("disapprove_user_id").Update(&upVal)
			return err
		}
		// 均没有就点赞
		appRList = append(appRList, v.UserId)
		upVal := model.FightApprove{
			ApproveUserId: trimComma(strings.Join(appRList, ",")),
		}
		_, err = config.DB.Table("pcr_arena_result_like").Where("result_relation_id = ?", v.FightResultId).Update(&upVal)
	} else {
		if isDiss { // 已有，在点就取消
			utils.RemoveItemArray(&disAppRList, dissIndex)
			upVal := model.FightApprove{
				DisapproveUserId: trimComma(strings.Join(disAppRList, ",")),
			}
			_, err = config.DB.Table("pcr_arena_result_like").Where("result_relation_id = ?", v.FightResultId).Cols("disapprove_user_id").Update(&upVal)
			return err
		}
		if isOk { // 已有，在就转为反对
			utils.RemoveItemArray(&appRList, okIndex)
			disAppRList = append(disAppRList, v.UserId)
			upVal := model.FightApprove{
				DisapproveUserId: trimComma(strings.Join(disAppRList, ",")),
				ApproveUserId:    trimComma(strings.Join(appRList, ",")),
			}
			_, err = config.DB.Table("pcr_arena_result_like").Where("result_relation_id = ?", v.FightResultId).Cols("approve_user_id").Cols("disapprove_user_id").Update(&upVal)
			return err
		}
		// 均没有就点赞
		disAppRList = append(disAppRList, v.UserId)
		upVal := model.FightApprove{
			DisapproveUserId: trimComma(strings.Join(disAppRList, ",")),
		}
		_, err = config.DB.Table("pcr_arena_result_like").Where("result_relation_id = ?", v.FightResultId).Update(&upVal)
	}
	return err
}

func trimComma(str string) (trimStr string) {
	trimStr = strings.TrimSuffix(str, ",")
	trimStr = strings.TrimPrefix(trimStr, ",")
	return trimStr
}
