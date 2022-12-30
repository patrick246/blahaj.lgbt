package prometheus

import (
	"context"
	"errors"
	"fmt"
	"github.com/patrick246/blahaj.lgbt/server/internal/datasources"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"go.uber.org/zap"
	"time"
)

const (
	prometheusQuery = "ikea_blahaj_count * on(store) group_left(name, country, lat, lon) ikea_info"
	countryQuery    = `sum by (country) (ikea_blahaj_count * on(store) group_left(name, country, lat, lon) ikea_info)`
	storeID         = "store"
	storeName       = "name"
	storeCountry    = "country"
	storeLat        = "lat"
	storeLon        = "lon"
)

var (
	ErrUnexpectedResult = errors.New("result is unexpected")
)

type Datasource struct {
	api v1.API
	log *zap.SugaredLogger
}

func NewDatasource(api v1.API, log *zap.SugaredLogger) *Datasource {
	return &Datasource{
		api: api,
		log: log,
	}
}

func (d *Datasource) GlobalAvailability(ctx context.Context) ([]datasources.CountryAvailability, error) {
	val, warnings, err := d.api.Query(ctx, countryQuery, time.Now())
	if err != nil {
		return nil, err
	}

	for _, warning := range warnings {
		d.log.Warnw("prometheus api warning", "message", warning)
	}

	vector, ok := val.(model.Vector)
	if !ok {
		return nil, fmt.Errorf("%w: result type is not a vector", ErrUnexpectedResult)
	}

	results := make([]datasources.CountryAvailability, 0, len(vector))

	for _, sample := range vector {
		results = append(results, datasources.CountryAvailability{
			Country: string(sample.Metric[storeCountry]),
			Number:  int64(sample.Value),
		})
	}

	return results, nil
}

func (d *Datasource) Availability(ctx context.Context) ([]datasources.BlahajAvailability, error) {
	/*reqCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()*/

	val, warnings, err := d.api.Query(ctx, prometheusQuery, time.Now())
	if err != nil {
		return nil, err
	}

	for _, warning := range warnings {
		d.log.Warnw("prometheus api warning", "message", warning)
	}

	vector, ok := val.(model.Vector)
	if !ok {
		return nil, fmt.Errorf("%w: result type is not a vector", ErrUnexpectedResult)
	}

	results := make([]datasources.BlahajAvailability, 0, len(vector))

	for _, sample := range vector {
		results = append(results, datasources.BlahajAvailability{
			StoreID:      string(sample.Metric[storeID]),
			StoreName:    string(sample.Metric[storeName]),
			StoreCountry: string(sample.Metric[storeCountry]),
			Location: datasources.GeoCoordinates{
				Latitude:  string(sample.Metric[storeLat]),
				Longitude: string(sample.Metric[storeLon]),
			},
			Number: int64(sample.Value),
		})
	}

	return results, nil
}
