package main

import (
	"fmt"
	"reflect"
)


type ServiceProvider interface {
	Register()
}

type Provider struct {

}

func (p *Provider)Register() {

}

func main() {
	pro := &Provider{}
	pro1 := &Provider{}
	fmt.Println(reflect.TypeOf(pro) == reflect.TypeOf(pro1))
	fmt.Printf("%#v\n",reflect.TypeOf(pro))
	fmt.Printf("%#v",reflect.TypeOf(pro1))
}