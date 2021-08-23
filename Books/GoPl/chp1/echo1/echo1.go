/**
 * File              : echo1.go
 * Author            : sanwich <122079260@qq.com>
 * Date              : 2021-02-07 15:12:42
 * Last Modified Date: 2021-02-08 14:34:58
 * Last Modified By  : sanwich <122079260@qq.com>
 */

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
