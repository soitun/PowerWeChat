package providers

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"

func RegisterConfigProvider(app kernel.ApplicationInterface) *kernel.Config {

	container := app.GetContainer()
	// 初始化配置文件，拍平
	container.InitConfig()

	return kernel.NewConfig(container.GetConfig())
}
