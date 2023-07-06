package persistence

import (
	"gorm.io/gorm"
	"myapp-go/internal/domain"
	"myapp-go/internal/infrustructure/rk/log"
)

// CosRepository 凭证仓储接口
type _CosRepository struct {
	db *gorm.DB
}

func NewCosRepository(db *gorm.DB) *_CosRepository {
	return &_CosRepository{db: db}
}
func (gr *_CosRepository) Save(req *domain.WxCos) error {
	mgr := domain.WxCosMgr(gr.db)
	if tx := mgr.Save(req); tx.Error != nil {
		log.Errorf("_CosRepository add error %v", tx.Error)
		return tx.Error
	}
	return nil
}

func (gr *_CosRepository) FetchUniqueByWxCosURL(url string) (rsp domain.WxCos, err error) {
	mgr := domain.WxCosMgr(gr.db)
	rsp, err = mgr.FetchUniqueByWxCosURL(url)
	return
}
