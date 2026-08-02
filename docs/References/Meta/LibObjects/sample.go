//go:build ignore

// This file is an illustrative sample, not part of the build.
package example_object

import (
	"strconv"

	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
)

// ExampleObjectMethodFactory returns the closure that fills
// api.ExampleLibObject.ExampleObjectMethod, closed over o, so it reads the
// object's own properties at call time.
func ExampleObjectMethodFactory(o *api.ExampleLibObject) func() string {
	return func() string {
		propsConcatenated := strconv.Itoa(o.FirstProp) + o.SecondProp
		return "Props Concatenated: " + propsConcatenated
	}
}

// New builds an api.ExampleLibObject from its properties and runs every
// object factory over it to fill its function fields. Adding a function
// field to api.ExampleLibObject means adding its factory call here.
func New(firstProp int, secondProp string) api.ExampleLibObject {
	o := api.ExampleLibObject{
		FirstProp:  firstProp,
		SecondProp: secondProp,
	}
	o.ExampleObjectMethod = ExampleObjectMethodFactory(&o)
	return o
}

// ---------------------------------------------------------------
// The constructor factory below lives in sandbox/internal/lib/, where
// the object is created on behalf of the parent lib.
// ---------------------------------------------------------------

// NewExampleObjectFactory returns the closure that fills
// api.Lib.NewExampleObject, creating an ExampleLibObject.
func NewExampleObjectFactory(l *api.Lib) func(firstProp int, secondProp string) api.ExampleLibObject {
	return func(firstProp int, secondProp string) api.ExampleLibObject {
		return example_object.New(firstProp, secondProp)
	}
}
