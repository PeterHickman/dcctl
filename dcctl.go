package main

import (
	"flag"
	"fmt"
	toolbox "github.com/PeterHickman/toolbox"
	"os"
	"strings"
)

var possible = []string{"docker-compose.yaml", "docker-compose.yml", "compose.yaml", "compose.yml"}

func find_compose_file() string {
	for _, e := range possible {
		if toolbox.FileExists(e) {
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

func main() {
	var given_file = flag.String("file", "", "Use this docker compose file")
	flag.Parse()

	var compose_file string
	if *given_file == "" {
		compose_file = find_compose_file()
	} else if toolbox.FileExists(*given_file) {
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
		toolbox.Command("docker compose --file " + compose_file + " up --detach")
	case "down":
		toolbox.Command("docker compose --file " + compose_file + " down")
	default:
		usage(c)
	}
}
