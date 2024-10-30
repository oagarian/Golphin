package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golphin/src/utils/environment"
	"io"
	"net/http"
	"os"
)

func GenerateContent(prompt string) (string, error){
    environmentData, err := environment.Get()
	if err!= nil {
        return "", err
    }

    url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash-latest:generateContent?key=" + environmentData.GeminiToken
    requestBody, _ := json.Marshal(map[string]interface{}{
        "contents": []map[string]interface{}{
            {
                "parts": []map[string]string{
                    {"text": prompt},
                },
            },
        },
    })

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
    if err != nil {
        fmt.Println("Error: ", err)
        os.Exit(1)
    }
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error: ", err)
        os.Exit(1)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Printf("Status code: %d\n", resp.StatusCode)
        body, _ := io.ReadAll(resp.Body)
        fmt.Println("Error: ", string(body))
        os.Exit(1)
    }

    var result map[string]interface{}
    err = json.NewDecoder(resp.Body).Decode(&result)
    if err != nil {
        fmt.Println("Error: ", err)
        os.Exit(1)
    }

	if candidates, ok := result["candidates"].([]interface{}); ok && len(candidates) > 0 {
        candidate := candidates[0].(map[string]interface{})
        if content, ok := candidate["content"].(map[string]interface{}); ok {
            if parts, ok := content["parts"].([]interface{}); ok && len(parts) > 0 {
                part := parts[0].(map[string]interface{})
                if text, ok := part["text"].(string); ok {
                    return text, nil
                }
            }
        }
    }
	return "", nil
}
