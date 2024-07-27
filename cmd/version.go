package cmd

import (
	"fmt"
)

func RunVersion(technology string) {
	switch technology {
	case "golang":
		versionGolang()
	case "laravel":
		versionLaravel()
	case "java":
		versionJava()
	case "nodejs":
		versionNodejs()
	case "reactjs":
		versionReactjs()
	case "php":
		versionPHP()
	default:
		fmt.Println("Unsupported technology")
	}
}

func versionGolang() {
	runCommand("go", "version")
}

func versionLaravel() {
	runCommand("laravel", "--version")
}

func versionJava() {
	runCommand("java", "-version")
}

func versionNodejs() {
	runCommand("node", "--version")
}

func versionReactjs() {
	runCommand("npm", "list", "-g", "create-react-app")
}

func versionPHP() {
	runCommand("php", "-v")
}
