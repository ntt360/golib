/**
* @file shell_test.go
* @brief test exec shell
* @author ligang
* @version 1.0
* @date 2014-12-31
 */

package shell

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	result := RunCmd("lsa")
	fmt.Println(result)
}
