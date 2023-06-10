package main

import (
	"fmt"
	"os"

	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	flag "github.com/spf13/pflag"
	"github.com/zerodha/logf"
)

func initConfig() {

	f := flag.NewFlagSet("config", flag.ContinueOnError)
	f.Usage = func() {
		fmt.Println(f.FlagUsages())
		os.Exit(0)
	}

	f.StringSlice("config", []string{"config.toml"}, "Path to one or more TOML config files")
	f.Bool("version", false, "show Build version")
	f.Parse(os.Args[1:])

	if ok, _ := f.GetBool("version"); ok {
		fmt.Printf("Build version: %s\n", buildString)
		os.Exit(0)
	}

	cFiles, _ := f.GetStringSlice("config")
	for _, f := range cFiles {
		fmt.Printf("Reading config: %s\n", f)
		if err := ko.Load(file.Provider(f), toml.Parser()); err != nil {
			fmt.Printf("Error reading config: %v\n", err)
		}
	}
}

func initLogger(debug bool) logf.Logger {
	opts := logf.Opts{EnableCaller: true}
	if debug {
		opts.Level = logf.DebugLevel
	}

	return logf.New(opts)
}
