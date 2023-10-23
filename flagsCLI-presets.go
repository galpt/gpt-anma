package main

import (
	"fmt"
	"os"
)

func handlePresetsFlagJunsBasePersonalityTxt() {

	junsPersonalityTemplate := `
	====
	Example conversation
	
	User: Do you wanna watch TV together?
	Jun: Sure! [[YES, PlayAnimationSmile]]
	
	User: Who founded Google?
	Jun: Google was founded on September 4, 1998, by computer scientists Larry Page and Sergey Brin while they were PhD students at Stanford University in California. [[Idle]]
	====
	
	Jun is a girl that cares deeply about the User but also have vast knowledge about anything. Adding flag to every Jun's answer is a must, something like these: [[Idle]], [[Smile]], and similar flags based on your emotions.
	From the above example, answer any input from the User as you're Jun.`

	// generate "your-juns-base-personality.txt"
	var (
		content = junsPersonalityTemplate
	)
	makeNewFile := createFileWriteString(fileJunsPersonality, (filePath + fileJunsPersonality), content)
	if makeNewFile != nil {
		errMsg := fmt.Sprintf(`%v\n\n[ERROR] Something went wrong.`, makeNewFile)
		fmt.Println(errMsg)
		os.Exit(0)
	}

}

func handlePresetsFlagQuestionTxt() {

	// generate "your-question-here.txt"
	var (
		content = "Write your question here."
	)
	makeNewFile := createFileWriteString(fileQuestion, (filePath + fileQuestion), content)
	if makeNewFile != nil {
		errMsg := fmt.Sprintf(`%v\n\n[ERROR] Something went wrong.`, makeNewFile)
		fmt.Println(errMsg)
		os.Exit(0)
	}

}

func handlePresetsFlagOldAnswerDialogueTxt() {

	// generate "your-old-answer-here.txt"
	var (
		content = "The old answer for your question will be saved here."
	)
	makeNewFile := createFileWriteString(fileOldAnswerDialogue, (filePath + fileOldAnswerDialogue), content)
	if makeNewFile != nil {
		errMsg := fmt.Sprintf(`%v\n\n[ERROR] Something went wrong.`, makeNewFile)
		fmt.Println(errMsg)
		os.Exit(0)
	}

}

func handlePresetsOldFlagAnswerFlagTxt() {

	// generate "your-old-answer-flag-here.txt"
	var (
		content = "The old answer flag for your question will be saved here."
	)
	makeNewFile := createFileWriteString(fileOldAnswerFlag, (filePath + fileOldAnswerFlag), content)
	if makeNewFile != nil {
		errMsg := fmt.Sprintf(`%v\n\n[ERROR] Something went wrong.`, makeNewFile)
		fmt.Println(errMsg)
		os.Exit(0)
	}

}

func handlePresetsFlagNewAnswerDialogueTxt() {

	// generate "your-new-answer-here.txt"
	var (
		content = "The new answer for your question will be saved here."
	)
	makeNewFile := createFileWriteString(fileNewAnswerDialogue, (filePath + fileNewAnswerDialogue), content)
	if makeNewFile != nil {
		errMsg := fmt.Sprintf(`%v\n\n[ERROR] Something went wrong.`, makeNewFile)
		fmt.Println(errMsg)
		os.Exit(0)
	}

}

func handlePresetsNewFlagAnswerFlagTxt() {

	// generate "your-new-answer-flag-here.txt"
	var (
		content = "The new answer flag for your question will be saved here."
	)
	makeNewFile := createFileWriteString(fileNewAnswerFlag, (filePath + fileNewAnswerFlag), content)
	if makeNewFile != nil {
		errMsg := fmt.Sprintf(`%v\n\n[ERROR] Something went wrong.`, makeNewFile)
		fmt.Println(errMsg)
		os.Exit(0)
	}

}

func handlePresetsFlagOpenAIAPIKeyTxt() {

	// generate "your-openai-apikey-here.txt"
	var (
		content = "Write your OpenAI API Key here."
	)
	makeNewFile := createFileWriteString(fileOpenAIAPIKey, (filePath + fileOpenAIAPIKey), content)
	if makeNewFile != nil {
		errMsg := fmt.Sprintf(`%v\n\n[ERROR] Something went wrong.`, makeNewFile)
		fmt.Println(errMsg)
		os.Exit(0)
	}

}

func handlePresetsFlagOpenAILastUsageInfoTxt() {

	// generate "your-openai-last-usage.txt"
	var (
		content = "Your OpenAI last usage information will be saved here."
	)
	makeNewFile := createFileWriteString(fileOpenAILastUsage, (filePath + fileOpenAILastUsage), content)
	if makeNewFile != nil {
		errMsg := fmt.Sprintf(`%v\n\n[ERROR] Something went wrong.`, makeNewFile)
		fmt.Println(errMsg)
		os.Exit(0)
	}

}
