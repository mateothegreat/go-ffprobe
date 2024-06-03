package ffprobe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
)

type FFProbeOutput struct {
	Streams []Stream `json:"streams"`
	Format  Format   `json:"format"`
}

// {
//   "streams": [
//     {
//       "index": 0,
//       "codec_name": "hevc",
//       "codec_long_name": "H.265 / HEVC (High Efficiency Video Coding)",
//       "profile": "Main",
//       "codec_type": "video",
//       "codec_tag_string": "[0][0][0][0]",
//       "codec_tag": "0x0000",
//       "width": 3840,
//       "height": 2160,
//       "coded_width": 3840,
//       "coded_height": 2160,
//       "closed_captions": 0,
//       "film_grain": 0,
//       "has_b_frames": 0,
//       "sample_aspect_ratio": "1:1",
//       "display_aspect_ratio": "16:9",
//       "pix_fmt": "yuvj420p",
//       "level": 153,
//       "color_range": "pc",
//       "chroma_location": "left",
//       "refs": 1,
//       "r_frame_rate": "15000/1001",
//       "avg_frame_rate": "0/0",
//       "time_base": "1/90000",
//       "start_pts": 6030,
//       "start_time": "0.067000",
//       "extradata_size": 87,
//       "disposition": {
//         "default": 0,
//         "dub": 0,
//         "original": 0,
//         "comment": 0,
//         "lyrics": 0,
//         "karaoke": 0,
//         "forced": 0,
//         "hearing_impaired": 0,
//         "visual_impaired": 0,
//         "clean_effects": 0,
//         "attached_pic": 0,
//         "timed_thumbnails": 0,
//         "non_diegetic": 0,
//         "captions": 0,
//         "descriptions": 0,
//         "metadata": 0,
//         "dependent": 0,
//         "still_image": 0
//       }
//     },
//     {
//       "index": 1,
//       "codec_name": "pcm_mulaw",
//       "codec_long_name": "PCM mu-law / G.711 mu-law",
//       "codec_type": "audio",
//       "codec_tag_string": "[0][0][0][0]",
//       "codec_tag": "0x0000",
//       "sample_fmt": "s16",
//       "sample_rate": "8000",
//       "channels": 1,
//       "channel_layout": "mono",
//       "bits_per_sample": 8,
//       "initial_padding": 0,
//       "r_frame_rate": "0/0",
//       "avg_frame_rate": "0/0",
//       "time_base": "1/8000",
//       "start_pts": 0,
//       "start_time": "0.000000",
//       "bit_rate": "64000",
//       "disposition": {
//         "default": 0,
//         "dub": 0,
//         "original": 0,
//         "comment": 0,
//         "lyrics": 0,
//         "karaoke": 0,
//         "forced": 0,
//         "hearing_impaired": 0,
//         "visual_impaired": 0,
//         "clean_effects": 0,
//         "attached_pic": 0,
//         "timed_thumbnails": 0,
//         "non_diegetic": 0,
//         "captions": 0,
//         "descriptions": 0,
//         "metadata": 0,
//         "dependent": 0,
//         "still_image": 0
//       }
//     }
//   ],
//   "format": {
//     "filename": "rtsp://admin:asdf@192.168.1.97:554",
//     "nb_streams": 2,
//     "nb_programs": 0,
//     "format_name": "rtsp",
//     "format_long_name": "RTSP input",
//     "start_time": "0.000000",
//     "probe_score": 100,
//     "tags": {
//       "title": "Media Presentation"
//     }
//   }
// }

type Stream struct {
	Index         int    `json:"index"`
	CodecName     string `json:"codec_name"`
	CodecType     string `json:"codec_type"`
	Width         int    `json:"width,omitempty"`
	Height        int    `json:"height,omitempty"`
	BitRate       string `json:"bit_rate,omitempty"`
	SampleRate    string `json:"sample_rate,omitempty"`
	Channels      int    `json:"channels,omitempty"`
	ChannelLayout string `json:"channel_layout,omitempty"`
	Duration      string `json:"duration"`
	Tags          Tags   `json:"tags"`
}

type Format struct {
	Filename       string  `json:"filename"`
	NBStreams      int     `json:"nb_streams"`
	NBPrograms     int     `json:"nb_programs"`
	FormatName     string  `json:"format_name"`
	FormatLongName string  `json:"format_long_name"`
	StartTime      string  `json:"start_time"`
	Duration       float32 `json:"duration"` // doesn't exist on rtsp streams
	Size           string  `json:"size"`
	BitRate        string  `json:"bit_rate"`
	Tags           Tags    `json:"tags"`
}

type Tags struct {
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
	Genre    string `json:"genre"`
	Comment  string `json:"comment"`
	Encoder  string `json:"encoder"`
	Language string `json:"language"`
}

func (f *Format) UnmarshalJSON(data []byte) error {
	type Alias Format
	aux := &struct {
		Duration string `json:"duration"`
		*Alias
	}{
		Alias: (*Alias)(f),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	if aux.Duration != "" {
		duration, err := strconv.ParseFloat(aux.Duration, 32)
		if err != nil {
			return fmt.Errorf("failed to parse duration: %v", err)
		}
		f.Duration = float32(duration)
	}

	return nil
}

func (o *FFProbeOutput) GetStreamType(t string) ([]Stream, error) {
	var streams []Stream

	for _, stream := range o.Streams {
		if stream.CodecType == t {
			streams = append(streams, stream)
		}
	}

	if len(streams) == 0 {
		return nil, fmt.Errorf("no streams found with codec type %s", t)
	}

	return streams, nil
}

func Probe(filePath string) (*FFProbeOutput, error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-print_format", "json", "-show_format", "-show_streams", filePath)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run ffprobe: %v, stderr: %s", err, stderr.String())
	}

	var probeOutput FFProbeOutput
	if err := json.Unmarshal(stdout.Bytes(), &probeOutput); err != nil {
		return nil, fmt.Errorf("failed to parse ffprobe output: %v", err)
	}

	return &probeOutput, nil
}
