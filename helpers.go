package main

import (
	"fmt"
	"os"
	"os/exec"
)

func startPlaying(videoId string, isAudioOnly bool) error {
	// Define options.
	var noVideo string
	if isAudioOnly {
		noVideo = "--no-video"
	}

	// Clear current session (delete socket file).
	_ = os.RemoveAll("/tmp/wotchy-mpv-socket")

	// Define URL.
	url := fmt.Sprintf("https://invidious.snopyta.org/watch?v=%s", videoId)

	// Define command line.
	cmd := exec.Command("mpv", noVideo, url, "--input-ipc-server=/tmp/wotchy-mpv-socket")

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

func togglePause() error {
	// Define command line.
	cmd := exec.Command("bash", "-c", "echo 'cycle pause' | socat - /tmp/wotchy-mpv-socket")

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

func stopPlaying() error {
	// Define command line.
	cmd := exec.Command("pkill", "mpv")

	// Start executing command.
	if errStart := cmd.Start(); errStart != nil {
		return errStart
	}

	// Clear current session (delete socket file).
	_ = os.RemoveAll("/tmp/wotchy-mpv-socket")

	// Wait for executing command.
	if errWait := cmd.Wait(); errWait != nil {
		return errWait
	}

	return nil
}
