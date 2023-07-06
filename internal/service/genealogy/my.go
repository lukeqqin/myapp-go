package genealogy

import (
	"myapp-go/internal/domain"
	"myapp-go/internal/infrustructure/persistence"
)

func MyGenealogy() (rsp []*domain.WxGenealogy, err error) {

	myRsp, err := persistence.GenealogyMembersRepository.FindByUserId(1)
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
