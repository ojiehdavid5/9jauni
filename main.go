package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type uniRequest struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	WebsiteLink  string `json:"website_link"`
}
type uniARequest struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	WebsiteLink  string `json:"website_link"`
}
type Abbreviation struct {
	Abbreviation string `json:"abbreviation"`
}
type errorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}


func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://nigerian-university-abbreviation-fo.vercel.app")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}


func apiError(w http.ResponseWriter, errResponse errorResponse) error {
	w.WriteHeader(errResponse.Code)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(errResponse); err != nil {
		return fmt.Errorf("failed to encode error response: %w", err)
	}

	return nil
}


func convertUniStruct() []uniRequest {

	uni, err := os.ReadFile("json/uni.json")

	if err != nil {
		log.Fatalf("an error occured: %v", err)
	}
	// fmt.Println(string(uni))

	var uniRequests []uniRequest
	if err := json.Unmarshal(uni, &uniRequests); err != nil {
		log.Fatalf("failed to unmarshal JSON: %v", err)
	}

	return uniRequests

}

func getAllUni(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		apiError(w, errorResponse{
			Message: "Invalid request method",
			Code:    http.StatusMethodNotAllowed,
		})
		return
	}
	result := convertUniStruct()

	// Set the response header to indicate JSON content
	w.Header().Set("Content-Type", "application/json")

	// Encode the result directly to the response writer
	if err := json.NewEncoder(w).Encode(result); err != nil {
		apiError(w, errorResponse{
			Message: "Failed to encode response",
			Code:    http.StatusInternalServerError,
		})
		return
	}
}

func getAUni(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		apiError(w, errorResponse{
			Message: "Invalid request method",
			Code:    http.StatusMethodNotAllowed,
		})
		return
	}
	decoder := json.NewDecoder(r.Body)

	var myUni uniARequest

	decoder = json.NewDecoder(r.Body)
	if err := decoder.Decode(&myUni); err != nil {
		apiError(w, errorResponse{
			Message: "Failed to decode request",
			Code:    http.StatusBadRequest,
		})
		return
	}
	fmt.Println(myUni)
	for _, uni := range convertUniStruct() {
		if uni.Name == myUni.Name {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(uni)
			return
		}
		// else{
		// 	http.Error(w, "University not found", http.StatusNotFound)
		// 	return
		// }

	}

}
func getAUniByAB(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		apiError(w, errorResponse{
			Message: "Method not allowed",
			Code:    http.StatusMethodNotAllowed,
		})
		return
	}

	// Get the abbreviation from the query parameters
	q := r.URL.Query()
	abbr := q.Get("abbreviation")
	fmt.Println(abbr)

	// Check if the abbreviation is provided
	if abbr == "" {
		apiError(w, errorResponse{
			Message: "Abbreviation is required",
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Retrieve the list of universities
	universities := convertUniStruct()

	// Find the university matching the abbreviation
	for _, uni := range universities {
		if uni.Abbreviation == abbr {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(uni)
			return
		}
	}

	// If not found, return a 404 error
	apiError(w, errorResponse{
		Message: "University not found",
		Code:    http.StatusNotFound,
	})
	return
}

func main() {
	convertUniStruct()
	fmt.Println("hello")
	http.HandleFunc("/", withCORS(getAllUni))
	http.HandleFunc("/search", withCORS(getAUni))
	http.HandleFunc("/searchab", withCORS(getAUniByAB))
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server started at :8080")
}
