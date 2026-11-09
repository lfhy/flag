package flag_test

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lfhy/flag"
)

var _ func(...string) *flag.ConfigFileBuilder = flag.ConfigFiles

func layerFixture(t *testing.T, directory, name, content string) string {
	t.Helper()
	path := filepath.Join(directory, name)
	if err := os.WriteFile(path, []byte(content), 0600); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestConfigFilesMergeDefaultsExplicitFileAndCLI(t *testing.T) {
	dir := t.TempDir()
	base := layerFixture(t, dir, "config.toml", "[server]\nport=7000\nhost='base'\n[extra]\nbase=true\n")
	shared := layerFixture(t, dir, "gateway.toml", "[server]\nport=7420\nhost='shared'\n[extra]\nshared=true\n")
	instance := layerFixture(t, dir, "gateway-7421.toml", "[server]\nport=7421\n[extra]\ninstance=true\n")
	override := layerFixture(t, dir, "selected.toml", "[server]\nhost='explicit'\n[extra]\nselected=true\n")
	for _, cli := range []bool{false, true} {
		t.Run(map[bool]string{false: "files", true: "CLI"}[cli], func(t *testing.T) {
			fs := flag.NewFlagSet("layers", flag.ContinueOnError)
			fs.ConfigFiles(base, shared, instance).Name("config").Alias("c")
			port := fs.New("port").Alias("p").Default(1).Config("server.port").Int()
			host := fs.New("host").Default("default").Config("server.host").String()
			args := []string{"-c", override}
			if cli {
				args = append(args, "-p", "8000", "-host", "cli")
			}
			if err := fs.Parse(args); err != nil {
				t.Fatal(err)
			}
			if cli {
				if *port != 8000 || *host != "cli" {
					t.Fatalf("CLI lost precedence: port=%d host=%q", *port, *host)
				}
			} else if *port != 7421 || *host != "explicit" {
				t.Fatalf("file precedence: port=%d host=%q", *port, *host)
			}
			cfg := fs.GetConfig()
			for _, key := range []string{"extra.base", "extra.shared", "extra.instance", "extra.selected"} {
				if !cfg.GetBool(key) {
					t.Errorf("merged config lost %s", key)
				}
			}
			if cfg.GetConfigPath() != override || cfg.GetInt("server.port") != 7421 {
				t.Fatalf("effective file config: path=%q port=%d", cfg.GetConfigPath(), cfg.GetInt("server.port"))
			}
		})
	}
}

func TestConfigFilesEnvironmentOverridesMergedFiles(t *testing.T) {
	path := layerFixture(t, t.TempDir(), "base.toml", "[server]\nport=7420\n")
	t.Setenv("FLAG_LAYERS_PORT", "9000")
	fs := flag.NewFlagSet("environment", flag.ContinueOnError)
	fs.ConfigFiles(path)
	port := fs.New("port").Default(1).Config("server.port").Env("FLAG_LAYERS_PORT").Int()
	if err := fs.Parse(nil); err != nil || *port != 9000 {
		t.Fatalf("environment precedence: port=%d err=%v", *port, err)
	}
}

func TestConfigFilesArraysReplaceAndPreserveElements(t *testing.T) {
	dir := t.TempDir()
	base := layerFixture(t, dir, "base.toml", "[gateway]\npeers=['old-a','old-b']\nports=[1,2]\nempty=['old']\n")
	overlay := layerFixture(t, dir, "instance.toml", "[gateway]\npeers=['http://node:7420/path,a','http://other:7420']\nports=[7,8]\nempty=[]\n")
	fs := flag.NewFlagSet("arrays", flag.ContinueOnError)
	fs.ConfigFiles(base, overlay)
	peers := fs.StringsConfig("peers", "gateway", "peers", "builtin", "peer list")
	ports := fs.IntsConfig("ports", "gateway", "ports", "99", "ports")
	empty := fs.StringsConfig("empty", "gateway", "empty", "builtin", "empty list")
	var reflected []int
	fs.New("reflected").Default([]int{99}).Config("gateway.ports").Var(&reflected)
	if err := fs.Parse(nil); err != nil {
		t.Fatal(err)
	}
	if want := []string{"http://node:7420/path,a", "http://other:7420"}; !reflect.DeepEqual(*peers, want) {
		t.Fatalf("array appended or elements split: got=%q want=%q", *peers, want)
	}
	if !reflect.DeepEqual(*ports, []int{7, 8}) || !reflect.DeepEqual(reflected, []int{7, 8}) || len(*empty) != 0 {
		t.Fatalf("array replacement: ports=%v reflected=%v empty=%q", *ports, reflected, *empty)
	}
	if got := fs.GetConfig().GetStringSlice("gateway.peers"); !reflect.DeepEqual(got, *peers) {
		t.Fatalf("GetConfig array differs from flag: %q", got)
	}
}

func TestConfigFilesMissingOptionalAndStrictReadFailures(t *testing.T) {
	dir := t.TempDir()
	missing := filepath.Join(dir, "missing.toml")
	valid := layerFixture(t, dir, "valid.toml", "[server]\nport=7420\n")
	malformed := layerFixture(t, dir, "malformed.toml", "[server\n")
	for _, tc := range []struct {
		name  string
		files []string
		args  []string
		fail  bool
		want  int
	}{
		{name: "optional missing", files: []string{missing, valid}, want: 7420},
		{name: "all optional missing", files: []string{missing}, want: 1},
		{name: "explicit missing", files: []string{valid}, args: []string{"-c", missing}, fail: true},
		{name: "malformed default", files: []string{malformed, valid}, fail: true},
		{name: "directory", files: []string{dir}, fail: true},
		{name: "empty disables all", files: []string{malformed, valid}, args: []string{"-c="}, want: 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			fs := flag.NewFlagSet("file errors", flag.ContinueOnError)
			fs.SetOutput(new(strings.Builder))
			fs.ConfigFiles(tc.files...)
			port := fs.New("port").Default(1).Config("server.port").Int()
			err := fs.Parse(tc.args)
			if (err != nil) != tc.fail {
				t.Fatalf("Parse error=%v wantFailure=%v", err, tc.fail)
			}
			if !tc.fail && *port != tc.want {
				t.Errorf("port=%d want=%d", *port, tc.want)
			}
		})
	}
	if _, err := os.Stat(missing); !os.IsNotExist(err) {
		t.Fatalf("configuration read created a missing file: %v", err)
	}
}

