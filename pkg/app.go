package pkg

import (
	"reflect"
	"sync"
)

var (
	app              *BaseApplication
	appOnce          sync.Once
	baseProvider     *BaseProvider
	baseProviderOnce sync.Once
)

type Provider interface {
	Register()
	Boot()
}

type Application interface {
	Bootstrap(app *BaseApplication)
}

type BaseApplication struct {
	*BaseContainer
	booted       bool
	providers    map[reflect.Type]Provider
	providerLock sync.RWMutex
	boxes        map[string]Application
	boxLock      sync.RWMutex
}

type BaseProvider struct {
	app *BaseApplication
}

func (app *BaseApplication) Mount(name string, subApp Application) {
	app.boxLock.Lock()
	defer app.boxLock.Unlock()
	if _, ok := app.boxes[name]; !ok {
		app.boxes[name] = subApp
	}
}

// Service provider register
func (app *BaseApplication) Register(providers ...Provider) {
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
func (app *BaseApplication) Boot() {
	if app.booted {
		return
	}
	for _, provider := range app.providers {
		provider.Boot()
	}
	app.booted = true
}

func (app *BaseApplication) Bootstrap(boxes ...string) {
	// Boot all providers
	app.Boot()

	for _, name := range boxes {
		if box, ok := app.boxes[name]; ok {
			box.Bootstrap(app)
		}
	}
}

// ============================= Application ================================

// Get app instance - singleton
func App() *BaseApplication {
	if app != nil {
		return app
	}

	appOnce.Do(func() {
		app = &BaseApplication{
			BaseContainer: NewContainer(),
			booted:        false,
			providers:     make(map[reflect.Type]Provider, 0),
			boxes:         make(map[string]Application, 0),
		}
	})

	return app
}

// ========================== BaseServiceProvider ===========================

// Get base service provider instance - singleton
func NewBaseProvider(app *BaseApplication) *BaseProvider {
	if baseProvider != nil {
		return baseProvider
	}

	baseProviderOnce.Do(func() {
		baseProvider = &BaseProvider{
			app: app,
		}
	})

	return baseProvider
}

// Get provider application instance
func (p *BaseProvider) App() *BaseApplication {
	return p.app
}
