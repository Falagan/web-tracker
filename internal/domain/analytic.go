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
	AnalyticNoData = pkg.Error{Code: "EANALYTICNODATA", Message: "Analytic url has no data"}
)

func NewAnalytic(url string, count int) (*Analytic, error) {
	validURL, err := NewURL(url)
	if err != nil {
		return nil, err
	}
	
	validCount := NewURLCount(count)
	
	return &Analytic{
		URL:   validURL,
		Count: validCount,
	}, nil
}

func (a *Analytic) Validate() error {
	if err := a.URL.Validate(); err != nil {
		return &AnalyticInvalidURLError
	}
	
	if !a.Count.IsValid() {
		return &AnalyticInvalidCountError
	}
	
	return nil
}

func (a *Analytic) IsValid() bool {
	return a.Validate() == nil
}

type AnalyticRepository interface {
	IncreaseVisitedURLCount(ctx context.Context, url URL) error
	GetVisitedURLCount(ctx context.Context, url URL) (*URLCount, error)
}
