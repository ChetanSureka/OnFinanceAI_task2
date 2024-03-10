package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct to represent the request body
type SearchRequest struct {
	Keyword string `json:"keyword"`
}

// Struct to represent the response body
type SearchResult struct {
	OriginalKeyword string   `json:"original_keyword"`
	NewKeywords     []string `json:"new_keywords"`
	MatchedChunks   []string `json:"matched_chunks"`
	Summary         string   `json:"summary"`
}

// Function to search documents for the given keyword and generate summary
func searchAndGenerateSummary(keyword string) SearchResult {
	// Example static data
	// Replace with actual document search and summarization logic
	newKeywords := []string{"new_keyword1", "new_keyword2"}
	matchedChunks := []string{"matched_chunk1", "matched_chunk2"}
	summary := "Generated summary based on the search results"

	return SearchResult{
		OriginalKeyword: keyword,
		NewKeywords:     newKeywords,
		MatchedChunks:   matchedChunks,
		Summary:         summary,
	}
}

// Handler function for the search endpoint
func searchHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure that only GET requests are accepted
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body
	var reqBody SearchRequest
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Check if keyword is provided in the request body
	if reqBody.Keyword == "" {
		http.Error(w, "Keyword is required", http.StatusBadRequest)
		return
	}

	// Call the searchAndGenerateSummary function to get search results
	searchResult := searchAndGenerateSummary(reqBody.Keyword)

	// Encode the response as JSON and write it to the response writer
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(searchResult)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func main() {
	// Register the searchHandler function to handle requests to the /search endpoint
	http.HandleFunc("/search", searchHandler)

	// Start the HTTP server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
