package webdev

import (
	"reflect"
	"testing"
)

func TestInitWebDav(t *testing.T) {
	type args struct {
		url      string
		userName string
		passWord string
		FilePath string
	}
	tests := []struct {
		name string
		args args
		want *WebDav
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitWebDav(tt.args.url, tt.args.userName, tt.args.passWord, tt.args.FilePath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitWebDav() = %v, want %v", got, tt.want)
			}
		})
	}
}
