package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// Load .env file, ignore error if not found (for production environments)
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or failed to load")
	}
}

// ChatRequest represents the structure of incoming chat requests
type ChatRequest struct {
	Message string `json:"message"`
	Context string `json:"context"`
}

// GeminiRequest represents the structure of requests to Gemini API
type GeminiRequest struct {
	Contents []GeminiContent `json:"contents"`
}

// GeminiContent represents a message in the Gemini API request
type GeminiContent struct {
	Parts []GeminiPart `json:"parts"`
	Role  string       `json:"role"`
}

// GeminiPart represents a part of a message
type GeminiPart struct {
	Text string `json:"text"`
}

// GeminiResponse represents the structure of responses from Gemini API
type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

// ChatResponse is the response sent back to the client
type ChatResponse struct {
	Response string `json:"response"`
}

// ChatHandler handles chat requests and communicates with Gemini API
func ChatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var chatReq ChatRequest
	if err := json.NewDecoder(r.Body).Decode(&chatReq); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Println("GEMINI_API_KEY environment variable not set")
		http.Error(w, "Server configuration error", http.StatusInternalServerError)
		return
	}

	systemPrompt := `You are a helpful assistant for a security service company.
Your job is to help users learn about and potentially book security services.

Available services include:
- Club Guards: Security personnel specialized for nightclubs and entertainment venues
- Event Security: Guards for special events, concerts, and gatherings
- Personal Security: Bodyguards and personal protection services
- Property Guards: Security for residential and commercial properties
- Corporate Security: Comprehensive security solutions for businesses

Pricing information:
- Base cost per guard: ₹1000
- Service charge: ₹1000
- GST: 18%
- Optional add-ons:
  - Camera surveillance: ₹500
  - Security vehicle: ₹2500
  - First aid training: ₹150
  - Walkie-talkie equipment: ₹500
  - Bulletproof vests: ₹2000
  - Fire safety training: ₹750

Keep responses brief, friendly, and professional. If users express interest in booking, suggest they start the booking process.`

	geminiReq := GeminiRequest{
		Contents: []GeminiContent{
			{
				Parts: []GeminiPart{
					{Text: systemPrompt},
				},
				Role: "model",
			},
			{
				Parts: []GeminiPart{
					{Text: chatReq.Context},
				},
				Role: "user",
			},
			{
				Parts: []GeminiPart{
					{Text: chatReq.Message},
				},
				Role: "user",
			},
		},
	}

	reqBody, err := json.Marshal(geminiReq)
	if err != nil {
		log.Printf("Error marshaling request: %v", err)
		http.Error(w, "Error preparing request", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("POST", "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent?key="+apiKey, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		http.Error(w, "Error preparing request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request to Gemini API: %v", err)
		http.Error(w, "Error communicating with AI service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		http.Error(w, "Error processing AI response", http.StatusInternalServerError)
		return
	}

	var geminiResp GeminiResponse
	if err := json.Unmarshal(respBody, &geminiResp); err != nil {
		log.Printf("Error unmarshaling response: %v", err)
		http.Error(w, "Error processing AI response", http.StatusInternalServerError)
		return
	}

	var responseText string
	if len(geminiResp.Candidates) > 0 && len(geminiResp.Candidates[0].Content.Parts) > 0 {
		responseText = geminiResp.Candidates[0].Content.Parts[0].Text
	} else {
		responseText = "I'm sorry, I couldn't process your request. Please try again."
	}

	chatResp := ChatResponse{Response: responseText}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chatResp)
}
