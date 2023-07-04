package persistence

import (
	"fmt"
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

func (gr *_GenealogyMembersRepository) Paging(req *valobj.GenealogyMembersPagingReq) (rsp *valobj.GenealogyMembersPagingRsp, err error) {
	var total int64
	mgr := domain.WxGenealogyMemberMgr(gr.db)
	var tx *gorm.DB
	if req.GenealogyId != 0 {
		tx = mgr.Where("genealogy_id = ?", req.GenealogyId)
	}
	if req.Name != "" {
		tx = mgr.Where("name like ?", fmt.Sprint("%", req.Name, "%"))
	}
	if tx = mgr.Count(&total); tx.Error != nil {
		log.Errorf("_GenealogyMembersRepository count error %v", tx.Error)
		return nil, tx.Error
	}
	var genealogyMembers []*domain.WxGenealogyMember
	if tx := mgr.Limit(req.PagingReq.Limit).Offset(req.PagingReq.Offset).Find(&genealogyMembers); tx.Error != nil {
		log.Errorf("_GenealogyMembersRepository paging error %v", tx.Error)
		return nil, tx.Error
	}
	rsp = new(valobj.GenealogyMembersPagingRsp)
	rsp.Total = total
	rsp.GenealogyMembers = genealogyMembers
	return
}

func (gr *_GenealogyMembersRepository) FindByUserId(userId int64) (rsp []*domain.WxGenealogyMember, err error) {
	var genealogyMembers []*domain.WxGenealogyMember
	mgr := domain.WxGenealogyMemberMgr(gr.db)
	if tx := mgr.Where("user_id = ?", userId).Find(&genealogyMembers); tx.Error != nil {
		log.Errorf("_GenealogyMembersRepository MyGenealogy error %v", tx.Error)
		return nil, tx.Error
	}
	return genealogyMembers, nil
}

func (gr *_GenealogyMembersRepository) Save(req *domain.WxGenealogyMember) error {
	mgr := domain.WxGenealogyMemberMgr(gr.db)
	if tx := mgr.Save(req); tx.Error != nil {
		log.Errorf("_GenealogyMembersRepository Save error %v", tx.Error)
		return tx.Error
	}
	return nil
}
