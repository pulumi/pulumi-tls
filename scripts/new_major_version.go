// go run scripts/new_major_version.go -ver v5 -name tls goes part of the way to udpate all Go source references to a
// new major version, but this automation is not complete. In particular here may be dangling references in:
//
// - go.mod
// - README.md
// - Makefile
// - goreleaser.*yml
package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	nameFlag := flag.String("name", "", `Short provider name such as "tls"`)
	verFlag := flag.String("ver", "", `Major version to upgrade to, such as "v5"`)
	flag.Parse()
	if verFlag == nil || *verFlag == "" {
		log.Fatal("-ver is required")
		return
	}
	if nameFlag == nil || *nameFlag == "" {
		log.Fatal("-name is required")
		return
	}
	ver := *verFlag
	name := *nameFlag
	prov := currentProvider(name)
	fixupGoFiles(prov, ver)
}

func fixupGoFiles(prov provider, ver string) {
	for _, file := range findGoFiles() {
		fixupGoFile(prov, ver, file)
	}
}

func fixupGoFile(prov provider, ver, file string) {
	replace := fmt.Sprintf("%q -> %q",
		fmt.Sprintf("%s/%s", prov.prefix, prov.major),
		fmt.Sprintf("%s/%s", prov.prefix, ver))
	gofmtr(replace, file)
	replaceVersionPkg := fmt.Sprintf("%q -> %q",
		fmt.Sprintf("%s/%s/pkg/version", prov.prefix, prov.major),
		fmt.Sprintf("%s/%s/pkg/version", prov.prefix, ver))
	gofmtr(replaceVersionPkg, file)

	sdk := strings.TrimSuffix(prov.prefix, "/provider") + "/sdk"
	replaceGoSdkPkg := fmt.Sprintf("%q -> %q",
		fmt.Sprintf("%s/%s/go/%s", sdk, prov.major, prov.name),
		fmt.Sprintf("%s/%s/go/%s", sdk, ver, prov.name))
	gofmtr(replaceGoSdkPkg, file)
}

func gofmtr(replace, file string) {
	cmd := exec.Command("gofmt", "-r", replace, "-w", file)
	fmt.Printf("gofmt -r %s -w %s\n", replace, file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func findGoFiles() []string {
	cmd := exec.Command("git", "ls-files", "-z", "**/*.go")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	fs := strings.Split(stdout.String(), string(rune(0)))
	result := []string{}
	for _, f := range fs {
		if f == "" {
			continue
		}
		result = append(result, f)
	}
	return result
}

type provider struct {
	name   string
	prefix string
	major  string
}

func currentProvider(name string) provider {
	gomod, _ := os.ReadFile("provider/go.mod")
	firstline := strings.Trim(strings.Split(string(gomod), "\n")[0], "\r\n")
	frags := strings.Split(firstline, "/")
	prefix := strings.TrimPrefix(strings.Join(frags[0:len(frags)-1], "/"), "module ")
	ver := frags[len(frags)-1]
	return provider{name, prefix, ver}
}
