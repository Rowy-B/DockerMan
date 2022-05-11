package main

import (
	"fmt"
	"os/exec"
)

func main() {

	_, Output := exec.Command("docker", "ps").Output()
	//error := Cmd.Run()
	fmt.Println(Output)
	//if error != nil {
	//	panic(error)
	//}

}
