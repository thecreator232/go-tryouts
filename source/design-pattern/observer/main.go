package main

import (
	"fmt"
	"time"
)

type Observer interface {
	Update(value float64)
	GetObserverName() string
}

type Subject interface {
	NotifyObservers()
	Register(ob Observer)
	DeRegister(ob Observer)
}

type ShareScrip struct {
	observers  []Observer
	ScripName  string
	ScripValue float64
}

func NewScrip(scripName string, scripValue float64) *ShareScrip {
	return &ShareScrip{
		ScripName:  scripName,
		ScripValue: scripValue,
	}
}

func (s *ShareScrip) Register(ob Observer) {
	s.observers = append(s.observers, ob)
	return
}

func (s *ShareScrip) DeRegister(ob Observer) {
	obName := ob.GetObserverName()

	for i, val := range s.observers {
		if val.GetObserverName() == obName {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ShareScrip) NotifyObservers() {
	for _, val := range s.observers {
		val.Update(s.ScripValue)
	}
}

func (s *ShareScrip) SetValue(f float64) {
	s.ScripValue = f
	s.NotifyObservers()
}

type LineGraph struct {
	values       map[string]float64
	observerName string
}

func NewLineGraphObject() *LineGraph {
	return &LineGraph{
		observerName: "LineGraph",
		values:       map[string]float64{},
	}
}

func (l *LineGraph) GetObserverName() string {
	return l.observerName
}

func (l *LineGraph) Update(value float64) {

	now := time.Now()

	l.values[now.String()] = value

}

func (l *LineGraph) RepresentLineGraph() {

	for key, val := range l.values {
		fmt.Println(key, " --> ", val)
	}

}

func main() {
	lineGraph := NewLineGraphObject()
	ss1 := NewScrip("itc", 400.0)
	ss1.Register(lineGraph)

	lineGraph2 := NewLineGraphObject()
	ss2 := NewScrip("cochinShipyard", 1000.0)
	ss2.Register(lineGraph2)

	ss1.SetValue(404.0)
	ss1.SetValue(401.0)
	lineGraph.RepresentLineGraph()

	ss2.SetValue(300.0)
	lineGraph2.RepresentLineGraph()
	ss2.DeRegister(lineGraph2)

}
