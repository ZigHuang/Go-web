package gee

import "testing"

func TestNestedGroup(t *testing.T) {
	r := New()
	api1 := r.Group("/api1")
	api2 := api1.Group("/api2")
	api3 := api2.Group("/api3")
	if api2.prefix != "/api1/api2" {
		t.Fatal("api2 prefix should be /api1/api2")
	}
	if api3.prefix != "/api1/api2/api3" {
		t.Fatal("api2 prefix should be /api1/api2")
	}
}
