package main

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/afero"
)

const (
	Gigabyte = 1 << 30
	Megabyte = 1 << 20
	Kilobyte = 1 << 10

	appVer = "GPT-ANMA version 2023.5.5.16"

	// file names
	filePath              = "./GPT-ANMA/"
	fileAudio             = "question.mp3"
	fileJunsPersonality   = "your-juns-base-personality.txt"
	fileQuestion          = "your-question-here.txt"
	fileOldAnswerDialogue = "your-old-answer-here.txt"
	fileOldAnswerFlag     = "your-old-answer-flag-here.txt"
	fileNewAnswerDialogue = "your-new-answer-here.txt"
	fileNewAnswerFlag     = "your-new-answer-flag-here.txt"
	fileOpenAIAPIKey      = "your-openai-apikey-here.txt"
	fileOpenAILastUsage   = "your-openai-last-usage.txt"
)

var (
	osFS = afero.NewOsFs()
	ctx  = context.Background()

	// VARS FOR OPENAI
	openaiAPIKey    = ""
	openaiClient    *openai.Client
	openaiMinRange  = float64(0.1)
	openaiMaxRange  = float64(0.9)
	openaiPresPen   = float32(0.0)
	openaiNVal      = 1
	openaiMaxTokens = 1000 // 100 tokens ~= 75 words
)
