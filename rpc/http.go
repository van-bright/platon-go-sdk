// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	contentType             = "application/json"
	httpMethod              = http.MethodPost
)

type HttpClient struct {
	endPoint string
	client *http.Client
	nextId uint64
}

func NewHttpClient(url string) (*HttpClient) {
	return &HttpClient{url, &http.Client{}, 1}
}

func (c *HttpClient) Post(method string, p interface{}) ([]byte, error) {
	params := JsonRequest{
		Version: "2.0",
		Id: c.nextId,
		Method:  method,
		Payload: p,
	}

	c.nextId++

	jsonStr, err := json.Marshal(params)
	if err != nil {
		fmt.Println("marsh error: ", err)
		return []byte(`""`),  err
	}

	req, err := http.NewRequest(httpMethod, c.endPoint, bytes.NewBuffer(jsonStr))
	if err != nil {
		return []byte(`""`), err
	}

	req.Header.Set("Content-Type", contentType)

	resp, err := c.client.Do(req)
	if err != nil {
		return []byte(`""`), err
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode
	if statusCode != http.StatusOK {
		return []byte(`""`), fmt.Errorf("Status Code: %d, Message: %s", statusCode, resp.Status)
	}

	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		return []byte(`""`), err
	} else {
		return body, nil
	}
}
