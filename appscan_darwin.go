package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func ScanApps() {
	packages := InstalledPackages()

	/*  FIX: App name lookup in macOS

	    displayNames := AppDisplayNames(packages)

	    for _, displayName := range displayNames {
	       fmt.Println(displayName)
	    }
	*/

	for _, pkg := range packages {
		fmt.Println(pkg)
	}

}

func InstalledPackages() []string {

	// TODO: Add more sources of installed packages in macOS

	// https://stackoverflow.com/a/19255125

	command := exec.Command("pkgutil", "--pkgs")

	output, err := command.Output()
	if err != nil {
		fmt.Println("Failed to execute command:", err)
		return nil
	}

	packages := strings.Split(string(output), "\n")
	return packages
}

func AppDisplayNames(packageIDs []string) []string {
	var displayNames []string

	for _, packageID := range packageIDs {
		command := exec.Command("mdfind", "kMDItemCFBundleIdentifier=="+packageID)

		output, err := command.Output()
		if err != nil {
			fmt.Println("Failed to execute command:", err)
			continue
		}

		result := strings.TrimSpace(string(output))
		if result != "" {
			displayNames = append(displayNames, result)
		}
	}

	return displayNames
}
