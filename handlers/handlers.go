//Package handlers is a bunch of handlers that support needed function for routes/api.go
package handlers

// SuccessFunc will be executed if everything went well
type SuccessFunc func(string)

//FailFunc handles the opposite result, that is, when something goes wrong
type FailFunc func(string, error)

//ExecuteFunc is a type that defines the operation we want to perform
type ExecuteFunc func() (string, error)

//DO is structure of Future/promise pattern
type DO struct {
	successFunc SuccessFunc
	failFunc    FailFunc
}

// Success accepts successFunc
func (t *DO) Success(f SuccessFunc) *DO {
	t.successFunc = f
	return t
}

// Fail accepts FailFunc
func (t *DO) Fail(f FailFunc) *DO {
	t.failFunc = f
	return t
}

//Execute accepts chain of do functions
func (t *DO) Execute(f ExecuteFunc) {
	go func(t *DO) {
		strVal, err := f()
		if err != nil {
			t.failFunc(strVal, err)
		} else {
			t.successFunc(strVal)
		}
	}(t)
}
