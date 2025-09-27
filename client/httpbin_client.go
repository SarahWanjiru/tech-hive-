package client

import (
	"context"
	"github.com/tech-hive/ecommerce/model"
)

type HttpBinClient interface {
	PostMethod(ctx context.Context, requestBody *model.HttpBin, response *map[string]interface{})
}
