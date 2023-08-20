package main

import (
	"fmt"
)

type Thermostats struct {
	id    string
	value float64
}

func (t *Thermostats) Value() float64 {
	return t.value
}

func (t *Thermostats) ID() string {
	return t.id
}

func (t *Thermostats) Set(value float64) {
	t.value = value
}

func (t *Thermostats) Kind() string {
	return "Thermostat"
}

type Camera struct {
	id string
}

func (c *Camera) ID() string {
	return c.id
}

func (*Camera) Kind() string {
	return "camera"
}

type Sensor interface {
	ID() string
	Kind() string
}

func printAll(sensor []Sensor) {
	for _, s := range sensor {
		fmt.Println("%s <%s>\n", s.ID(), s.Kind())
	}
}

func main() {
	t := Thermostats{"Living Room", 16.2}
	c := Camera{"Baby Room"}

	sensors := []Sensor{&t, &c}
	printAll(sensors)

}
