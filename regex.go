package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "preach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("preach"))

	fmt.Println(r.FindString("preach patch"))

	fmt.Println("Start and End Indexes of match: ", r.FindStringIndex("pinch pouch"))

	fmt.Println(r.FindStringSubmatch("poach pitch"))

	fmt.Println(r.FindStringSubmatchIndex("punch"))

	fmt.Println(r.FindAllString("parch patch pitch", -1))

	fmt.Println("Indexes of all matches and submatches: ", r.FindAllStringSubmatch("potch pooch porch", -1))

	fmt.Println(r.FindAllString("prelaunch postlaunch pitch", 2))

	fmt.Println(r.Match([]byte("pinch")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp: ", r)

	fmt.Println(r.ReplaceAllString("pinch it!", "hurt"))

	in := []byte("The prelaunch")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
