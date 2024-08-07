package utils

import (
	"testing"
)

func TestMd5(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{

		{
			name: "Calculate the md5 hash of a string",
			args: args{
				s: "http://example.com",
			},
			want: "687474703a2f2f6578616d706c652e636f6dd41d8cd98f00b204e9800998ecf8427e",
		},
		{
			name: "Should return the same hash for the same string",
			args: args{
				s: "http://example.com",
			},
			want: "687474703a2f2f6578616d706c652e636f6dd41d8cd98f00b204e9800998ecf8427e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5(tt.args.s); got != tt.want {
				t.Errorf("Md5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want Result
	}{

		{
			name: "Should return true for a valid URL",
			args: args{
				url: "https://mzunino.com.uy",
			},
			want: Result{IsValid: true, Message: "https://mzunino.com.uy"},
		},
		{
			name: "Should return false for an invalid URL",
			args: args{
				url: "https://nonexistent.mzunino.com.uy",
			},
			want: Result{IsValid: false, Message: "We couldn't reach the server"},
		},
		{
			name: "Should return false for an invalid URL",
			args: args{
				url: "https://httpstat.us/404",
			},
			want: Result{IsValid: false, Message: "Not found"},
		},
		{
			name: "Should return false for an empty URL",
			args: args{
				url: "",
			},
			want: Result{IsValid: false, Message: "URL cannot be empty"},
		},
		{
			name: "Should return false for an empty URL",
			args: args{
				url: "--??",
			},
			want: Result{IsValid: false, Message: "Invalid URL provided"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidURL(tt.args.url); got.IsValid != tt.want.IsValid || got.Message != tt.want.Message {
				t.Errorf("IsValidURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateShortKey(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Should generate a short key",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateShortKey()
			got2 := GenerateShortKey()
			if got == got2 {
				t.Errorf("GenerateShortKey() = %v, want %v", got, got2)
			}
		})
	}
}
