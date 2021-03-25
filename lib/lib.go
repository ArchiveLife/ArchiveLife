package lib

import (
	"github.com/ArchiveLife/core/adapter"
	"github.com/ArchiveLife/weibo/provision"
)

func GetServices() []adapter.ArchiveService {
	weiboProvision := provision.WeiboServiceProvision{}
	return weiboProvision.ProvideServices()
}
