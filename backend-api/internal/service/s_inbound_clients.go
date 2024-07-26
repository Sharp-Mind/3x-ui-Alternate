package service

import (
	"fmt"
	"x-ui/backend-api/internal/repository"
	"x-ui/database/model"
)

type InboundClientService struct {
	InboundClientRepo repository.InboundClient
}

func NewInboundClientService(inboundClientRepo repository.InboundClient) *InboundClientService {
	return &InboundClientService{
		InboundClientRepo: inboundClientRepo,
	}
}

func (s *InboundClientService) GetInboundClients(inboundId int) ([]model.Client, error) {
	return s.InboundClientRepo.Get(inboundId)
}

func (s *InboundClientService) AddInboundClient(inboundId int, newClient *model.Client) (string, bool, error) {
	res, err := s.InboundClientRepo.Add(inboundId, newClient)
	if err != nil {
		return "", false, err
	}

	//ipaddr := string(*net.IPAddr)
	key := fmt.Sprintf("vless://" + newClient.ID + "@185.225.201.103:443?type=tcp&security=reality&pbk=hR_tQr8FVdSOM-k7pt4oSGtjct6FfPvKNQMDzIDvjB8&fp=chrome&sni=microsoft.com&sid=a5ee0427&spx=%2F&flow=xtls-rprx-vision#Alpha_Proxy-" + newClient.Email)

	return key, res, nil
}
