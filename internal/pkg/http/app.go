package http

import (
	"github.com/blog/internal/document"
	"github.com/blog/internal/pkg/routes"
	"github.com/blog/pkg/utils"
	"github.com/kataras/iris"
	"strconv"
	"strings"
	"sync"
)

var (
	webApp  *iris.Application
	webOnce sync.Once
)

type ErrorOption struct {
	Status  int
	Message string
	View    string
	//irisFunc func(ctx iris.Context)
}

type Application struct {
	iris        *iris.Application
	isBootstrap bool
}

func (app *Application) Run(addr string) {
	app.iris.Run(iris.Addr(addr))
}

func (app *Application) New() {
	app.iris = iris.New()
}

func (app *Application) Bootstrap(addr string) {
	if app.isBootstrap == true {
		return
	}

	// init
	app.defaultConfigure()

	// template view
	app.defaultView()

	// default global errors
	app.defaultErrorHandler()

	// @todo 只写到这里，路由还没改
	// Load routing
	//routes.RegisterRoutes(appebApp, document.Register)
}

// Quickly convert all other configurations
// @todo waiting testing
func (app *Application) ConfigFromOther(key string) interface{}  {
	value := app.iris.ConfigurationReadOnly().GetOther()[key]

	if strings.Index(key,`.`) != -1 {
		keys := strings.Split(key,`.`)

		mapValue := value.(map[string]interface{})
		length := len(keys)
		for i,k := range keys {
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
	views := app.ConfigFromOther(`views`).(map[string]string)

	path := views[`path`]
	if path == `` {
		path = utils.CurrentRelativePath("../../web/views")
	}

	extension := views[`extension`]
	if extension == `` {
		extension = `.html`
	} else {
		extension = `.` + extension
	}

	app.RegisterView(path, extension)
}

// Register template view
func (app *Application) RegisterView(extension, path string) {
	// Register template
	app.iris.RegisterView(iris.HTML(path, extension))
}

func (app *Application) RegisterRoutes(routes ...Routes) {

}

// Load default config
func (app *Application) defaultConfigure() *Web {
	app.iris.Configure(iris.WithConfiguration(iris.YAML(utils.CurrentRelativePath("../../config/app.yml"))))
}

func (app *Application) defaultErrorHandler() {
	app.RegisterErrorHandler(ErrorOption{
		Status:  iris.StatusNotFound,
		Message: "Not Found",
		View:    "errors/404",
	})
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

func (app *Application) RegisterHandler() {
	Status404Page(app)
}

//func App(app *Application) *iris.Application {
//	return app.iris
//}

func Instance() {
	//return
}

func NewWeb() {
	if webApp != nil {
		return webApp
	}

	webOnce.Do(func() {

		webApp.Get("/product-problem", newProductProblemRender)
	})

	return webApp
}

func RunWeb(addr string) {

}

func newProductProblemRender(ctx iris.Context) {
	ctx.Problem(newProductProblem("abc", "ef"), iris.ProblemOptions{
		//JSON: iris.JSON{
		//},
	})
}
