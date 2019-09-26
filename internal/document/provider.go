package document

import "github.com/blog/internal/pkg/http"

type ServiceProvider struct {
	*http.BaseServiceProvider
}

func (s *ServiceProvider) Register()  {

}

func (s *ServiceProvider) Boot()  {
	s.App().RegisterRoutes(Register)
}

func NewServiceProvider(app *http.Application) *ServiceProvider {
	return &ServiceProvider{
		BaseServiceProvider:http.BaseProvider(app),
	}
}