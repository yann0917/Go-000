package controllers

import (
	"fmt"
	"homework/internal/models"
	"math/rand"
	"strconv"
	"time"

	. "github.com/wlxpkg/base"
	. "github.com/wlxpkg/zwyd"

	"github.com/gin-gonic/gin"
)

func AddAd(c *gin.Context) {
	ctl := Controller{Ctx: c}
	var AdModel models.Ad
	if err := c.ShouldBindJSON(&AdModel); err == nil {

		id, err := AdModel.AddNewAd()

		if err != nil {
			ctl.Error(err)
			return
		}
		data := make(map[string]uint, 1)
		data["id"] = id
		ctl.Success(data)
	} else {
		ctl.Error(err)
		return
	}
}

func AdList(c *gin.Context) {
	ctl := Controller{Ctx: c}

	typeValue, err := strconv.ParseUint(c.Param("type"), 10, 0)
	if err != nil {
		ctl.Error(err)
		return
	}
	page, err := strconv.ParseUint(c.DefaultQuery("page", "1"), 10, 0)
	if err != nil {
		ctl.Error(err)
		return
	}
	limit, err := strconv.ParseUint(c.DefaultQuery("limit", "10"), 10, 0)
	if err != nil {
		ctl.Error(err)
		return
	}

	var AdModel models.Ad
	AdModel.Type = uint(typeValue)
	result, count, err := AdModel.GetList(int(page), int(limit))

	for key, val := range result {
		var str float64
		if val.Pv > 0 {
			float32Click, _ := strconv.ParseFloat(Int2String(val.Clicks), 32)
			float32Pv, _ := strconv.ParseFloat(Int2String(val.Pv), 32)
			str = float32Click / float32Pv * 100
		} else {
			str = 0.00
		}
		result[key].ClickRate = fmt.Sprintf("%.2f", str) + "%"

		// 循环统计PV,Uv，点击率
		AdModel.StatisticPvUv(val.ID)
	}

	if err != nil {
		ctl.Error(err)
		return
	}

	ctl.Success(gin.H{
		"list":  result,
		"count": count,
	})
}

func EditAd(c *gin.Context) {
	ctl := Controller{Ctx: c}
	var AdModel models.Ad
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		ctl.Error(err)
	}
	if c.ShouldBindJSON(&AdModel) == nil {

		result, err := AdModel.Edit(uint(id))
		if err != nil {
			ctl.Error(err)
			return
		}

		ctl.Success(result)
	} else {
		ctl.Error(err)
		return
	}

}

func AdInfo(c *gin.Context) {
	ctl := Controller{Ctx: c}

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)

	if err != nil {
		ctl.Error(err)
		return
	}

	var AdModel models.Ad

	AdModel.ID = uint(id)
	result, err := AdModel.GetInfo()

	if err != nil {
		ctl.Error(err)
		return
	}

	ctl.Success(result)
}

// DelCourseType 删除广告
func DelAd(c *gin.Context) {
	ctl := Controller{Ctx: c}
	var AdModel models.Ad
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		ctl.Error(err)
		return
	}
	AdModel.ID = uint(id)
	result, err := AdModel.Delete()
	if err != nil {
		ctl.Error(err)
		return
	}

	ctl.Success(result)
}

func AdFrontList(c *gin.Context) {
	ctl := Controller{Ctx: c}

	typeValue, err := strconv.ParseUint(c.Param("type"), 10, 0)
	if err != nil {
		ctl.Error(err)
		return
	}
	var codeValue = c.DefaultQuery("code", "")
	var AdModel models.Ad
	var status uint = 1
	AdModel.Type = uint(typeValue)
	AdModel.Code = codeValue
	AdModel.Status = &status
	result, err := AdModel.GetListForFront(1, -1)

	if err != nil {
		ctl.Error(err)
		return
	}

	// 当是广告  或者 弹窗 时，随机展示一个
	if typeValue == 1 || typeValue == 2 {
		var returnList [1]models.Ad

		if len(result) >= 1 {
			rand.Seed(time.Now().UnixNano())
			returnList[0] = result[rand.Intn(len(result))]
		} else {
			returnList[0] = AdModel
		}

		ctl.Success(returnList)

	} else {
		ctl.Success(result)
	}

}

func AdFrontOperateView(c *gin.Context) {
	ctl := NewController(c)
	id := String2Uint(c.Param("id"))

	var AdOpModel models.AdOperate
	AdOpModel.UserID = Int642String(ctl.UserID)
	AdOpModel.ClientID = ctl.ClientID
	AdOpModel.AdID = id

	err := AdOpModel.AddAdOperateView()
	if err != nil {
		ctl.Error(err)
		return
	}
	ctl.Success("success")
}

func AdFrontOperateClick(c *gin.Context) {
	ctl := NewController(c)
	id := String2Uint(c.Param("id"))

	var AdOpModel models.AdOperate
	AdOpModel.UserID = Int642String(ctl.UserID)
	AdOpModel.ClientID = ctl.ClientID
	AdOpModel.AdID = id

	err := AdOpModel.AddAdOperateClick()
	if err != nil {
		ctl.Error(err)
		return
	}
	ctl.Success("success")
}
