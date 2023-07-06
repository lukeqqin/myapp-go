package routes

import (
	"github.com/gin-gonic/gin"
	"myapp-go/internal/infrustructure/rk/log"
	"myapp-go/internal/response"
	"myapp-go/internal/service/genealogy"
)

// MyGenealogy handler
// @Summary 查询自己的所有族谱
// @Tags Genealogy
// @version 1.0
// @produce application/json
// @Success 200 {object} response.Result{Data=[]domain.WxGenealogy}
// @Router /myapp/api/v1/genealogy/my [post]
func MyGenealogy(c *gin.Context) {
	results, err := genealogy.MyGenealogy()
	if err != nil {
		log.Errorf("MyGenealogy err %v", err)
		response.Fail(c, err.Error())
	} else {
		response.Success(c, results)
	}
}
