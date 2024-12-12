package main

import (
	"fmt"

	"m7s.live/plugin/gb28181/v4/utils"
)

func main() {
	m := NewManager()

	m.Add(MatchContain, "a", "a", 12121)
	for i := range 100 {
		t := utils.RandString(12)
		m.Add(MatchContain, t, t, i+1)
	}
	m.Add(MatchContain, "ai", "ai", 102)
	m.Add(MatchContain, "main", "main", 103)
	m.Add(MatchPrefix, "mai", "mai", 1000)
	for i := range 100 {
		t := utils.RandString(12)
		m.Add(MatchPrefix, t, t, i+1001)
	}
	m.Add(MatchPrefix, "ma", "ma", 1102)
	m.Add(MatchPrefix, "main", "main", 1103)
	m.Add(MatchSuffix, "in", "in", 2000)
	for i := range 100 {
		t := utils.RandString(12)
		m.Add(MatchSuffix, t, t, i+2001)
	}
	m.Add(MatchSuffix, "ain", "ain", 2102)
	m.Add(MatchSuffix, "main", "main", 0)

	fmt.Println(m.Find("main"))
}
