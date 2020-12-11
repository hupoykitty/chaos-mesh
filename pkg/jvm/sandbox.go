// Copyright 2020 Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package jvm

import (
	"bytes"
	"fmt"
	"net/http"
)

const (
	BaseURL    = "http://%s:%d/sandbox/default/module/http/"
	ActiveURL  = BaseURL + "sandbox-module-mgr/active?ids=chaosblade"
	InjectURL  = BaseURL + "chaosblade/create"
	RecoverURL = BaseURL + "chaosblade/destroy"
)

func ActiveSandbox(host string, port int) error {
	url := fmt.Sprintf(ActiveURL, host, port)
	_, err := http.Get(url)
	if err != nil {
		return err
	}
	return nil
}

func InjectChaos(host string, port int, body []byte) error {
	client := &http.Client{}
	reqBody := bytes.NewBuffer([]byte(body))
	url := fmt.Sprintf(InjectURL, host, port)
	request, _ := http.NewRequest("POST", url, reqBody)
	request.Header.Set("Content-type", "application/json")
	_, err := client.Do(request)
	if err != nil {
		return err
	}
	return nil
}

func RecoverChaos(host string, port int, body []byte) error {
	client := &http.Client{}
	reqBody := bytes.NewBuffer([]byte(body))
	url := fmt.Sprintf(RecoverURL, host, port)
	request, _ := http.NewRequest("POST", url, reqBody)
	request.Header.Set("Content-type", "application/json")
	_, err := client.Do(request)
	if err != nil {
		return err
	}
	return nil
}
