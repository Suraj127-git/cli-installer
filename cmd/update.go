package cmd

import (
	"fmt"
)

func RunUpdate(technology string) {
	switch technology {
	case "golang":
		updateGolang()
	case "laravel":
		updateLaravel()
	case "java":
		updateJava()
	case "nodejs":
		updateNodejs()
	case "reactjs":
		updateReactjs()
	default:
		fmt.Println("Unsupported technology")
	}
}

func updateGolang() {
	runCommand("go", "get", "-u", "golang.org/dl/go")
	fmt.Println("Golang updated successfully")
}

func updateLaravel() {
	runCommand("composer", "global", "update", "laravel/installer")
	fmt.Println("Laravel updated successfully")
}

func updateJava() {
	runCommand("sudo", "apt", "update")
	runCommand("sudo", "apt", "upgrade", "default-jdk", "-y")
	fmt.Println("Java updated successfully")
}

func updateNodejs() {
	runCommand("sudo", "npm", "install", "-g", "n")
	runCommand("sudo", "n", "latest")
	fmt.Println("Node.js updated successfully")
}

func updateReactjs() {
	runCommand("npm", "install", "-g", "create-react-app")
	fmt.Println("React.js updated successfully")
}
