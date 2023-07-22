package youtube

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/sugar-cat7/vspo-common-api/constants"
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/util"
)

// API is a YouTube API client.
type API struct {
	Client *http.Client
}

// NewAPI creates a new YouTube API client.
func NewAPI(client *http.Client) *API {
	if client == nil {
		client = http.DefaultClient
	}
	return &API{Client: client}
}

// GetVideos returns a slice of YoutubeVideoListResponses.
func (api *API) GetVideos(videoIDs []string) ([]entities.YTVideoListResponse, error) {
	// Define a slice to hold all YoutubeVideoListResponses
	var data []entities.YTVideoListResponse

	// Define the number of videoIDs per request
	chunkSize := 50

	// Use the Chunk function to split the videoIDs slice into chunks
	videoIDChunks, err := util.Chunk(videoIDs, chunkSize)
	if err != nil {
		return nil, fmt.Errorf("error splitting videoIDs into chunks: %v", err)
	}

	// Loop through the chunks
	for _, chunk := range videoIDChunks {
		// Join the current chunk of videoIDs with commas
		videoIDsJoined := strings.Join(chunk, ",")

		// Create the request URL
		requestURL := "https://www.googleapis.com/youtube/v3/videos?id=" + videoIDsJoined + "&part=snippet,statistics&key=" + os.Getenv("YOUTUBE_API_KEY")

		// Make the request
		resp, err := api.Client.Get(requestURL)
		if err != nil {
			return nil, fmt.Errorf("error making HTTP request to YouTube API: %v", err)
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body: %v", err)
		}

		// Define a variable to hold the current YoutubeVideoListResponse
		var currentData entities.YTVideoListResponse

		// Unmarshal the JSON response into currentData
		err = json.Unmarshal(body, &currentData)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling JSON response: %v", err)
		}

		// Add currentData to the data slice
		data = append(data, currentData)
	}

	// Return the data slice and nil error
	return data, nil
}

// GetPlaylists returns a slice of YoutubePlaylistResponses.
func (api *API) GetPlaylists() ([]entities.YTYouTubePlaylistResponse, error) {
	var responses []entities.YTYouTubePlaylistResponse

	for _, playlistID := range constants.YTPlaylistIDs {
		var data entities.YTYouTubePlaylistResponse
		requestURL := "https://www.googleapis.com/youtube/v3/playlistItems?playlistId=" + playlistID + "&key=" + os.Getenv("YOUTUBE_API_KEY") + "&part=snippet&maxResults=50"

		resp, err := api.Client.Get(requestURL)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(body, &data)
		if err != nil {
			return nil, err
		}

		responses = append(responses, data)
	}

	return responses, nil
}

// GetChannels returns a slice of channels.
func (api *API) GetChannels(channelIDs []string) ([]entities.Channel, error) {
	var data []entities.Channel

	chunkSize := 50

	channelIDChunks, err := util.Chunk(channelIDs, chunkSize)
	if err != nil {
		return nil, fmt.Errorf("error splitting channelIDs into chunks: %v", err)
	}

	for _, chunk := range channelIDChunks {
		channelIDsJoined := strings.Join(chunk, ",")

		requestURL := "https://www.googleapis.com/youtube/v3/channels?part=snippet,statistics&id=" + channelIDsJoined + "&key=" + os.Getenv("YOUTUBE_API_KEY")

		resp, err := api.Client.Get(requestURL)
		if err != nil {
			return nil, fmt.Errorf("error making HTTP request to YouTube API: %v", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body: %v", err)
		}

		var currentData struct {
			Items []entities.Channel `json:"items"`
		}

		err = json.Unmarshal(body, &currentData)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling JSON response: %v", err)
		}

		data = append(data, currentData.Items...)
	}

	return data, nil
}
