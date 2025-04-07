package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"unicode"
)

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
	Format string `json:"format"`
}

// OllamaResponseStream is used for streaming responses, which Gemma doesn't support in the same way as llama2
type OllamaResponseStream struct {
	Model              string `json:"model"`
	Response           string `json:"response"`
	Done               bool   `json:"done"`
	Context            []int  `json:"context"`
	TotalDuration      int    `json:"total_duration"`
	LoadDuration       int    `json:"load_duration"`
	PromptEvalCount    int    `json:"prompt_eval_count"`
	PromptEvalDuration int    `json:"prompt_eval_duration"`
	EvalCount          int    `json:"eval_count"`
	EvalDuration       int    `json:"eval_duration"`
}

// TranslationResponse is the struct for the desired JSON output
type TranslationResponse struct {
	Translation string `json:"translation"`
}

// isChinese checks if a string contains Chinese characters.
func isChinese(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}

func main() {
	var (
		model     string
		ollamaURL string
		input     string
	)

	// Set the default model to gemma3:4b
	flag.StringVar(&model, "model", "gemma3:4b", "Ollama model name")
	flag.StringVar(&ollamaURL, "url", "http://localhost:11434/api/generate", "Ollama API URL")
	flag.StringVar(&input, "input", "", "Input string to translate")
	flag.Parse()

	if input == "" {
		fmt.Fprintf(os.Stderr, "Error: Input string is required. Use -input \"your text\"\n")
		flag.Usage()
		os.Exit(1)
	}

	// Determine the source language and construct the appropriate prompt
	var prompt string
	if isChinese(input) {
		prompt = fmt.Sprintf("Translate the following Chinese text to English concisely:\n%s", input)
	} else {
		prompt = fmt.Sprintf("Translate the following English text to Chinese concisely:\n%s", input)
	}

	requestBody := OllamaRequest{
		Model:  model,
		Prompt: prompt,
		Stream: false, // Set stream to false for Gemma
		Format: "json",
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	resp, err := http.Post(ollamaURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error sending request to Ollama: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Ollama request failed with status: %s\n", resp.Status)
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading response body: %v\n", err)
		os.Exit(1)
	}

	// Gemma returns a stream of JSON objects, even when stream is false.
	// We need to concatenate the "response" fields until "done" is true.
	var fullResponse strings.Builder
	decoder := json.NewDecoder(bytes.NewReader(body))
	for {
		var streamResponse OllamaResponseStream
		if err := decoder.Decode(&streamResponse); err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "Error decoding stream JSON: %v\n", err)
			os.Exit(1)
		}
		fullResponse.WriteString(streamResponse.Response)
		if streamResponse.Done {
			break
		}
	}

	// Remove leading and trailing whitespace from the response
	trimmedResponse := strings.TrimSpace(fullResponse.String())

	// Check if the response is a valid JSON string, if so, extract the translation
	// First, check if the response itself is a JSON string
	if strings.HasPrefix(trimmedResponse, "{") && strings.HasSuffix(trimmedResponse, "}") {
		var tempMap map[string]string
		if err := json.Unmarshal([]byte(trimmedResponse), &tempMap); err == nil {
			if val, ok := tempMap["translation"]; ok {
				trimmedResponse = val
			} else {
				// If there is no "translation" key, it may be a stringified JSON
				// Try to unmarshal the value of the "response" key
				if val, ok := tempMap["response"]; ok {
					var innerMap map[string]string
					if err := json.Unmarshal([]byte(val), &innerMap); err == nil {
						if innerVal, innerOk := innerMap["translation"]; innerOk {
							trimmedResponse = innerVal
						} else {
							trimmedResponse = val
						}
					} else {
						trimmedResponse = val
					}
				}
			}
		}
	}

	// Create the TranslationResponse struct
	translationResponse := TranslationResponse{
		Translation: trimmedResponse,
	}

	// Marshal the struct to JSON
	jsonOutput, err := json.Marshal(translationResponse)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON output: %v\n", err)
		os.Exit(1)
	}

	// Write the JSON output to stdout
	_, err = io.WriteString(os.Stdout, string(jsonOutput))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to output: %v\n", err)

		os.Exit(1)
	}
}
