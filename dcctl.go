package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var possible = []string{"docker-compose.yaml", "docker-compose.yml", "compose.yaml", "compose.yml"}

func fileExists(filename string) bool {
	ok := true

	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if errors.Is(err, os.ErrNotExist) {
		ok = false
	} else {
		file.Close()
	}

	return ok
}

func find_compose_file() string {
	for _, e := range possible {
		if fileExists(e) {
			return e
		}
	}

	return ""
}

func usage(cmd string) {
	fmt.Println()

	if cmd != "" {
		fmt.Println("Unknown command [" + cmd + "]")
		fmt.Println()
	}

	fmt.Println("dcctl [--file my-compose.yaml] up|down")
	fmt.Println("    up - bring up the services in docker-compose.yaml")
	fmt.Println("    down - take down the services in docker-compose.yaml")

	os.Exit(1)
}

func cmd_up(compose_file string) {
	cmd := exec.Command("docker", "compose", "--file", compose_file, "up", "--detach")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

func cmd_down(compose_file string) {
	cmd := exec.Command("docker", "compose", "--file", compose_file, "down")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

func main() {
	var given_file = flag.String("file", "", "Use this docker compose file")
	flag.Parse()

	var compose_file string
	if *given_file == "" {
		compose_file = find_compose_file()
	} else if fileExists(*given_file) {
		compose_file = *given_file
	} else {
		compose_file = ""
	}

	if compose_file == "" {
		fmt.Println("Compose file not given or found")
		os.Exit(1)
	}

	s := fmt.Sprintf("Using %s", compose_file)
	fmt.Println(s)

	if len(flag.Args()) != 1 {
		fmt.Println("One command must be supplied")
		os.Exit(1)
	}

	c := strings.ToLower(flag.Arg(0))
	switch c {
	case "up":
		cmd_up(compose_file)
	case "down":
		cmd_down(compose_file)
	default:
		usage(c)
	}
}
