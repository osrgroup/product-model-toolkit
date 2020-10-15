// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package scanning

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/osrgroup/product-model-toolkit/pkg/client/http/rest"
	"github.com/osrgroup/product-model-toolkit/pkg/client/scanner"
)

// Run executes a scan with a scanner tool for a given directory.
func Run(cfg *scanner.Config, c *rest.Client) {
	logServerVersion(c)
	log.Printf("[Scanner] Selected : %v", cfg.Tool.String())
	log.Printf("[Scanner] Input directory: %v", cfg.InDir)
	log.Printf("[Scanner] Result directory: %v", cfg.ResultDir)

	err := execDockerCall(cfg)
	if err != nil {
		log.Printf("[Scanner] Error during Docker execution: %v", err.Error())
		return
	}
	files := findResultFiles(cfg)

	switch cfg.Tool.Name {
	case scanner.Composer.Name:
		sendComposerResults(cfg.ResultDir, files, c)
	default:
		checkResults(cfg.ResultDir, files)
	}
}

func execDockerCall(cfg *scanner.Config) error {
	dockerCmd := execStr(cfg)
	log.Println("[Docker] ", dockerCmd)

	_, err := exec.Command("/bin/sh", "-c", dockerCmd).CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}

// execStr returns a command string for a Docker execution by the OS.
func execStr(cfg *scanner.Config) string {
	return fmt.Sprintf("docker run --rm -v %s:/input -v %s:/result %s %s", cfg.InDir, cfg.ResultDir, cfg.Tool.DockerImg, cfg.Tool.Cmd)
}

type fileName string

func findResultFiles(cfg *scanner.Config) []fileName {
	infos, err := ioutil.ReadDir(cfg.ResultDir)
	if err != nil {
		log.Printf("[Docker] Error during checking files: %v", err)
		return nil
	}

	var names = make([]fileName, 0)
	for _, e := range infos {
		names = append(names, fileName(e.Name()))
	}

	expected := cfg.Tool.Results
	return findFiles(names, expected)
}

// findFiles returns an array of file names that are both presented in names and expected args.
func findFiles(names []fileName, expected []string) []fileName {
	found := make([]fileName, 0)

	for _, n := range names {
		if contains(expected, string(n)) {
			found = append(found, n)
		}
	}

	log.Printf("[Scanner] Found %v of %v expected result files: %v", len(found), len(expected), found)

	return found
}

func checkResults(resDir string, files []fileName) {
	for _, f := range files {
		log.Printf("[Scanner] Content of result file %v", f)
		path := filepath.Join(resDir, string(f))
		data, _ := ioutil.ReadFile(path)
		log.Printf("\n%s", data)
	}
}

// contains return true if the given value is present in the given array.
func contains(slice []string, val string) bool {
	for _, e := range slice {
		if e == val {
			return true
		}
	}
	return false
}

func logServerVersion(c *rest.Client) {
	v, err := c.GetServerVersion()
	if err != nil {
		log.Printf("[REST-Client] Unable to read server version: %s", err)
		return
	}

	log.Printf("[REST-Client] Server version: %s", v)
}

func sendComposerResults(resDir string, files []fileName, c *rest.Client) {
	for _, f := range files {
		path := filepath.Join(resDir, string(f))
		resFile, err := os.Open(path)
		if err != nil {
			log.Printf("[Scanner] Error while reading Composer result files: %s", err)
			return
		}
		defer resFile.Close()

		loc, err := c.PostResult("/products/composer", resFile)
		if err != nil {
			log.Printf("[Scanner] Error while sending Composer results to server: %s", err)
			return
		}

		log.Printf("[Scanner] Successfully sent Composer results to server. Path to created resource: %s", loc)
	}
}
