package persistence

import (
	"gorm.io/gorm"
	"myapp-go/internal/domain"
	"myapp-go/internal/domain/valobj"
	"myapp-go/internal/infrustructure/rk/log"
)

// GenealogyMembersRepository 凭证仓储接口
type _GenealogyMembersRepository struct {
	db *gorm.DB
}

func newGenealogyMembersRepository(db *gorm.DB) *_GenealogyMembersRepository {
	return &_GenealogyMembersRepository{db: db}
}

func (gr *_GenealogyMembersRepository) Paging(req *valobj.GenealogyMembersPagingReq) (rsp *valobj.GenealogyMembersPagingRes, err error) {
	var total int64
	mgr := domain.WxGenealogyMembersMgr(gr.db)
	var tx *gorm.DB
	if req.GenealogyId != 0 {
		tx = mgr.Where("genealogy_id = ?", req.GenealogyId)
	}
	if tx = mgr.Count(&total); tx.Error != nil {
		log.Errorf("_GenealogyMembersRepository count error %v", tx.Error)
		return nil, tx.Error
	}
	var genealogyMembers []*domain.WxGenealogyMembers
	if tx := mgr.Limit(req.PagingReq.Limit).Offset(req.PagingReq.Offset).Find(&genealogyMembers); tx.Error != nil {
		log.Errorf("_GenealogyMembersRepository paging error %v", tx.Error)
		return nil, tx.Error
	}
	rsp = new(valobj.GenealogyMembersPagingRes)
	rsp.Total = total
	rsp.GenealogyMembers = genealogyMembers
	return
}
