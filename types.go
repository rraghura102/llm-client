package main

// Author: Rayan Raghuram
// Cpyright @ 2025 Rayan Raghuram. All rights reserved.
//
// Data structures use by the llm-client

import (
    "encoding/json"
    "fmt"
)

type BaseChatRequest struct {
    CachePrompt       bool        `json:"cache_prompt"`
    FrequencyPenalty  float64     `json:"frequency_penalty"`
    ImageData         interface{} `json:"image_data"`
    MainGPU           int         `json:"main_gpu"`
    MinP              int         `json:"min_p"`
    Mirostat          float64     `json:"mirostat"`
    MirostatEta       float64     `json:"mirostat_eta"`
    MirostatTau       float64     `json:"mirostat_tau"`
    NKeep             int         `json:"n_keep"`
    NPredict          int         `json:"n_predict"`
    PenalizeNL        bool        `json:"penalize_nl"`
    PresencePenalty   float64     `json:"presence_penalty"`
    Prompt            string      `json:"prompt"`
    RepeatLastN       int         `json:"repeat_last_n"`
    RepeatPenalty     float64     `json:"repeat_penalty"`
    Seed              int         `json:"seed"`
    Stop              interface{} `json:"stop"`
    Stream            bool        `json:"stream"`
    Temperature       float64     `json:"temperature"`
    TopK              int         `json:"top_k"`
    TopP              float64     `json:"top_p"`
    TypicalP          float64     `json:"typical_p"`
}

type SecureChatRequest struct {
    Role                 string `json:"role"`
    EncryptedPrompt      string `json:"EncryptedPrompt"`
    EncryptedSymmetricKey string `json:"encryptedSymmetricKey"`
}

// --- Interface ---
type ChatRequest interface {
    ToJSON() []byte
}

// --- Implement Interface ---
func (s SecureChatRequest) ToJSON() []byte {
    b, _ := json.Marshal(s)
    return b
}

func (b BaseChatRequest) ToJSON() []byte {
    bts, _ := json.Marshal(b)
    return bts
}

const promptFormat = "<|start_header_id|>system<|end_header_id|>\n\n" +
    "Cutting Knowledge Date: December 2023\n\n" +
    "<|eot_id|><|start_header_id|>user<|end_header_id|>\n\n" +
    "%s" +
    "<|eot_id|><|start_header_id|>assistant<|end_header_id|>\n\n"

func NewChatRequest(prompt string) ChatRequest {
    prompt = fmt.Sprintf(promptFormat, prompt)
    return BaseChatRequest{
        CachePrompt:      true,
        FrequencyPenalty: 0,
        ImageData:        nil,
        MainGPU:          0,
        MinP:             0,
        Mirostat:         0,
        MirostatEta:      0.1,
        MirostatTau:      5,
        NKeep:            4,
        NPredict:         -1,
        PenalizeNL:       true,
        PresencePenalty:  0,
        Prompt:           prompt,
        RepeatLastN:      64,
        RepeatPenalty:    1.1,
        Seed:             -1,
        Stop:             nil,
        Stream:           true,
        Temperature:      0.8,
        TopK:             40,
        TopP:             0.9,
        TypicalP:         1,
    }
}

func NewSecureChatRequest(prompt string) ChatRequest {
    prompt = fmt.Sprintf(promptFormat, prompt)
    return SecureChatRequest{
        Role:                 "user",
        EncryptedPrompt:      prompt,
        EncryptedSymmetricKey: "",
    }
}

type ChatResponse struct {
    Content string `json:"content"`
    Stop    bool   `json:"stop"`
    Timings struct {
        PredictedN  int `json:"predicted_n"`
        PredictedMs int `json:"predicted_ms"`
        PromptN     int `json:"prompt_n"`
        PromptMs    int `json:"prompt_ms"`
    } `json:"timings"`
    StoppedLimit bool `json:"stopped_limit"`
}
