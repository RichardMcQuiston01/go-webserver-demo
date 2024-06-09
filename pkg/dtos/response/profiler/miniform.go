package profiler

import "go-webserver/pkg/dtos/response"

type MiniFormResponse struct {
	response.Pagination
	link           []response.PaginationLink
	answer_options []response.AnswerOption
	question       []response.InterestProfilerQuestion
}
