package scanning

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/osrgroup/product-model-toolkit/pkg/client/scanner"
)

// Run executes a scan with a scanner tool for a given directory.
func Run(cfg *scanner.Config) {
	log.Printf("[Scanner] Selected : %v", cfg.Tool.String())
	log.Printf("[Scanner] Input directory: %v", cfg.InDir)
	log.Printf("[Scanner] Result directory: %v", cfg.ResultDir)

	err := execDockerCall(cfg)
	if err != nil {
		log.Printf("[Scanner] Error during Docker execution: %v", err.Error())
		return
	}
	files := findResultFiles(cfg)
	checkResults(cfg.ResultDir, files)
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

	expected := cfg.Tool.Results

	return findFiles(infos, expected)
}

func findFiles(infos []os.FileInfo, expected []string) []fileName {
	found := make([]fileName, 0)
	others := make([]fileName, 0)

	for _, f := range infos {
		if contains(expected, f.Name()) {
			found = append(found, fileName(f.Name()))
		} else {
			others = append(others, fileName(f.Name()))
		}
	}
	
	log.Printf("[Scanner] Found %v of %v expected result files: %v", len(found), len(expected), found)
	log.Printf("[Scanner] Found %v other files in result folder: %v", len(others), others)
	
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

func contains(slice []string, val string) bool {
	for _, e := range slice {
		if e == val {
			return true
		}
	}
	return false
}
