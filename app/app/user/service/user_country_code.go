package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	adminService "go-admin/app/admin/sys/service"
	"go-admin/app/app/user/models"
	"go-admin/app/app/user/service/dto"
	baseLang "go-admin/config/base/lang"
	cDto "go-admin/core/dto"
	"go-admin/core/dto/service"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"gorm.io/gorm"
	"time"
)

type UserCountryCode struct {
	service.Service
}

// NewUserCountryCodeService app-实例化国家区号管理
func NewUserCountryCodeService(s *service.Service) *UserCountryCode {
	var srv = new(UserCountryCode)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage app-获取国家区号管理分页列表
func (e *UserCountryCode) GetPage(c *dto.UserCountryCodeQueryReq, p *middleware.DataPermission) ([]models.UserCountryCode, int64, int, error) {
	var data models.UserCountryCode
	var list []models.UserCountryCode
	var count int64

	err := e.Orm.Order("created_at desc").Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			middleware.Permission(data.TableName(), p),
		).Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return list, count, baseLang.SuccessCode, nil
}

// Get app-获取国家区号管理详情
func (e *UserCountryCode) Get(id int64, p *middleware.DataPermission) (*models.UserCountryCode, int, error) {
	if id <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.UserCountryCode{}
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).First(data, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return data, baseLang.SuccessCode, nil
}

// QueryOne app-获取国家区号管理一条记录
func (e *UserCountryCode) QueryOne(queryCondition *dto.UserCountryCodeQueryReq, p *middleware.DataPermission) (*models.UserCountryCode, int, error) {
	data := &models.UserCountryCode{}
	err := e.Orm.Scopes(
		cDto.MakeCondition(queryCondition.GetNeedSearch()),
		middleware.Permission(data.TableName(), p),
	).First(data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return data, baseLang.SuccessCode, nil
}

// Count admin-获取国家区号管理数据总数
func (e *UserCountryCode) Count(queryCondition *dto.UserCountryCodeQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.UserCountryCode{}).
		Scopes(
			cDto.MakeCondition(queryCondition.GetNeedSearch()),
		).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return 0, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return count, baseLang.SuccessCode, nil
}

// Insert app-新增国家区号管理
func (e *UserCountryCode) Insert(c *dto.UserCountryCodeInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.Country == "" {
		return 0, baseLang.UserCountryEmptyCode, lang.MsgErr(baseLang.UserCountryEmptyCode, e.Lang)
	}
	if c.Code == "" {
		return 0, baseLang.UserCountryCodeEmptyCode, lang.MsgErr(baseLang.UserCountryCodeEmptyCode, e.Lang)
	}
	if c.Status == "" {
		return 0, baseLang.UserCountryStatusEmptyCode, lang.MsgErr(baseLang.UserCountryStatusEmptyCode, e.Lang)
	}

	//检测国家名称是否存在
	reqName := dto.UserCountryCodeQueryReq{}
	reqName.CountryInner = c.Country
	count, respCode, err := e.Count(&reqName)
	if err != nil && respCode != baseLang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, baseLang.UserCountryHasExistCode, lang.MsgErr(baseLang.UserCountryHasExistCode, e.Lang)
	}

	//检测国家区号是否存在
	reqCode := dto.UserCountryCodeQueryReq{}
	reqCode.Code = c.Code
	count, respCode, err = e.Count(&reqCode)
	if err != nil && respCode != baseLang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, baseLang.UserCountryCodeHasExistCode, lang.MsgErr(baseLang.UserCountryCodeHasExistCode, e.Lang)
	}

	now := time.Now()
	var data models.UserCountryCode
	data.Country = c.Country
	data.Code = c.Code
	data.Status = c.Status
	data.CreateBy = c.CurrUserId
	data.UpdateBy = c.CurrUserId
	data.CreatedAt = &now
	data.UpdatedAt = &now
	err = e.Orm.Create(&data).Error
	if err != nil {
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return data.Id, baseLang.SuccessCode, nil
}

// Update app-更新国家区号管理
func (e *UserCountryCode) Update(c *dto.UserCountryCodeUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.Country == "" {
		return false, baseLang.UserCountryEmptyCode, lang.MsgErr(baseLang.UserCountryEmptyCode, e.Lang)
	}
	if c.Code == "" {
		return false, baseLang.UserCountryCodeEmptyCode, lang.MsgErr(baseLang.UserCountryCodeEmptyCode, e.Lang)
	}
	if c.Status == "" {
		return false, baseLang.UserCountryStatusEmptyCode, lang.MsgErr(baseLang.UserCountryStatusEmptyCode, e.Lang)
	}
	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	updates := map[string]interface{}{}
	if c.Country != "" && data.Country != c.Country {
		req := dto.UserCountryCodeQueryReq{}
		req.CountryInner = c.Country
		resp, respCode, err := e.QueryOne(&req, nil)
		if err != nil && respCode != baseLang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == baseLang.SuccessCode && resp.Id != data.Id {
			return false, baseLang.UserCountryHasExistCode, lang.MsgErr(baseLang.UserCountryHasExistCode, e.Lang)
		}
		updates["country"] = c.Country
	}
	if c.Code != "" && data.Code != c.Code {
		req := dto.UserCountryCodeQueryReq{}
		req.Code = c.Code
		resp, respCode, err := e.QueryOne(&req, nil)
		if err != nil && respCode != baseLang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == baseLang.SuccessCode && resp.Id != data.Id {
			return false, baseLang.UserCountryCodeHasExistCode, lang.MsgErr(baseLang.UserCountryCodeHasExistCode, e.Lang)
		}
		updates["code"] = c.Code
	}
	if c.Status != "" && data.Status != c.Status {
		updates["status"] = c.Status
	}
	if len(updates) > 0 {
		updates["updated_at"] = time.Now()
		updates["update_by"] = c.CurrUserId
		err = e.Orm.Model(&data).Where("id=?", data.Id).Updates(&updates).Error
		if err != nil {
			return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
		}
		return true, baseLang.SuccessCode, nil
	}
	return false, baseLang.SuccessCode, nil
}

// Delete app-删除国家区号管理
func (e *UserCountryCode) Delete(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	var data models.UserCountryCode
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return baseLang.SuccessCode, nil
}

// Export app-导出国家区号管理
func (e *UserCountryCode) Export(list []models.UserCountryCode) ([]byte, error) {
	sheetName := "UserCountryCode"
	xlsx := excelize.NewFile()
	no, _ := xlsx.NewSheet(sheetName)
	_ = xlsx.SetColWidth(sheetName, "A", "L", 25)
	_ = xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{
		"编号", "状态"})
	dictService := adminService.NewSysDictDataService(&e.Service)
	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)
		status := dictService.GetLabel("admin_sys_status", item.Status)

		//按标签对应输入数据
		_ = xlsx.SetSheetRow(sheetName, axis, &[]interface{}{
			item.Id, status,
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}
