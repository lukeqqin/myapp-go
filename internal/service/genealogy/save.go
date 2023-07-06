package genealogy

import (
	"myapp-go/internal/domain"
	"myapp-go/internal/infrustructure/persistence"
)

func Save(req *domain.WxGenealogy) error {
	err := persistence.GenealogyRepository.Save(req)
	if err != nil {
		return err
	}
	if req.Cover == "" {
		return nil
	}
	cover, err := persistence.CosRepository.FetchUniqueByWxCosURL(req.Cover)
	if err != nil {
		return err
	}
	cover.Status = 1
	return persistence.CosRepository.Save(&cover)
}
