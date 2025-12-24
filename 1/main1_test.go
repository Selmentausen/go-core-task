package main

import "testing"

func ExampleDiscernTypes() {
	DiscernTypes(42, "hello", true)

	// Output:
	// Type: int	Value: 42
	// Type: string	Value: hello
	// Type: bool	Value: true
}

func TestConcatenateValues(t *testing.T) {
	v1 := 42
	v2 := 3.14
	v3 := "test"

	expected := "423.14test"
	result := concatenateValues(v1, v2, v3)
	if result != expected {
		t.Errorf("concatenateValues returned %s, expected %s", result, expected)
	}
}

func TestHashWithSalt(t *testing.T) {
	input := []rune("123456")
	hash := hashWithSalt(input)
	if len(hash) != 64 {
		t.Errorf("Expected SHA256 hash length of 64, got %d", len(hash))
	}
}
