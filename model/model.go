package model

import (
	"time"
)

// DateTimeForecast provides a format for timestamps provided by openweathermap.org
const DateTimeForecast = "1494505756"

// Forecast contains all forecast data
type Forecast struct {
	OneCall OneCall
}

// OneCall contains a response for /onecall
type OneCall struct {
	Latitude       float64    `json:"lat"`
	Longitude      float64    `json:"lon"`
	Timezone       string     `json:"timezone"`
	TimezoneOffset int        `json:"timezone_offset"`
	Current        Current    `json:"current"`
	Minutely       []Minutely `json:"minutely"`
	Hourly         []Hourly   `json:"hourly"`
	Daily          []Daily    `json:"daily"`
	Alerts         []Alerts   `json:"alerts"`
}

type Current struct {
	TimeRaw     int64 `json:"dt"`
	Time        time.Time
	SunriseRaw  int64 `json:"sunrise"`
	Sunrise     time.Time
	SunsetRaw   int64 `json:"sunset"`
	Sunset      time.Time
	Temperature float64          `json:"temp"`
	FeelsLike   float64          `json:"feels_like"`
	Pressure    int              `json:"pressure"`
	Humidity    int              `json:"humidity"`
	DewPoint    float64          `json:"dew_point"`
	UVI         float64          `json:"uvi"`
	Clouds      int              `json:"clouds"`
	Visibility  int              `json:"visibility"`
	WindSpeed   float64          `json:"wind_speed"`
	WindInt     float64          `json:"wind_deg"`
	Weather     []OneCallWeather `json:"weather"`
}

type OneCallWeather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Minutely struct {
	TimeRaw       int64 `json:"dt"`
	Time          time.Time
	Precipitation float64 `json:"precipitation"`
}

type Hourly struct {
	TimeRaw                    int64 `json:"dt"`
	Time                       time.Time
	Temperature                float64          `json:"temp"`
	FeelsLike                  float64          `json:"feels_like"`
	Pressure                   int              `json:"pressure"`
	Humidity                   int              `json:"humidity"`
	DewPoint                   float64          `json:"dew_point"`
	UVI                        float64          `json:"uvi"`
	Clouds                     int              `json:"clouds"`
	Visibility                 int              `json:"visibility"`
	WindSpeed                  float64          `json:"wind_speed"`
	WindInt                    float64          `json:"wind_deg"`
	WindGust                   float64          `json:"wind_gust"`
	Weather                    []OneCallWeather `json:"weather"`
	ProbabilityOfPrecipitation float64          `json:"pop"`
}

type Daily struct {
	TimeRaw                    int64 `json:"dt"`
	Time                       time.Time
	SunriseRaw                 int64 `json:"sunrise"`
	Sunrise                    time.Time
	SunsetRaw                  int64 `json:"sunset"`
	Sunset                     time.Time
	MoonriseRaw                int64 `json:"moonrise"`
	Moonrise                   time.Time
	MoonsetRaw                 int64 `json:"moonset"`
	Moonset                    time.Time
	MoonPhase                  float64            `json:"moon_phase"`
	Temperature                OneCallTemperature `json:"temp"`
	FeelsLike                  OneCallFeelsLike   `json:"feels_like"`
	Pressure                   int                `json:"pressure"`
	Humidity                   int                `json:"humidity"`
	DewPoint                   float64            `json:"dew_point"`
	WindSpeed                  float64            `json:"wind_speed"`
	WindInt                    float64            `json:"wind_deg"`
	WindGust                   float64            `json:"wind_gust"`
	Clouds                     int                `json:"clouds"`
	ProbabilityOfPrecipitation float64            `json:"pop"`
	Rain                       float64            `json:"rain"`
	UVI                        float64            `json:"uvi"`
}

type OneCallTemperature struct {
	Day     float64 `json:"day"`
	Min     float64 `json:"min"`
	Max     float64 `json:"max"`
	Night   float64 `json:"night"`
	Evening float64 `json:"eve"`
	Morning float64 `json:"morn"`
}

type OneCallFeelsLike struct {
	Day     float64 `json:"day"`
	Night   float64 `json:"night"`
	Evening float64 `json:"eve"`
	Morning float64 `json:"morn"`
}

type Alerts struct {
	SenderName  string `json:"sender_name"`
	Event       string `json:"event"`
	StartRaw    int64  `json:"start"`
	Start       time.Time
	EndRaw      int64 `json:"end"`
	End         time.Time
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

func (o *OneCall) ParseTime() {

	o.Current.Time = time.Unix(o.Current.TimeRaw, 0)
	o.Current.Sunrise = time.Unix(o.Current.SunriseRaw, 0)
	o.Current.Sunset = time.Unix(o.Current.SunsetRaw, 0)

	for i, minute := range o.Minutely {
		o.Minutely[i].Time = time.Unix(minute.TimeRaw, 0)
	}

	for i, hour := range o.Hourly {
		o.Hourly[i].Time = time.Unix(hour.TimeRaw, 0)
	}

	for i, day := range o.Daily {
		o.Daily[i].Time = time.Unix(day.TimeRaw, 0)
		o.Daily[i].Sunrise = time.Unix(day.SunriseRaw, 0)
		o.Daily[i].Sunset = time.Unix(day.SunsetRaw, 0)
		o.Daily[i].Moonrise = time.Unix(day.MoonriseRaw, 0)
		o.Daily[i].Moonset = time.Unix(day.MoonsetRaw, 0)
	}

	for i, alert := range o.Alerts {
		o.Alerts[i].Start = time.Unix(alert.StartRaw, 0)
		o.Alerts[i].End = time.Unix(alert.EndRaw, 0)
	}

}
