package newrelic

import "fmt"

type errWrapper struct {
	err interface{}
}

func (e errWrapper) Error() string {
	return fmt.Sprintf("error %v", e.err)
}
