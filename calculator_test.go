package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldAdd(t *testing.T) {
	c := new(Calculator)

	c.Add(3)

	assert.Equal(t, 3, c.Result())
}

func TestShouldSub(t *testing.T) {
	c := new(Calculator)
	c.Press(10)

	c.Sub(3)

	assert.Equal(t, 7, c.Result())
}

func TestPress(t *testing.T) {
	c := new(Calculator)

	c.Press(10)

	assert.Equal(t, 10, c.Result())
}

//Sample Add Unit test with auto generate function
func TestCalculator_Add(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name        string
		c           *Calculator
		args        args
		want        int
		intialValue int
	}{
		// TODO: Add test cases.
		{
			args: args{
				x: 5,
			},
			c:           new(Calculator),
			intialValue: 10,
			want:        15,
		},
		{
			args: args{
				x: 5,
			},
			c:           new(Calculator),
			intialValue: 0,
			want:        5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Press(tt.intialValue)
			tt.c.Add(tt.args.x)
			if got := tt.c.Result(); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
