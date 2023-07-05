package main

import (
	"fmt"
	"sort"
)

func main() {

	states := make(map[string]string)

	states["WA"] = "Washington"
	states["NY"] = "New York"
	states["CA"] = "Caliornia"

	keys := make([]string, len(states))
	i := 0
	for key := range states {
		keys[i] = key
		i++
	}

	fmt.Println("Order of keys before sorting: ", keys)
	sort.Strings(keys)
	fmt.Println("Order of keys after forting: ", keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
