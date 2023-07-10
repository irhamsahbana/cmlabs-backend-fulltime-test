package questions_test

import (
	"fmt"
	"testing"

	"github.com/irhamsahbana/cmlabs-backend-fulltime-test/questions"
)

type fizzBuzzTest struct {
	name string
	arg  int
	want string
}

// test fizzbuzz with test table from 1 to 100
func TestFizzBuzzFrom1To100(t *testing.T) {
	tests := []fizzBuzzTest{}

	for i := 1; i <= 100; i++ {
		tests = append(tests, fizzBuzzTest{
			name: fmt.Sprintf("FizzBuzz %d", i),
			arg:  i,
			want: questions.FizzBuzz(i),
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := questions.FizzBuzz(tt.arg); got != tt.want {
				t.Errorf("FizzBuzz() = %v, want %v", got, tt.want)
			}

			// print result
			fmt.Printf("%s = %s\n", tt.name, tt.want)
		})
	}
}
