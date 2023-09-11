package zpkg2

type Foo int

type WithUnexportedFoo interface {
	foo()
}

type WithFoo struct{}

func (a WithFoo) foo() {}
