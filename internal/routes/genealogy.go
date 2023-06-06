package routes

import (
	"github.com/gin-gonic/gin"
	"myapp-go/internal/domain/valobj"
	"myapp-go/internal/infrustructure/rk/log"
	"myapp-go/internal/response"
	"myapp-go/internal/service/genealogy"
)

// Assemble handler
// @Summary 分页聚合查询家谱数据
// @Tags Genealogy
// @version 1.0
// @produce application/json
// @Param request body valobj.GenealogyPagingReq true "query params"
// @Success 200 {object} response.Result{Data=[]genealogy.AssembleRes}
// @Router /myapp/api/v1/genealogy/assemble [post]
func Assemble(c *gin.Context) {
	req := new(valobj.GenealogyPagingReq)
	if err := c.ShouldBindJSON(req); err != nil {
		response.Fail(c, err.Error())
		return
	}
	results, err := genealogy.Assemble(req)
	if err != nil {
		log.Errorf("FindAll err %v", err)
		response.Fail(c, err.Error())
	} else {
		response.Success(c, results)
	}
}
