package main

import (
	"fmt"
	"sort"
)

//County is a simple map
type County struct {
	Borough map[string][]string
}

//List returns map list
func (c County) List() {
	for i, v := range c.Borough {
		val := append(v, "Lewisham")
		sort.Strings(val)
		fmt.Println(i, val)
	}
}

//Add add's new map as list
func (c County) Add(borough string, values ...string) {
	old, ok := c.Borough[borough]

	if ok {
		c.Borough[borough] = append(old, values...)
		return
	}
	c.Borough[borough] = values
}

//Delete removes item from map
func (c County) Delete(borough string, value string) {
	old, ok := c.Borough[borough]

	if ok {
		var values []string

		for _, check := range old {
			if check != value {
				values = append(values, check)
			}
		}
		c.Borough[borough] = values
	}
}

func main() {
	county := County{Borough: map[string][]string{"London": {"Bexley", "Catford", "Camberwell"}}}
	county.List()
	county.Delete("London", "Catford")
	county.Add("Kent", "Dover", "Canterbury", "Maidstone")
	county.List()
}
