package flag_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lfhy/flag"
)

var _ func(string) *flag.ConfigFileBuilder = flag.ConfigFile

func writeConfigFileFixture(t *testing.T, dir, name, value string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte("[title]\nkey = \""+value+"\"\n"), 0600); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestConfigFileDefaultPathLoadsWithoutCommandLineFlag(t *testing.T) {
	path := writeConfigFileFixture(t, t.TempDir(), "default.toml", "from-default-file")
	f := flag.NewFlagSet("default file", flag.ContinueOnError)
	f.ConfigFile(path)
	value := f.New("value").Default("fallback").Config("title.key").String()

	if err := f.Parse(nil); err != nil {
		t.Fatal(err)
	}
	if *value != "from-default-file" {
		t.Fatalf("value = %q, want default file's value", *value)
	}
	if got := f.ConfigFlagName(); got != "c" {
		t.Fatalf("ConfigFlagName() = %q, want c", got)
	}
	if config := f.Lookup("c"); config == nil || config.Name != "c" || config.Value.String() != path {
		t.Fatalf("default config flag = %v, want canonical -c with value %q", config, path)
	}
}

func TestConfigFileNameReplacesShortFlag(t *testing.T) {
	dir := t.TempDir()
	defaultPath := writeConfigFileFixture(t, dir, "default.toml", "default")
	overridePath := writeConfigFileFixture(t, dir, "override.toml", "override")
	for _, tc := range []struct {
		name string
		args []string
		want string
		path string
	}{
		{name: "default path", want: "default", path: defaultPath},
		{name: "long flag overrides default", args: []string{"-config", overridePath}, want: "override", path: overridePath},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := flag.NewFlagSet("renamed file", flag.ContinueOnError)
			f.ConfigFile(defaultPath).Name("config")
			value := f.New("value").Default("fallback").Config("title.key").String()
			if err := f.Parse(tc.args); err != nil {
				t.Fatal(err)
			}
			if *value != tc.want {
				t.Fatalf("value = %q, want %q", *value, tc.want)
			}
			if got := f.ConfigFlagName(); got != "config" {
				t.Fatalf("ConfigFlagName() = %q, want config", got)
			}
			if config := f.Lookup("config"); config == nil || config.Name != "config" || config.Value.String() != tc.path {
				t.Fatalf("renamed config flag = %v, want -config with value %q", config, tc.path)
			}
			if short := f.Lookup("c"); short != nil {
				t.Fatalf("Name retained unexpected -c flag: %v", short)
			}
		})
	}

	t.Run("old short flag is rejected", func(t *testing.T) {
		f := flag.NewFlagSet("renamed file", flag.ContinueOnError)
		f.SetOutput(new(strings.Builder))
		f.ConfigFile(defaultPath).Name("config")
		if err := f.Parse([]string{"-c", overridePath}); err == nil {
			t.Fatal("Parse accepted -c after replacing it with -config")
		}
	})
}

func TestConfigFileAliasKeepsShortAndLongFlags(t *testing.T) {
	dir := t.TempDir()
	defaultPath := writeConfigFileFixture(t, dir, "default.toml", "default")
	overridePath := writeConfigFileFixture(t, dir, "override.toml", "override")
	for _, tc := range []struct {
		name string
		args []string
		want string
		path string
	}{
		{name: "default path", want: "default", path: defaultPath},
		{name: "short flag overrides default", args: []string{"-c", overridePath}, want: "override", path: overridePath},
		{name: "long alias overrides default", args: []string{"-config", overridePath}, want: "override", path: overridePath},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := flag.NewFlagSet("aliased file", flag.ContinueOnError)
			f.ConfigFile(defaultPath).Alias("config")
			value := f.New("value").Default("fallback").Config("title.key").String()
			if err := f.Parse(tc.args); err != nil {
				t.Fatal(err)
			}
			if *value != tc.want {
				t.Fatalf("value = %q, want %q", *value, tc.want)
			}
			short, long := f.Lookup("c"), f.Lookup("config")
			if short == nil || long != short || short.Name != "c" || short.Value.String() != tc.path {
				t.Fatalf("short flag = %v, long flag = %v; want one canonical -c with path %q", short, long, tc.path)
			}
		})
	}
}

func TestConfigFileDefaultPathSourcePrecedence(t *testing.T) {
	path := writeConfigFileFixture(t, t.TempDir(), "default.toml", "from-file")
	const envName = "LFHY_FLAG_CONFIG_FILE_VALUE"
	for _, tc := range []struct {
		name string
		env  string
		args []string
		want string
	}{
		{name: "file", want: "from-file"},
		{name: "environment beats file", env: "from-env", want: "from-env"},
		{name: "command line beats environment and file", env: "from-env", args: []string{"-value=from-cli"}, want: "from-cli"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv(envName, tc.env)
			if tc.env == "" {
				if err := os.Unsetenv(envName); err != nil {
					t.Fatal(err)
				}
			}
			f := flag.NewFlagSet("sources", flag.ContinueOnError)
			f.ConfigFile(path)
			value := f.New("value").Default("fallback").Config("title.key").Env(envName).String()
			if err := f.Parse(tc.args); err != nil {
				t.Fatal(err)
			}
			if *value != tc.want {
				t.Fatalf("value = %q, want %q", *value, tc.want)
			}
		})
	}
}

func TestConfigFileMissingDefaultDoesNotFailParse(t *testing.T) {
	path := filepath.Join(t.TempDir(), "absent.toml")
	f := flag.NewFlagSet("missing file", flag.ContinueOnError)
	f.ConfigFile(path)
	value := f.New("value").Default("fallback").Config("title.key").String()
	if err := f.Parse(nil); err != nil {
		t.Fatalf("Parse with missing default file: %v", err)
	}
	if *value != "fallback" {
		t.Fatalf("value = %q, want fallback", *value)
	}
	_ = f.GetConfig()
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("missing default config was created or stat failed unexpectedly: %v", err)
	}
}

