package factory

type Test interface {
	test(string)
}

type CreateFactory struct {
	Kind int //0 表示TestA   1标识TestB
}

func (CreateFactory) NewFactory(kind int) Test {
	switch kind {
	case 1:
		return TestA{}
	case 2:
		return TestB{}
	default:
		return nil
	}
}

type TestA struct {
}

func (TestA) test(a string) {
	println("this is testA", a)
}

type TestB struct {
}

func (TestB) test(a string) {
	println("this is testB", a)
}
