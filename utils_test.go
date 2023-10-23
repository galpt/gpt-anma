package main

import (
	"reflect"
	"testing"

	openai "github.com/sashabaranov/go-openai"
)

func Test_getOpenAIAPIkey(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getOpenAIAPIkey()
		})
	}
}

func Test_getUserIP(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getUserIP(); got != tt.want {
				t.Errorf("getUserIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_randFloats(t *testing.T) {
	type args struct {
		min float64
		max float64
		n   int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randFloats(tt.args.min, tt.args.max, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("randFloats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createFileWriteString(t *testing.T) {
	type args struct {
		filename string
		path     string
		content  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createFileWriteString(tt.args.filename, tt.args.path, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("createFileWriteString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_readFileToStr(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readFileToStr(tt.args.path); got != tt.want {
				t.Errorf("readFileToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_openaiModeration(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := openaiModeration(tt.args.input); got != tt.want {
				t.Errorf("openaiModeration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_openaiSaveChatGPTResp(t *testing.T) {
	type args struct {
		cgptResp openai.ChatCompletionResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := openaiSaveChatGPTResp(tt.args.cgptResp); (err != nil) != tt.wantErr {
				t.Errorf("openaiSaveChatGPTResp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_openaiLastUsageInfo(t *testing.T) {
	type args struct {
		usrQuestion string
		cgptResp    openai.ChatCompletionResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := openaiLastUsageInfo(tt.args.usrQuestion, tt.args.cgptResp); (err != nil) != tt.wantErr {
				t.Errorf("openaiLastUsageInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
