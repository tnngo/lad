package ladglobal

import (
	"testing"

	"github.com/tnngo/lad"
)

func TestBuildDefaultFile(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BuildDefaultFile()
			lad.L().Debug("test")
		})
	}
}
