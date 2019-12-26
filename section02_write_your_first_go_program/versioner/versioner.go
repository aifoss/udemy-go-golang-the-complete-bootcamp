package versioner

import "runtime"

/**
 * Created by sofia on 2019-12-23.
 */

/**
 * Source: Udemy Go (Golang) The Complete Bootcamp Course
 */

// Version returns Go version
func Version() string {
	return runtime.Version()
}
