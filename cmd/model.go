package cmd

import "strings"

// Gurume data type
type Gurume struct {
	Category []Category `json:"category"`
	Town     string     `json:"town,omitempty"`
	Station  []Station  `json:"station,omitempty"`
	Name     string     `json:"name"`
	Note     string     `json:"note,omitempty"`
}

// NewGurume is to init gurume
func NewGurume() *Gurume {
	c := make([]Category, 0)
	s := make([]Station, 0)
	return &Gurume{
		Category: c,
		Station:  s,
	}
}

// SetCategory is
func (g *Gurume) SetCategory(categoryName string) *Gurume {
	for _, c := range strings.Split(categoryName, ",") {
		c = strings.TrimSpace(c)
		if c != "N/A" && c != "" {
			g.Category = append(g.Category, Category{c})
		}
	}
	return g
}

// SetStation is
func (g *Gurume) SetStation(stationName string) *Gurume {
	for _, s := range strings.Split(stationName, ",") {
		s = strings.TrimSpace(s)
		if s != "N/A" && s != "" {
			g.Station = append(g.Station, Station{s})
		}
	}
	return g
}

// SetTown is
func (g *Gurume) SetTown(town string) *Gurume {
	town = strings.TrimSpace(town)
	g.Town = town
	return g
}

// SetName is
func (g *Gurume) SetName(name string) *Gurume {
	name = strings.TrimSpace(name)
	g.Name = name
	return g
}

// SetNote is
func (g *Gurume) SetNote(note string) *Gurume {
	note = strings.TrimSpace(note)
	g.Note = note
	return g
}

// Station data type
type Station struct {
	Name string `json:"name,omitempty"`
}

// Category data type
type Category struct {
	Name string `json:"name,omitempty"`
}
