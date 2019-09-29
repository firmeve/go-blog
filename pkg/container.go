package pkg

import (
	"fmt"
	"github.com/blog/pkg/utils"
	"reflect"
	"strings"
	"sync"
)

type prototypeFunc func(container Container, params ...interface{}) interface{}

type Container interface {
	Has(name string) bool
	Get(name string) interface{}
	Bind(name string, prototype interface{}, options ...utils.OptionFunc)
	Resolve(abstract interface{}, params ...interface{}) interface{}
	Remove(name string)
}

type BaseContainer struct {
	bindings map[string]*binding
	types    map[reflect.Type]string
}

type binding struct {
	name        string
	share       bool
	instance    interface{}
	prototype   prototypeFunc
	reflectType reflect.Type
}

type bindingOption struct {
	name      string
	share     bool
	cover     bool
	prototype interface{}
}

var (
	typeMutex      sync.Mutex
	bindingRwMutex sync.RWMutex
)

// Create a new container instance
func NewContainer() *BaseContainer {
	return &BaseContainer{
		bindings: make(map[string]*binding),
		types:    make(map[reflect.Type]string),
	}
}

// Determine whether the specified name object is included in the container
func (f *BaseContainer) Has(name string) bool {
	bindingRwMutex.RLock()
	defer bindingRwMutex.RUnlock()
	if _, ok := f.bindings[strings.ToLower(name)]; ok {
		return true
	}

	return false
}

// Get a object from container
func (f *BaseContainer) Get(name string) interface{} {
	if !f.Has(name) {
		panic(fmt.Errorf("object[%s] that does not exist", name))
	}

	return f.bindings[strings.ToLower(name)].resolvePrototype(f)
}

// Bind method `share` param
func WithBindShare(share bool) utils.OptionFunc {
	return func(option utils.Option) {
		option.(*bindingOption).share = share
	}
}

// Bind method `cover` param
func WithBindCover(cover bool) utils.OptionFunc {
	return func(option utils.Option) {
		option.(*bindingOption).cover = cover
	}
}

// Bind a object to container
func (f *BaseContainer) Bind(name string, prototype interface{}, options ...utils.OptionFunc) { //, value interface{}
	// Parameter analysis
	bindingOption := utils.ApplyOption(newBindingOption(name, prototype), options...).(*bindingOption)

	// Coverage detection
	if _, ok := f.bindings[bindingOption.name]; ok && !bindingOption.cover {
		panic(`binding alias type already exists`)
	}

	// set binding item
	f.setBindingItem(newBinding(bindingOption))
}

// Parsing various objects
func (f *BaseContainer) Resolve(abstract interface{}, params ...interface{}) interface{} {
	reflectType := reflect.TypeOf(abstract)
	kind := reflectType.Kind()

	if kind == reflect.Func {
		newParams := make([]reflect.Value, 0)
		if len(params) != 0 {
			for param := range params {
				newParams = append(newParams, reflect.ValueOf(param))
			}
		} else {
			for i := 0; i < reflectType.NumIn(); i++ {
				if name, ok := f.types[reflectType.In(i)]; ok {
					newParams = append(newParams, reflect.ValueOf(f.Get(name)))
				} else {
					panic(`unable to find reflection parameter`)
				}
			}
		}

		results := reflect.ValueOf(abstract).Call(newParams)

		resultsInterface := make([]interface{}, 0)
		for _, result := range results {
			resultInterface := result.Interface()
			if err, ok := resultInterface.(error); ok && err != nil {
				panic(err.Error())
			}

			resultsInterface = append(resultsInterface, resultInterface)
		}

		if reflectType.NumOut() == 1 {
			return resultsInterface[0]
		} else {
			return resultsInterface
		}
	} else if kind == reflect.Ptr {
		if reflectType.Elem().Kind() == reflect.Struct {
			// Get the struct value corresponding to the pointer structure
			// Only the strcut value can get the number of fields, and set the field
			// is not setting the pointer address

			// reflect.Prt type, only get the pointer value through Elem()
			// The value to be modified is the value of the pointer, not the address of the pointer itself.

			// The Interface type that finally gets the address of the modified value returns

			// 取得指针结构体对应的struct值
			// 只有strcut值才能取得字段数，以及设置字段
			// 并不是设置指针地址

			// reflect.Prt类型，只有通过Elem()来获取指针值
			// 要修改的是指针对应的值，而不是指针本身的地址

			// 最后取得修改值的地址的Interface类型返回
			return f.parseStruct(reflectType.Elem(), reflect.ValueOf(abstract).Elem()).Addr().Interface()
		}
	} else if kind == reflect.Struct {
		return f.parseStruct(reflectType, reflect.ValueOf(abstract)).Interface()
	} else if kind == reflect.String {
		return f.Get(abstract.(string))
	}

	panic(fmt.Errorf("%s %#v", `unsupported type`, abstract))
}

