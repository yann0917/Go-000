/*
 * @Author: yubb
 * @Date: 2019-11-18 15:27:22
 * @Last Modified by: qiuling
 * @Last Modified time: 2019-12-04 16:13:38
 */

package models

import (
	. "github.com/wlxpkg/base"
)

// AdOperate 广告位操作记录的 model
type AdOperate struct {
	ID        int       `json:"id"`
	UserID    string    `json:"user_id"`
	ClientID  string    `json:"client_id"`
	AdID      uint      `json:"ad_id"`
	View      int       `json:"view"`
	Click     int       `json:"click"`
	CreatedAt JSONTime  `json:"created_at"`
	UpdatedAt JSONTime  `json:"updated_at"`
	DeletedAt *JSONTime `json:"-"`
}

type StatisticPUV struct {
	View  int
	Click int
}

func (AdOperate) TableName() string {
	return "common_ad_operate"
}

// AddAdOperateView 新增广告位 查看记录
func (AdOpModel *AdOperate) AddAdOperateView() (err error) {

	adOne, err := AdOpModel.GetAdOperateInfo()

	if err != nil {
		AdOpModel.View = 1
		err = DB.Create(&AdOpModel).Error

	} else {
		var params = make(map[string]interface{})
		params["view"] = adOne.View + 1
		err = AdOpModel.UpdateInfo(params)
	}
	return
}

// AddAdOperateClick 新增广告位 点击记录
func (AdOpModel *AdOperate) AddAdOperateClick() (err error) {

	adOne, err := AdOpModel.GetAdOperateInfo()

	if err != nil {
		AdOpModel.Click = 1
		err = DB.Create(&AdOpModel).Error

	} else {
		var params = make(map[string]interface{})
		params["click"] = adOne.Click + 1
		err = AdOpModel.UpdateInfo(params)
	}
	return
}

// 获取单条操作记录 GetAdOperateInfo
func (AdOpModel *AdOperate) GetAdOperateInfo() (adOperate AdOperate, err error) {
	err = DB.Where("client_id = ? and ad_id = ?", AdOpModel.ClientID, AdOpModel.AdID).First(&adOperate).Error
	return
}

// UpdateInfo 更新广告操作
func (AdOpModel *AdOperate) UpdateInfo(params map[string]interface{}) (err error) {
	err = DB.Model(AdOpModel).
		Where("client_id = ? and ad_id = ?", AdOpModel.ClientID, AdOpModel.AdID).
		Update(params).Error
	return
}

// GetStatisticPvUv 统计PV，Uv,点击率
func (AdOpModel *AdOperate) GetStatisticPvUv() (uv int, data StatisticPUV) {

	DB.Model(AdOpModel).Select("sum(view) as view,sum(click) as click").
		Where("ad_id = ?", AdOpModel.AdID).
		Scan(&data).
		Count(&uv)
	return
}
