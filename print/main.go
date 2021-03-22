package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	ctx := context.Background()
	args := []string{"-i", "./BigBuckBunny.mp4", "-ss", "10", "-to", "30", "-s", "426x240", "./output.mp4"}
	cmd := exec.CommandContext(ctx, "ffmpeg", args...)
	stdout, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	go trace(stdout)
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

func trace(stdout io.ReadCloser) {
	r := bufio.NewReader(stdout)
	line, _, err := r.ReadLine()
	fmt.Printf("line: %s err %s\n", line, err)
}