// Remove a binding
func (f *BaseContainer) Remove(name string) {
	typeMutex.Lock()
	defer typeMutex.Unlock()

	name = strings.ToLower(name)

	delete(f.bindings, name)

	for key, v := range f.types {
		if v == name {
			delete(f.types, key)
			break
		}
	}
}

// Set a item to types and bindings
func (f *BaseContainer) setBindingItem(b *binding) {
	bindingRwMutex.Lock()
	defer bindingRwMutex.Unlock()

	// Set binding
	f.bindings[b.name] = b

	// Set type
	// Only support prt,struct and func type
	// No support string,float,int... scalar type
	originalKind := b.reflectType.Kind()
	if originalKind == reflect.Ptr || originalKind == reflect.Struct {
		f.types[b.reflectType] = b.name
	} else if originalKind == reflect.Func {
		// This is mainly used as a non-singleton type, using function execution, each time returning a different instance
		// When it is a function, parse the function, get the current real type, only support one parameter, the function must have only one return value
		f.types[reflect.TypeOf(b.resolvePrototype(f))] = b.name
	}
}

// Parse struct fields and auto binding field
func (f *BaseContainer) parseStruct(reflectType reflect.Type, reflectValue reflect.Value) reflect.Value {
	for i := 0; i < reflectType.NumField(); i++ {
		tag := reflectType.Field(i).Tag.Get("inject")
		if tag != `` && reflectValue.Field(i).CanSet() {
			if _, ok := f.bindings[tag]; ok {
				result := f.Resolve(f.Get(tag))
				// Non-same type of direct skip
				if reflect.TypeOf(result).Kind() == reflectType.Field(i).Type.Kind() {
					reflectValue.Field(i).Set(reflect.ValueOf(result))
				}
			}
		}
	}

	return reflectValue
}

// ---------------------------- bindingOption ------------------------

// Create a new binding option struct
func newBindingOption(name string, prototype interface{}) *bindingOption {
	return &bindingOption{share: false, cover: false, name: strings.ToLower(name), prototype: prototype}
}

// ---------------------------- binding ------------------------

// Create a new binding struct
func newBinding(option *bindingOption) *binding {
	binding := &binding{
		name:        option.name,
		reflectType: reflect.TypeOf(option.prototype),
	}
	binding.share = binding.getShare(option.share)
	binding.prototype = binding.getPrototypeFunc(option.prototype)

	return binding
}

// Get share, If type kind is not func type
func (b *binding) getShare(share bool) bool {
	if b.reflectType.Kind() != reflect.Func {
		b.share = true
	}

	return share
}

// Parse package prototypeFunc type
func (b *binding) getPrototypeFunc(prototype interface{}) prototypeFunc {
	var prototypeFunction prototypeFunc

	if b.reflectType.Kind() == reflect.Func {
		prototypeFunction = func(container Container, params ...interface{}) interface{} {
			return container.Resolve(prototype)
		}
	} else {
		prototypeFunction = func(container Container, params ...interface{}) interface{} {
			return prototype
		}
	}

	return prototypeFunction
}

// Parse binding object prototype
func (b *binding) resolvePrototype(container Container, params ...interface{}) interface{} {
	if b.share && b.instance != nil {
		return b.instance
	}

	prototype := b.prototype(container, params...)
	if b.share {
		b.instance = prototype
	}

	return prototype
}
