package base

type ValidateRequest struct {
	UserID      string
	Title       string
	Description string
}

func Validate(req *ValidateRequest) []string {
	res := make([]string, 0)
	if req == nil || req.UserID == "" || req.Title == "" || req.Description == "" {
		return append(res, "validate error")
	}
	return append(res, "validate message")
}
