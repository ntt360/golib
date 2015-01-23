/**
* @file client_test.go
* @brief test http client
* @author ligang
* @version 1.0
* @date 2014-12-31
 */

package http

import (
	"testing"
)

func TestQuery(t *testing.T) {
	contents, _ := Get("http://www.job360.com/test.php", "3s", 3, "127.0.0.1")
	println(contents)
}
