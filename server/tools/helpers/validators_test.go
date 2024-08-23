package helpers

import (
	"testing"
)

func TestMacAddressValidator(t *testing.T) {
	tests := []struct {
		name    string
		mac     string
		want    string
		wantErr bool
	}{
		{"valid MAC address with colon", "00:11:22:33:44:55", "00:11:22:33:44:55", false},
		{"valid MAC address with dash", "00-11-22-33-44-55", "00:11:22:33:44:55", false},
		{"valid MAC address without separator", "001122334455", "00:11:22:33:44:55", false},
		{"invalid MAC address length", "0011223344556", "", true},
		{"invalid MAC address characters", "00:11:22:33:44:ZZ", "", true},
		{"empty MAC address", "", "", true},
		{"nil MAC address", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MacAddressValidator(tt.mac)
			if (err != nil) != tt.wantErr {
				t.Errorf("MacAddressValidator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MacAddressValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}
