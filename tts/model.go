package tts

type Voice struct {
	VoiceId    string      `json:"voice_id"`
	Name       string      `json:"name"`
	Samples    interface{} `json:"samples"`
	Category   string      `json:"category"`
	FineTuning struct {
		IsAllowedToFineTune bool `json:"is_allowed_to_fine_tune"`
		State               struct {
			ElevenMultilingualV2    string `json:"eleven_multilingual_v2,omitempty"`
			ElevenTurboV25          string `json:"eleven_turbo_v2_5,omitempty"`
			ElevenTurboV2           string `json:"eleven_turbo_v2,omitempty"`
			ElevenMultilingualStsV2 string `json:"eleven_multilingual_sts_v2,omitempty"`
		} `json:"state"`
		VerificationFailures        []interface{} `json:"verification_failures"`
		VerificationAttemptsCount   int           `json:"verification_attempts_count"`
		ManualVerificationRequested bool          `json:"manual_verification_requested"`
		Language                    string        `json:"language"`
		Progress                    struct {
		} `json:"progress"`
		Message struct {
			ElevenMultilingualV2    string `json:"eleven_multilingual_v2,omitempty"`
			ElevenTurboV25          string `json:"eleven_turbo_v2_5,omitempty"`
			ElevenTurboV2           string `json:"eleven_turbo_v2,omitempty"`
			ElevenMultilingualStsV2 string `json:"eleven_multilingual_sts_v2,omitempty"`
		} `json:"message"`
		DatasetDurationSeconds interface{} `json:"dataset_duration_seconds"`
		VerificationAttempts   interface{} `json:"verification_attempts"`
		SliceIds               interface{} `json:"slice_ids"`
		ManualVerification     interface{} `json:"manual_verification"`
	} `json:"fine_tuning"`
	Labels struct {
		Accent      string `json:"accent"`
		Description string `json:"description"`
		Age         string `json:"age"`
		Gender      string `json:"gender"`
		UseCase     string `json:"use_case"`
	} `json:"labels"`
	Description             interface{}   `json:"description"`
	PreviewUrl              string        `json:"preview_url"`
	AvailableForTiers       []interface{} `json:"available_for_tiers"`
	Settings                interface{}   `json:"settings"`
	Sharing                 interface{}   `json:"sharing"`
	HighQualityBaseModelIds []string      `json:"high_quality_base_model_ids"`
	SafetyControl           interface{}   `json:"safety_control"`
	VoiceVerification       struct {
		RequiresVerification      bool          `json:"requires_verification"`
		IsVerified                bool          `json:"is_verified"`
		VerificationFailures      []interface{} `json:"verification_failures"`
		VerificationAttemptsCount int           `json:"verification_attempts_count"`
		Language                  interface{}   `json:"language"`
		VerificationAttempts      interface{}   `json:"verification_attempts"`
	} `json:"voice_verification"`
	PermissionOnResource interface{} `json:"permission_on_resource"`
	IsOwner              bool        `json:"is_owner"`
	IsLegacy             bool        `json:"is_legacy"`
	IsMixed              bool        `json:"is_mixed"`
}

type VoiceResponse struct {
	Voices []Voice `json:"voices"`
}
