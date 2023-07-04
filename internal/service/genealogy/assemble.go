package genealogy

import (
	"myapp-go/internal/domain/valobj"
	"myapp-go/internal/infrustructure/persistence"
)

type AssembleRsp struct {
	Total       int64
	Genealogies []*wxGenealogyAssemble
}

type wxGenealogyAssemble struct {
	ID       int64
	Title    string
	Cover    string
	CreateBy string
	Avatar   string
}

func Assemble(req *valobj.GenealogyPagingReq) (rsp *AssembleRsp, err error) {
	grsp, err := persistence.GenealogyRepository.Paging(req)
	if err != nil {
		return nil, err
	}
	rsp = new(AssembleRsp)
	rsp.Total = grsp.Total
	var wxGenealogyAssembles []*wxGenealogyAssemble
	for i := range grsp.Genealogies {
		v := grsp.Genealogies[i]
		u, err := persistence.UserRepository.FetchByPrimaryKey(v.CreateBy)
		if err != nil {
			return nil, err
		}
		wxGenealogyAssemble := &wxGenealogyAssemble{
			ID:       v.ID,
			Title:    v.Title,
			Cover:    v.Cover,
			CreateBy: u.Nickname,
			Avatar:   u.Avatar,
		}
		wxGenealogyAssembles = append(wxGenealogyAssembles, wxGenealogyAssemble)
	}
	rsp.Genealogies = wxGenealogyAssembles
	return
}
