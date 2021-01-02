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
	"github.com/spf13/viper"
)

// BaseURL endpoint.
func BaseURL() string {
	base := viper.GetString("baseURL")
	return base
}

// Profile endpoint.
func Profile(handle string) string {
	return BaseURL() + handle
}

// Followers endpoint.
func Followers(handle string) string {
	return Profile(handle) + "/followers"
}

// Following endpoint.
func Following(handle string) string {
	return Profile(handle) + "/following"
}

// Gists endpoint.
func Gists(handle string) string {
	return Profile(handle) + "/gists"
}

// Subscriptions endpoint.
func Subscriptions(handle string) string {
	return Profile(handle) + "/subscriptions"
}

// Orgs endpoint.
func Orgs(handle string) string {
	return Profile(handle) + "/orgs"
}

// Repos endpoint.
func Repos(handle string) string {
	return Profile(handle) + "/repos"
}
