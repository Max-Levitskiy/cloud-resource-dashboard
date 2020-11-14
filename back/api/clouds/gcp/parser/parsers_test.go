package parser

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want GcpResourceName
	}{
		{
			name: "Should parse name",
			args: args{name: "projects/some-project-123/locations/europe-west3/functions/some-function-name"},
			want: GcpResourceName{
				ProjectId:    "some-project-123",
				Location:     "europe-west3",
				ResourceType: "functions",
				ResourceName: "some-function-name",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseName() = %v, want %v", got, tt.want)
			}
		})
	}
}
