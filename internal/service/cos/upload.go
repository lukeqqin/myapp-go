package cos

import (
	"fmt"
	"myapp-go/internal/domain"
	"myapp-go/internal/domain/valobj"
	"myapp-go/internal/infrustructure/cos"
	"myapp-go/internal/infrustructure/persistence"
)

func Upload(req *valobj.UploadReq) (rsp *valobj.UploadRsp, err error) {
	userId := 1
	name := fmt.Sprint(userId, "/", req.Name)
	err = cos.Put(name, req.File)
	if err != nil {
		return nil, err
	}
	url := cos.GetUrl(name)
	r := &domain.WxCos{
		URL:      url,
		CreateBy: 1,
		UpdateTy: 1,
	}

	err = persistence.CosRepository.Save(r)
	if err != nil {
		return nil, err
	}
	rsp = &valobj.UploadRsp{
		Url: url,
	}
	return
}
