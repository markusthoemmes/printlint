package printlint

import (
	"errors"
	"fmt"
	"testing"
)

func BenchmarkSprint(b *testing.B) {
	number := 10
	floating := 1.3333
	err := errors.New("this is a test error")
	str := "test"
	obj := struct {
		str    string
		number int
	}{
		str:    str,
		number: number,
	}

	tests := []struct {
		format string
		value  interface{}
	}{
		{"%d", number},
		{"%f", floating},
		{"%v", err},
		{"%s", str},
		{"%v", obj},
	}

	const logline = "Test this is a somewhat longer line, as in prod"

	for _, t := range tests {
		b.Run(fmt.Sprintf("single-format-%s/%v", t.format, t.value), func(b *testing.B) {
			format := logline + " " + t.format
			for i := 0; i < b.N; i++ {
				_ = fmt.Sprintf(format, t.value)
			}
		})

		b.Run(fmt.Sprintf("plain-%s/%v", t.format, t.value), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = fmt.Sprint(logline, t.value)
			}
		})
	}
}
