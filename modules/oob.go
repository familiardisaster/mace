package modules

import (
	tools "/pkg"
	"regexp"
	"strings"
)

func OOB(target string) {
	targets := []string{}
	payloads, _ := tools.ReadLines("/payloads/oob.txt")
	regex := regexp.MustCompile(`\=(.)\&?`)
	matches := regex.FindAllStringSubmatch(target, -1)
	for _, match := range matches {
		for _, payload := range payloads {
			targets = append(targets, strings.ReplaceAll(target, match[1], payload))
		}
	}

	if len(targets) > 0 {
		tools.Fuzz(targets)
	}
}
