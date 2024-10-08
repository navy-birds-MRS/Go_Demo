package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func main() {
	cmd := exec.Command("netsh", "wlan", "show", "profile")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}

	re := regexp.MustCompile(`All User Profile\s*:\s*(.+?)\n`)
	matches := re.FindAllStringSubmatch(string(output), -1)
	for _, match := range matches {
		if len(match) < 2 {
			continue
		}
		profileName := strings.TrimSpace(match[1])

		// Command to show profile details
		cmd := exec.Command("netsh", "wlan", "show", "profile", "name="+profileName, "key=clear")

		output, err := cmd.CombinedOutput()
		if err != nil {

			continue
		}

		// Regular expressions to extract SSID and password
		passwordRe := regexp.MustCompile(`Key Content\s*:\s*(.+)`)
		passwordMatch := passwordRe.FindStringSubmatch(string(output))

		if len(passwordMatch) > 1 {

			if len(passwordMatch) > 1 {
				fmt.Println("SSID:", profileName, "|", "Password:", passwordMatch[1])
			}
		} else {
			fmt.Println("SSID:", profileName, "|", "Password:", "Open Network")
		}
		time.Sleep(250 * time.Millisecond)
	}

}
