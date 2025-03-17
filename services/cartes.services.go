package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type YugiohCard struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Level     int    `json:"level"`
	ATK       int    `json:"atk"`
	DEF       int    `json:"def"`
	Race      string `json:"race"`
	Attribute string `json:"attribute"`
	Images    []struct {
		Url string `json:"image_url_small"`
	} `json:"card_images"`
}

type YugiohResponse struct {
	Data []YugiohCard `json:"data"`
}

func GetYugiohData() (YugiohResponse, int, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, "https://db.ygoprodeck.com/api/v7/cardinfo.php", nil)
	if err != nil {
		return YugiohResponse{}, http.StatusInternalServerError, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return YugiohResponse{}, http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return YugiohResponse{}, resp.StatusCode, errors.New("la requête a échoué avec le code : " + resp.Status)
	}

	var yugiohResponse YugiohResponse
	if err := json.NewDecoder(resp.Body).Decode(&yugiohResponse); err != nil {
		return YugiohResponse{}, http.StatusInternalServerError, err
	}

	return yugiohResponse, http.StatusOK, nil
}
