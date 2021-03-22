package main

import (
	"context"
	"io"
	"log"
	"os"
	"os/exec"

	"golang.org/x/sync/errgroup"
)

func main() {
	src, err := os.Open("./BigBuckBunny.mp4")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	dst1, err := os.Create("./output_480p.mp4")
	if err != nil {
		log.Fatal(err)
	}
	defer dst1.Close()
	dst2, err := os.Create("./output_240p.mp4")
	if err != nil {
		log.Fatal(err)
	}
	defer dst2.Close()
	src2, srcw2 := io.Pipe()
	src1 := io.TeeReader(src, srcw2)
	ctx := context.Background()
	wg, ectx := errgroup.WithContext(ctx)
	wg.Go(func() error {
		return run(ectx, "720x480", src1, dst1)
	})
	wg.Go(func() error {
		return run(ectx, "426x240", src2, dst2)
	})
	if err := wg.Wait(); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, size string, r io.Reader, w io.Writer) error {
	args := []string{"-i", "pipe:0", "-ss", "10", "-to", "30", "-s", size, "-f", "mp4", "-movflags", "frag_keyframe+empty_moov", "pipe:1"}
	cmd := exec.CommandContext(ctx, "ffmpeg", args...)
	cmd.Stdin = r
	cmd.Stdout = w
	if err := cmd.Start(); err != nil {
		return err
	}
	return cmd.Wait()
}
