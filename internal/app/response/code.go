package response

type ResultCode int

const (
	Success               ResultCode = iota
	Failed                           = -1
	BadRequest                       = 400
	Unauthorized                     = 401
	Forbidden                        = 403
	NotFound                         = 404
	StatusTooManyRequests            = 429
	InternalServerError              = 500
)
