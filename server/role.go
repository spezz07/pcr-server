package server

import (
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"pcrweb/config"
	"pcrweb/impl"
	"pcrweb/model"
	"pcrweb/response"
	"strings"
)

type roleServer struct{}

var RoleServer impl.BaseRequestsImpl = roleServer{}

type rolePageList struct {
}

func (r roleServer) GetPageList(v interface{}) (page response.Pagination, data interface{}, err error) {
	vv, _ := v.(model.PcrRolePageList)
	var pcrRolePageList []model.PcrRolePageInfo
	params := map[string]interface{}{
		"RoleName":  vv.RoleName,
		"RoleAlias": vv.RoleAlias,
		"Limit":     vv.Rows, "Offset": vv.Rows * (vv.PageNo - 1),
	}
	//err = config.DB.Table("pcr_role").Limit(limit,offset).Find(&list)
	err = config.DB.SqlTemplateClient("selectRole.stpl", &params).Find(&pcrRolePageList)
	_, _, page = response.PaginationInfo(vv, len(pcrRolePageList))
	return page, pcrRolePageList, err
}

func (r roleServer) GetList(v interface{}) (data interface{}, err error) {
	return nil, nil
}

func (r roleServer) Add(v interface{}) (data interface{}, err error) {
	panic("implement me")
}

func (r roleServer) Update(v interface{}) (data interface{}, err error) {
	panic("implement me")
}

func (r roleServer) Delete(v interface{}) (data interface{}, err error) {
	panic("implement me")
}

type key struct {
	RoleName string
	Id       int
}

func RoleImgUpload() {
	// 七牛云上传
}

func RoleDataUpdata() {
	var roleName []key
	_ = config.DB.SQL("select role_name,id from pcr_role").Find(&roleName)

	for _, v := range roleName {
		var uData model.UnitData
		_, err := config.DB.Table("unit_data").Where("role_name = ?", v.RoleName).Get(&uData)
		if err != nil {
			fmt.Printf("Read:error%v\n", err)
		}
		uData.Pos = uData.SearchAreaWidth
		_, err = config.DB.Table("pcr_role").Where("id = ?", v.Id).AllCols().Update(&uData)
		_, err = config.DB.Table("pcr_role").Where("id = ?", v.Id).Update(map[string]interface{}{"pos": uData.Pos})
		if err != nil {
			fmt.Printf("Updata:Err:%v\n", err)
			panic(err.Error())
		}
	}

}

func RoleAvatarImport() {
	var fileList []*model.PcrRole
	imgFileDir := "./static"
	imgFileDir_N := "/static/N"
	imgFileDir_SR := "/static/SR"
	imgFileDir_SSR := "/static/SSR"
	dir, err := ioutil.ReadDir(imgFileDir)
	if err != nil {
		fmt.Errorf("%s", err)
	}
	for _, fileDir := range dir {
		if fileDir.Name() == "N" {
			nFile, err := ioutil.ReadDir("./" + imgFileDir_N)
			if err != nil {
				fmt.Errorf("%s", err)
			}
			for _, file := range nFile {
				if file.Name() == ".DS_Store" {
					continue
				}
				imgId := strings.Replace(uuid.New().String(), "-", "", -1)
				fileMap := &model.PcrRole{
					RoleImgName: file.Name(),
					RoleType:    1,
					RoleStar:    1,
					RoleName:    file.Name()[0:strings.Index(file.Name(), ".")],
					RoleImgId:   imgId[0:10],
					RoleImgUrl:  imgFileDir_N + "/" + file.Name(),
				}
				fileList = append(fileList, fileMap)
			}
		} else if fileDir.Name() == "SR" {
			nFile, err := ioutil.ReadDir("./" + imgFileDir_SR)
			if err != nil {
				fmt.Errorf("%s", err)
			}
			for _, file := range nFile {
				if file.Name() == ".DS_Store" {
					continue
				}
				imgId := strings.Replace(uuid.New().String(), "-", "", -1)

				fileMap := &model.PcrRole{
					RoleImgName: file.Name(),
					RoleType:    1,
					RoleStar:    2,
					RoleName:    file.Name()[0:strings.Index(file.Name(), ".")],
					RoleImgId:   imgId[0:10],
					RoleImgUrl:  imgFileDir_SR + "/" + file.Name(),
				}
				fileList = append(fileList, fileMap)
			}
		} else if fileDir.Name() == "SSR" {
			nFile, err := ioutil.ReadDir("./" + imgFileDir_SSR)
			if err != nil {
				fmt.Errorf("%s", err)
			}
			for _, file := range nFile {
				if file.Name() == ".DS_Store" {
					continue
				}
				imgId := strings.Replace(uuid.New().String(), "-", "", -1)

				fileMap := model.PcrRole{
					RoleImgName: file.Name(),
					RoleType:    1,
					RoleStar:    3,
					RoleName:    file.Name()[0:strings.Index(file.Name(), ".")],
					RoleImgId:   imgId[0:10],
					RoleImgUrl:  imgFileDir_SSR + "/" + file.Name(),
				}
				fileList = append(fileList, &fileMap)
			}
		}
		//fmt.Println(fileList)
	}
	aff, err := config.DB.Table("pcr_role").Insert(fileList)
	fmt.Println(aff)
	fmt.Println(err)
}