func TestConfigFileMissingExplicitOverrideFailsParse(t *testing.T) {
	defaultPath := writeConfigFileFixture(t, t.TempDir(), "default.toml", "default")
	missingPath := filepath.Join(t.TempDir(), "absent.toml")
	for _, tc := range []struct {
		name  string
		alias bool
		args  []string
	}{
		{name: "short flag", args: []string{"-c", missingPath}},
		{name: "long alias", alias: true, args: []string{"-config", missingPath}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := flag.NewFlagSet("missing override", flag.ContinueOnError)
			f.SetOutput(new(strings.Builder))
			config := f.ConfigFile(defaultPath)
			if tc.alias {
				config.Alias("config")
			}
			value := f.New("value").Default("fallback").Config("title.key").String()
			if err := f.Parse(tc.args); err == nil {
				t.Fatalf("Parse(%q) accepted an explicitly missing config file", tc.args)
			}
			if *value == "default" {
				t.Fatal("failed explicit override silently loaded the default file")
			}
		})
	}
}

func TestConfigFileRejectsUnreadableOrMalformedFile(t *testing.T) {
	dir := t.TempDir()
	malformed := filepath.Join(dir, "malformed.toml")
	if err := os.WriteFile(malformed, []byte("[title\nkey = broken\n"), 0600); err != nil {
		t.Fatal(err)
	}
	valid := writeConfigFileFixture(t, dir, "valid.toml", "valid")
	for _, tc := range []struct {
		name        string
		defaultPath string
		args        []string
	}{
		{name: "malformed default", defaultPath: malformed},
		{name: "directory default", defaultPath: dir},
		{name: "malformed explicit override", defaultPath: valid, args: []string{"-c", malformed}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := flag.NewFlagSet("invalid config", flag.ContinueOnError)
			f.SetOutput(new(strings.Builder))
			f.ConfigFile(tc.defaultPath)
			f.New("value").Default("fallback").Config("title.key").String()
			if err := f.Parse(tc.args); err == nil {
				t.Fatal("Parse silently ignored an unreadable or malformed config file")
			}
		})
	}
}

func TestConfigFileExplicitEmptyPathSuppressesDefault(t *testing.T) {
	defaultPath := writeConfigFileFixture(t, t.TempDir(), "default.toml", "default")
	f := flag.NewFlagSet("empty override", flag.ContinueOnError)
	f.ConfigFile(defaultPath)
	value := f.New("value").Default("fallback").Config("title.key").String()
	if err := f.Parse([]string{"-c="}); err != nil {
		t.Fatal(err)
	}
	if *value != "fallback" {
		t.Fatalf("value = %q, want fallback when explicit -c= disables the default", *value)
	}
}

func TestConfigFileRespectsSetConfigFlagName(t *testing.T) {
	dir := t.TempDir()
	defaultPath := writeConfigFileFixture(t, dir, "default.toml", "default")
	overridePath := writeConfigFileFixture(t, dir, "override.toml", "override")
	for _, tc := range []struct {
		name string
		args []string
		want string
	}{
		{name: "default path", want: "default"},
		{name: "custom flag overrides default", args: []string{"-conf", overridePath}, want: "override"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := flag.NewFlagSet("custom name", flag.ContinueOnError)
			f.SetConfigFlagName("conf")
			f.ConfigFile(defaultPath)
			value := f.New("value").Default("fallback").Config("title.key").String()
			if err := f.Parse(tc.args); err != nil {
				t.Fatal(err)
			}
			if *value != tc.want || f.Lookup("conf") == nil || f.Lookup("c") != nil {
				t.Fatalf("value = %q, -conf = %v, -c = %v; want %q and only -conf", *value, f.Lookup("conf"), f.Lookup("c"), tc.want)
			}
		})
	}
}

func TestConfigFileReconfigurationBeforeParse(t *testing.T) {
	dir := t.TempDir()
	firstPath := writeConfigFileFixture(t, dir, "first.toml", "first")
	secondPath := writeConfigFileFixture(t, dir, "second.toml", "second")
	f := flag.NewFlagSet("reconfigured", flag.ContinueOnError)
	first := f.ConfigFile(firstPath).Alias("long")
	second := f.ConfigFile(secondPath)
	if first != second {
		t.Fatal("ConfigFile did not reuse its FlagSet builder")
	}
	first.Name("old")
	f.SetConfigFlagName("new") // Last naming operation wins.
	value := f.New("value").Default("fallback").Config("title.key").String()
	if err := f.Parse(nil); err != nil {
		t.Fatal(err)
	}
	if *value != "second" || f.Lookup("new") == nil || f.Lookup("long") != f.Lookup("new") || f.Lookup("old") != nil {
		t.Fatalf("value = %q; aliases old=%v new=%v long=%v", *value, f.Lookup("old"), f.Lookup("new"), f.Lookup("long"))
	}
}

func TestConfigFileAliasRejectsConflictingName(t *testing.T) {
	path := writeConfigFileFixture(t, t.TempDir(), "default.toml", "default")
	f := flag.NewFlagSet("alias conflict", flag.ContinueOnError)
	f.SetOutput(new(strings.Builder))
	f.New("occupied").String()
	f.ConfigFile(path).Alias("occupied")
	if err := f.Parse(nil); err == nil {
		t.Fatal("Parse accepted a config alias already used as another canonical flag")
	}
}
