package _go

import (
	"github.com/go-playground/assert/v2"
	"go-learn/go/mocks"
	"testing"
)

func Greeter(name string) string {
	return "hi " + name
}

func Test_Greeter_when_param_is_elliot_get_hi_Elliot(t *testing.T) {
	name := "elliot"
	greet := Greeter(name)
	expectedGreetMsg := "hi elliot"
	assert.Equal(t, expectedGreetMsg, greet)
}

func Test_Greeter_name_greetMsg(t *testing.T) {
	name := "elliot"
	greet := Greeter(name)
	assert.Equal(t, "hi elliot", greet)
}

func isLargerThanTen(num int) bool {
	return num > 10
}

func TestIsLargerThanTen_All(t *testing.T) {
	var tests = []struct {
		name     string
		num      int
		expected bool
	}{
		{
			name:     "test_larger_than_ten",
			num:      11,
			expected: true,
		},
		{
			name:     "test_less_than_ten",
			num:      9,
			expected: false,
		},
		{
			name:     "test_equal_than_ten",
			num:      10,
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := isLargerThanTen(test.num)
			assert.Equal(t, test.expected, res)
		})
	}
}

func TestMan_All(t *testing.T) {
	man := mocks.Man{}
	man.On("GetName").Return("Elliot").On("IsHandSomeBoy").Return(true)
	assert.Equal(t, "Elliot", man.GetName())
	assert.Equal(t, true, man.IsHandSomeBoy())
}
