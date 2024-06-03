package ffprobe

import (
	"fmt"
	"log"
	"testing"
)

func TestFile(t *testing.T) {
	probe, err := Probe("/Users/matthewdavis/workspace/media/bunny.mp4")
	if err != nil {
		t.Fatalf("failed to probe: %v", err)
	}
	fmt.Println(probe)

	if probe.Format.FormatName != "mov,mp4,m4a,3gp,3g2,mj2" {
		t.Errorf("unexpected format name: %s", probe.Format.FormatName)
	}

	if probe.Format.Duration != 10 {
		t.Errorf("unexpected duration: %f", probe.Format.Duration)
	}

	log.Printf("Duration: %f", probe.Format.Duration)
	if probe.Format.Tags.Title != "Big Buck Bunny, Sunflower version" {
		t.Errorf("unexpected title: %s", probe.Format.Tags.Title)
	}

	videos, err := probe.GetStreamType("video")
	if err != nil {
		t.Fatalf("failed to get video stream: %v", err)
	}

	if len(videos) != 1 {
		t.Fatalf("unexpected number of video streams: %d", len(videos))
	}

	if videos[0].CodecName != "h264" {
		t.Errorf("unexpected codec name: %s", videos[0].CodecName)
	}

	if videos[0].CodecType != "video" {
		t.Errorf("unexpected codec type: %s", videos[0].CodecType)
	}

	if videos[0].Width != 1920 {
		t.Errorf("unexpected width: %d", videos[0].Width)
	}

	if videos[0].Height != 1080 {
		t.Errorf("unexpected height: %d", videos[0].Height)
	}
}

func TestRTSP(t *testing.T) {
	probe, err := Probe("rtsp://admin:admin@192.168.1.97:554")
	if err != nil {
		t.Fatalf("failed to probe: %v", err)
	}

	if probe.Format.FormatName != "rtsp" {
		t.Errorf("unexpected format name: %s", probe.Format.FormatName)
	}

	if probe.Format.NBStreams != 2 {
		t.Errorf("unexpected number of streams: %d", probe.Format.NBStreams)
	}

	videos, err := probe.GetStreamType("video")
	if err != nil {
		t.Fatalf("failed to get video stream: %v", err)
	}

	if len(videos) != 1 {
		t.Fatalf("unexpected number of video streams: %d", len(videos))
	}

	audios, err := probe.GetStreamType("audio")
	if err != nil {
		t.Fatalf("failed to get audio stream: %v", err)
	}

	if len(audios) != 1 {
		t.Fatalf("unexpected number of audio streams: %d", len(audios))
	}
}
