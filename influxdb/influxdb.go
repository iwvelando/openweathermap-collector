package influxdb

import (
	"context"
	"crypto/tls"
	"fmt"
	influx "github.com/influxdata/influxdb-client-go/v2"
	influxAPI "github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/iwvelando/openweathermap-collector/config"
	"github.com/iwvelando/openweathermap-collector/model"
)

// Connect establishes an InfluxDB client
func Connect(config *config.Configuration) (influx.Client, influxAPI.WriteAPIBlocking, error) {
	var auth string
	if config.InfluxDB.Token != "" {
		auth = config.InfluxDB.Token
	} else if config.InfluxDB.Username != "" && config.InfluxDB.Password != "" {
		auth = fmt.Sprintf("%s:%s", config.InfluxDB.Username, config.InfluxDB.Password)
	} else {
		auth = ""
	}

	var writeDest string
	if config.InfluxDB.Bucket != "" {
		writeDest = config.InfluxDB.Bucket
	} else if config.InfluxDB.Database != "" && config.InfluxDB.RetentionPolicy != "" {
		writeDest = fmt.Sprintf("%s/%s", config.InfluxDB.Database, config.InfluxDB.RetentionPolicy)
	} else {
		return nil, nil, fmt.Errorf("must configure at least one of bucket or database/retention policy")
	}

	options := influx.DefaultOptions().
		SetTLSConfig(&tls.Config{
			InsecureSkipVerify: config.InfluxDB.SkipVerifySsl,
		})
	client := influx.NewClientWithOptions(config.InfluxDB.Address, auth, options)

	writeAPI := client.WriteAPIBlocking(config.InfluxDB.Organization, writeDest)

	return client, writeAPI, nil
}

// WriteAll performs final processing and formatting of raw data before submitting to InfluxDB
func WriteAll(config *config.Configuration, writeAPI influxAPI.WriteAPIBlocking, forecast model.Forecast) error {

	// Write minutely precipitation data

	for _, minute := range forecast.OneCall.Minutely {
		p := influx.NewPoint(
			config.InfluxDB.MeasurementPrefix+"weather_forecast",
			map[string]string{
				"site_latitude":  fmt.Sprintf("%f", forecast.OneCall.Latitude),
				"site_longitude": fmt.Sprintf("%f", forecast.OneCall.Longitude),
			},
			map[string]interface{}{
				"precipitation_mm": minute.Precipitation,
			},
			minute.Time)
		err := writeAPI.WritePoint(context.Background(), p)
		if err != nil {
			return err
		}
	}

	return nil

}
