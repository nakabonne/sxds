package config

import (
	"os"
	"strconv"
	"testing"
)

const (
	testProduction = "true"
	testAdsMode    = "true"
	testXdsPort    = "8888"
	testCacherPort = "9999"
)

func TestNew(t *testing.T) {
	reset := setenvs(t, map[string]string{
		"SXDS_PRODUCTION":  testProduction,
		"SXDS_ADS_MODE":    testAdsMode,
		"SXDS_XDS_PORT":    testXdsPort,
		"SXDS_CACHER_PORT": testCacherPort,
	})
	defer reset()

	conf, err := New()
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	if got, want := conf.Production, testProduction; strconv.FormatBool(got) != want {
		t.Fatalf("got %v, want %v", got, want)
	}
	if got, want := conf.AdsMode, testAdsMode; strconv.FormatBool(got) != want {
		t.Fatalf("got %v, want %v", got, want)
	}
	if got, want := conf.Xds.Port, testXdsPort; strconv.Itoa(got) != want {
		t.Fatalf("got %v, want %v", got, want)
	}
	if got, want := conf.Cacher.Port, testCacherPort; strconv.Itoa(got) != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func setenvs(t *testing.T, kv map[string]string) func() {
	t.Helper()

	resetFs := make([]func(), 0, len(kv))
	for k, v := range kv {
		resetF := setenv(t, k, v)
		resetFs = append(resetFs, resetF)
	}

	return func() {
		for _, resetF := range resetFs {
			resetF()
		}
	}
}

func setenv(t *testing.T, k, v string) func() {
	t.Helper()

	prev := os.Getenv(k)
	if err := os.Setenv(k, v); err != nil {
		t.Fatal(err)
	}

	return func() {
		if prev == "" {
			os.Unsetenv(k)
		} else {
			if err := os.Setenv(k, prev); err != nil {
				t.Fatal(err)
			}
		}
	}
}
