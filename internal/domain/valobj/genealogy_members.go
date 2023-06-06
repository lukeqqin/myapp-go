package valobj

import "myapp-go/internal/domain"

type GenealogyMembersPagingReq struct {
	GenealogyMembersReq
	PagingReq
}
type GenealogyMembersReq struct {
	GenealogyId int64
}
type GenealogyMembersPagingRes struct {
	Total            int64
	GenealogyMembers []*domain.WxGenealogyMembers
}
