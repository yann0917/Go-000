package models

import (
	"errors"
	"time"

	. "github.com/wlxpkg/base"

	"github.com/jinzhu/gorm"
)

// AdType 4个固定的广告位类型
var AdType = [4]uint{1, 2, 3, 4}

// Ad 广告位的 model
type Ad struct {
	ID        uint      `json:"id"`
	Type      uint      `json:"type"`
	Code      string    `json:"code"`
	Title     string    `json:"title"`
	Image     string    `json:"image"`
	Link      *string   `json:"link"`
	Stime     string    `gorm:"-" json:"stime"`
	Etime     string    `gorm:"-" json:"etime"`
	StartTime JSONTime  `json:"start_time"`
	EndTime   JSONTime  `json:"end_time"`
	Sort      *uint     `gorm:"default:0" json:"sort"`
	Status    *uint     `gorm:"default:0" json:"status"`
	Pv        int       `json:"pv"`
	Uv        int       `json:"uv"`
	Clicks    int       `json:"clicks"`
	ClickRate string    `gorm:"-" json:"click_rate"`
	CreatedAt JSONTime  `json:"created_at"`
	UpdatedAt JSONTime  `json:"updated_at"`
	DeletedAt *JSONTime `json:"-"`
}

type AdTwo struct {
	ID        uint   `json:"id"`
	Type      uint   `json:"type"`
	Code      string `json:"code"`
	Title     string `json:"title"`
	Image     string `json:"image"`
	Link      string `json:"link"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Sort      *uint  `gorm:"default:0" json:"sort"`
	Status    *uint  `gorm:"default:0" json:"status"`
	Pv        int    `json:"pv"`
	Uv        int    `json:"uv"`
	Clicks    int    `json:"clicks"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// ===================缓存参数==============
var Ttl = 60 * 60 * 24
var AdDetailCache = "ad:detail:"

func (Ad) TableName() string {
	return "common_ad"
}

func (AdTwo) TableName() string {
	return "common_ad"
}

func (AdModel *Ad) BeforeCreate(scope *gorm.Scope) (err error) {

	err = scope.SetColumn("CreatedAt", time.Now())
	err = scope.SetColumn("Status", 1)
	err = scope.SetColumn("Sort", 0)

	loc, _ := time.LoadLocation("Local")
	startTime, _ := time.ParseInLocation(TimeFormat, AdModel.Stime, loc)
	endTime, _ := time.ParseInLocation(TimeFormat, AdModel.Etime, loc)

	err = scope.SetColumn("StartTime", startTime)
	err = scope.SetColumn("EndTime", endTime)

	// err = scope.SetColumn("StartTime", String2Time(AdModel.Stime))
	// err = scope.SetColumn("EndTime", String2Time(AdModel.Etime))
	return
}

func (AdModel *Ad) BeforeUpdate(scope *gorm.Scope) (err error) {
	err = scope.SetColumn("UpdatedAt", time.Now())
	if AdModel.Stime != "" {
		err = scope.SetColumn("StartTime", String2Time(AdModel.Stime))
	}

	if AdModel.Etime != "" {
		err = scope.SetColumn("EndTime", String2Time(AdModel.Etime))
	}
	return
}

// AddNewAd 新增广告位
func (AdModel *Ad) AddNewAd() (id uint, err error) {

	// fmt.Printf("%+v", AdModel)

	sum := 0

	for _, typeValue := range AdType {

		if AdModel.Type == typeValue {
			sum++
		}
	}

	if sum == 0 {
		err = errors.New("广告位类型错误")
		return
	}

	if AdModel.Title == "" {
		err = errors.New("标题不能为空")
		return
	}

	result := DB.Create(&AdModel)

	if result.Error != nil {
		err = result.Error
		return
	}
	id = AdModel.ID
	return
}

// GetList 查看所有的广告位(后台专用)
func (AdModel *Ad) GetList(page int, limit int) (adList []Ad, count uint, err error) {

	if page < 1 {
		page = 1
	}
	var offset = (page - 1) * limit
	if err = DB.Where(AdModel).Order("sort, created_at desc").Limit(limit).Offset(offset).Find(&adList).Limit(-1).Offset(-1).Count(&count).Error; err != nil {
		return
	}
	return
}

// GetListForFront 查看所有的广告位（前端专用）
func (AdModel *Ad) GetListForFront(page int, limit int) (adList []Ad, err error) {

	if page < 1 {
		page = 1
	}
	var offset = (page - 1) * limit
	if err = DB.Where(AdModel).
		Where("(start_time < ? AND end_time > ?) OR (start_time < ? AND end_time IS NULL) OR (end_time > ? AND start_time IS NULL) OR (start_time IS NULL AND end_time IS NULL)", time.Now(), time.Now(), time.Now(), time.Now()).
		Order("sort, created_at desc").Limit(limit).Offset(offset).Find(&adList).Error; err != nil {
		return
	}
	return
}

// Edit 修改广告位
func (AdModel *Ad) Edit(id uint) (updateAd Ad, err error) {

	sum := 0

	for _, typeValue := range AdType {

		if AdModel.Type == typeValue {
			sum++
		}
	}

	if sum == 0 && AdModel.Type != 0 {
		err = errors.New("广告位类型错误")
		return
	}

	if err = DB.Select([]string{"id"}).First(&updateAd, id).Error; err != nil {
		return
	}
	if err = DB.Model(&updateAd).Updates(&AdModel).Error; err == nil {
		cache.Del(AdDetailCache + Uint642String((uint64)(id)))
	}
	DB.First(&updateAd, id)
	return
}

// GetInfo 获取一个广告位的信息
func (AdModel *Ad) GetInfo() (cacheData map[string]interface{}, err error) {

	var infoAd Ad

	var infoAdTwo AdTwo
	// 获取缓存
	err = cache.Get(AdDetailCache+Uint642String((uint64)(AdModel.ID)), &infoAdTwo)

	if infoAdTwo.ID > 0 {

		cacheData = infoAdTwo.FormatCacheData()
		return
	}

	if err = DB.Where(AdModel).First(&infoAd).Error; err == nil {
		// 设置config信息缓存 缓存一天
		cache.Set(AdDetailCache+Uint642String((uint64)(AdModel.ID)), infoAd, Ttl)

	}

	cacheData = infoAd.FormatCacheData()

	return cacheData, err
}

// Delete 删除广告位
func (AdModel *Ad) Delete() (Result uint, err error) {
	if err = DB.Where(AdModel).First(&Ad{}).Error; err != nil {
		return
	}
	if err = DB.Delete(&AdModel).Error; err == nil {
		cache.Del(AdDetailCache + Uint642String((uint64)(AdModel.ID)))
	}
	Result = AdModel.ID
	return
}

func (Admodel *Ad) StatisticPvUv(ad_id uint) {
	var adOpModel AdOperate
	adOpModel.AdID = ad_id
	uv, data := adOpModel.GetStatisticPvUv()
	Admodel.Uv = uv
	Admodel.Pv = data.View
	Admodel.Clicks = data.Click
	Admodel.Edit(ad_id)
	return
}

func (AdModel *Ad) FormatCacheData() (info map[string]interface{}) {

	info = make(map[string]interface{})

	info["id"] = AdModel.ID
	info["type"] = AdModel.Type
	info["code"] = AdModel.Code
	info["title"] = AdModel.Title
	info["image"] = AdModel.Image
	info["link"] = AdModel.Link
	info["start_time"] = AdModel.StartTime
	info["end_time"] = AdModel.EndTime
	info["sort"] = AdModel.Sort
	info["status"] = AdModel.Status
	info["created_at"] = AdModel.CreatedAt
	info["updated_at"] = AdModel.UpdatedAt
	return
}

func (data *AdTwo) FormatCacheData() (info map[string]interface{}) {

	info = make(map[string]interface{})

	info["id"] = data.ID
	info["type"] = data.Type
	info["code"] = data.Code
	info["title"] = data.Title
	info["image"] = data.Image
	info["link"] = data.Link
	info["start_time"] = data.StartTime
	info["end_time"] = data.EndTime
	info["sort"] = data.Sort
	info["status"] = data.Status
	info["created_at"] = data.CreatedAt
	info["updated_at"] = data.UpdatedAt
	return
}
