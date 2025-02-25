package video

import (
	"os/exec"

	"github.com/MetaerMarket/goconf"
)

type VideoDownload struct {
	C      chan string
	AudioC chan []interface{}
	MsgC   chan string
	path   string
}

func NewVideoDownload() *VideoDownload {
	return &VideoDownload{
		C:      make(chan string, 3),
		MsgC:   make(chan string, 3),
		AudioC: make(chan []interface{}, 3),
		path:   goconf.VarStringOrDefault("/tmp/aio-tgbot/video/", "video", "path"),
	}
}

func (t *VideoDownload) Cut(file, start, end, output string) error {
	args := []string{"-i", file, "-ss", start, "-to", end, "-c", "copy", output}
	cmd := exec.Command("ffmpeg", args...)
	err := cmd.Run()

	return err
}
