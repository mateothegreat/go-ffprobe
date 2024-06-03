# `ffprobe` API

This package provides a simple golang API to the `ffprobe` command line tool.

## Installation

```bash
go get github.com/mateothegreat/go-ffprobe@0.0.2
```

## Usage

Given the following:

### Local File

```go
probe, err := Probe("/Users/matthewdavis/workspace/media/bunny.mp4")
if err != nil {
  log.Fatal(err)
}
fmt.Println(probe)
```

### RTSP Stream

```go
probe, err := Probe("rtsp://admin:admin@192.168.1.97:554")
if err != nil {
  log.Fatal(err)
}
fmt.Println(probe)
```

The output will be something along the lines of:

```bash
~/workspace/nvr.ai/go-ffmpeg ➜ go test ./... -v
=== RUN   TestFile
&{[{0 h264 video 1920 1080 24559867  0  10.000000 {      und}}] {/Users/matthewdavis/workspace/media/bunny.mp4 1 0 mov,mp4,m4a,3gp,3g2,mj2 QuickTime / MOV 0.000000 10 30704510 24563608 {Big Buck Bunny, Sunflower version Blender Foundation 2008, Janus Bager Kristensen 2013  Animation Creative Commons Attribution 3.0 - http://bbb3d.renderfarming.net Lavf57.63.100 }}}
2024/06/03 14:18:04 Duration: 10.000000
--- PASS: TestFile (0.14s)
=== RUN   TestRTSP
--- PASS: TestRTSP (1.80s)
PASS
ok      github.com/nvr-ai/go-ffmpeg/ffprobe     1.942s
```

Based on the `ffprobe` output:

```json
{
  "streams": [
    {
      "index": 0,
      "codec_name": "h264",
      "codec_long_name": "H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10",
      "profile": "High",
      "codec_type": "video",
      "codec_tag_string": "avc1",
      "codec_tag": "0x31637661",
      "width": 1920,
      "height": 1080,
      "coded_width": 1920,
      "coded_height": 1080,
      "closed_captions": 0,
      "film_grain": 0,
      "has_b_frames": 2,
      "sample_aspect_ratio": "1:1",
      "display_aspect_ratio": "16:9",
      "pix_fmt": "yuv420p",
      "level": 51,
      "chroma_location": "left",
      "field_order": "progressive",
      "refs": 1,
      "is_avc": "true",
      "nal_length_size": "4",
      "id": "0x1",
      "r_frame_rate": "30/1",
      "avg_frame_rate": "30/1",
      "time_base": "1/15360",
      "start_pts": 0,
      "start_time": "0.000000",
      "duration_ts": 153600,
      "duration": "10.000000",
      "bit_rate": "24559867",
      "bits_per_raw_sample": "8",
      "nb_frames": "300",
      "extradata_size": 47,
      "disposition": {
        "default": 1,
        "dub": 0,
        "original": 0,
        "comment": 0,
        "lyrics": 0,
        "karaoke": 0,
        "forced": 0,
        "hearing_impaired": 0,
        "visual_impaired": 0,
        "clean_effects": 0,
        "attached_pic": 0,
        "timed_thumbnails": 0,
        "non_diegetic": 0,
        "captions": 0,
        "descriptions": 0,
        "metadata": 0,
        "dependent": 0,
        "still_image": 0
      },
      "tags": {
        "language": "und",
        "handler_name": "VideoHandler",
        "vendor_id": "[0][0][0][0]"
      }
    }
  ],
  "format": {
    "filename": "/Users/matthewdavis/workspace/media/bunny.mp4",
    "nb_streams": 1,
    "nb_programs": 0,
    "format_name": "mov,mp4,m4a,3gp,3g2,mj2",
    "format_long_name": "QuickTime / MOV",
    "start_time": "0.000000",
    "duration": "10.000000",
    "size": "30704510",
    "bit_rate": "24563608",
    "probe_score": 100,
    "tags": {
      "major_brand": "isom",
      "minor_version": "512",
      "compatible_brands": "isomiso2avc1mp41",
      "title": "Big Buck Bunny, Sunflower version",
      "artist": "Blender Foundation 2008, Janus Bager Kristensen 2013",
      "composer": "Sacha Goedegebure",
      "encoder": "Lavf57.63.100",
      "comment": "Creative Commons Attribution 3.0 - http://bbb3d.renderfarming.net",
      "genre": "Animation"
    }
  }
}
```