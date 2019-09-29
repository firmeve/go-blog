package document

import (
	"github.com/blog/pkg"
	"github.com/kataras/iris"
)

type Provider struct {
	*pkg.BaseProvider
}

func (p *Provider) Register()  {
}

func (p *Provider) Boot()  {
	newRoute(p.App().Get(`iris`).(*iris.Application)).registerRoutes()
}

func NewProvider(app *pkg.BaseApplication) *Provider {
	return &Provider{
		BaseProvider:pkg.NewBaseProvider(app),
	}
}