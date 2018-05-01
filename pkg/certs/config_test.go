package certs

import (
	"reflect"
	"testing"
	"strings"
)

func TestNewConfig(t *testing.T) {
	t.Run("Test read config", func(t *testing.T) {
		pki := PKI{Organization: "PLEASE", OrganizationalUnit: "STANDUP", Locality: "Saint Joseph", Province: "Missouri", Country: "US"}
		expected := &Config{
			Template: "Please Standup",
			CA:       "will.therealpki.com",
			PKI:      pki,
		}

		actual, err := NewConfig([]string{"../../"})

		if err != nil {
			t.Errorf("err: %s", err)
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Actual: %s, \nExpected: %s", actual, expected)
		}
	})
	t.Run("Test when no conf provided", func(t *testing.T) {
		_, err := NewConfig([]string{"."})
		if !strings.Contains(err.Error(), "\"config\" Not Found in") {
			t.Errorf("err not raised correctly or nil: %s", err)
		}
	})
}
