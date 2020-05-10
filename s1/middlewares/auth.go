package middlewares

import (
	"errors"
	"net/http"
	"time"

	"github.com/itcuihao/staging/s1/common"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	AuthKeyAccount  = "account"
	AuthKeyUserId   = "userId"
	AuthKeyUserRole = "userRole"
)

const (
	UserRoleGuest      = 0
	UserRoleUser       = 1
	UserRoleAdmin      = 8 // 普通管理员
	UserRoleSuperAdmin = 9 // 网站管理员
)

var (
	SigningKey = []byte("OM_SIGNING_KEY")
)

var WhiteUserIds = map[int]bool{
	4: true,
}

type CustomClaims struct {
	Account  string `json:"account"`
	UserId   int    `json:"user_id"`
	UserRole string `json:"user_role"`
	ExpireAt int64  `json:"expire_at"`
}

func (c *CustomClaims) Valid() error {
	if WhiteUserIds[c.UserId] {
		// 白名单用户真按时不校验超时
		common.Log.Warningf("token expired. user id:%d", c.UserId)
		return nil
	}

	if c.ExpireAt < time.Now().Unix() {
		return errors.New("token expired")
	}
	return nil
}

func (m Middleware) AuthMiddleware(c *gin.Context) {
	if SkipHandler(c, m.Skippers...) {
		c.Next()
		return
	}
	accessToken := c.Request.Header.Get("access_token")
	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  common.ErrCodeUnauthorized,
			"error": "empty token",
		})
		c.Abort()
		return
	}

	token, err := jwt.ParseWithClaims(accessToken, &CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return SigningKey, nil
		})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  common.ErrCodeUnauthorized,
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  common.ErrCodeUnauthorized,
			"error": "token invalid",
		})
		c.Abort()
		return
	}

	c.Set(AuthKeyAccount, claims.Account)
	c.Set(AuthKeyUserId, claims.UserId)
	c.Set(AuthKeyUserRole, claims.UserRole)
	c.Next()
}

func GenToken(userId int, account string, role string) (string, int64) {
	expireAt := time.Now().Add(7 * 24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomClaims{
		UserId:   userId,
		Account:  account,
		UserRole: role,
		ExpireAt: expireAt,
	})

	newToken, _ := token.SignedString(SigningKey)
	return newToken, expireAt
}

func RoleUserMiddleware(c *gin.Context) {
	userRole := c.GetInt(AuthKeyUserRole)
	if userRole < UserRoleUser {
		c.JSON(http.StatusForbidden, gin.H{
			"code":  common.ErrCodePermissionDenied,
			"error": "游客无权限",
		})
		c.Abort()
		return
	}

	c.Next()
}

func RoleAdminMiddleware(c *gin.Context) {
	userRole := c.GetInt(AuthKeyUserRole)
	if userRole < UserRoleAdmin {
		c.JSON(http.StatusForbidden, gin.H{
			"code":  common.ErrCodePermissionDenied,
			"error": "仅限管理员操作",
		})
		c.Abort()
		return
	}

	c.Next()
}
