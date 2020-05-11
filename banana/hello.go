package banana

import "fmt"

func Hello(s string) string {
	ret := fmt.Sprintf("hello ! %s. THIS IS FUN" , s)
	fmt.Println(ret)
	return ret
}
