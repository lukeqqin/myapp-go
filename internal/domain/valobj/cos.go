package valobj

import (
	"io"
)

type UploadReq struct {
	Name string
	File io.Reader
}

type UploadRsp struct {
	Url string
}

type UpdateStatusReq struct {
	Url    string
	Status int8
}
