package member

import (
	"myapp-go/internal/domain"
	"myapp-go/internal/infrustructure/persistence"
)

func Save(req *domain.WxGenealogyMember) error {
	err := persistence.GenealogyMembersRepository.Save(req)
	return err
}
