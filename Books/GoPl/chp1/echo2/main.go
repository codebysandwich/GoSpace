/**
 * File              : main.go
 * Author            : sanwich <122079260@qq.com>
 * Date              : 2021-08-13 18:02:34
 * Last Modified Date: 2021-08-13 18:14:45
 * Last Modified By  : sanwich <122079260@qq.com>
 */

package main

import (
	"fmt"
	"os"
)

func main() {
	sep, s := "", ""
	for _, arg := range os.Args[:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
