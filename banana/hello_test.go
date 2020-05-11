package banana

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/assert"

func Test_hello(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "greet bob", args: args{s: "bob"}},
		{name: "greet boris", args: args{s: "boris"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValues(t, fmt.Sprintf("hello ! %s.", tt.args.s), hello(tt.args.s))
		})
	}
}
