package genealogy

import (
	"myapp-go/internal/domain"
	"myapp-go/internal/infrustructure/persistence"
)

func Save(req *domain.WxGenealogy) error {
	if req.ID == 0 {
		return persistence.GenealogyRepository.Save(req)
	}
	//if req {
	//
	//}
	//rsp, err := persistence.GenealogyRepository.FindById(req.ID)
	//if err != nil {
	//	return err
	//}
	//err := persistence.GenealogyRepository.Save(req)
	return nil
}
