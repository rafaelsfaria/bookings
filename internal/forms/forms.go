package forms

import (
	"net/http"
	"net/url"
)

// Form simple form struct
type Form struct {
	url.Values
	Errors errors
}

// New initializes form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors{},
	}
}

// Has checks if form field is in post or not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)

	if x == "" {
		f.Errors.Add(field, "This field cannot be empty")
		return false
	}

	return true
}

// Valid checks true if there are no errors
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
