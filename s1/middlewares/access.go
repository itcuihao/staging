package middlewares

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/itcuihao/staging/s1/common"

	"github.com/gin-gonic/gin"
)

type LogResponseWriter struct {
	gin.ResponseWriter
	rspBody *bytes.Buffer
}

func (r *LogResponseWriter) Write(p []byte) (int, error) {
	r.rspBody.Write(p)
	return r.ResponseWriter.Write(p)
}

func AccessMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var startTime = time.Now()
		var reqBody []byte
		logRspWriter := &LogResponseWriter{c.Writer, bytes.NewBufferString("")}
		c.Writer = logRspWriter
		if c.Request.Method != http.MethodGet {
			reqBody, _ = ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		}

		c.Next()

		elapsed := float64(time.Now().Sub(startTime).Nanoseconds()) / 1e6
		common.Log.Debugf("%s %s %.3fms %s %s %d %s", c.Request.Method, c.Request.URL.String(),
			elapsed, c.ClientIP(), reqBody, logRspWriter.Status(), logRspWriter.rspBody)
	}
}
