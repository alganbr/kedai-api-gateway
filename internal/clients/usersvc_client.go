package clients

import (
	"github.com/alganbr/kedai-api-gateway/configs"
	"github.com/alganbr/kedai-usersvc-client/client"
	"github.com/mercadolibre/golang-restclient/rest"
)

func NewUserSvcClient(cfg *configs.Config) client.IUserSvcClient {
	return &client.UserSvcClient{
		HttpClient: &rest.RequestBuilder{
			BaseURL: cfg.Outbound.UserSvcClient,
			Headers: map[string][]string{
				headerXCallerId: {cfg.Server.Name},
				headerXClientId: {cfg.Server.Name},
			},
		},
	}
}
