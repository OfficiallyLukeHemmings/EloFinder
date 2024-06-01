package elo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// TODO: check type structures
type FaceitResponse struct {
	ActivatedAt        string            `json:"activated_at"`
	Avatar             string            `json:"avatar"`
	Country            string            `json:"country"`
	CoverFeaturedImage string            `json:"cover_featured_image"`
	CoverImage         string            `json:"cover_image"`
	FaceitURL          string            `json:"faceit_url"`
	FriendsIDs         []string          `json:"friends_ids"`
	Games              map[string]Game   `json:"games"`
	Infractions        map[string]string `json:"infractions"`
	MembershipType     string            `json:"membership_type"`
	Memberships        []string          `json:"memberships"`
	NewSteamID         string            `json:"new_steam_id"`
	Nickname           string            `json:"nickname"`
	Platforms          map[string]string `json:"platforms"`
	PlayerID           string            `json:"player_id"`
	Settings           map[string]string `json:"settings"`
	SteamID64          string            `json:"steam_id_64"`
	SteamNickname      string            `json:"steam_nickname"`
	Verified           bool              `json:"verified"`
}

// TODO: check type structures
type Game struct {
	FaceitElo       int               `json:"faceit_elo"`
	GamePlayerID    string            `json:"game_player_id"`
	GamePlayerName  string            `json:"game_player_name"`
	GameProfileID   string            `json:"game_profile_id"`
	Region          string            `json:"region"`
	Regions         map[string]string `json:"regions"`
	SkillLevel      int               `json:"skill_level"`
	SkillLevelLabel string            `json:"skill_level_label"`
}

func GetElo(player string) int {
	// Structuring the Request
	apiKey := GetAPIKey()
	url := "https://open.faceit.com/data/v4/players?nickname=Xais_"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Setting headers (api key)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	// Making the HTTP call
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	// Default elo value
	elo := -1
	// Handling the response
	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Error reading response body: %v", err)
		}

		// Deconstructing the response
		var data FaceitResponse
		if err := json.Unmarshal(body, &data); err != nil {
			log.Fatalf("Error unmarshaling JSON: %v", err)
		}
		// Extracting elo from cs2 game
		// (note - games referring to [cs2, csgo]
		cs2Data := data.Games["cs2"]
		if cs2Data.FaceitElo != 0 {
			elo = cs2Data.FaceitElo
		}
	}

	return elo
}
