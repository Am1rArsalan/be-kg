package response

type WeatherApiResponse struct {
	Latitude             float64     `json:"latitude"`
	Longitude            float64     `json:"longitude"`
	GenerationTimeMs     float64     `json:"generationtime_ms"`
	UtcOffsetSeconds     int         `json:"utc_offset_seconds"`
	Timezone             string      `json:"timezone"`
	TimezoneAbbreviation string      `json:"timezone_abbreviation"`
	Elevation            float64     `json:"elevation"`
	HourlyUnits          HourlyUnits `json:"hourly_units"`
	Hourly               Hourly      `json:"hourly"`
	DailyUnits           DailyUnits  `json:"daily_units"`
	Daily                Daily       `json:"daily"`
}

type HourlyUnits struct {
	Time               string `json:"time"`
	Temperature2m      string `json:"temperature_2m"`
	RelativeHumidity2m string `json:"relativehumidity_2m"`
}

type Hourly struct {
	Time               []string  `json:"time"`
	Temperature2m      []float64 `json:"temperature_2m"`
	RelativeHumidity2m []float64 `json:"relativehumidity_2m"`
}

type DailyUnits struct {
	Time                   string `json:"time"`
	WeatherCode            string `json:"weathercode"`
	Temperature2mMax       string `json:"temperature_2m_max"`
	Temperature2mMin       string `json:"temperature_2m_min"`
	ApparentTemperatureMax string `json:"apparent_temperature_max"`
	ApparentTemperatureMin string `json:"apparent_temperature_min"`
	Sunrise                string `json:"sunrise"`
	Sunset                 string `json:"sunset"`
}

type Daily struct {
	ApparentTemperatureMax   []float64 `json:"apparent_temperature_max"`
	ApparentTemperatureMin   []float64 `json:"apparent_temperature_min"`
	Sunrise                  []string  `json:"sunrise"`
	Sunset                   []string  `json:"sunset"`
	Time                     []string  `json:"ime"`
	WeatherCode              []int     `json:"weathercode"`
	Temperature2mMax         []float64 `json:"temperature_2m_max"`
	Temperature2mMin         []float64 `json:"temperature_2m_min"`
	WindDirection10mDominant []float64 `json:"winddirection_10m_dominant,omitempty"`
}
