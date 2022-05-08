package connect

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/iwvelando/openweathermap-collector/config"
	"github.com/iwvelando/openweathermap-collector/model"
	"io/ioutil"
	"net/http"
)

const expectedHTTPStatus = 200

// Client provides an HTTP client
func Client(config *config.Configuration) *http.Client {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: config.OpenWeatherMap.SkipVerifySsl}
	client := &http.Client{}

	return client
}

// GetEndpoint retrieves JSON formatted data from a specific endpoint of openweathermap.org
func GetEndpoint(config *config.Configuration, client *http.Client, endpoint string, data *model.OneCall) error {
	req, err := http.NewRequest("GET", config.OpenWeatherMap.BaseURL+endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	status := resp.StatusCode
	if status != expectedHTTPStatus {
		err = fmt.Errorf("expected %d HTTP status code but got %d; raw body %s", expectedHTTPStatus, resp.StatusCode, body)
		return err
	}

	err = json.Unmarshal(body, data)

	if err != nil {
		err = fmt.Errorf("%w; raw body %s", err, body)
		return err
	}

	return nil
}

// GetAll queries all required endpoints from openweathermap.org for forecast data
func GetAll(config *config.Configuration, client *http.Client) (model.Forecast, error) {

	forecasts := model.Forecast{}

	endpoint := fmt.Sprintf("/onecall?lat=%f&lon=%f&units=%s&lang=%s&appid=%s",
		config.OpenWeatherMap.Latitude, config.OpenWeatherMap.Longitude,
		config.OpenWeatherMap.Units, config.OpenWeatherMap.Language,
		config.OpenWeatherMap.ApiKey)
	err := GetEndpoint(config, client, endpoint, &forecasts.OneCall)

	if err != nil {
		return forecasts, err
	}

	return forecasts, nil

}
