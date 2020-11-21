// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/osrgroup/product-model-toolkit/pkg/client/http/rest"
	"github.com/osrgroup/product-model-toolkit/pkg/client/scanning"
	"github.com/osrgroup/product-model-toolkit/pkg/plugin"
	"github.com/osrgroup/product-model-toolkit/pkg/version"
)

var gitCommit string

const serverBaseURL = "http://localhost:8081/api/v1"

type flags struct {
	scanner string
	inDir   string
}

func main() {
	flg, abort := checkFlags()
	if abort {
		os.Exit(0)
	}

	if plugin.NoPlugins() {
		log.Println("[Core] No scanner plugins available")
		os.Exit(0)
	}

	scn, found := plugin.FromStr(flg.scanner)
	if flg.scanner == "" || !found {
		scn = plugin.Default
		log.Printf("[Core] Scanner plugin not specified or not found, default scanner plugin %v is selected\n", scn.Name)
	}

	cfg := &plugin.Config{Plugin: scn, InDir: flg.inDir, ResultDir: "/tmp/pm/"}

	scanning.Run(
		cfg,
		rest.NewClient(serverBaseURL),
	)
}

func checkFlags() (flags, bool) {
	version := flag.Bool("v", false, "show version")

	lstScanner := flag.Bool("l", false, "list all available scanner")

	scanner := flag.String("s", "", "scanner to to use from list of available scanner")

	wd, _ := os.Getwd()
	inDir := flag.String("i", wd, "input dir to scan. Default is current working directory")

	flag.Parse()

	if *version {
		printVersion()
	}

	if *lstScanner {
		listScanner()
	}

	abortAfterFlags := *version || *lstScanner

	return flags{
			*scanner,
			*inDir,
		},
		abortAfterFlags
}

func printVersion() {
	fmt.Printf(
		"PMT Client\n----------\nVersion: %s\nGit commit: %s\n",
		version.Name(),
		gitCommit,
	)
}

func listScanner() {
	fmt.Println("Available license scanner:")
	for _, scn := range plugin.Available {
		fmt.Printf("----------\nName:    %s\nVersion: %s\nImage:   %s\n", scn.Name, scn.Version, scn.DockerImg)
	}
	fmt.Printf("----------\n")
}
