package routes

import (
	"github.com/gin-gonic/gin"
	"myapp-go/internal/domain/valobj"
	"myapp-go/internal/infrustructure/rk/log"
	"myapp-go/internal/response"
	"myapp-go/internal/service/genealogy"
)

// MyGenealogy handler
// @Summary 查询自己的所有族谱
// @Tags Genealogy
// @version 1.0
// @produce application/json
// @Param request body valobj.MyGenealogyReq true "query params"
// @Success 200 {object} response.Result{Data=[]domain.WxGenealogy}
// @Router /myapp/api/v1/genealogy/my [post]
func MyGenealogy(c *gin.Context) {
	req := new(valobj.MyGenealogyReq)
	if err := c.ShouldBindJSON(req); err != nil {
		response.Fail(c, err.Error())
		return
	}
	results, err := genealogy.MyGenealogy(req)
	if err != nil {
		log.Errorf("MyGenealogy err %v", err)
		response.Fail(c, err.Error())
	} else {
		response.Success(c, results)
	}
}
