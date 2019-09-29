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
	registerRoutes(p.App().Get(`iris`).(*iris.Application))
}

func NewProvider(app *pkg.BaseApplication) *Provider {
	return &Provider{
		BaseProvider:pkg.NewBaseProvider(app),
	}
}