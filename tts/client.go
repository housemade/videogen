package tts

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

const (
	XI_API_KEY = ""
	MODEL_ID   = "eleven_multilingual_v2"
)

type Client struct {
	client *resty.Client
}

func NewTTSClient() Client {
	return Client{
		client: resty.New(),
	}
}

func (c Client) GetModels() error {
	resp, err := c.client.R().
		SetHeader("Accept", "audio/mpeg").
		SetHeader("Content-Type", "application/json").
		SetHeader("xi-api-key", XI_API_KEY).
		Get("https://api.elevenlabs.io/v1/models")

	if err != nil {
		return errors.Wrap(err, "error making request")
	}

	if resp.IsError() {
		return errors.New(fmt.Sprintf("error response: %s", resp.String()))
	}

	return nil
}

func (c Client) ListVoices() (VoiceResponse, error) {

	var response VoiceResponse
	resp, err := c.client.R().
		SetHeader("Accept", "audio/mpeg").
		SetHeader("Content-Type", "application/json").
		SetHeader("xi-api-key", XI_API_KEY).
		SetResult(&response).
		Get("https://api.elevenlabs.io/v1/voices")

	if err != nil {
		return VoiceResponse{}, errors.Wrap(err, "error making request")
	}

	if resp.IsError() {
		return VoiceResponse{}, errors.New(fmt.Sprintf("error response: %s", resp.String()))
	}

	return response, nil
}

func (c Client) TextToSpeech(voiceID, text, outputPath string) error {
	ttsURL := fmt.Sprintf("https://api.elevenlabs.io/v1/text-to-speech/%s", voiceID)

	resp, err := c.client.R().
		SetHeader("Accept", "audio/mpeg").
		SetHeader("Content-Type", "application/json").
		SetHeader("xi-api-key", XI_API_KEY).
		SetBody(map[string]interface{}{
			"text":     text,
			"model_id": "eleven_multilingual_v2",
			"voice_settings": map[string]interface{}{
				"stability":        0.5,
				"similarity_boost": 0.5,
				"style":            0.5,
			},
		}).
		SetOutput(outputPath).
		Post(ttsURL)

	if err != nil {
		return errors.Wrap(err, "error making request")
	}

	if resp.IsError() {
		return errors.New(fmt.Sprintf("error response: %s", resp.String()))
	}

	return nil
}
