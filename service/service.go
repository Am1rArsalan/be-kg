package service

import (
	"github.com/Am1rArsalan/kelvin-green/graph/model"
	"github.com/Am1rArsalan/kelvin-green/repo"
	"github.com/Am1rArsalan/kelvin-green/response"
)

type ServiceI interface {
	GetWeatherData(latitude, longitude string) (*model.Weather, error)
}

type Service struct {
	repo repo.OpenMeteoRepoI
}

func NewService(repo repo.OpenMeteoRepoI) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetWeatherData(latitude, longitude string) (*model.Weather, error) {
	weatherData, err := s.repo.GetWeatherFromOpenMeteo(latitude, longitude)
	if err != nil {
		return nil, err
	}

	return s.convertWeatherRespToModel(weatherData), nil
}

func (s *Service) convertWeatherRespToModel(weatherResp response.WeatherApiResponse) *model.Weather {
	weatherModel := &model.Weather{
		Latitude:             &weatherResp.Latitude,
		Longitude:            &weatherResp.Longitude,
		GenerationTimeMs:     &weatherResp.GenerationTimeMs,
		UtcOffsetSeconds:     &weatherResp.UtcOffsetSeconds,
		Timezone:             &weatherResp.Timezone,
		TimezoneAbbreviation: &weatherResp.TimezoneAbbreviation,
		Elevation:            &weatherResp.Elevation,
		HourlyUnits: &model.HourlyUnits{
			Time:               &weatherResp.HourlyUnits.Time,
			Temperature2m:      &weatherResp.HourlyUnits.Temperature2m,
			RelativeHumidity2m: &weatherResp.HourlyUnits.RelativeHumidity2m,
		},
		Hourly: &model.Hourly{
			Time:               convertStringsToPointers(weatherResp.Hourly.Time),
			Temperature2m:      convertFloatsToPointers(weatherResp.Hourly.Temperature2m),
			RelativeHumidity2m: convertFloatsToPointers(weatherResp.Hourly.RelativeHumidity2m),
		},
		DailyUnits: &model.DailyUnits{
			Time:                   &weatherResp.DailyUnits.Time,
			WeatherCode:            &weatherResp.DailyUnits.WeatherCode,
			Temperature2mMax:       &weatherResp.DailyUnits.Temperature2mMax,
			Temperature2mMin:       &weatherResp.DailyUnits.Temperature2mMin,
			ApparentTemperatureMax: &weatherResp.DailyUnits.ApparentTemperatureMax,
			ApparentTemperatureMin: &weatherResp.DailyUnits.ApparentTemperatureMin,
			Sunrise:                &weatherResp.DailyUnits.Sunrise,
			Sunset:                 &weatherResp.DailyUnits.Sunset,
		},
		Daily: &model.Daily{
			ApparentTemperatureMax:   convertFloatsToPointers(weatherResp.Daily.ApparentTemperatureMax),
			ApparentTemperatureMin:   convertFloatsToPointers(weatherResp.Daily.ApparentTemperatureMin),
			Sunrise:                  convertStringsToPointers(weatherResp.Daily.Sunrise),
			Sunset:                   convertStringsToPointers(weatherResp.Daily.Sunset),
			Time:                     convertStringsToPointers(weatherResp.Daily.Time),
			WeatherCode:              convertIntsToPointers(weatherResp.Daily.WeatherCode),
			Temperature2mMax:         convertFloatsToPointers(weatherResp.Daily.Temperature2mMax),
			Temperature2mMin:         convertFloatsToPointers(weatherResp.Daily.Temperature2mMin),
			WindDirection10mDominant: convertFloatsToPointers(weatherResp.Daily.WindDirection10mDominant),
		},
	}

	return weatherModel
}

func convertStringsToPointers(strings []string) []*string {
	pointers := make([]*string, len(strings))
	for i, str := range strings {
		pointers[i] = &str
	}
	return pointers
}

func convertFloatsToPointers(floats []float64) []*float64 {
	pointers := make([]*float64, len(floats))
	for i, f := range floats {
		pointers[i] = &f
	}
	return pointers
}

func convertIntsToPointers(ints []int) []*int {
	pointers := make([]*int, len(ints))
	for i, iVal := range ints {
		pointers[i] = &iVal
	}
	return pointers
}
