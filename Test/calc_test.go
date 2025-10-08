package main

import (
	"fmt"
	"os"
	"testing"
)

func setup() {
	fmt.Println("测试前")
}

func teardown() {
	fmt.Println("测试后")
}

func TestAdd(t *testing.T) {
	fmt.Println("测试中")
	cases := []struct {
		name           string
		a, b, expected int
	}{
		{"Add1", 1, 2, 3},
		{"Add2", 1, -1, 0},
		{"Add3", 1, -2, -1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := Add(c.a, c.b)
			if result != c.expected {
				t.Errorf("Expected %d, but got %d", c.expected, result)
			}
		})
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
