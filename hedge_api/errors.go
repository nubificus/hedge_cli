package hedge_api

import (
	"fmt"
)

type VMParamError struct {
	Param string
	Value string
}

func (e *VMParamError) Error() string {
	return fmt.Sprintf("Value '%s' not valid for %s parameter", e.Value, e.Param)
}
