package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr error
	}{
		{
			name:    "valid header",
			headers: http.Header{"Authorization": []string{"ApiKey abc123"}},
			want:    "abc123",
			wantErr: nil,
		},
		{
			name:    "missing header",
			headers: http.Header{},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
			if err != tt.wantErr {
				t.Errorf("got err %v, want %v", err, tt.wantErr)
			}
		})
	}
}
