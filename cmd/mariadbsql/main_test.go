package main

import "testing"

func TestMain(t *testing.T) {
	reset()
	var result int
	exit = func(code int) {
		result = code
	}
	main()
	if result != 0 {
		t.Errorf("Expected exit code of 0, but got %d", result)
	}
}

func TestMainNoConnection(t *testing.T) {
	dataSourceName = "bad_datasource"
	var result int
	exit = func(code int) {
		result = code
	}
	main()
	if result != 1 {
		t.Errorf("Expected exit code of 1, but got %d", result)
	}
}
