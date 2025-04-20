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
type Abbreviation struct{
	Abbreviation string `json:"abbreviation"`

}

func convertUniStruct() []uniRequest {

	uni, err := os.ReadFile("/Users/user/go/src/9jauni/json/uni.json")
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
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	result := convertUniStruct()

	// Set the response header to indicate JSON content
	w.Header().Set("Content-Type", "application/json")

	// Encode the result directly to the response writer
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getAUni(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	decoder := json.NewDecoder(r.Body)

	var myUni uniARequest

	decoder = json.NewDecoder(r.Body)
	if err := decoder.Decode(&myUni); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	decoder := json.NewDecoder(r.Body)

	var myUni uniARequest

	decoder = json.NewDecoder(r.Body)
	if err := decoder.Decode(&myUni); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(myUni)
	for _, uni := range convertUniStruct() {
		if uni.Abbreviation == myUni.Abbreviation{
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

func main() {
	convertUniStruct()
	fmt.Println("hello")
	http.HandleFunc("/", getAllUni)
	http.HandleFunc("/search", getAUni)
	http.HandleFunc("/searchab", getAUniByAB)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server started at :8080")
}
