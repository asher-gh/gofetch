package main

import "fmt"

func main() {
	packages := InstalledPackages()

	ScanApps() // scan for applications based on OS

	fmt.Printf("found %d packages installed\n", len(packages))
}
