// Copyright 2016 Google Inc. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to writing, software distributed
// under the License is distributed on a "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.

package events

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"golang.org/x/net/context"

	"google.golang.org/appengine/urlfetch"
)

const (
	apiURL          = "http://api.openweathermap.org/data/2.5/weather"
	iconURLTemplate = "http://openweathermap.org/img/w/%s.png"
)

func weather(ctx context.Context, location string) (*Weather, error) {
	// TODO: check if the weather for the location is in memcache.
	// If it is return it directly, if not just log the error if it is not a cache miss.

	// Prepare the request to the weather API.
	values := make(url.Values)
	values.Set("APPID", os.Getenv("WEATHER_API_KEY"))
	values.Set("q", location)
	url := apiURL + "?" + values.Encode()

	res, err := urlfetch.Client(ctx).Get(url)
	if err != nil {
		return nil, fmt.Errorf("could not get weather: %v", err)
	}

	// We need to close the body of the API response to avoid leaks.
	defer res.Body.Close()

	// We need to decode the list of weathers and the error message.
	var data struct {
		Weather []Weather
		Message string
	}
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("could not decode weather: %v", err)
	}

	// If the error message is not empty, something bad happened.
	if data.Message != "" {
		return nil, fmt.Errorf("no weather found: %s", data.Message)
	}

	// Check whether we received any weather.
	if len(data.Weather) == 0 {
		return nil, fmt.Errorf("no weather found")
	}

	// We just take the first value for the weather.
	weather := data.Weather[0]
	// And make the icon a complete url.
	weather.Icon = fmt.Sprintf(iconURLTemplate, weather.Icon)

	// TODO: cache the weather in memcache for later.
	// If there's an error just log it and return the weather.

	return &weather, nil
}
