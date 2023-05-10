package persistence

import (
	"gorm.io/gorm"
	"myapp-go/internal/domain"
	"myapp-go/internal/domain/valobj"
	"myapp-go/internal/infrustructure/rk/log"
)

// GenealogyRepository 凭证仓储接口
type _GenealogyRepository struct {
	db *gorm.DB
}

// NewGenealogyRepository 构造函数
func NewGenealogyRepository(db *gorm.DB) *_GenealogyRepository {
	return &_GenealogyRepository{db: db}
}

func (gr *_GenealogyRepository) Paging(req *valobj.GenealogyPagingReq) (rsp *valobj.GenealogyPagingRes, err error) {
	rsp = new(valobj.GenealogyPagingRes)
	var total int64
	mgr := domain.WxGenealogyMgr(gr.db)

	if tx := mgr.Where(req.GenealogyReq).Count(&total); tx.Error != nil {
		log.Errorf("_GenealogyRepository count error %v", tx.Error)
		return nil, tx.Error
	}
	var genealogies []*domain.WxGenealogy
	if tx := mgr.Limit(req.PagingReq.Limit).Offset(req.PagingReq.Offset).Find(&genealogies); tx.Error != nil {
		log.Errorf("_GenealogyRepository paging error %v", tx.Error)
		return nil, tx.Error
	}
	rsp.Total = total
	rsp.Genealogies = genealogies
	return
}