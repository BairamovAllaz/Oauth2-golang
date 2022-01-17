package handler

import (

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	// "os"
)

var (
	Randomkey = "randomkey"
)




type Handler struct{ 
	GoogleAuth *oauth2.Config
}

func NewHandler(Googleauth *oauth2.Config) *Handler { 
	return &Handler{GoogleAuth: Googleauth}
}

func(h *Handler) InitRoutes() *gin.Engine {
	

	router := gin.Default();
	router.GET("/",h.Home);
	router.POST("/login",h.Login)
	router.POST("/callback",h.Callback)
	return router;
}


