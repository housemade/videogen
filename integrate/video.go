package integrate

import (
	"fmt"
	"github.com/pkg/errors"
	"homemade/tts"
	"homemade/youtube"
	"os"
	"os/exec"
)

type Integrator struct {
	ttsClient     tts.Client
	youtubeClient youtube.Client
}

func NewIntegrator(ttsClient tts.Client, youtubeClient youtube.Client) Integrator {
	return Integrator{
		ttsClient:     ttsClient,
		youtubeClient: youtubeClient,
	}
}

func (i Integrator) GetSpeechFile(voiceID, text string) error {
	tmpFile, err := os.CreateTemp("", "output-*.mp3")
	if err != nil {
		return errors.Wrap(err, "error creating temporary file")
	}
	defer tmpFile.Close()

	err = i.ttsClient.TextToSpeech(voiceID, text, tmpFile.Name())
	if err != nil {
		return errors.Wrap(err, "error getting speech file")
	}

	// TODO: filename rename 필요
	err = os.Rename(tmpFile.Name(), "test.mp3")
	if err != nil {
		return errors.Wrap(err, "error renaming temporary file to test.mp3")
	}
	return nil
}

func (i Integrator) UploadVideo() error {
	err := convertMP3ToMP4("test.mp3", "test.mp4")
	if err != nil {
		return errors.Wrap(err, "error renaming temporary file to test.mp3")
	}
	return nil
}

func convertMP3ToMP4(inputPath, outputPath string) error {
	cmd := exec.Command("ffmpeg", "-i", inputPath, outputPath)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error converting file: %w", err)
	}
	return nil
}
