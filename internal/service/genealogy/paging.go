package genealogy

import (
	"myapp-go/internal/domain/valobj"
	"myapp-go/internal/infrustructure/persistence"
)

func Paging(req *valobj.GenealogyPagingReq) (rsp *valobj.GenealogyPagingRes, err error) {
	rsp, err = persistence.GenealogyRepository.Paging(req)
	return
}
