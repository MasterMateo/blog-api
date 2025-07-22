package blog

import (
	"encoding/json"
	"log"
	"os"
)

type PostsData struct {
	Posts []Post `json:"posts"`
}

func LoadPostsFromJSON(path string) []Post {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Error opening JSON file: %v", err)
		return nil
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Printf("Error closing JSON file: %v", cerr)
		}
	}()

	var data PostsData
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		return nil
	}

	return data.Posts
}
