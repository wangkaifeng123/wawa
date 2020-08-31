package factory

import "testing"

func TestFactory(t *testing.T) {
	new(CreateFactory).NewFactory(1).test("ok")
	new(CreateFactory).NewFactory(2).test("ok")
}
