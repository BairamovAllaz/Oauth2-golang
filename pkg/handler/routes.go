package handler

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	// "golang.org/x/oauth2"
)

func (h *Handler) Home(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{
		"link" : "http://localhost:8000/login",
	})
}
func(h *Handler)Login(c *gin.Context) {
	url := 	h.GoogleAuth.AuthCodeURL(Randomkey);
	fmt.Println("url: " , url);
	c.Redirect(http.StatusMovedPermanently,url);
}

func(h *Handler)Callback(c *gin.Context) {
	if c.Request.FormValue("state") != Randomkey {
		log.Fatalf("state is not valid")
		c.Redirect(http.StatusMovedPermanently,"/"); 
		return;
	}

	token,err := h.GoogleAuth.Exchange(context.TODO(),c.Request.FormValue("code"));

	if err != nil { 
		log.Fatalf("Could not get token %s", err.Error());
		c.Redirect(http.StatusMovedPermanently,"/"); 
		return;
	}
	resp,err := http.Get("https://googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken);
	if err != nil { 
		log.Fatalf("Could not get request %s", err.Error());
		c.Redirect(http.StatusMovedPermanently,"/"); 
		return;
	}

	defer resp.Body.Close();

	user,err := ioutil.ReadAll(resp.Body);
	if err != nil { 
		log.Fatalf("Could not get user %s", err.Error());
		c.Redirect(http.StatusMovedPermanently,"/"); 
		return;
	}

	c.JSON(http.StatusOK,gin.H{ 
		"user" : user,
	})

}
