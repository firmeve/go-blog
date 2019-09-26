package http

import (
	"github.com/blog/pkg/utils"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

var (
	app              *Application
	appOnce          sync.Once
	baseProvider     *BaseServiceProvider
	baseProviderOnce sync.Once
)

type IrisFunc func(app *iris.Application)
type IrisMiddleware = context.Handler

type ServiceProvider interface {
	Register()
	Boot()
}

type ErrorOption struct {
	Status  int
	Message string
	View    string
}

type Application struct {
	iris         *iris.Application
	booted       bool
	providers    map[reflect.Type]ServiceProvider
	providerLock sync.Mutex
}

type appOption struct {
	providers        []ServiceProvider
	beforeMiddleware []IrisMiddleware
	afterMiddleware  []IrisMiddleware
	errors           []ErrorOption
}

type BaseServiceProvider struct {
	app *Application
}

// ============================= Application ================================

// Get app instance - singleton
func App() *Application {
	if app != nil {
		return app
	}

	appOnce.Do(func() {
		app = &Application{
			iris:      iris.New(),
			booted:    false,
			providers: make(map[reflect.Type]ServiceProvider, 0),
		}
	})

	return app
}

// Func: providers params
func WithProviders(providers ...ServiceProvider) utils.OptionFunc {
	return func(option utils.Option) {
		option.(*appOption).providers = providers
	}
}

// Func: errors params
func WithErrors(errors ...ErrorOption) utils.OptionFunc {
	return func(option utils.Option) {
		option.(*appOption).errors = errors
	}
}

// Func: before middleware params
func WithBeforeMiddleware(middleware ...IrisMiddleware) utils.OptionFunc {
	return func(option utils.Option) {
		option.(*appOption).beforeMiddleware = middleware
	}
}

// Func: after middleware params
func WithAfterMiddleware(middleware ...IrisMiddleware) utils.OptionFunc {
	return func(option utils.Option) {
		option.(*appOption).afterMiddleware = middleware
	}
}

// ========== Application struct =========

// Application bootstrap
func (app *Application) Bootstrap(options ...utils.OptionFunc) {
	// Load default run parameters
	app.Default()

	option := utils.ApplyOption(&appOption{}, options...).(*appOption)

	if len(option.errors) > 0 {
		app.RegisterErrorHandler(option.errors...)
	}

	if len(option.beforeMiddleware) > 0 {
		app.RegisterMiddleware(true, option.beforeMiddleware...)
	}
	if len(option.afterMiddleware) > 0 {
		app.RegisterMiddleware(false, option.afterMiddleware...)
	}

	if len(option.providers) > 0 {
		app.Register(option.providers...)
	}

	// Boot all providers
	app.Boot()
}

// Get iris instance
func (app *Application) Iris() *iris.Application {
	return app.iris
}

// Logger
func (app *Application) Logger() *golog.Logger {
	return app.iris.Logger()
}

// Run app server
func (app *Application) Run(addr string) {
	app.iris.Run(iris.Addr(addr))
	//err := app.iris.Run(iris.Addr(addr))
	//if err != nil {
	//	panic(err)
	//}
}

// Register template view
func (app *Application) RegisterView(path, extension string) {
	// Register template
	app.iris.RegisterView(iris.HTML(path, extension))
}

// Register error handler
func (app *Application) RegisterErrorHandler(options ...ErrorOption) {
	for _, option := range options {
		if option.View != `` {
			app.errorView(option)
		} else {
			// json
		}
	}
}

// Register global middleware
func (app *Application) RegisterMiddleware(before bool, middleware ...IrisMiddleware) {
	if before {
		app.iris.UseGlobal(middleware...)
	} else {
		app.iris.DoneGlobal(middleware...)
	}
}

// Register routes
func (app *Application) RegisterRoutes(routes ...IrisFunc) {
	for _, route := range routes {
		route(app.iris)
	}
}

// Service provider register
func (app *Application) Register(providers ...ServiceProvider) {
	app.providerLock.Lock()
	defer app.providerLock.Unlock()

	for _, provider := range providers {
		providerType := reflect.TypeOf(provider)
		if _, ok := app.providers[providerType]; !ok {
			provider.Register()
			if app.booted {
				provider.Boot()
			}
			app.providers[providerType] = provider
		}
	}
}

// Service provider boot
func (app *Application) Boot() {
	if app.booted {
		return
	}
	for _, provider := range app.providers {
		provider.Boot()
	}
	app.booted = true
}

// Quickly convert all other configurations
// @todo waiting testing
func (app *Application) ConfigFromOther(key string) interface{} {
	if strings.Index(key, `.`) != -1 {
		keys := strings.Split(key, `.`)
		var mapValue map[interface{}]interface{}
		length := len(keys)
		for i, k := range keys {
			//last
			if length-1 == i {
				return mapValue[k]
			} else if i == 0 {
				mapValue = app.iris.ConfigurationReadOnly().GetOther()[k].(map[interface{}]interface{})
			} else {
				mapValue = mapValue[k].(map[interface{}]interface{})
			}
		}
	}

	return app.iris.ConfigurationReadOnly().GetOther()[key]
}

// Quickly convert all other configurations
// If is nil return default value
func (app *Application) ConfigFromOtherDefault(key string, defaultValue interface{}) interface{} {
	value := app.ConfigFromOther(key)

	// empty string conversion return defaultValue
	if v, ok := value.(string); ok && v == "" {
		return defaultValue
	}

	return value
}

// Default init
func (app *Application) Default() {
	// Default global config
	app.defaultConfigure()

	// Default template view
	app.defaultView()

	// Default global errors
	// If you want to override, please use the RegisterErrorHandler method directly
	app.defaultErrorHandler()

	// Default middleware
	app.defaultMiddleware()
}

// Load default config
func (app *Application) defaultConfigure() {
	app.iris.Configure(iris.WithConfiguration(iris.YAML(utils.CurrentRelativePath("../../config/app.yml"))))
}

// Load default error handler
func (app *Application) defaultErrorHandler() {
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
		app.RegisterErrorHandler(option)
	}
}

// Default template view
func (app *Application) defaultMiddleware() {
}

// Default template view
func (app *Application) defaultView() {
	views := app.ConfigFromOther(`views`).(map[interface{}]interface{})

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
	app.RegisterView(path, extension)
}

// error view
func (app *Application) errorView(option ErrorOption) {
	app.iris.OnErrorCode(option.Status, func(ctx iris.Context) {
		statusText := strconv.Itoa(option.Status)
		info := map[string]string{
			"status":  statusText,
			"message": option.Message,
		}
		//@todo 这里先这样写吧.html还没测试呢，测试成功后统一使用后缀
		ctx.View(strings.Join([]string{option.View, `.` + app.ConfigFromOtherDefault(`views.extension`, `html`).(string)}, ``), info)
	})
}

// ========================== BaseServiceProvider ===========================

// Get base service provider instance - singleton
func BaseProvider(app *Application) *BaseServiceProvider {
	if baseProvider != nil {
		return baseProvider
	}

	baseProviderOnce.Do(func() {
		baseProvider = &BaseServiceProvider{
			app: app,
		}
	})

	return baseProvider
}

// Get provider application instance
func (s *BaseServiceProvider) App() *Application {
	return app
}
