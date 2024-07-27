package cmd

import (
	"fmt"
	"os/exec"
	"runtime"
)

func RunInstall(technology string) {
	switch technology {
	case "golang":
		installGolang()
	case "laravel":
		installLaravel()
	case "java":
		installJava()
	case "nodejs":
		installNodejs()
	case "reactjs":
		installReactjs()
	case "php":
		installPHP()
	default:
		fmt.Println("Unsupported technology")
	}
}

func runCommand(command string, args ...string) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing %s: %v\n", command, err)
		return
	}
	fmt.Println(string(output))
}

func installGolang() {
	var url string
	switch runtime.GOOS {
	case "linux":
		url = "https://golang.org/dl/go1.16.4.linux-amd64.tar.gz"
	case "darwin":
		url = "https://golang.org/dl/go1.16.4.darwin-amd64.pkg"
	case "windows":
		url = "https://golang.org/dl/go1.16.4.windows-amd64.msi"
	}
	runCommand("wget", url, "-O", "go.tar.gz")
	runCommand("tar", "-C", "/usr/local", "-xzf", "go.tar.gz")
	fmt.Println("Golang installed successfully")
}

func installLaravel() {
	runCommand("composer", "global", "require", "laravel/installer")
	fmt.Println("Laravel installed successfully")
}

func installJava() {
	runCommand("sudo", "apt", "install", "default-jdk", "-y")
	fmt.Println("Java installed successfully")
}

func installNodejs() {
	runCommand("curl", "-fsSL", "https://deb.nodesource.com/setup_14.x", "|", "sudo", "-E", "bash", "-")
	runCommand("sudo", "apt-get", "install", "-y", "nodejs")
	fmt.Println("Node.js installed successfully")
}

func installReactjs() {
	runCommand("npm", "install", "-g", "create-react-app")
	fmt.Println("React.js installed successfully")
}

func installPHP() {
	if runtime.GOOS == "darwin" {
		runCommand("brew", "install", "php")
		fmt.Println("PHP installed successfully")
	} else {
		fmt.Println("PHP installation for this OS is not supported yet")
	}
}
