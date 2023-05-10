package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rookie-ninja/rk-boot/v2"
	"github.com/rookie-ninja/rk-gin/v2/boot"
	"myapp-go/internal/infrustructure/rk"
	_ "myapp-go/internal/infrustructure/rk"
	"myapp-go/internal/routes"
	"net/http"
)

func main() {
	// Create a new boot instance.
	boot := rkboot.NewBoot()
	// Bootstrap
	boot.Bootstrap(context.Background())
	rk.Init()
	ginEntry := rkgin.GetGinEntry("myapp")
	myapp := ginEntry.Router.Group("/myapp/api/v1")
	{
		genealogy := myapp.Group("/genealogy")
		{
			genealogy.POST("/paging", routes.Paging)
			genealogy.GET("/test", Greeter)
		}
	}
	boot.WaitForShutdownSig(context.Background())

}

// Greeter handler
// @Summary Greeter
// @Id 1
// @Tags Hello
// @version 1.0
// @Param name query string true "name"
// @produce application/json
// @Success 200 {object} GreeterResponse
// @Router /v1/greeter [get]
func Greeter(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &GreeterResponse{
		Message: fmt.Sprintf("Hello %s!", ctx.Query("name")),
	})
}

type GreeterResponse struct {
	Message string
}
