package deploy

import "testing"

func Test_deployMain(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deployMain()
		})
	}
}
