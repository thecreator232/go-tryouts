package main

import (
	"errors"
	"fmt"
)

func main() {
	e := NewEngine()
	j := 123
	// j := "thecreator232"
	// j := NewEngine()
	e.SetValue("t1", j)
	val, err := e.GetValue("t1")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(val)
}

type Engine struct {
	Data map[string]interface{}
}

func NewEngine() *Engine {
	return &Engine{
		Data: make(map[string]interface{}),
	}
}

func (e *Engine) GetValue(key string) (interface{}, error) {
	var value interface{}
	var ok bool
	if value, ok = e.Data[key]; ok != true {
		fmt.Println("Key not found")
		return nil, errors.New("key not found")
	}
	return value, nil
}

func (e *Engine) SetValue(key string, value interface{}) error {
	e.Data[key] = value
	return nil
}
