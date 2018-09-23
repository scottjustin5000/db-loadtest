package main

// ResponseResult represents the QueryResponse time including the queryIndex
type ResponseResult struct {
	Index          int
	ResponseTimeMs int
	ErrorMessage   string
}
