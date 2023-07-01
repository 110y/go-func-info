package package1

type MyType struct{}

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
