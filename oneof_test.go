package oneof

import "testing"

var falseFuncs []func() bool = []func() bool{
	func() bool { return false },
	func() bool { return false },
}

var trueFuncs []func() bool = []func() bool{
	func() bool { return true },
	func() bool { return true },
}

var mixFuncs []func() bool = []func() bool{
	func() bool { return false },
	func() bool { return true },
	func() bool { return false },
	func() bool { return true },
}

func TestOneOf(t *testing.T) {
	tests := []struct {
		input  []func() bool
		expect bool
	}{
		{falseFuncs, false},
		{trueFuncs, true},
		{mixFuncs, true},
	}

	for i, test := range tests {
		actual := OneOf(test.input...)
		if actual != test.expect {
			t.Errorf("%d: expected %v got %v", i, test.expect, actual)
		}
	}
}
