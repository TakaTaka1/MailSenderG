package main

import (
	// "net/http"
    // "github.com/gin-gonic/gin"      
    // "encoding/json"
    // "log"
    "github.com/joho/godotenv"
    "fmt"
    "os"    
)

func main() {
	err := godotenv.Load(".env")
		
	if err != nil {
		fmt.Printf("Couldn't load : %v", err)
	} 
	
    key := os.Getenv("ENV")
    fmt.Println(key)
}