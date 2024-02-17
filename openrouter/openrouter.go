package openrouter

import (
	"encoding/json"
	"net/http"
)

type OpenRouterRequest struct {
	Prompt string
	Seed int
}

func Ask(question string) {
	requestBody := OpenRouterRequest{
		Prompt: question,
	}

	client := &http.Client{}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return error
	}

	req, err := http.NewRequest(http.MethodPost, "https://openrouter.ai/api/v1/chat/completions", bytes)
}