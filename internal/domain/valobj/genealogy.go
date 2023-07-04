package valobj

import "myapp-go/internal/domain"

type GenealogyPagingReq struct {
	GenealogyReq
	PagingReq
}
type GenealogyReq struct {
	Title string
}
type GenealogyPagingRsp struct {
	Total       int64
	Genealogies []*domain.WxGenealogy
}

type MyGenealogyReq struct {
	UserId int64
}
