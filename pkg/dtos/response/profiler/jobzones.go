package profiler

import "go-webserver/pkg/dtos/response"

type JobZonesResponse struct {
	response.Pagination
	job_zone []response.JobZone
}
