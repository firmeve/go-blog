package http

import (
	iris2 "github.com/blog/internal/pkg/iris"
	"github.com/blog/pkg"
	"github.com/blog/pkg/utils"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"strconv"
	"strings"
)

type IrisFunc func(app *iris.Application)
type IrisMiddleware = context.Handler

type ErrorOption struct {
	Status  int
	Message string
	View    string
}

type bootOption struct {
	beforeMiddleware []IrisMiddleware
	afterMiddleware  []IrisMiddleware
	errors           []ErrorOption
}

type Provider struct {
	*pkg.BaseProvider
	bootOption *bootOption
	iris *iris.Application
}

// Func: errors params
func WithErrors(errors ...ErrorOption) utils.OptionFunc {
	return func(option utils.Option) {
		option.(*bootOption).errors = errors
	}
}

// Func: before middleware params
func WithBeforeMiddleware(middleware ...IrisMiddleware) utils.OptionFunc {
	return func(option utils.Option) {
		option.(*bootOption).beforeMiddleware = middleware
	}
}

// Func: after middleware params
func WithAfterMiddleware(middleware ...IrisMiddleware) utils.OptionFunc {
	return func(option utils.Option) {
		option.(*bootOption).afterMiddleware = middleware
	}
}

func (p *Provider) SetBootOption(options ...utils.OptionFunc) *Provider {
	p.bootOption = utils.ApplyOption(&bootOption{}, options...).(*bootOption)
	return p
}

// Application bootstrap
func (p *Provider) Register() {

}

// Application bootstrap
func (p *Provider) Boot() {
	// Load default run parameters
	p.Default()

	if len(p.bootOption.errors) > 0 {
		p.RegisterErrorHandler(p.bootOption.errors...)
	}

	if len(p.bootOption.beforeMiddleware) > 0 {
		p.RegisterMiddleware(true, p.bootOption.beforeMiddleware...)
	}
	if len(p.bootOption.afterMiddleware) > 0 {
		p.RegisterMiddleware(false, p.bootOption.afterMiddleware...)
	}
}

// Register template view
func (p *Provider) RegisterView(path, extension string) {
	// Register template
	p.iris.RegisterView(iris.HTML(path, extension))
}

// Register error handler
func (p *Provider) RegisterErrorHandler(options ...ErrorOption) {
	for _, option := range options {
		if option.View != `` {
			p.errorView(option)
		} else {
			// json
		}
	}
}

// Register global middleware
func (p *Provider) RegisterMiddleware(before bool, middleware ...IrisMiddleware) {
	if before {
		p.iris.UseGlobal(middleware...)
	} else {
		p.iris.DoneGlobal(middleware...)
	}
}

// Default init
func (p *Provider) Default() {
	// Default global config
	//p.defaultConfigure()

	// Default template view
	p.defaultView()

	// Default global errors
	// If you want to override, please use the RegisterErrorHandler method directly
	p.defaultErrorHandler()

	// Default middleware
	p.defaultMiddleware()
}

// Load default error handler
func (p *Provider) defaultErrorHandler() {
	errorOptions := []ErrorOption{
		{
			Status:  iris.StatusNotFound,
			Message: "Not Found",
			View:    "errors/404",
		},
		{
			Status:  iris.StatusForbidden,
			Message: "Forbidden",
			View:    "errors/403",
		},
	}

	for _, option := range errorOptions {
		p.RegisterErrorHandler(option)
	}
}

// Default template view
func (p *Provider) defaultMiddleware() {
}

// Default template view
func (p *Provider) defaultView() {
	views := p.App().Get(`config`).(*iris2.Config).Get(`views`).(map[interface{}]interface{})

	path := views[`path`].(string)
	if path == `` {
		path = utils.CurrentRelativePath("../../web/views")
	}

	extension := views[`extension`].(string)
	if extension == `` {
		extension = `.html`
	} else {
		extension = `.` + extension
	}
	p.RegisterView(path, extension)
}

// error view
func (p *Provider) errorView(option ErrorOption) {
	p.iris.OnErrorCode(option.Status, func(ctx iris.Context) {
		statusText := strconv.Itoa(option.Status)
		info := map[string]string{
			"status":  statusText,
			"message": option.Message,
		}
		//@todo 这里先这样写吧.html还没测试呢，测试成功后统一使用后缀
		ctx.View(strings.Join([]string{option.View, `.` + p.App().Get(`config`).(*iris2.Config).GetDefault(`views.extension`, `html`).(string)}, ``), info)
	})
}

func NewProvider(app *pkg.BaseApplication) *Provider {
	return &Provider{
		BaseProvider: pkg.NewBaseProvider(app),
		iris:         app.Get(`iris`).(*iris.Application),
	}
}
