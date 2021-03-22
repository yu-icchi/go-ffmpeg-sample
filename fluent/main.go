package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	fluentffmpeg "github.com/modfy/fluent-ffmpeg"
)

func main() {
	file, err := os.Create("output_copy.mp4")
	if err != nil {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	cmd := fluentffmpeg.NewCommand("")
	err = cmd.InputPath("./output_480p.mp4").
		OutputFormat("mp4").
		PipeOutput(file).
		OutputLogs(buf).
		OutputOptions("-movflags", "frag_keyframe+empty_moov").
		Run()
	fmt.Println(err)
	fmt.Println(buf.String())
}
