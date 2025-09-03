package domain

import (
	"context"

	"github.com/Falagan/web-tracker/pkg"
)

type Analytic struct {
	URL   URL
	Count URLCount
}

var (
	AnalyticInvalidURLError   = pkg.Error{Code: "EDANALYTICURL", Message: "Analytic has invalid URL"}
	AnalyticInvalidCountError = pkg.Error{Code: "EDANALYTICCOUNT", Message: "Analytic has invalid count"}
	AnalyticNoData            = pkg.Error{Code: "EANALYTICNODATA", Message: "Analytic url has no data"}
)

func NewAnalytic(url string, count int) (*Analytic, error) {
	validURL, err := NewURL(url)
	if err != nil {
		return nil, err
	}

	validCount, err := NewURLCount(count)

	if err != nil {
		return nil, err
	}

	return &Analytic{
		URL:   validURL,
		Count: validCount,
	}, nil
}

type AnalyticRepository interface {
	IncreaseVisitedURLCount(ctx context.Context, url string) error
	GetVisitedURLCount(ctx context.Context, url string) (*URLCount, error)
}
