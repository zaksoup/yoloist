package main

import "fmt"
import "os/exec"

func main() {
	fmt.Println("Yoloist - Soloist for risk takers")
	xcodeInstall := exec.Command("xcode-select", "--install")
	err := xcodeInstall.Run()
	if err != nil {
		panic(err)
	}
}
