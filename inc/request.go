/*
Copyright © 2021 Gerald Villorente <geraldvillorente@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package inc

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// QueryArray Github API.
func QueryArray(endpoint string) []interface{} {
	var data []interface{}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(res), &data)

	return data
}

// Query Github API.
func Query(endpoint string) map[string]interface{} {
	var data map[string]interface{}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(res), &data)

	return data
}
