// http is sample package
package http

import "github.com/go-resty/resty/v2"

type httpLogger struct{ l resty.Logger }

// Errorf nop
func (l httpLogger) Errorf(format string, v ...interface{}) {}

// Warnf nop
func (l httpLogger) Warnf(format string, v ...interface{}) {}

// Debugf nop
func (l httpLogger) Debugf(format string, v ...interface{}) {}
