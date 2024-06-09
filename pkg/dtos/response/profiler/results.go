package profiler

import "go-webserver/pkg/dtos/response"

type InterestProfilerResultsResponse struct {
	answers string
	response.Pagination
	result []response.InterestProfilerResult
}
