// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"fmt"
	"github.com/osrgroup/product-model-toolkit/pkg/db/postgraph"
	"github.com/osrgroup/product-model-toolkit/pkg/services/diff"
	"github.com/osrgroup/product-model-toolkit/pkg/services/importing"
	"github.com/osrgroup/product-model-toolkit/pkg/services/querying"
	"github.com/osrgroup/product-model-toolkit/pkg/services/version"
	"os"

	"github.com/osrgroup/product-model-toolkit/pkg/http/rest"
)

var gitCommit string

func main() {
	if checkFlags() {
		os.Exit(0)
	}

	//repo := new(memory.DB)
	//repo.AddSampleData()

	// Use Postgraphile as DB backend
	repo := postgraph.NewRepo("http://localhost:5433/graphql")

	r := rest.NewSrv(
		"127.0.0.1:8081",
		diff.NewService(repo),
		querying.NewService(repo),
		importing.NewService(repo),
	)
	go r.Start()
	defer r.Shutdown()
}

func checkFlags() bool {
	version := flag.Bool("v", false, "show version")

	flag.Parse()

	if *version {
		printVersion()
	}

	return flag.NFlag() > 0
}

func printVersion() {
	fmt.Println("PMT Server")
	fmt.Println("----------")
	fmt.Println("Version: " + version.Name())
	fmt.Println("Git commit: " + gitCommit)
	fmt.Println("----------")
}
