package errors

type ErrorValidation struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
	Value  string `json:"reason_value,omitempty"`
}

func (p *ErrorValidation) Error() string {
	return p.Field
}

type ErrorValidations []*ErrorValidation

func (p ErrorValidations) Error() string {
	str := ""
	if len(p) > 0 {
		str = p[0].Error()
	}

	return str
}
