//go:build ignore

// This file is an illustrative sample, not part of the build.
package api

// ExampleLibObject is an object handed back by the library, created
// through Lib.NewExampleObject. Its function fields are filled by the
// factories in sandbox/internal/example_object/.
type ExampleLibObject struct {
	// FirstProp and SecondProp are plain data fields, set at construction.
	FirstProp  int
	SecondProp string
	// ExampleObjectMethod is filled by example_object.ExampleObjectMethodFactory.
	ExampleObjectMethod func() string
}

// Lib is the entry point handed back by lib.New.
type Lib struct {
	// Multiplier is plain data, set at construction from lib.New's argument.
	Multiplier int
	// NewExampleObject is filled by lib.NewExampleObjectFactory.
	NewExampleObject func(firstProp int, secondProp string) ExampleLibObject
	// ExampleFunction is filled by lib.ExampleFunctionFactory.
	ExampleFunction func(i int) int
}
