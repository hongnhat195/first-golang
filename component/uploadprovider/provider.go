package uploadprovider

import (
	"context"

	"github.com/hongnhat195/first-golang/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
