package uploadprovider

import "context"

type UploadFileProvider interface {
	UploadFile(context.Context, string, []byte) (string, error)
}
