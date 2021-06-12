package main

import "fmt"

func main() {
	x:= map[string][]string{
		`last_name`: []string{`somenicethingstodo`},
		`bond_james`: []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Sceince`},
		`no_dr`: []string{`Being evil`, `Ice cream`, `Sunsets`}}

	for i,v := range x {
		for j := range v{
			fmt.Println(i,j,v[j])
		}
	}

	//adding a record

	x[`somenewguy`] =[]string {`money`, `fame`, `respect`}

	for i,v := range x {
		for j := range v{
			fmt.Println(i,j,v[j])
		}
	}
}




