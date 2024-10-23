package pagination

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		pageStr := c.Query("page")
		limitStr := c.Query("limit")

		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}

		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit < 1 {
			limit = 10 // مقدار پیش‌فرض
		}

		// ذخیره اطلاعات pagination در context
		c.Set("page", page)
		c.Set("limit", limit)

		c.Next() // ادامه به handler بعدی
	}
}
