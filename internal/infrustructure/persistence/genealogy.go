package persistence

import (
	"fmt"
	"gorm.io/gorm"
	"myapp-go/internal/domain"
	"myapp-go/internal/domain/valobj"
	"myapp-go/internal/infrustructure/rk/log"
)

// GenealogyRepository 凭证仓储接口
type _GenealogyRepository struct {
	db *gorm.DB
}

func newGenealogyRepository(db *gorm.DB) *_GenealogyRepository {
	return &_GenealogyRepository{db: db}
}

func (gr *_GenealogyRepository) Paging(req *valobj.GenealogyPagingReq) (rsp *valobj.GenealogyPagingRsp, err error) {
	var total int64
	mgr := domain.WxGenealogyMgr(gr.db)
	var tx *gorm.DB
	if req.GenealogyReq.Title != "" {
		tx = mgr.Where("title like ?", fmt.Sprint("%", req.GenealogyReq.Title, "%"))
	}
	if tx = mgr.Count(&total); tx.Error != nil {
		log.Errorf("_GenealogyRepository count error %v", tx.Error)
		return nil, tx.Error
	}
	var genealogies []*domain.WxGenealogy
	if tx := mgr.Limit(req.PagingReq.Limit).Offset(req.PagingReq.Offset).Find(&genealogies); tx.Error != nil {
		log.Errorf("_GenealogyRepository paging error %v", tx.Error)
		return nil, tx.Error
	}
	rsp = new(valobj.GenealogyPagingRsp)
	rsp.Total = total
	rsp.Genealogies = genealogies
	return
}

func (gr *_GenealogyRepository) FindByIds(ids []int64) (rsp []*domain.WxGenealogy, err error) {
	mgr := domain.WxGenealogyMgr(gr.db)
	rsp, err = mgr.GetBatchFromID(ids)
	if err != nil {
		log.Errorf("_GenealogyRepository FindByIds error %v", err)
		return nil, err
	}
	return
}

func (gr *_GenealogyRepository) FindById(ids int64) (rsp *domain.WxGenealogy, err error) {
	mgr := domain.WxGenealogyMgr(gr.db)
	result, err := mgr.FetchByPrimaryKey(ids)
	if err != nil {
		log.Errorf("_GenealogyRepository FetchByPrimaryKey error %v", err)
		return nil, err
	}
	rsp = &result
	return
}

func (gr *_GenealogyRepository) Save(req *domain.WxGenealogy) error {
	mgr := domain.WxGenealogyMgr(gr.db)
	if tx := mgr.Save(req); tx.Error != nil {
		log.Errorf("_GenealogyRepository add error %v", tx.Error)
		return tx.Error
	}
	return nil
}
