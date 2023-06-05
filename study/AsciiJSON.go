/*
	JSON是一种轻量级的数据交换格式，全称为JavaScript Object Notation，
	它具有易于读写、易于解析等特点。JSON通常用于客户端和服务器之间的数据传输。
*/

// 通过JSON方法向网页传输ASCII码信息

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
