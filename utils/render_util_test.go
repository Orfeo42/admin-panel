package utils

import (
	"testing"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func TestRender(t *testing.T) {
	type args struct {
		component templ.Component
		context   echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Render(tt.args.component, tt.args.context); (err != nil) != tt.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
