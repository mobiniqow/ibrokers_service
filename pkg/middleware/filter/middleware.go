package filter

import (
	"fmt"
	"ibrokers_service/pkg/middleware/filter/operators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func QueryFilterMiddleware(filterMapper Mapper) gin.HandlerFunc {
	return func(c *gin.Context) {
		var blocks []operators.FilterBlock
		if c.Request.Method == http.MethodGet {
			var filters []string
			for key, values := range c.Request.URL.Query() {
				if len(values) > 0 {
					filters = append(filters, fmt.Sprintf("%s=%s", key, values[0]))
				}
			}
			blocks = filterMapper.Convert(filters)
		}
		c.Set("filters", blocks)
		c.Next()
	}
}
