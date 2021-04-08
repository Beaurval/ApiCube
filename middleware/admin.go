package middleware

import (
	"github.com/gin-gonic/gin"
)

var Authorizator = func(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*User); ok && v.RangID >= 3 {
		return true
	}
	return false
}
