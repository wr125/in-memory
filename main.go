package main

import (
	"fmt"
	"sort"
)

//Countie is a simple map
type Countie struct {
	Borough map[string][]string
}

//List returns map list
func (c Countie) List() {
	for i, v := range c.Borough {
		val := append(v, "Lewisham")
		sort.Strings(val)
		fmt.Println(i, val)
	}
}

//Add add's new map as list
func (c Countie) Add(borough string, values ...string) {
	old, ok := c.Borough[borough]

	if ok {
		c.Borough[borough] = append(old, values...)
		return
	}
	c.Borough[borough] = values
}

//Delete removes item from map
func (c Countie) Delete(borough string, value string) {
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
	countie := Countie{Borough: map[string][]string{"London": {"Bexley", "Catford", "Camberwell"}}}
	countie.List()
	countie.Delete("London", "Catford")
	countie.Add("Kent", "Dover", "Canterbury", "Maidstone")

	countie.List()
}
