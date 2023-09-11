package lang

import (
	"learn-go/std/lang/zpkg1"
	"learn-go/std/lang/zpkg2"
	"testing"
)

func Test_defined_types(t *testing.T) {
	var pkg1FooVar zpkg1.Foo
	var pkg2FooVar zpkg2.Foo

	_, ok1 := any(pkg1FooVar).(zpkg2.Foo)
	if ok1 {
		panic("impossible")
	}
	_, ok2 := any(pkg2FooVar).(zpkg1.Foo)
	if ok2 {
		panic("impossible")
	}
}

func Test_unexported_method_in_interface(t *testing.T) {
	var pkg1IfVar zpkg1.WithUnexportedFoo
	var pkg2IfVar zpkg2.WithUnexportedFoo

	_, ok1 := any(pkg1IfVar).(zpkg2.WithUnexportedFoo)
	if ok1 {
		panic("impossible")
	}
	_, ok2 := any(pkg2IfVar).(zpkg1.WithUnexportedFoo)
	if ok2 {
		panic("impossible")
	}
}
