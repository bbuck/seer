package config

import "testing"

func TestDefaultConfigAfterLoadingIsValid(t *testing.T) {
	LoadOrCreate()

	if Get().DB.Host != "localhost:7474" {
		t.Errorf("Expected [database] host = %q but found %q", "localhost:7474", Get().DB.Host)
	}
}
