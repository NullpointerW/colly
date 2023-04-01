package main

import (
	"colly/cache"
	"fmt"
	

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/v1/:sc", func(c *gin.Context) {
		t := c.Param("sc")
		hs := cache.Rows.Search(t)
		if len := len(hs); len != 0 {
			c.Header("Content-Type", "text/html; charset=utf-8")
			if len > 1 {
				var html string
				for _, r := range hs {
					html += fmt.Sprintf(`<a href="%s">%s</a><br>`, r.Link, r.Title)
				}
				c.String(200, html)
			} else {
				c.Redirect(301, hs[0].Link)
			}
		} else {
			msg := fmt.Sprintf("tilte \"%s\" :not found", t)
			c.JSON(200, gin.H{
				"msg": msg,
			})
		}

	})
	r.Run(":8964") 
}
