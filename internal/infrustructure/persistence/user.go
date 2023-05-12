package persistence

import (
	"gorm.io/gorm"
	"myapp-go/internal/domain"
	"myapp-go/internal/infrustructure/rk/log"
)

// _UserRepository 凭证仓储接口
type _UserRepository struct {
	db *gorm.DB
}

func newUserRepository(db *gorm.DB) *_UserRepository {
	return &_UserRepository{db: db}
}

func (gr *_UserRepository) FetchByPrimaryKey(id int64) (rsp *domain.WxUser, err error) {
	rsp = new(domain.WxUser)
	mgr := domain.WxUserMgr(gr.db)
	user, err := mgr.FetchByPrimaryKey(id)
	if err != nil {
		log.Errorf("_UserRepository FetchByPrimaryKey error %v", err)
		return nil, err
	}
	rsp = &user
	return
}
