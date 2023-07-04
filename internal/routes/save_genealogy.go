package routes

import (
	"github.com/gin-gonic/gin"
	"myapp-go/internal/domain"
	"myapp-go/internal/infrustructure/rk/log"
	"myapp-go/internal/response"
	"myapp-go/internal/service/genealogy"
)

// SaveGenealogy handler
// @Summary 保存家谱
// @Tags Genealogy
// @version 1.0
// @produce application/json
// @Param request body domain.WxGenealogy true "query params"
// @Success 200 {object} response.Result
// @Router /myapp/api/v1/genealogy/save [post]
func SaveGenealogy(c *gin.Context) {
	req := new(domain.WxGenealogy)
	if err := c.ShouldBindJSON(req); err != nil {
		response.Fail(c, err.Error())
		return
	}
	err := genealogy.Save(req)
	if err != nil {
		log.Errorf("Save err %v", err)
		response.Fail(c, err.Error())
	} else {
		response.Success(c, nil)
	}
}
