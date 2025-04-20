package main

import (
	"fmt"
	"os"
	"net/http"
	"log"
	"encoding/json"
)

type uniRequest struct {
	short string `json:"abbreviation"`
}

func convertUniStruct() []uniRequest {

	uni,err:= os.ReadFile("/Users/user/go/src/9jauni/json/uni.json")
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

}

func getAUni(w http.ResponseWriter, r *http.Request) {

}

func main() {
	convertUniStruct();
	fmt.Println("hello")
	http.HandleFunc("/",getAllUni)
	http.HandleFunc("/search",getAUni)
	http.ListenAndServe(":8080", nil);
	fmt.Println("Server started at :8080")
}
