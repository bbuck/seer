package config

import (
	"os"
	"testing"
)

func TestLoadOrCreateWillCreateMissingFile(t *testing.T) {
	// Setup
	os.Remove(configFileName)

	_, err := os.Open(configFileName)
	if !os.IsNotExist(err) {
		t.Errorf("Expected %s not to exist, got %s instead.", configFileName, err)
	}

	LoadOrCreate()

	_, err = os.Open(configFileName)
	if err != nil {
		t.Errorf("Expected no error, instead received %s", err)
	}

	// Cleanup
	os.Remove(configFileName)
}

func TestDefaultConfigAfterLoadingIsValid(t *testing.T) {
	// Setup
	LoadOrCreate()

	if Get().DB.Host != "localhost:7474" {
		t.Errorf("Expected [database] host = %q but found %q", "localhost:7474", Get().DB.Host)
	}

	// Cleanup
	os.Remove(configFileName)
}
