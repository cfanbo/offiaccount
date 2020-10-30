package util

import (
	"io/ioutil"

	"github.com/levigross/grequests"
)

func HttpGet(url string, params map[string]string) ([]byte, error) {
	requestOptions := &grequests.RequestOptions{
		Params: params,
	}

	resp, err := grequests.Get(url, requestOptions)
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	return ioutil.ReadAll(resp.RawResponse.Body)
}

func HttpPost(url string, params map[string]string, data map[string]string) ([]byte, error) {
	requestOptions := &grequests.RequestOptions{
		Params: params,
		Data:   data,
	}
	resp, err := grequests.Post(url, requestOptions)
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	return ioutil.ReadAll(resp.RawResponse.Body)
}
