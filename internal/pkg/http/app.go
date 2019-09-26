package http

import (
	"github.com/blog/pkg/utils"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

var (
	app     *Application
	appOnce sync.Once
)

type ErrorOption struct {
	Status  int
	Message string
	View    string
	//irisFunc func(ctx iris.Context)
}

type ServiceProvider interface {
	Register()

	Boot()
}

type IrisFunc func(app *iris.Application)

type Application struct {
	iris         *iris.Application
	booted       bool
	providers    map[reflect.Type]ServiceProvider
	providerLock sync.Mutex
}

func (app *Application) Run(addr string) {
	app.iris.Run(iris.Addr(addr))
}

func (app *Application) Default(providers ...ServiceProvider) {
	// Default global config
	app.defaultConfigure()

	// Default template view
	app.defaultView()

	// Default global errors @todo 还未测试错误覆盖
	app.defaultErrorHandler()

	// Default middleware

}

func (app *Application) Logger() *golog.Logger {
	return app.iris.Logger()
}

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
	value := app.iris.ConfigurationReadOnly().GetOther()[key]

	if strings.Index(key, `.`) != -1 {
		keys := strings.Split(key, `.`)

		mapValue := value.(map[string]interface{})
		length := len(keys)
		for i, k := range keys {
			//last
			if length-1 == i {
				return mapValue
			} else {
				mapValue = mapValue[k].(map[string]interface{})
			}
		}
	}

	return value
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

// Register template view
func (app *Application) RegisterView(path, extension string) {
	// Register template
	app.iris.RegisterView(iris.HTML(path, extension))
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

func (app *Application) RegisterErrorHandler(options ...ErrorOption) {
	for _, option := range options {
		if option.View != `` {
			app.errorView(option)
		} else {
			// json
		}
	}
}

func (app *Application) RegisterRoutes(routes ...IrisFunc) {
	for _, route := range routes {
		route(app.iris)
	}
}

func (app *Application) errorView(option ErrorOption) {
	app.iris.OnErrorCode(option.Status, func(ctx iris.Context) {
		statusText := strconv.Itoa(option.Status)
		info := map[string]string{
			"status":  statusText,
			"message": option.Message,
		}
		//@todo 这里先这样写吧.html还没测试呢，测试成功后统一使用后缀
		ctx.View(strings.Join([]string{option.View, ".html"}, ``), info)
	})
}

func (app *Application) Iris() *iris.Application {
	return app.iris
}

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
