package router

import (
	"encoding/gob"

	"github.com/ddiogoo/trade/tree/master/auth-service/platform/authenticator"
	"github.com/ddiogoo/trade/tree/master/auth-service/platform/middleware"
	"github.com/ddiogoo/trade/tree/master/auth-service/web/app/callback"
	"github.com/ddiogoo/trade/tree/master/auth-service/web/app/home"
	"github.com/ddiogoo/trade/tree/master/auth-service/web/app/login"
	"github.com/ddiogoo/trade/tree/master/auth-service/web/app/logout"
	"github.com/ddiogoo/trade/tree/master/auth-service/web/app/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// New registers the routes and returns the router.
func New(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "web/static")
	router.LoadHTMLGlob("web/template/*")

	router.GET("/", home.Handler)
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/user", middleware.IsAuthenticated, user.Handler)
	router.GET("/logout", logout.Handler)

	return router
}
