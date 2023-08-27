package main

import (
	"fmt"
)

type Thermostat struct {
	ID    string
	value float64
}

func (t *Thermostat) Value() float64 {
	return t.value
}

func (t *Thermostat) Set(value float64) {
	t.value = value
}

func (t *Thermostat) Kind() string {
	return "thermostat"
}

func main() {
	t := Thermostat{"Living Room", 16.2}
	fmt.Printf("%s Before: %.2f\n", t.ID, t.Value())
	t.Set(18)
	fmt.Printf("%s After: %.2f\n", t.ID, t.Value())

	t2 := Thermostat{"Bed Room", 14.8}
	fmt.Println(t2)
}
