package valobj

import "myapp-go/internal/domain"

type GenealogyMembersPagingReq struct {
	GenealogyReq
	PagingReq
}
type GenealogyMembersReq struct {
	GenealogyId string
}
type GenealogyPagingRes struct {
	Total       int64
	Genealogies []*domain.WxGenealogy
}
