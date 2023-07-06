package routes

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"myapp-go/internal/domain/valobj"
	"myapp-go/internal/infrustructure/rk/log"
	"myapp-go/internal/response"
	"myapp-go/internal/service/cos"
)

// CosUpload handler
// @Summary 上传cos
// @Tags Genealogy
// @version 1.0
// @produce application/json
// @Success 200 {object} response.Result{Data=[]valobj.UploadRsp}
// @Router /myapp/api/v1/cos/upload [post]
func CosUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Errorf("CosUpload err %v", err)
		response.Fail(c, err.Error())
	} else {
		f, err := file.Open()
		defer func(f multipart.File) {
			err := f.Close()
			if err != nil {
				response.Fail(c, err.Error())
			}
		}(f)
		if err != nil {
			log.Errorf("CosUpload err %v", err)
			response.Fail(c, err.Error())
		}
		req := &valobj.UploadReq{
			Name: file.Filename,
			File: f,
		}
		rsp, err := cos.Upload(req)
		if err != nil {
			log.Errorf("CosUpload err %v", err)
			response.Fail(c, err.Error())
		}
		response.Success(c, rsp)
	}

}
