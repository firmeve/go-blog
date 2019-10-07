package iris

import (
	"github.com/blog/internal/pkg"
	pkg2 "github.com/blog/pkg"
	"github.com/blog/pkg/utils"
	"github.com/kataras/iris"
)

type Provider struct {
	*pkg.BaseProvider
	iris *iris.Application
}

func (p *Provider) Register() {
	p.defaultConfigure()

	p.App().Bind("iris", p.iris, pkg2.WithBindShare(true))

	p.App().Bind("config", NewConfig(p.iris),pkg2.WithBindShare(true))

	p.App().Bind("logger", p.iris.Logger(),pkg2.WithBindShare(true))
}

func (p *Provider) Boot() {

}

// Load default config
func (p *Provider) defaultConfigure() {
	p.iris.Configure(iris.WithConfiguration(iris.YAML(utils.CurrentRelativePath("../../config/app.yml"))))
}

func NewProvider(app *pkg.BaseApplication) *Provider {
	return &Provider{
		BaseProvider: pkg.NewBaseProvider(app),
		iris:         iris.New(),
	}
}
