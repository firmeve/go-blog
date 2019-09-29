package database

import (
	iris2 "github.com/blog/internal/pkg/iris"
	"github.com/blog/pkg"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
)

var (
	DB *gorm.DB
)

type Provider struct {
	*pkg.BaseProvider
}

func (p *Provider) Register() {
	config := p.App().Resolve(`config`).(*iris2.Config)
	driver := config.GetDefault("databases.default", `mysql`).(string)

	var err error
	DB, err = gorm.Open(driver,
		config.Get("databases." + driver + `.addr`).(string))

	if err != nil {
		panic(err)
	}

	iris.RegisterOnInterrupt(func() {
		DB.Close()
	})

	p.App().Bind(`db`, DB, pkg.WithBindShare(true))
}

func (p *Provider) Boot() {

}

func NewProvider(app *pkg.BaseApplication) *Provider {
	return &Provider{
		BaseProvider: pkg.NewBaseProvider(app),
	}
}
