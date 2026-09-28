package flag_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lfhy/flag"
)

// Existing typed registrations must resolve the same sources as fluent ones.
func TestLegacyFlagSourcePriority(t *testing.T) {
	configPath := filepath.Join(t.TempDir(), "app.toml")
	if err := os.WriteFile(configPath, []byte("[title]\nfollow = true\n"), 0600); err != nil {
		t.Fatal(err)
	}

	const envName = "LFHY_FLAG_TEST_LEGACY_FOLLOW"
	t.Setenv(envName, "false")

	t.Run("environment overrides file", func(t *testing.T) {
		f := flag.NewFlagSet("test", flag.ContinueOnError)
		follow := f.BoolFull("follow", "title", "follow", envName, true, "")
		if err := f.Parse([]string{"-c", configPath}); err != nil {
			t.Fatal(err)
		}
		if *follow {
			t.Fatal("explicit false in environment should override true in file")
		}
		if got := f.NFlag(); got != 2 { // -c and follow
			t.Fatalf("NFlag() = %d, want 2", got)
		}
	})

	t.Run("command line overrides environment without changing it", func(t *testing.T) {
		f := flag.NewFlagSet("test", flag.ContinueOnError)
		follow := f.BoolFull("follow", "title", "follow", envName, false, "")
		if err := f.Parse([]string{"-c", configPath, "--follow=true"}); err != nil {
			t.Fatal(err)
		}
		if !*follow {
			t.Fatal("explicit true on command line should override false in environment")
		}
		if got := os.Getenv(envName); got != "false" {
			t.Fatalf("environment changed to %q, want false", got)
		}
	})
}

func TestInvalidLegacyEnvironmentValue(t *testing.T) {
	const envName = "LFHY_FLAG_TEST_INVALID_LEGACY_BOOL"
	t.Setenv(envName, "not-a-bool")

	t.Run("invalid environment is an error", func(t *testing.T) {
		f := flag.NewFlagSet("test", flag.ContinueOnError)
		f.SetOutput(new(strings.Builder))
		f.BoolEnv("follow", envName, false, "")
		if err := f.Parse(nil); err == nil {
			t.Fatal("invalid environment value should fail parsing")
		}
	})

	t.Run("explicit flag bypasses invalid environment", func(t *testing.T) {
		f := flag.NewFlagSet("test", flag.ContinueOnError)
		follow := f.BoolEnv("follow", envName, true, "")
		if err := f.Parse([]string{"-follow=false"}); err != nil {
			t.Fatal(err)
		}
		if *follow {
			t.Fatal("command-line false did not take precedence")
		}
	})
}

func TestConfigFlagNameCannotBeShadowedByAlias(t *testing.T) {
	f := flag.NewFlagSet("test", flag.ContinueOnError)
	f.SetOutput(new(strings.Builder))
	f.New("follow").Alias("c").Default(false)
	if err := f.Parse(nil); err == nil {
		t.Fatal("reserved config flag name used as alias should return an error")
	}
}
