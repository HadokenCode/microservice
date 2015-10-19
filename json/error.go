package json

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *Error) Error() string {
	return e.Err.Error()
}
