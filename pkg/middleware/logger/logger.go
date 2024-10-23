package logger

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grafana/loki-client-go/loki"
	"github.com/prometheus/common/model"
)

func Logger(lokiClient *loki.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latency := endTime.Sub(startTime)
		statusCode := c.Writer.Status()

		log.Printf("| %3d | %13v | %-7s %s\n",
			statusCode,
			latency,
			c.Request.Method,
			c.Request.URL.Path,
		)

		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()

		var deviceType string
		if strings.Contains(strings.ToLower(userAgent), "mobile") {
			deviceType = "mobile"
		} else {
			deviceType = "desktop"
		}

		labels := model.LabelSet{
			"job":     "gin-server",
			"status":  model.LabelValue(fmt.Sprintf("%d", statusCode)),
			"method":  model.LabelValue(c.Request.Method),
			"path":    model.LabelValue(c.Request.URL.Path),
			"ip":      model.LabelValue(clientIP),
			"device":  model.LabelValue(deviceType),
			"latency": model.LabelValue(latency.String()),
		}

		// ایجاد لاگ برای ارسال به Loki
		line := fmt.Sprintf("| %3d | %13v | %-7s %s", statusCode, latency, c.Request.Method, c.Request.URL.Path)

		// Send the log to Loki using Push method
		err := lokiClient.Handle(labels, time.Now(), line)
		if err != nil {
			log.Printf("Error sending log to Loki: %v", err)
		} else {
			log.Printf("Log successfully sent to Loki")
		}
	}
}
