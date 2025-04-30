package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 定义一个接口
	r.Any("/debug", func(c *gin.Context) {
		// 打印请求方法
		fmt.Println("Method:", c.Request.Method)

		// 打印请求URL
		fmt.Println("URL:", c.Request.URL.String())

		// 打印所有请求头
		fmt.Println("Headers:")
		for k, v := range c.Request.Header {
			fmt.Printf("  %s: %v\n", k, v)
		}

		// 打印Query参数 (?后面的参数)
		fmt.Println("Query Parameters:")
		for k, v := range c.Request.URL.Query() {
			fmt.Printf("  %s: %v\n", k, v)
		}

		// 打印Form参数 (POST表单、x-www-form-urlencoded)
		c.Request.ParseForm() // 重要
		fmt.Println("Form Parameters:")
		for k, v := range c.Request.Form {
			fmt.Printf("  %s: %v\n", k, v)
		}

		// 打印Body内容
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println("Failed to read body:", err)
		} else {
			fmt.Println("Body:")
			fmt.Println(string(bodyBytes))
		}

		// 返回简单响应
		c.JSON(http.StatusOK, gin.H{
			"message": "请求数据已经打印到后台控制台！",
		})
	})

	r.Run(":8080") // 启动在 8080 端口
}
