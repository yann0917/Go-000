/*
 * @Author: yubb
 * @Date: 2019-11-26 19:41:11
 * @Last Modified by: yubb
 * @Last Modified time: 2020-02-21 11:49:57
 */
package controllers

import (
	"homework/internal/models"
	"strconv"

	. "github.com/wlxpkg/base"
	. "github.com/wlxpkg/zwyd"

	"github.com/gin-gonic/gin"
)

// GetAppList App列表
func GetAppList(c *gin.Context) {
	ctl := NewController(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	appType := c.Param("type")

	var AppModel models.CommonApp
	AppModel.Type, _ = String2Int(appType)
	appList, count, err := AppModel.GetList(page, limit)

	if err != nil {
		ctl.Error(err)
		return
	}

	data := Paginator(page, count, appList)

	ctl.Success(data)
}

// 新增App
func AddCommonApp(c *gin.Context) {
	ctl := NewController(c)
	var appModel models.CommonApp

	if err := c.ShouldBindJSON(&appModel); err != nil {
		ctl.Error(err)
		return
	}
	err := appModel.AddApp()

	if err != nil {
		ctl.Error(err)
		return
	}
	ctl.Success(1)
}

// UpdateCommonApp 更新App信息
func UpdateCommonApp(c *gin.Context) {
	ctl := NewController(c)

	id := c.Param("id")

	var appModel models.CommonApp
	var params = make(map[string]interface{})
	params["id"] = id
	appModel.ID, _ = String2Int(id)
	_, err := appModel.GetInfo(params)

	if err != nil {
		ctl.Error(Excp("ERR_PARAM"))
		return
	}
	_ = c.ShouldBindJSON(&appModel)
	err2 := appModel.UpdateCommonApp()

	if err2 != nil {
		ctl.Error(err2)
		return
	}

	ctl.Success(1)
}
