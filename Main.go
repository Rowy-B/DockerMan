package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	webbie()
}

func dockerPS() string {
	out, err := exec.Command("docker", "ps", "-a").Output()

	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}

func dockerImage() string {
	out, err := exec.Command("docker", "image", "ls").Output()

	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func dockerVolume() string {
	out, err := exec.Command("docker", "volume", "ls").Output()

	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func webbie() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "List of containers")
		fmt.Fprintln(w, dockerPS())
		fmt.Fprintln(w, "List of images")
		fmt.Fprintln(w, dockerImage())
		fmt.Fprintln(w, "List of volumes")
		fmt.Fprintln(w, dockerVolume())
	})

	port := ":5000"
	fmt.Println("Server is running on port" + port)

	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, nil))
}
