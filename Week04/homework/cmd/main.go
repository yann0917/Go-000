/*
 * @Description:
 * @Author: maxinyu
 * @Date: 2019-06-14 14:52:52
 * @LastEditTime: 2019-07-02 15:35:31
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"homework/internal/router"

	"github.com/gookit/config/v2"
	. "github.com/wlxpkg/base"
)

func main() {
	defer DB.Close()

	r := router.InitRouter()
	// Listen and Server in 0.0.0.0:8080
	_ = r.Run(":" + config.Getenv("PORT", "8000"))
}
