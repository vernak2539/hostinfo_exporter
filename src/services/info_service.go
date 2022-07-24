package services

import (
	"context"
	"github.com/matishsiao/goInfo"
	"github.com/vernak2539/hostinfo_exporter/src/config"
	"github.com/vernak2539/hostinfo_exporter/src/utils"
)

type HostInfoService interface {
	GetHostInfo(ctx context.Context) (config.HostInfo, error)
}

type hostInfoService struct{}

func CreateHostInfoService() HostInfoService {
	return &hostInfoService{}
}

//This information only get populated during start up of exporter
func (h hostInfoService) GetHostInfo(ctx context.Context) (config.HostInfo, error) {
	var config config.HostInfo
	gi, err := goInfo.GetInfo()
	if err != nil {
		return config, err
	}
	config.OS = gi.GoOS
	config.Arch = gi.Platform
	config.Hostname = gi.Hostname
	config.ExternalIp = utils.GetExternalIP()

	outboundIP, err := utils.GetOutboundIP()
	if err != nil {
		return config, err
	}
	config.VpcIp = outboundIP.String()

	return config, err
}
