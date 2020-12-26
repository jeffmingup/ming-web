package main

import (
	"ming"
	"net/http"
)

func main() {
	r := ming.New()
	r.GET("/", func(c *ming.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *ming.Context) {
		// expect /hello?name=mingktutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *ming.Context) {
		c.JSON(http.StatusOK, ming.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
