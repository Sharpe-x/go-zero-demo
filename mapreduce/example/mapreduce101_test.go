package example

import "testing"

func Test_MapStrToStr(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "runMapStrToStr",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runMapStrToStr()
		})
	}
}

func Test_MapStrToInt(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "runMapStrToInt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runMapStrToInt()
		})
	}
}

func Test_runReduce(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "runReduce",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runReduce()
		})
	}
}

func Test_runFilter(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "runFilter",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runFilter()
		})
	}
}
