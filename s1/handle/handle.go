package handle

import (
	"net/http"
	"strconv"

	"github.com/itcuihao/staging/s1/common"
	"github.com/itcuihao/staging/s1/dao"

	"github.com/gin-gonic/gin"
)

type handle struct {
	DB *dao.Dao
}

func newHandle(db *dao.Dao) *handle {
	return &handle{DB: db.New()}
}

func (s *handle) ClientErr(c *gin.Context, code int, msg string) {
	s.Err(c, http.StatusBadRequest, code, msg)
}

func (s *handle) ClientBadParam(c *gin.Context, msg string) {
	s.Err(c, http.StatusBadRequest, common.ErrCodeBadParam, msg)
}

func (s *handle) AuthFailed(c *gin.Context) {
	s.Err(c, http.StatusForbidden, common.ErrCodePermissionDenied, "无权限访问")
}

func (s *handle) ServerErr(c *gin.Context, code int, msg string) {
	s.Err(c, http.StatusInternalServerError, code, msg)
}

func (s *handle) ServerBusy(c *gin.Context, msg string) {
	s.Err(c, http.StatusInternalServerError, common.ErrCodeServiceBusy, msg)
}

func (s *handle) Err(c *gin.Context, statusCode, code int, msg string) {
	c.JSON(statusCode, gin.H{
		"code":  code,
		"error": msg,
	})
}

func (s *handle) JsonReply(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data": obj,
	})
}

func (s *handle) JsonWithPaging(c *gin.Context, obj, paging interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data":   obj,
		"paging": paging,
	})
}

func (s *handle) JsonOk(c *gin.Context) {
	s.JsonReply(c, gin.H{
		"msg": "ok",
	})
}

func (s *handle) Paging(c *gin.Context) (limit, offset int) {
	p, _ := strconv.Atoi(c.Query("page"))
	ps, _ := strconv.Atoi(c.Query("page_size"))

	if p <= 0 {
		return 10, 0
	}
	if ps <= 0 || ps >= 50 {
		ps = 20
	}
	limit = ps
	offset = (p - 1) * ps
	return
}
