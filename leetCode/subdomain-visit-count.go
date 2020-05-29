package main

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {

	res := []string{}
	m := make(map[string]int)

	for _, str := range cpdomains {
		// split domain and visits
		slice := strings.Fields(str)
		visits, domain := slice[0], "."+slice[1]

		// create a map of domains and visits
		for i := len(domain) - 1; i >= 0; i-- {
			if domain[i] == '.' {
				key := domain[i+1:]
				val, _ := strconv.Atoi(visits)
				m[key] += val
			}
		}
	}

	// generate res from map
	for k, v := range m {
		res = append(res, strconv.Itoa(v)+" "+k)
	}
	return res
}
