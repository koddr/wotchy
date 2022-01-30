package main

import (
	"fmt"
	"os"
	"os/exec"
)

var PID = make(chan int, 1)

func buildVideoLink(videoId string) string {
	host := "https://invidious.snopyta.org"

	if os.Getenv("INVIDOUS_HOST") != "" {
		host = os.Getenv("INVIDOUS_HOST")
	}

	return fmt.Sprintf("%s/watch?v=%s", host, videoId)
}

func startPlaying(m model, isAudioOnly bool) error {
	// Define options.
	var noVideo string
	if isAudioOnly {
		noVideo = "--no-video"
	}

	// Define command line.
	cmd := exec.Command("mpv", noVideo, buildVideoLink(m.selected))

	// Start executing command.
	if errStart := cmd.Start(); errStart != nil {
		return errStart
	}

	// Send mpv pid to channel for output in state.
	PID <- cmd.Process.Pid

	// Wait for executing command.
	if errWait := cmd.Wait(); errWait != nil {
		return errWait
	}

	return nil
}

func stopPlaying() error {
	// Define command line.
	cmd := exec.Command("pkill", "mpv")

	// Start executing command.
	if errStart := cmd.Start(); errStart != nil {
		return errStart
	}

	// Wait for executing command.
	if errWait := cmd.Wait(); errWait != nil {
		return errWait
	}

	return nil
}
