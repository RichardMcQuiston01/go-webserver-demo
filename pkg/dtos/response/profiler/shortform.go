package profiler

import "go-webserver/pkg/dtos/response"

type ShortFormResponse struct {
	response.Pagination
	link           []response.PaginationLink
	answer_options []response.AnswerOption
	question       []response.InterestProfilerQuestion
}
