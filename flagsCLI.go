package main

import (
	"flag"
	"fmt"
	"os"
)

func cliFlags() {

	versionFlag := flag.Bool("version", false, "Print current program version.")
	presetsFlag := flag.Bool("presets", false, "Generate the required files for GPT-ANMA.")
	chatgptFlag := flag.Bool("chatgpt", false, "Process data inside the 'GPT-ANMA' folder using OpenAI's ChatGPT.")
	whisperFlag := flag.Bool("whisper", false, "Process data inside the 'GPT-ANMA' folder using OpenAI's Whisper.")

	flag.Parse()

	// handle print version flag
	if *versionFlag {
		fmt.Println()
		fmt.Println(appVer)
		os.Exit(0)
	}

	// handle presets flag
	if *presetsFlag {
		osFS.RemoveAll("./GPT-ANMA/")
		osFS.MkdirAll("./GPT-ANMA/", 0777)
		handlePresetsFlagJunsBasePersonalityTxt()
		handlePresetsFlagQuestionTxt()
		handlePresetsFlagOldAnswerDialogueTxt()
		handlePresetsOldFlagAnswerFlagTxt()
		handlePresetsFlagNewAnswerDialogueTxt()
		handlePresetsNewFlagAnswerFlagTxt()
		handlePresetsFlagOpenAIAPIKeyTxt()
		handlePresetsFlagOpenAILastUsageInfoTxt()
		fmt.Println("Successfully generated the required files to 'GPT-ANMA' folder.")
		os.Exit(0)
	}

	// handle openai -chatgpt flag
	if *chatgptFlag {
		handleChatGPT()
		fmt.Println("[ChatGPT] Done.")
		os.Exit(0)
	}

	// handle openai -whisper flag
	if *whisperFlag {
		handleWhisper()
		fmt.Println("[Whisper] Done.")
		os.Exit(0)
	}

}
