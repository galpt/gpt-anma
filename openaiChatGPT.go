package main

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func handleChatGPT() {

	// prevent panic
	getOpenAIAPIkey()

	// required vars for chatgpt client
	floats := randFloats(openaiMinRange, openaiMaxRange, 2)
	floatstopp := randFloats(0.85, 1.0, 1)
	freq := float32(floats[0])
	temp := float32(floats[1])
	ptop := float32(floatstopp[0])

	// read user input data from the text files
	readJunsTemplate := readFileToStr(filePath + fileJunsPersonality)
	readJunsLastResp := readFileToStr(filePath + fileOldAnswerDialogue)
	readUsrQuestionTxt := readFileToStr(filePath + fileQuestion)

	// reconstruct user question
	usrQuestion := fmt.Sprintf("%v\n\nUser: %v", readJunsTemplate, readUsrQuestionTxt)

	// use moderation endpoint
	chkMsg := openaiModeration(usrQuestion)
	if chkMsg {
		fmt.Println("Your message didn't pass the Moderations API.")
		os.Exit(0)
	}

	// get user's IP address
	userIP := getUserIP()

	// reconstruct openai client
	resp, err := openaiClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:            openai.GPT3Dot5Turbo,
			MaxTokens:        openaiMaxTokens,
			Temperature:      temp,
			TopP:             ptop,
			N:                openaiNVal,
			PresencePenalty:  openaiPresPen,
			FrequencyPenalty: freq,
			User:             fmt.Sprintf("%v", userIP),
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: readJunsTemplate,
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: readJunsLastResp,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("\n\nUser: %v", readUsrQuestionTxt),
				},
			},
		},
	)
	if err != nil {
		errMsg := fmt.Sprintf(`%v\n\n[ChatGPT API] Something went wrong.`, err)
		fmt.Println(errMsg)
		os.Exit(0)
	}

	saveResponse := openaiLastUsageInfo(readUsrQuestionTxt, resp)
	if saveResponse != nil {
		errMsg := fmt.Sprintf(`%v\n\n[ChatGPT Response Handler] Something went wrong.`, saveResponse)
		fmt.Println(errMsg)
		os.Exit(0)
	}

}
