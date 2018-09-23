package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//Config represents the settings and inputs for the LoadTest
type Config struct {
	Concurrency       int
	RequestsPerSecond int
	ConnectionString  string
	Random            bool
	Sequential        bool
	File              string
	Duration          int
	EOF               bool
}

// Queries reads queries from file into an array
func (c *Config) Queries() []string {
	var queries []string
	data, err := ioutil.ReadFile(c.File)
	if err != nil {
		fmt.Println(err)
	}
	_ = json.Unmarshal(data, &queries)
	return queries
}
