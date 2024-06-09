package response

type JobZone struct {
	value        int32
	title        string
	experience   string
	education    string
	job_training string
	examples     string
	svp_range    string
}

type CareerIndustry struct {
	href  string
	code  int32
	title string
}

type Pagination struct {
	start int32
	end   int32
	total int32
}

type Link struct {
	href  string
	title string
}

type PaginationLink struct {
	href string
	rel  string
}

type AnswerOption struct {
	value int32
	name  string
}

type InterestProfilerQuestion struct {
	index int32
	area  string
	text  string
}

type InterestProfilerResult struct {
	area        string
	score       int32
	description string
}

type Tags struct {
	bright_outlook bool
	green          bool
	apprenticeship bool
}

type MatchingCareer struct {
	href  string
	fit   string
	code  string
	title string
	Tags
}
