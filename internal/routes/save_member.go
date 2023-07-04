package routes

import (
	"github.com/gin-gonic/gin"
	"myapp-go/internal/domain"
	"myapp-go/internal/infrustructure/rk/log"
	"myapp-go/internal/response"
	"myapp-go/internal/service/member"
)

// SaveGenealogyMember handler
// @Summary 保存家谱成员
// @Tags Genealogy
// @version 1.0
// @produce application/json
// @Param request body domain.WxGenealogyMember true "query params"
// @Success 200 {object} response.Result
// @Router /myapp/api/v1/genealogy/member/save [post]
func SaveGenealogyMember(c *gin.Context) {
	req := new(domain.WxGenealogyMember)
	if err := c.ShouldBindJSON(req); err != nil {
		response.Fail(c, err.Error())
		return
	}
	err := member.Save(req)
	if err != nil {
		log.Errorf("Save err %v", err)
		response.Fail(c, err.Error())
	} else {
		response.Success(c, nil)
	}
}
