package genealogy

import (
	"myapp-go/internal/domain/valobj"
	"myapp-go/internal/infrustructure/persistence"
)

type GenealogyAssembleRes struct {
	Total       int64
	Genealogies []*wxGenealogyAssemble
}

type wxGenealogyAssemble struct {
	ID       int64
	Title    string
	Cover    string
	CreateBy string
	Avatar   string
	Tags     []string
}

func Assemble(req *valobj.GenealogyPagingReq) (rsp *GenealogyAssembleRes, err error) {
	grsp, err := persistence.GenealogyRepository.Paging(req)
	if err != nil {
		return nil, err
	}
	rsp = new(GenealogyAssembleRes)
	rsp.Total = grsp.Total
	var wxGenealogyAssembles []*wxGenealogyAssemble
	for i := range grsp.Genealogies {
		v := grsp.Genealogies[i]
		u, err := persistence.UserRepository.FetchByPrimaryKey(v.CreateBy)
		if err != nil {
			return nil, err
		}
		tags, err := persistence.GenealogyTagRepository.FetchIndexByWxGenealogyTagsGenealogyIDIndex(v.ID)
		if err != nil {
			return nil, err
		}
		var Tags []string
		for _, tag := range tags {
			Tags = append(Tags, tag.Name)
		}
		wxGenealogyAssemble := &wxGenealogyAssemble{
			ID:       v.ID,
			Title:    v.Title,
			Cover:    v.Title,
			CreateBy: u.Nickname,
			Avatar:   u.Avatar,
			Tags:     Tags,
		}
		wxGenealogyAssembles = append(wxGenealogyAssembles, wxGenealogyAssemble)
	}
	rsp.Genealogies = wxGenealogyAssembles
	return
}
