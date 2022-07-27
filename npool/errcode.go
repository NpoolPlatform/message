package npool

const (
	ErrService = 10000 + iota
	ErrParams
	ErrAlreadyExists
	ErrPermissionDenied
	ErrNotFound
)

const (
	ErrMsgServiceErr = "ServiceErr"
)
