package routes

import (
	"github.com/gin-gonic/gin"
	"myapp-go/internal/domain/valobj"
	"myapp-go/internal/infrustructure/rk/log"
	"myapp-go/internal/response"
	"myapp-go/internal/service/genealogy"
)

// Paging handler
// @Summary 分页查询家谱数据
// @Tags Genealogy
// @version 1.0
// @produce application/json
// @Param request body valobj.GenealogyPagingReq true "query params"
// @Success 200 {object} response.Result{Data=[]valobj.GenealogyPagingRes}
// @Router /myapp/api/v1/genealogy/paging [post]
func Paging(c *gin.Context) {
	req := new(valobj.GenealogyPagingReq)
	if err := c.ShouldBindJSON(req); err != nil {
		response.Fail(c, err.Error())
		return
	}
	results, err := genealogy.Paging(req)
	if err != nil {
		log.Errorf("FindAll err %v", err)
		response.Fail(c, err.Error())
	} else {
		response.Success(c, results)
	}
}
