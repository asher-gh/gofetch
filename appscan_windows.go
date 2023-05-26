/*
* Only compiles on windows
 */

package main

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
)

func ScanApps() {
	packages := InstalledPackages()
	displayNames := DisplayName(packages)

	for _, displayName := range displayNames {
		fmt.Println(displayName)
	}

}

// Retrieves a list of installed package IDs from the Windows Registry.
func InstalledPackages() []string {
	const uninstallKey = `SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall`

	key, err := registry.OpenKey(registry.LOCAL_MACHINE, uninstallKey, registry.ENUMERATE_SUB_KEYS|registry.QUERY_VALUE)
	if err != nil {
		fmt.Println("Failed to open registry key:", err)
		return nil
	}
	defer key.Close()

	subKeyNames, err := key.ReadSubKeyNames(-1)
	if err != nil {
		fmt.Println("Failed to read subkey names:", err)
		return nil
	}

	return subKeyNames
}

// Find the corresponding display name from the registry
func DisplayName(packageIDs []string) []string {
	var displayNames []string

	for _, packageID := range packageIDs {
		key, err := registry.OpenKey(registry.LOCAL_MACHINE, fmt.Sprintf(`SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\%s`, packageID), registry.QUERY_VALUE)
		if err != nil {
			fmt.Println("Failed to open subkey:", err)
			continue
		}

		displayName, _, err := key.GetStringValue("DisplayName")
		if err == nil && displayName != "" {
			displayNames = append(displayNames, displayName)
		}

		key.Close()
	}

	return displayNames
}
