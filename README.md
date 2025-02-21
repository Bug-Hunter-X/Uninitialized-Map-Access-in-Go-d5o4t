# Uninitialized Map Access in Go

This example demonstrates a common, yet subtle, issue in Go: accessing elements in uninitialized maps. Unlike some languages that throw an exception upon accessing a key in an uninitialized map, Go returns the zero value for the map's value type.