package package1

type MyType struct{}

type MyInterface interface{}

func F1(x, y int) int {
	return x + y
}

func F2() {
	_ = func() {
		_ = 1
	}
}

var F3 = func() {
	_ = 1
}

func (t *MyType) F4() {
	_ = 1
}

func F5() MyType {
	return MyType{}
}

func F6() (int, error) {
	return 0, nil
}

func F7() (MyInterface, error) {
	return nil, nil
}

func F8() (*MyType, error) {
	return nil, nil
}

func F9() (MyType, error) {
	return MyType{}, nil
}

func (m MyType) F10() error {
	return nil
}

type MyGenerictype[T any] struct{}
