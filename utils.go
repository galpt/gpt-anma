package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/afero"
)

// get OpenAI API key
func getOpenAIAPIkey() {
	openaiAPIKey = readFileToStr(filePath + fileOpenAIAPIKey) // value will be filled automatically
	openaiClient = openai.NewClient(openaiAPIKey)
}

// get user's IP to prevent abuses while using OpenAI's API
func getUserIP() string {
	req, err := http.Get("https://0ms.dev/get-ip")
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}

	return string(body)
}

// support for random float numbers generation
func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}

// make a new file and write string data to it
func createFileWriteString(filename, path, content string) error {

	createFile, err := osFS.Create(path)
	if err != nil {
		errMsg := fmt.Sprintf(`%v\n\n[CREATE FILE] Failed to create "%v"`, err, filename)
		fmt.Println(errMsg)
		os.Exit(0)
	}
	_, err2 := createFile.WriteString(content)
	if err2 != nil {
		errMsg := fmt.Sprintf(`%v\n\n[WRITE FILE] Failed to write data to "%v"`, err2, filename)
		fmt.Println(errMsg)
		os.Exit(0)
	}
	createFile.Close()

	return err
}

// read file based on the given path
func readFileToStr(path string) string {

	readFile, err := afero.ReadFile(osFS, path)
	if err != nil {
		errMsg := fmt.Sprintf(`%v\n\n[ERROR] Failed to read "%v"`, err, path)
		fmt.Println(errMsg)
		os.Exit(0)
	}

	return string(readFile)
}

// use OpenAI Moderations API
func openaiModeration(input string) bool {

	moderationReq := openai.ModerationRequest{
		Model: openai.ModerationTextLatest,
		Input: input,
	}

	resp, err := openaiClient.Moderations(ctx, moderationReq)
	if err != nil {
		errMsg := fmt.Sprintf(`%v\n\n[ERROR] Something went wrong.`, err)
		fmt.Println(errMsg)
		os.Exit(0)
	}

	return resp.Results[0].Flagged
}

// get ChatGPT last usage data
func openaiSaveChatGPTResp(cgptResp openai.ChatCompletionResponse) error {

	// split response flag from the dialogue
	var (
		respDialogue = ""
		respFlag     = ""
	)
	if strings.Contains(cgptResp.Choices[0].Message.Content, "[[") {
		split := strings.Split(cgptResp.Choices[0].Message.Content, "[[")
		respDialogue = split[0]
		respFlag = fmt.Sprintf("[[%v", split[1])
	}

	// save new response dialogue
	os.RemoveAll(filePath + fileNewAnswerDialogue)
	saveNewDialogue := createFileWriteString(fileNewAnswerDialogue, (filePath + fileNewAnswerDialogue), respDialogue)
	if saveNewDialogue != nil {
		return saveNewDialogue
	}

	// save new response flag
	os.RemoveAll(filePath + fileNewAnswerFlag)
	saveNewFlag := createFileWriteString(fileNewAnswerFlag, (filePath + fileNewAnswerFlag), respFlag)
	if saveNewFlag != nil {
		return saveNewFlag
	}

	return saveNewFlag
}

// get ChatGPT last usage data
func openaiLastUsageInfo(usrQuestion string, cgptResp openai.ChatCompletionResponse) error {

	// parse and save new AI response
	saveResp := openaiSaveChatGPTResp(cgptResp)
	if saveResp != nil {
		return saveResp
	}

	// SAVE THIS TO LAST USAGE INFO FILE
	// read last AI response (old)
	readOldAIRespDialogue := readFileToStr(filePath + fileOldAnswerDialogue)
	readOldAIRespFlag := readFileToStr(filePath + fileOldAnswerFlag)

	lastUsageInfoTemplate := fmt.Sprintf(`
	====
	[TOKEN USAGE]
	Input: %v
	Output: %v
	Total: %v
	====
	
	====
	[HISTORY]
	Last User Input:
	%v
	Last AI Output (New):
	%v
	Last AI Output (Old):
	%v %v
	====
	`, cgptResp.Usage.PromptTokens, cgptResp.Usage.CompletionTokens, cgptResp.Usage.TotalTokens, usrQuestion, cgptResp.Choices[0].Message.Content, readOldAIRespDialogue, readOldAIRespFlag)

	makeNewFile := createFileWriteString(fileOpenAILastUsage, (filePath + fileOpenAILastUsage), lastUsageInfoTemplate)
	if makeNewFile != nil {
		return makeNewFile
	}

	// replace the old answer dialogue & flag files.
	// split response flag from the dialogue
	var (
		respDialogue = ""
		respFlag     = ""
	)
	if strings.Contains(cgptResp.Choices[0].Message.Content, "[[") {
		split := strings.Split(cgptResp.Choices[0].Message.Content, "[[")
		respDialogue = split[0]
		respFlag = fmt.Sprintf("[[%v", split[1])
	}

	// save new response dialogue
	os.RemoveAll(filePath + fileOldAnswerDialogue)
	saveNewDialogue := createFileWriteString(fileOldAnswerDialogue, (filePath + fileOldAnswerDialogue), respDialogue)
	if saveNewDialogue != nil {
		return saveNewDialogue
	}

	// save new response flag
	os.RemoveAll(filePath + fileOldAnswerFlag)
	saveNewFlag := createFileWriteString(fileOldAnswerFlag, (filePath + fileOldAnswerFlag), respFlag)
	if saveNewFlag != nil {
		return saveNewFlag
	}

	return makeNewFile
}
