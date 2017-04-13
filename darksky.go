// https://api.darksky.net/forecast/ba40ad71d8f85f4488d2f205c91ad174/52.5200,-13.4050
package weather

import "encoding/json"

type DarkSky struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Offset    int     `json:"offset"`
	Currently struct {
		Time                int     `json:"time"`
		Summary             string  `json:"summary"`
		Icon                string  `json:"icon"`
		PrecipIntensity     float64 `json:"precipIntensity"`
		PrecipProbability   float64 `json:"precipProbability"`
		PrecipType          string  `json:"precipType"`
		Temperature         float64 `json:"temperature"`
		ApparentTemperature float64 `json:"apparentTemperature"`
		DewPoint            float64 `json:"dewPoint"`
		Humidity            float64 `json:"humidity"`
		WindSpeed           float64 `json:"windSpeed"`
		WindBearing         int     `json:"windBearing"`
		Visibility          float64 `json:"visibility"`
		CloudCover          float64 `json:"cloudCover"`
		Pressure            float64 `json:"pressure"`
		Ozone               float64 `json:"ozone"`
	} `json:"currently"`
	Hourly struct {
		Summary string `json:"summary"`
		Icon    string `json:"icon"`
		Data    []struct {
			Time                int     `json:"time"`
			Summary             string  `json:"summary"`
			Icon                string  `json:"icon"`
			PrecipIntensity     float64 `json:"precipIntensity"`
			PrecipProbability   float64 `json:"precipProbability"`
			PrecipType          string  `json:"precipType,omitempty"`
			Temperature         float64 `json:"temperature"`
			ApparentTemperature float64 `json:"apparentTemperature"`
			DewPoint            float64 `json:"dewPoint"`
			Humidity            float64 `json:"humidity"`
			WindSpeed           float64 `json:"windSpeed"`
			WindBearing         int     `json:"windBearing"`
			Visibility          float64 `json:"visibility"`
			CloudCover          float64 `json:"cloudCover"`
			Pressure            float64 `json:"pressure"`
			Ozone               float64 `json:"ozone"`
		} `json:"data"`
	} `json:"hourly"`
	Daily struct {
		Summary string `json:"summary"`
		Icon    string `json:"icon"`
		Data    []struct {
			Time                       int     `json:"time"`
			Summary                    string  `json:"summary"`
			Icon                       string  `json:"icon"`
			SunriseTime                int     `json:"sunriseTime"`
			SunsetTime                 int     `json:"sunsetTime"`
			MoonPhase                  float64 `json:"moonPhase"`
			PrecipIntensity            float64 `json:"precipIntensity"`
			PrecipIntensityMax         float64 `json:"precipIntensityMax"`
			PrecipIntensityMaxTime     int     `json:"precipIntensityMaxTime,omitempty"`
			PrecipProbability          float64 `json:"precipProbability"`
			PrecipType                 string  `json:"precipType,omitempty"`
			TemperatureMin             float64 `json:"temperatureMin"`
			TemperatureMinTime         int     `json:"temperatureMinTime"`
			TemperatureMax             float64 `json:"temperatureMax"`
			TemperatureMaxTime         int     `json:"temperatureMaxTime"`
			ApparentTemperatureMin     float64 `json:"apparentTemperatureMin"`
			ApparentTemperatureMinTime int     `json:"apparentTemperatureMinTime"`
			ApparentTemperatureMax     float64 `json:"apparentTemperatureMax"`
			ApparentTemperatureMaxTime int     `json:"apparentTemperatureMaxTime"`
			DewPoint                   float64 `json:"dewPoint"`
			Humidity                   float64 `json:"humidity"`
			WindSpeed                  float64 `json:"windSpeed"`
			WindBearing                int     `json:"windBearing"`
			Visibility                 float64 `json:"visibility,omitempty"`
			CloudCover                 float64 `json:"cloudCover"`
			Pressure                   float64 `json:"pressure"`
			Ozone                      float64 `json:"ozone"`
		} `json:"data"`
	} `json:"daily"`
	Flags struct {
		Sources     []string `json:"sources"`
		IsdStations []string `json:"isd-stations"`
		Units       string   `json:"units"`
	} `json:"flags"`
}

const darkSkyURL = "https://api.darksky.net/forecast/ba40ad71d8f85f4488d2f205c91ad174/52.5200,-13.4050"

// DarkSkyTemp will fetch darsky temp
func DarkSkyTemp() (float64, error) {
	var data DarkSky
	content, err := Read(darkSkyURL)
	if err != nil {
		return 0.0, err
	}
	if err := json.Unmarshal(content, &data); err != nil {
		panic(err)
	}
	return data.Currently.Temperature, nil
}
