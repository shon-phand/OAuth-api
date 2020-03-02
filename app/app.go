package app

import (
	"github.com/gin-gonic/gin"
	"github.com/shon-phand/OAuth-api/domain/access_token"
	"github.com/shon-phand/OAuth-api/http"
	"github.com/shon-phand/OAuth-api/repository/db"
)

var (
	r = gin.Default()
)

func StartApplication() {
	atService := access_token.NewService(db.NewRepository())
	atHanlder := http.NewHandler(atService)

	r.GET("/oauth/access_token/:acess_token_id", atHanlder.GetById)

	r.Run(":8080")
}
