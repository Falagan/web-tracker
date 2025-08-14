package ingestvisitor

import (
	"github.com/Falagan/web-tracker/pkg"
)

var (
	InvalidEvent    = pkg.Error{Code: "EINGESTVISITORINVALIDEVENT", Message: "Ingest Visitor invalid event"}
	SaveIngestError = pkg.Error{Code: "EINGESTSAVE", Message: "Ingest Visitor can not be stored"}
	InvalidUID      = pkg.Error{Code: "EUID", Message: "Invalid UID"}
)
