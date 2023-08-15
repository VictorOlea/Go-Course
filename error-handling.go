package main

import (
	"errors"
	"fmt"
)

func function1(argument int) (int, error) {
	if argument == 12 {
		return -1, errors.New("12 is not an acceptable argument")
	}
	return argument + 3, nil
}

type argError struct {
	argument int
	problem  string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.argument, e.problem)
}

func function2(argument int) (int, error) {
	if argument == 12 {
		return -1, &argError{argument, "Unacceptable Arguement"}
	}
	return argument + 3, nil
}

func main() {
	for _, i := range []int{7, 12} {
		if r, e := function1(i); e != nil {
			fmt.Println("Funcion1 Failed to execute successfuly \nerror:", e)
		} else {
			fmt.Println("Funcion1 executed successfully: ", r)
		}
	}

	for _, i := range []int{70, 12} {
		if r, e := function2(i); e != nil {
			fmt.Println("Funcion2 Failed to execute successfuly \nerror:", e)
		} else {
			fmt.Println("Funcion2 executed successfully: ", r)
		}
	}

	_, e := function2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.argument)
		fmt.Println(ae.problem)
	}
}
