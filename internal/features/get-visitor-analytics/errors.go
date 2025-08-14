package getvisitoranalytics

import (
	"github.com/Falagan/web-tracker/pkg"
)

var (
	InvalidURL         = pkg.Error{Code: "EGETANALYTICSURL", Message: "Get Analytics invalid URL"}
	GetAnalyticsError  = pkg.Error{Code: "EGETANALYTICS", Message: "Analytics can not be retrieved"}
)