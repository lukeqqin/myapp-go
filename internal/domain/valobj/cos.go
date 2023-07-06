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
