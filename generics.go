package main

import (
	"fmt"
)

// LogLevel is a logging level
type LogLevel uint8

// possible log levels
const (
	DebugLevel LogLevel = iota + 1
	WarningLevel
	ErrorLevel
)

// string implements the fmt.Stringer interface
func (l LogLevel) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case WarningLevel:
		return "warning"
	case ErrorLevel:
		return "error"
	}
	return fmt.Sprintf("unknow log level: %d", l)
}

func main() {

	fmt.Println(WarningLevel)

	lvl_1 := LogLevel(1)
	fmt.Println(lvl_1)

	lvl_2 := LogLevel(2)
	fmt.Println(lvl_2)

	lvl_3 := LogLevel(3)
	fmt.Println(lvl_3)

	lvl_0 := LogLevel(10)
	fmt.Println(lvl_0)
}
