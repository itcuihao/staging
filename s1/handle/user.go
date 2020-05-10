package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/itcuihao/staging/s1/common"
	"github.com/itcuihao/staging/s1/dao"
	"github.com/itcuihao/staging/s1/middlewares"
	"time"
)

type UserHandle struct {
	*handle
}

func NewUserHandle(db *dao.Dao) *UserHandle {
	return &UserHandle{
		handle: newHandle(db),
	}
}

func (h *UserHandle) GetUser(c *gin.Context) {
	id := c.Param("id")
	data, err := h.DB.GetUserById(id)
	if err != nil {
		common.Log.Errorf("user error: %v", err)
		h.ClientBadParam(c, err.Error())
		return
	}
	h.JsonReply(c, data)
}

type ReqLogin struct {
	Account string `json:"account"`
	Role    string `json:"role"`
}

func (h *UserHandle) Login(c *gin.Context) {
	reqLogin := &ReqLogin{}
	if err := c.BindJSON(reqLogin); err != nil {
		h.ClientBadParam(c, err.Error())
		return
	}

	if len(reqLogin.Account) == 0 {
		h.ClientBadParam(c, "empty account or password")
		return
	}

	user, err := h.DB.CreateOrFirstUserByAccount(reqLogin.Account)
	if err != nil {
		h.ServerBusy(c, "get account error")
		return
	}
	userRole := "admin"
	if time.Unix(user.TokenExpireAt, 0).Before(time.Now()) {
		token, tokenExpire := middlewares.GenToken(user.Id, user.Account, userRole)
		if err := h.DB.UpdateUserToken(user.Id, token, tokenExpire); err != nil {
			h.ServerBusy(c, "update token error")
			return
		}
		user.AccessToken = token
		user.TokenExpireAt = tokenExpire
	}
	// 设置权限
	c.Set(middlewares.AuthKeyUserRole, userRole)
	result := gin.H{
		"id":           user.Id,
		"account":      user.Account,
		"access_token": user.AccessToken,
		"role":         1,
	}
	h.JsonReply(c, result)
}
