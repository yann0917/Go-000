/*
 * @Author: yubb
 * @Date: 2019-11-26 19:43:56
 * @Last Modified by: yubb
 * @Last Modified time: 2020-02-21 11:49:41
 */
package models

import (
	. "github.com/wlxpkg/base"
)

// App common_app的model(App管理-渠道管理)
type CommonApp struct {
	ID            int       `json:"id"`
	Type          int       `json:"type"`
	Name          string    `json:"name"`
	Code          string    `json:"code"`
	IosWechat     string    `json:"ios_wechat"`
	AndroidWechat string    `json:"android_wechat"`
	CreatedAt     JSONTime  `json:"created_at"`
	UpdatedAt     JSONTime  `json:"updated_at"`
	DeletedAt     *JSONTime `json:"-"`
}

// ===================缓存参数==============
var AppDetailCache = "app:detail:"

// AddApp 新增
func (AppModel *CommonApp) AddApp() (err error) {

	err = DB.Create(&AppModel).Error
	return
}

// UpdateCommonApp 更新App信息
func (AppModel *CommonApp) UpdateCommonApp() (err error) {
	err = DB.Model(&AppModel).Where("id = ?", AppModel.ID).Updates(&AppModel).Error
	return
}

// GetInfo 获取App信息
func (AppModel *CommonApp) GetInfo(params map[string]interface{}) (AppModels CommonApp, err error) {
	db := DB

	if _, ok := params["id"]; ok {
		db = db.Where("id = ?", params["id"])
	}
	if _, ok := params["code"]; ok {
		db = db.Where("code = ?", params["code"])
	}
	err = db.
		First(&AppModels).Error
	return
}

// GetList 查看所有的广告位(后台专用)
func (AppModel *CommonApp) GetList(page int, limit int) (appList []CommonApp, count int, err error) {

	if page < 1 {
		page = 1
	}
	var offset = (page - 1) * limit
	err = DB.Where(AppModel).Order("created_at asc").
		Limit(limit).Offset(offset).
		Find(&appList).
		Limit(-1).Offset(-1).Count(&count).Error
	return
}
