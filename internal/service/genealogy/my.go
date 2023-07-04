package genealogy

import (
	"myapp-go/internal/domain"
	"myapp-go/internal/domain/valobj"
	"myapp-go/internal/infrustructure/persistence"
)

func MyGenealogy(req *valobj.MyGenealogyReq) (rsp []*domain.WxGenealogy, err error) {
	myRsp, err := persistence.GenealogyMembersRepository.FindByUserId(req.UserId)
	if err != nil {
		return nil, err
	}
	var gids []int64
	for _, v := range myRsp {
		gids = append(gids, v.GenealogyID)
	}
	rsp, err = persistence.GenealogyRepository.FindByIds(gids)
	return
}
