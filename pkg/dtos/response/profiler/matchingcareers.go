package profiler

import "go-webserver/pkg/dtos/response"

type MatchingCareersResponse struct {
	answers string
	response.Pagination
	link   []response.PaginationLink
	career []response.MatchingCareer
}
