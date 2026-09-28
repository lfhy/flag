package flag_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lfhy/flag"
)

func TestFlagBuilderAliasesShareCanonicalFlag(t *testing.T) {
	for _, tc := range []struct {
		name string
		args []string
	}{
		{name: "single dash with equals", args: []string{"-p=first"}},
		{name: "double dash with separate value", args: []string{"--prim", "second"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := flag.NewFlagSet("aliases", flag.ContinueOnError)
			value := f.New("primary").Alias("p").Alias("prim").Usage("primary option").String()
			if err := f.Parse(tc.args); err != nil {
				t.Fatal(err)
			}
			if want := map[string]string{"single dash with equals": "first", "double dash with separate value": "second"}[tc.name]; *value != want {
				t.Fatalf("bound value = %q, want %q", *value, want)
			}
			canonical := f.Lookup("primary")
			if canonical == nil || f.Lookup("p") != canonical || f.Lookup("prim") != canonical {
				t.Fatal("aliases did not resolve to the same canonical flag")
			}
			if canonical.Name != "primary" || f.NFlag() != 1 {
				t.Fatalf("canonical name = %q, NFlag = %d", canonical.Name, f.NFlag())
			}
			var visited int
			f.VisitAll(func(*flag.Flag) { visited++ }) // Includes the implicit -c flag.
			if visited != 2 {
				t.Fatalf("VisitAll count = %d, want 2", visited)
			}
			var help strings.Builder
			f.SetOutput(&help)
			f.PrintDefaults()
			if !strings.Contains(help.String(), "-primary, -p, -prim") || strings.Count(help.String(), "primary option") != 1 {
				t.Fatalf("help does not show one canonical flag with both aliases: %q", help.String())
			}
		})
	}
}

func TestFlagBuilderSourcePrecedence(t *testing.T) {
	path := filepath.Join(t.TempDir(), "app.toml")
	if err := os.WriteFile(path, []byte("label = \"from-config\"\n[server]\nfollow = true\n"), 0600); err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		name       string
		config     bool
		env        bool
		cli        bool
		envLabel   string
		envFollow  string
		wantLabel  string
		wantFollow bool
	}{
		{name: "default", wantLabel: "fallback", wantFollow: false},
		{name: "config root and section", config: true, wantLabel: "from-config", wantFollow: true},
		{name: "environment empty and false override config", config: true, env: true, envLabel: "", envFollow: "false", wantLabel: "", wantFollow: false},
		{name: "command line empty and false override environment", config: true, env: true, cli: true, envLabel: "from-env", envFollow: "true", wantLabel: "", wantFollow: false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			const labelEnv = "LFHY_FLAG_BUILDER_LABEL"
			const followEnv = "LFHY_FLAG_BUILDER_FOLLOW"
			// Ensure even an inherited environment cannot affect default/config cases.
			t.Setenv(labelEnv, "")
			t.Setenv(followEnv, "")
			if tc.env {
				t.Setenv(labelEnv, tc.envLabel)
				t.Setenv(followEnv, tc.envFollow)
			} else {
				if err := os.Unsetenv(labelEnv); err != nil {
					t.Fatal(err)
				}
				if err := os.Unsetenv(followEnv); err != nil {
					t.Fatal(err)
				}
			}
			f := flag.NewFlagSet("sources", flag.ContinueOnError)
			label := f.New("label").Default("fallback").Config("label").Env(labelEnv).String()
			follow := f.New("follow").Default(false).Config("server.follow").Env(followEnv).Bool()
			var args []string
			if tc.config {
				args = append(args, "-c", path)
			}
			if tc.cli {
				args = append(args, "-label=", "-follow=false")
			}
			if err := f.Parse(args); err != nil {
				t.Fatal(err)
			}
			if *label != tc.wantLabel || *follow != tc.wantFollow {
				t.Fatalf("label = %q, follow = %t; want %q, %t", *label, *follow, tc.wantLabel, tc.wantFollow)
			}
		})
	}
}

func TestFlagBuilderTypedPointersAndVarBinding(t *testing.T) {
	f := flag.NewFlagSet("typed", flag.ContinueOnError)
	b := f.New("enabled").Default(false).Bool()
	i := f.New("count").Default(1).Int()
	i64 := f.New("large").Default(int64(2)).Int64()
	s := f.New("title").Default("before").String()
	var bound string
	f.New("bound").Default("original").Var(&bound)
	if *b || *i != 1 || *i64 != 2 || *s != "before" || bound != "original" {
		t.Fatalf("defaults: bool=%t int=%d int64=%d string=%q bound=%q", *b, *i, *i64, *s, bound)
	}
	if err := f.Parse([]string{"-enabled", "-count=7", "-large=4294967296", "-title=after", "-bound=updated"}); err != nil {
		t.Fatal(err)
	}
	if !*b || *i != 7 || *i64 != 4294967296 || *s != "after" || bound != "updated" {
		t.Fatalf("parsed bindings: bool=%t int=%d int64=%d string=%q bound=%q", *b, *i, *i64, *s, bound)
	}
}

