# GPT-ANMA

> GPT-ANMA is designed to comply with OpenAI's [usage policies](https://openai.com/policies/usage-policies). We will not be hold accountable for any modifications made to the code by users.

## How to compile the code

Download and install [The Go Programming Language](https://go.dev/).
Then, simply compile the code with the following commands:

```
cd ./GPT-ANMA/
go mod tidy
go build
```

## How to use the CLI

### Getting started

Copy the new compiled executable to a new empty folder, and open a new console terminal there, then use the following command to see all the available flags:

```
$ GPT-ANMA -h
```

Now, you can generate all the required files by the CLI by using the following command:

```
$ GPT-ANMA -presets
```

> A new `GPT-ANMA` directory will be created with all the required files in it. In case some of the files are missing, you can always regenerate them by using the `-presets` flag.

Inside the `GPT-ANMA` directory, find a file named `your-openai-apikey-here.txt`, fill it with your OpenAI's API key, and then save the file.

### Sending questions to OpenAI's ChatGPT

You might want to explore all the available text files and see what's inside them.

The `your-ai-base-personality.txt` file is used to tell ChatGPT how do you want it to response to your questions.

The `your-question-here.txt` file is used to store your questions before sending them to OpenAI's APIs.

The `your-openai-last-usage.txt` file is used to store your OpenAI's last usage information.

These files are used to store the old and new answers, with `answer dialogue` and `answer flag` separated to different files to make it easier for you to get the values from other application (e.g. if you want to integrate it with Unity):

- `your-new-answer-here.txt`
- `your-new-answer-flag-here.txt`
- `your-old-answer-here.txt`
- `your-old-answer-flag-here.txt`
