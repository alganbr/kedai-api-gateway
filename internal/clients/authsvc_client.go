package clients

import (
	"github.com/alganbr/kedai-api-gateway/configs"
	"github.com/alganbr/kedai-authsvc-client/client"
	"github.com/mercadolibre/golang-restclient/rest"
)

func NewAuthSvcClient(cfg *configs.Config) client.IAuthSvcClient {
	return &client.AuthSvcClient{
		HttpClient: &rest.RequestBuilder{
			BaseURL: cfg.Outbound.AuthSvcClient,
			Headers: map[string][]string{
				headerXCallerId: {cfg.Server.Name},
				headerXClientId: {cfg.Server.Name},
			},
		},
	}
}
