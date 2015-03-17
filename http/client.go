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

type T_Http_Response struct {
	Code     int
	Contents string
}

type t_http_client struct {
	client   *http.Client
	request  *http.Request
	response *T_Http_Response
}

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
func Get(rawurl string, timeout string, retry int, ip string) (*T_Http_Response, error) {
	http_client := newHttpClient()
	err := http_client.prepareRequest("GET", rawurl, timeout, ip)
	if nil != err {
		return nil, err
	}

	err = http_client.do(retry)
	return http_client.response, err
}

func newHttpClient() *t_http_client {
	http_client := new(t_http_client)
	http_client.client = new(http.Client)
	http_client.response = new(T_Http_Response)

	return http_client
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

	http_client.response.Code = response.StatusCode
	body, errb := ioutil.ReadAll(response.Body)
	if nil != errb {
		return errb
	}

	http_client.response.Contents = string(body)
	return nil
}
