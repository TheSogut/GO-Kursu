package for_range

import (
	"fmt"
)

func Demo3() {
	dict := map[string]string{"book": "kitap", "table": "masa"}

	for k, v := range dict {
		fmt.Println(k)
		fmt.Println(v)
	}
}
