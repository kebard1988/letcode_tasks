package reversestring

import "fmt"

func reverseString(s []byte) {
	for i := len(s)-1; i>=0; i-- {
		fmt.Println(s[i])
	}
}
