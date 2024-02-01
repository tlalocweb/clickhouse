//go:build debug
// +build debug

package clickhouse

import "fmt"

func debugf(format string, args ...interface{}) {
	fmt.Printf("clickhouse-go: "+format, args...)
}

// func debugf(format string, args ...interface{}) {

// }
