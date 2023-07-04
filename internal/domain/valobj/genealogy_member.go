package valobj

import "myapp-go/internal/domain"

type GenealogyMembersPagingReq struct {
	GenealogyMembersReq
	PagingReq
}
type GenealogyMembersReq struct {
	GenealogyId int64
	Name        string
}
type GenealogyMembersPagingRsp struct {
	Total            int64
	GenealogyMembers []*domain.WxGenealogyMember
}
