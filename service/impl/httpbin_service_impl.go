package impl

import (
	"context"
	"github.com/tech-hive/ecommerce/client"
	"github.com/tech-hive/ecommerce/common"
	"github.com/tech-hive/ecommerce/model"
	"github.com/tech-hive/ecommerce/service"
)

func NewHttpBinServiceImpl(httpBinClient *client.HttpBinClient) service.HttpBinService {
	return &httpBinServiceImpl{HttpBinClient: *httpBinClient}
}

type httpBinServiceImpl struct {
	client.HttpBinClient
}

func (h *httpBinServiceImpl) PostMethod(ctx context.Context) {
	httpBin := model.HttpBin{
		Name: "rizki",
	}
	var response map[string]interface{}
	h.HttpBinClient.PostMethod(ctx, &httpBin, &response)
	common.NewLogger().Info("log response service ", response)
}
