package config

import (
	"github.com/blog/pkg"
)

type Provider struct {
	*pkg.BaseProvider
}

func (p *Provider) Register() {

}

func (p *Provider) Boot() {

}

type Config struct {

}

func NewProvider(app *pkg.BaseApplication) *Provider {
	return &Provider{
		BaseProvider: pkg.NewBaseProvider(app),
	}
}
