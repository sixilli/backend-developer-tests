package fizzbuzz

import (
	"fmt"
	"strconv"
	"testing"
)

// Testing that FizzBuzz returns expected results
func TestFizzBuzz(t *testing.T) {
	total := int64(100)
	fizzAt := int64(10)
	buzzAt := int64(5)

	res := FizzBuzz(total, fizzAt, buzzAt)

	for i := int64(1); i <= total; i++ {
		if !(i%fizzAt == 0) && !(i%buzzAt == 0) {
			number, err := strconv.ParseInt(res[i-1], 10, 64)
			if err != nil {
				fmt.Println("Parse Int failed")
				t.Fail()
			}

			if number != i {
				fmt.Println("Number mismatch")
				t.Fail()
			}
			continue
		}

		if i%fizzAt == 0 && i%buzzAt == 0 {
			if res[i-1] != "FizzBuzz" {
				fmt.Println("Did not FizzBuzz")
				t.Fail()
			}
			continue
		}

		if i%fizzAt == 0 {
			if res[i-1] != "Fizz" {
				fmt.Println("Did not equal Fizz")
				fmt.Printf("%v", res[i-1])
				t.Fail()
			}
		}

		if i%buzzAt == 0 {
			if res[i-1] != "Buzz" {
				fmt.Println("Did not equal Buzz")
				t.Fail()
			}
		}
	}
}

func TestNegativeFizzBuzz(t *testing.T) {
	total := int64(-32)
	fizzAt := int64(10)
	buzzAt := int64(5)

	errMsg := "Total must be greater than 0"

	res := FizzBuzz(total, fizzAt, buzzAt)

	if res[0] != errMsg || len(res) > 1 {
		t.Fail()
	}
}
