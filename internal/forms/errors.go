package forms

type errors map[string][]string

// Add appends an error message
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get returns first error message by field
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) > 0 {
		return es[0]
	}
	return ""
}
