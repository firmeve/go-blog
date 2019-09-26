package database

import (
	"github.com/blog/internal/pkg/http"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

type ServiceProvider struct {
	*http.BaseServiceProvider
}

func (s *ServiceProvider) Register() {
	driver := s.App().ConfigFromOtherDefault("databases.default", `mysql`).(string)
	var err error
	DB, err = gorm.Open(driver,
		s.App().ConfigFromOther("databases." + driver + `.addr`).(string) )

	if err != nil {
		panic(err)
	}

	iris.RegisterOnInterrupt(func() {
		DB.Close()
	})
}

func (s *ServiceProvider) Boot() {

}

func NewServiceProvider(app *http.Application) *ServiceProvider {
	return &ServiceProvider{
		BaseServiceProvider: http.BaseProvider(app),
	}
}
