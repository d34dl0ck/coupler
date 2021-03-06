# Coupler - simple dependency injection for Go

[![codecov](https://codecov.io/gh/d34dl0ck/coupler/branch/master/graph/badge.svg?token=53UWLJSOE5)](https://codecov.io/gh/d34dl0ck/coupler)

# Overview

Coupler should be a simple but flexible dependency injection container for Go. It should be simple for use, just by chaining function calls, like option Go pattern. On the other side, Coupler shoud provide flexible architecture with a bunch of extension points, following the Open/Close principle from SOLID.

# Examples

You can find the examples in the coupler package of the current module. Anyway, it's considered to use two basic functions: Register and Resolve.
The following example showing the registration the struct of two fields as interface implementation (errors check was skipped):
```
// importing Coupler
import c "github.com/d34dl0ck/coupler/pkg/coupler"

// our struct, that implements some interface testInterface
type testStruct struct {
	SomeString string
	SomeInt    int
}

// register with Coupler
err := c.Register(c.ByInstance("some_string"))
err = c.Register(c.ByInstance(34))
err = c.Register(
	c.ByType[testStruct](),
	c.AsImplementationOf[testInterface]())

// check all dependencies can be resolved
err = c.Check()

// resolve with Coupler
instance, err := c.Resolve[testInterface]()
```

# Architecture
The following cases are supported by Coupler:

![Types](/docs/Cases.jpg)

This class diagram illustrates types:

![Types](/docs/Types.jpg)

More documentation will be placed in the docs directory and godoc documentation.

# Package structure

TBD

# Notes

Current project status - just started, work in progress