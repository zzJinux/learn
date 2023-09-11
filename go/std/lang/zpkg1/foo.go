package zpkg1

type Foo int

type WithUnexportedFoo interface {
	foo()
}

type WithFoo struct{}

func (a WithFoo) foo() {}