func TestConfigFilesCopiesPathsAndSupportsEarlyGetConfig(t *testing.T) {
	dir := t.TempDir()
	a := layerFixture(t, dir, "a.toml", "[server]\nport=7420\n")
	b := layerFixture(t, dir, "b.toml", "[server]\nport=7421\n")
	paths := []string{a, b}
	fs := flag.NewFlagSet("early config", flag.ContinueOnError)
	fs.ConfigFiles(paths...).Alias("config")
	paths[1] = filepath.Join(dir, "caller-mutated.toml")
	if got := fs.GetConfig().GetInt("server.port"); got != 7421 {
		t.Fatalf("early merged config port=%d", got)
	}
	port := fs.New("port").Default(1).Config("server.port").Int()
	if err := fs.Parse(nil); err != nil || *port != 7421 {
		t.Fatalf("Parse after GetConfig: port=%d err=%v", *port, err)
	}
}

func TestConfigFilesHigherLayerReplacesDifferentValueKinds(t *testing.T) {
	directory := t.TempDir()
	a := layerFixture(t, directory, "a.toml", "options={old=true}\nreverse='old'\narray={old=true}\n")
	b := layerFixture(t, directory, "b.toml", "options='new'\nreverse={new=true}\narray=['new']\n")
	fs := flag.NewFlagSet("kind changes", flag.ContinueOnError)
	fs.ConfigFiles(a, b)
	if err := fs.Parse(nil); err != nil {
		t.Fatal(err)
	}
	cfg := fs.GetConfig()
	if cfg.GetString("options") != "new" || cfg.IsSet("options.old") || !cfg.GetBool("reverse.new") || !reflect.DeepEqual(cfg.GetStringSlice("array"), []string{"new"}) {
		t.Fatalf("higher file lost priority on type change: %v", cfg.AllSettings())
	}
}
