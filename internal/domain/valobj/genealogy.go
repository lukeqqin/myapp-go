package valobj

import "myapp-go/internal/domain"

type GenealogyPagingReq struct {
	GenealogyReq
	PagingReq
}
type GenealogyReq struct {
	Title string
}
type GenealogyPagingRes struct {
	Total       int64
	Genealogies []*domain.WxGenealogy
}
