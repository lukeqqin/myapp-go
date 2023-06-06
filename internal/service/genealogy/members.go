package genealogy

import (
	"myapp-go/internal/domain/valobj"
	"myapp-go/internal/infrustructure/persistence"
)

func PagingMembers(req *valobj.GenealogyMembersPagingReq) (rsp *valobj.GenealogyMembersPagingRes, err error) {
	grsp, err := persistence.GenealogyMembersRepository.Paging(req)
	if err != nil {
		return nil, err
	}
	rsp = new(valobj.GenealogyMembersPagingRes)
	rsp.Total = grsp.Total
	rsp.GenealogyMembers = grsp.GenealogyMembers
	return
}
