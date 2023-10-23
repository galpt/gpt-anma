package main

import (
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func handleWhisper() {

	// prevent panic
	getOpenAIAPIkey()

	// reconstruct whisper request
	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: (filePath + fileAudio),
	}
	resp, err := openaiClient.CreateTranscription(ctx, req)
	if err != nil {
		errMsg := fmt.Sprintf(`%v\n\n[Whisper API] Something went wrong.`, err)
		fmt.Println(errMsg)
		os.Exit(0)
	}

	// save response to user question file
	makeNewFile := createFileWriteString(fileQuestion, (filePath + fileQuestion), resp.Text)
	if makeNewFile != nil {
		errMsg := fmt.Sprintf(`%v\n\n[Whisper Response Handler] Something went wrong.`, makeNewFile)
		fmt.Println(errMsg)
		os.Exit(0)
	}

}
