package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetPage(c *gin.Context)int{
	page, _ :=com.StrTo(c.Query("page")).Int()
	if page == 0 {
		page = 1
	}
	return page
}
