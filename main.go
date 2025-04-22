package main

// Author: Rayan Raghuram
// Cpyright @ 2025 Rayan Raghuram. All rights reserved.
//
// llm-client - Command Line Chat Client for LLM Server
//
// This Go program provides a command-line interface to interact with a locally hosted llama-server.
// Ref: https://github.com/rraghura102/llm-server. It reads user input from stdin, sends it to the 
// server at `http://localhost:60000/completion`, and streams the server's response back to the terminal. 
// A secure version is also provided.
//
// Functionality:
// - Interactive prompt using `bufio.Scanner`.
// - Maintains history of inputs and responses.
// - Sends JSON POST requests to a completion or secure endpoint.
// - Decodes streaming JSON responses until completion.
// - Gracefully handles server response and exit signal (Ctrl+D).
//
// Endpoints:
// - `/completion`: standard LLM chat completion.
// - `/secure/completion`: secure endpoint with an alternative request structure.
//
// Prerequisites:
// - Go (1.20+ recommended)
// - Running LLM server on `localhost:60000` supporting the `/completion` endpoint
//
// NOTE: This is a proof-of-concept and **not production-ready**. Error handling is minimal and hardcoded values are used.

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter CTRL+D to quit")

	var history []string
	for {

		fmt.Print("llm-client>>> ")
		if !scanner.Scan() {
			fmt.Println("Exiting chat")
			break
		}

		input := scanner.Text()
		respone := callServer(input)
		history = append(history, input)
		history = append(history, respone)
	}

	for _, input := range history {
		fmt.Println(input)
	}
}

func callServer(input string) string {

	fmt.Println("")

	request := NewChatRequest(input)
	requestBody, _ := json.Marshal(request)
	resp, _ := http.Post("http://localhost:60000/completion", "application/json", bytes.NewBuffer(requestBody))
	defer resp.Body.Close()

	var response string
	decoder := json.NewDecoder(resp.Body)
	for {

		var apiResponse ChatResponse
		err := decoder.Decode(&apiResponse)
		
		if err == io.EOF {
			break
		}
		
		if err != nil {
			fmt.Println("Error decoding JSON response")
		}

		fmt.Print(apiResponse.Content)
		response = response + apiResponse.Content

		if apiResponse.Stop || apiResponse.StoppedLimit {
			break
		}
	}

	fmt.Println("\n")
	return response
}

func callServerSecurely(input string) string {

	fmt.Println("")

	request := NewSecureChatRequest(input)
	requestBody, _ := json.Marshal(request)
	resp, _ := http.Post("http://localhost:60000/secure/completion", "application/json", bytes.NewBuffer(requestBody))
	defer resp.Body.Close()

	var response string
	decoder := json.NewDecoder(resp.Body)
	for {

		var apiResponse ChatResponse
		err := decoder.Decode(&apiResponse)
		
		if err == io.EOF {
			break
		}
		
		if err != nil {
			fmt.Println("Error decoding JSON response")
		}

		fmt.Print(apiResponse.Content)
		response = response + apiResponse.Content

		if apiResponse.Stop || apiResponse.StoppedLimit {
			break
		}
	}

	fmt.Println("\n")
	return response
}