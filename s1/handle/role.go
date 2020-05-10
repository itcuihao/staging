package handle

import (
	"strconv"

	"github.com/itcuihao/staging/s1/common"
	"github.com/itcuihao/staging/s1/dao"

	"github.com/gin-gonic/gin"
)

type RoleHandle struct {
	*handle
}

func NewRoleHandle(db *dao.Dao) *RoleHandle {
	return &RoleHandle{
		handle: newHandle(db),
	}
}

func (h *RoleHandle) GetRole(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId, _ := strconv.Atoi(userIdStr)
	if userId <= 0 {
		common.Log.Info(h)
		h.ClientBadParam(c, "user id empty")
		return
	}
	roleIds, err := h.DB.GetUserRoleIds(userId)
	if err != nil {
		h.ServerBusy(c, err.Error())
		return
	}
	roles, err := h.DB.GetRoleByIds(roleIds...)
	if err != nil {
		h.ServerBusy(c, err.Error())
		return
	}
	h.JsonReply(c, roles)
}
