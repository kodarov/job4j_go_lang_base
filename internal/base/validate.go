package base

type ValidateRequest struct {
	UserID      string
	Title       string
	Description string
}

func Validate(req *ValidateRequest) []string {
	res := make([]string, 0)
	if req == nil {
		return append(res, "ValidateRequest nil")
	}
	if req.UserID == "" {
		res = append(res, "ValidateRequest UserID is empty")
	}
	if req.Title == "" {
		res = append(res, "ValidateRequest Title is empty")
	}
	if req.Description == "" {
		res = append(res, "ValidateRequest Description is empty")
	}

	return res
}
