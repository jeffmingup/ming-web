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

	r.GET("/hello/:name", func(c *ming.Context) {
		// expect /hello/mingktutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *ming.Context) {
		c.JSON(http.StatusOK, ming.H{"filepath": c.Param("filepath")})
	})
	r.Run(":9999")
}
