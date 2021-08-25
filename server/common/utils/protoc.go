package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// copy from source: kratos/cmd/kratos/internal/base/mod.go
func ModuleVersion(path string) (string, error) {
	stdout := &bytes.Buffer{}
	fd := exec.Command("go", "mod", "graph")
	fd.Stdout = stdout
	fd.Stderr = stdout
	if err := fd.Run(); err != nil {
		return "", err
	}
	rd := bufio.NewReader(stdout)
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			return "", err
		}
		str := string(line)
		i := strings.Index(str, "@")
		if strings.Contains(str, path+"@") && i != -1 {
			return path + str[i:], nil
		}
	}
}

// copy from source: kratos/cmd/kratos/internal/base/mod.go
// KratosMod returns kratos mod.
func KratosMod() string {
	// go 1.15+ read from env GOMODCACHE
	cacheOut, _ := exec.Command("go", "env", "GOMODCACHE").Output()
	cachePath := strings.Trim(string(cacheOut), "\n")
	pathOut, _ := exec.Command("go", "env", "GOPATH").Output()
	gopath := strings.Trim(string(pathOut), "\n")
	if cachePath == "" {
		cachePath = filepath.Join(gopath, "pkg", "mod")
	}
	if path, err := ModuleVersion("github.com/go-kratos/kratos/v2"); err == nil {
		// $GOPATH/pkg/mod/github.com/go-kratos/kratos@v2
		return filepath.Join(cachePath, path)
	}
	// $GOPATH/src/github.com/go-kratos/kratos
	return filepath.Join(gopath, "src", "github.com", "go-kratos", "kratos")
}

func Generate() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if ext := filepath.Ext(path); ext != ".proto" {
			return nil
		}

		if strings.Contains(path, "vendor") || strings.Contains(path, "third_party") {
			return nil
		}

		execDir := filepath.Dir(dir)
		name := strings.ReplaceAll(path, execDir+string(filepath.Separator), "")
		input := []string{
			"--proto_path=.",
			//"--proto_path=" + filepath.Join(dir, "api"), // can import each .proto in api/, but itis error when protoc, why?
			"--proto_path=" + filepath.Join(os.Getenv("GOPATH"), "src"),
			"--proto_path=" + filepath.Join(KratosMod(), "api"),
			"--proto_path=" + filepath.Join(KratosMod(), "third_party"),
			"--go_out=paths=source_relative:.",
			"--go-grpc_out=paths=source_relative:.",
			"--go-http_out=paths=source_relative:.",
			"--go-errors_out=paths=source_relative:.",
			"--validate_out=lang=go,paths=source_relative:.",
			name,
		}

		fd := exec.Command("protoc", input...)
		fd.Stdout = os.Stdout
		fd.Stderr = os.Stderr
		fd.Dir = execDir
		if err := fd.Run(); err != nil {
			return err
		}
		fmt.Printf("proto: %s\n", name)
		return nil
	})
	if err != nil {
		return err
	}

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !strings.Contains(path, ".pb.go") {
			return nil
		}

		if strings.Contains(path, "vendor") {
			return nil
		}

		buf, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		content := string(buf)
		newContent := strings.Replace(content, ",omitempty", "", -1)
		err = ioutil.WriteFile(path, []byte(newContent), 0)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func GenSwagger() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if ext := filepath.Ext(path); ext != ".proto" {
			return nil
		}
		if strings.Contains(path, "vendor") {
			return nil
		}
		if !strings.Contains(path, "admin-server") &&
			!strings.Contains(path, "openai-server") {
			return nil
		}

		execDir := filepath.Dir(dir)
		name := strings.ReplaceAll(path, execDir+string(filepath.Separator), "")
		input := []string{
			"--proto_path=.",
			"--proto_path=" + filepath.Join(os.Getenv("GOPATH"), "src"),
			"--proto_path=" + filepath.Join(KratosMod(), "api"),
			"--proto_path=" + filepath.Join(KratosMod(), "third_party"),
			"--openapiv2_out",
			"./",
			"--openapiv2_opt",
			"logtostderr=true",
			name,
		}

		fd := exec.Command("protoc", input...)
		fd.Stdout = os.Stdout
		fd.Stderr = os.Stderr
		fd.Dir = execDir
		if err := fd.Run(); err != nil {
			return err
		}
		fmt.Printf("proto: %s\n", name)
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