func TestFlagBuilderPendingValuesAndConversions(t *testing.T) {
	f := flag.NewFlagSet("pending", flag.ContinueOnError)
	follow := f.New("follow").Default(false)
	count := f.New("count").Default(3)
	text := f.New("text").Default("before")
	if f.Lookup("follow") != nil || f.Lookup("count") != nil {
		t.Fatal("non-terminal flags registered before Parse")
	}
	if err := f.Parse([]string{"-follow", "-count=42", "-text=after"}); err != nil {
		t.Fatal(err)
	}
	if !follow.Value().Bool() || follow.Value().String() != "true" {
		t.Fatalf("bare boolean value = %q", follow.Value().String())
	}
	if count.Value().Int() != 42 || count.Value().Int64() != 42 || count.Value().String() != "42" {
		t.Fatalf("numeric value = %q", count.Value().String())
	}
	if text.Value().String() != "after" {
		t.Fatalf("text value = %q", text.Value().String())
	}
	if err := f.Set("count", "not-a-number"); err != nil {
		t.Fatal(err) // Pending values retain text; conversion is deferred to Value().Int().
	}
	assertBuilderPanic(t, func() { count.Value().Int() })
	assertBuilderPanic(t, func() { text.Value().Bool() })
}

func TestFlagBuilderPendingConfigAndEnv(t *testing.T) {
	path := filepath.Join(t.TempDir(), "app.toml")
	if err := os.WriteFile(path, []byte("[title]\nkey = true\ncount = 21\n"), 0600); err != nil {
		t.Fatal(err)
	}
	const envName = "LFHY_FLAG_TEST_PENDING_FOLLOW"
	t.Setenv(envName, "false")

	f := flag.NewFlagSet("pending sources", flag.ContinueOnError)
	follow := f.New("follow").Alias("f").Default(false).Config("title.key").Env(envName)
	count := f.New("count").Default("0").Config("title.count")
	if err := f.Parse([]string{"-c", path, "-count=42"}); err != nil {
		t.Fatal(err)
	}
	if follow.Value().Bool() || follow.Value().String() != "false" {
		t.Fatalf("environment false should override config true, got %q", follow.Value().String())
	}
	if count.Value().Int() != 42 {
		t.Fatalf("command line should override config 21, got %d", count.Value().Int())
	}
	if f.Lookup("f") != f.Lookup("follow") {
		t.Fatal("pending flag alias did not resolve to its canonical flag")
	}
}

func TestFlagBuilderRejectsAliasCollision(t *testing.T) {
	for _, tc := range []struct {
		name string
		add  func(*flag.FlagSet)
	}{
		{name: "alias duplicates another canonical name", add: func(f *flag.FlagSet) {
			f.New("taken").String()
			f.New("other").Alias("taken").String()
		}},
		{name: "canonical name duplicates alias", add: func(f *flag.FlagSet) {
			f.New("first").Alias("taken").String()
			f.New("taken").String()
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := flag.NewFlagSet("collisions", flag.PanicOnError)
			f.SetOutput(new(strings.Builder))
			assertBuilderPanic(t, func() { tc.add(f) })
		})
	}
}

func TestFlagBuilderPendingAliasCollisionReturnsParseError(t *testing.T) {
	f := flag.NewFlagSet("pending collisions", flag.ContinueOnError)
	f.SetOutput(new(strings.Builder))
	f.New("first").Alias("shared")
	f.New("second").Alias("shared")
	if err := f.Parse(nil); err == nil {
		t.Fatal("Parse succeeded after rejecting a pending alias collision")
	}
	if got := f.Lookup("second"); got != nil {
		t.Fatalf("rejected flag was partially registered: %v", got)
	}
}

func TestFlagBuilderInvalidPendingConfigReturnsParseError(t *testing.T) {
	f := flag.NewFlagSet("pending config", flag.ContinueOnError)
	f.SetOutput(new(strings.Builder))
	f.New("follow").Config("title.")
	if err := f.Parse(nil); err == nil {
		t.Fatal("Parse succeeded after rejecting an empty configuration key")
	}
}

func TestFlagVarAliasesAndDistinctFlagSets(t *testing.T) {
	first := flag.NewFlagSet("first", flag.ContinueOnError)
	second := flag.NewFlagSet("second", flag.ContinueOnError)
	var direct string
	first.Var(&flag.FlagVar{Name: "value", Aliases: []string{"v"}, DefaultValue: "one", Value: &direct})
	other := second.New("value").Alias("v").Default("two").String()
	if err := first.Parse([]string{"-v=first"}); err != nil {
		t.Fatal(err)
	}
	if err := second.Parse([]string{"--v", "second"}); err != nil {
		t.Fatal(err)
	}
	if direct != "first" || *other != "second" {
		t.Fatalf("independent values: first=%q second=%q", direct, *other)
	}
	if first.Lookup("v") != first.Lookup("value") || second.Lookup("v") != second.Lookup("value") || first.Lookup("v") == second.Lookup("v") {
		t.Fatal("alias resolution leaked between FlagSets")
	}
}

func assertBuilderPanic(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Fatal("expected a panic")
		}
	}()
	fn()
}
