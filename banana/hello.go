package banana

import "fmt"

func hello(s string) string {
	ret := fmt.Sprintf("hello ! %s.", s)
	fmt.Println(ret)
	return ret
}
