package bestpractices

type ErrZeroDivision struct {
	message string
}

func NewErrZeroDivision(message string) *ErrZeroDivision {
	return &ErrZeroDivision{
		message: message,
	}
}

func (e *ErrZeroDivision) Error() string {
	return e.message
}

func Divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, NewErrZeroDivision("Cannot divide by zero")
	}
	return a / b, nil
}
