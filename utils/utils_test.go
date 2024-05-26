package utils

import (
	"bytes"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestJsonDecoder(t *testing.T) {
	type args struct {
		r *http.Request
	}

	type TestStruct struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	tests := []struct {
		name    string
		args    args
		want    *TestStruct
		wantErr bool
	}{
		{
			name: "valid json",
			args: args{
				r: &http.Request{
					Body: io.NopCloser(
						bytes.NewBufferString(`{"name": "test", "value": "1"}`)),
				},
			},
			want: &TestStruct{
				Name:  "test",
				Value: "1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonDecoder[TestStruct](tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonDecoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonDecoder() = %v, want %v", *got, tt.want)
			}
		})
	}
}
