package main

import (
	"github.com/samber/lo"
	"homemade/integrate"
	"homemade/tts"
	"homemade/youtube"
)

func main() {
	ttsClient := tts.NewTTSClient()
	response, err := ttsClient.ListVoices()
	if err != nil {
		panic(err)
	}

	highQualityVoices := lo.Filter(response.Voices, func(voice tts.Voice, _ int) bool {
		return lo.Contains(voice.HighQualityBaseModelIds, tts.MODEL_ID)
	})
	lauraVoice, ok := lo.Find(highQualityVoices, func(voice tts.Voice) bool {
		return voice.Name == "Laura"
	})
	if !ok {
		panic("Laura voice not found")
	}

	youtubeClient := youtube.NewClient("client_secret.json")

	integrator := integrate.NewIntegrator(ttsClient, youtubeClient)
	err = integrator.GetSpeechFile(lauraVoice.VoiceId, "오늘 날씨가 너무 좋은 것 같아요")
	if err != nil {
		panic(err)
	}

	err = integrator.UploadVideo()
	if err != nil {
		panic(err)
	}
}
