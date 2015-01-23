/**
* @file client.go
* @brief http client
* @author ligang
* @version 1.0
* @date 2014-12-31
 */

package http

import (
	//     "fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type t_http_response struct {
	code     int
	contents string
}

type t_http_client struct {
	client   *http.Client
	request  *http.Request
	response *t_http_response
}

var _http_client *t_http_client

/**
* @brief http get
*
* @param string
* @param string eg: 3s, 100ms
* @param int
* @param string
*
* @return
 */
func Get(rawurl string, timeout string, retry int, ip string) (string, error) {
	http_client := getHttpClient()
	err := http_client.prepareRequest("GET", rawurl, timeout, ip)
	if nil != err {
		return "", err
	}

	err = http_client.do(retry)
	return http_client.response.contents, err
}

/**
* @brief 200、50x and so on
*
* @return
 */
func GetHttpCode() int {
	return getHttpClient().response.code
}

func getHttpClient() *t_http_client {
	if nil == _http_client {
		_http_client = new(t_http_client)
		_http_client.client = new(http.Client)
		_http_client.response = new(t_http_response)
	}
	return _http_client
}

func (http_client *t_http_client) prepareRequest(method string, rawurl string, timeout string, ip string) error {
	td, err := time.ParseDuration(timeout)
	if nil != err {
		return err
	}
	http_client.client.Timeout = td

	request, _ := http.NewRequest(method, rawurl, nil)
	request.Host = request.URL.Host

	if "" != ip {
		s := strings.Split(request.URL.Host, ":")
		s[0] = ip
		request.URL.Host = strings.Join(s, ":")
	}

	request.Header.Set("Accept", "*/*")
	request.Header.Set("Cache-Control", "no-cache")
	request.Header.Set("Connection", "keep-alive")

	http_client.request = request

	return nil
}

func (http_client *t_http_client) do(retry int) error {
	response, err := http_client.client.Do(http_client.request)
	if nil != err {
		for i := 0; i < retry; i++ {
			response, err = http_client.client.Do(http_client.request)
			if nil == err {
				break
			}
		}
	}

	if nil != err {
		return err
	}

	http_client.response.code = response.StatusCode
	body, errb := ioutil.ReadAll(response.Body)
	if nil != errb {
		return errb
	}

	http_client.response.contents = string(body)
	return nil
}
