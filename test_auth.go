package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	baseURL := "http://localhost:3000/api"

	// Test Register
	fmt.Println("=== Testing Register ===")
	registerData := map[string]string{
		"name":     "Test User",
		"email":    "test@example.com",
		"password": "fiber123",
	}

	registerJSON, _ := json.Marshal(registerData)
	resp, err := http.Post(baseURL+"/register", "application/json", bytes.NewBuffer(registerJSON))
	if err != nil {
		fmt.Println("Register Error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Register Response (%d): %s\n\n", resp.StatusCode, string(body))

	// Test Login
	fmt.Println("=== Testing Login ===")
	loginData := map[string]string{
		"email":    "test@example.com",
		"password": "fiber123",
	}

	loginJSON, _ := json.Marshal(loginData)
	resp, err = http.Post(baseURL+"/login", "application/json", bytes.NewBuffer(loginJSON))
	if err != nil {
		fmt.Println("Login Error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	fmt.Printf("Login Response (%d): %s\n\n", resp.StatusCode, string(body))

	// Extract token for profile test
	var loginResp map[string]interface{}
	json.Unmarshal(body, &loginResp)

	if token, ok := loginResp["token"].(string); ok && resp.StatusCode == 200 {
		// Test Profile
		fmt.Println("=== Testing Profile ===")
		req, _ := http.NewRequest("GET", baseURL+"/profile", nil)
		req.Header.Set("Authorization", token)

		client := &http.Client{}
		resp, err = client.Do(req)
		if err != nil {
			fmt.Println("Profile Error:", err)
			return
		}
		defer resp.Body.Close()

		body, _ = io.ReadAll(resp.Body)
		fmt.Printf("Profile Response (%d): %s\n", resp.StatusCode, string(body))
	}
}
