/*
Copyright © 2024 LitFill <marrazzy54@gmail.com>
*/
package main

import (
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "Nahwu",
		},
		{
			name: "Bahasa Arab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
