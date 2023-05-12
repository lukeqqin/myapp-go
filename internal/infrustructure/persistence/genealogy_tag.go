package persistence

import (
	"gorm.io/gorm"
	"myapp-go/internal/domain"
	"myapp-go/internal/infrustructure/rk/log"
)

// _GenealogyTagRepository 凭证仓储接口
type _GenealogyTagRepository struct {
	db *gorm.DB
}

func newGenealogyTagRepository(db *gorm.DB) *_GenealogyTagRepository {
	return &_GenealogyTagRepository{db: db}
}

func (gr *_GenealogyTagRepository) FetchIndexByWxGenealogyTagsGenealogyIDIndex(id int64) (rsp []*domain.WxGenealogyTags, err error) {
	mgr := domain.WxGenealogyTagsMgr(gr.db)
	rsp, err = mgr.FetchIndexByWxGenealogyTagsGenealogyIDIndex(id)
	if err != nil {
		log.Errorf("_UserRepository FetchByPrimaryKey error %v", err)
		return nil, err
	}
	return
}
