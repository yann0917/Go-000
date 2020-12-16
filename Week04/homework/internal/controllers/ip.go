/*
 * @Author: qiuling
 * @Date: 2019-09-16 18:19:04
 * @Last Modified by: qiuling
 * @Last Modified time: 2019-12-04 16:13:38
 */
package controllers

import (
	"errors"

	"github.com/wlxpkg/base/log"
	. "github.com/wlxpkg/zwyd"

	"github.com/gin-gonic/gin"
	"github.com/ipipdotnet/ipdb-go"
)

var ipDB *ipdb.City

func init() {
	var err error
	ipDB, err = ipdb.NewCity("../config/ipdb/ipipfree.ipdb")
	if err != nil {
		log.Err(err)
	}
}

func Ip2Address(ctx *gin.Context) {
	c := NewController(ctx)

	if !c.CheckSecret() {
		c.Error(errors.New("ERR_UNAUTHORIZED"))
		return
	}

	ip := c.Getd("ip", "")

	if ip == "" {
		c.Error(errors.New("ERR_PARAM"))
		return
	}

	city, err := ipDB.FindMap(ip, "CN")

	if err != nil {
		c.Error(err)
		return
	}
	// R(err, "err")
	// R(city, "city")

	c.Success(city)
}
