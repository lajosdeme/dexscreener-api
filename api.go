package dexscreenerapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// GetLatestTokenProfiles fetches the latest token profiles
func GetLatestTokenProfiles() ([]TokenProfile, error) {
	client := &http.Client{}
	req, err := http.NewRequest(GET, string(latestTokenProfiles), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var profiles []TokenProfile
	if err := json.Unmarshal(body, &profiles); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return profiles, nil
}

// GetLatestTokenProfiles fetches the latest token profiles
func GetLatestBoostedTokenProfiles() ([]BoostedTokenProfile, error) {
	client := &http.Client{}
	req, err := http.NewRequest(GET, string(latestBoostedTokens), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var profiles []BoostedTokenProfile
	if err := json.Unmarshal(body, &profiles); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return profiles, nil
}

func GetTokensWithMostActiveBoosts() ([]BoostedTokenProfile, error) {
	client := &http.Client{}
	req, err := http.NewRequest(GET, string(tokensWithMostActiveBoosts), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var profiles []BoostedTokenProfile
	if err := json.Unmarshal(body, &profiles); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return profiles, nil
}

// GetTokenPairs fetches pair information for the given token addresses
func GetTokenPairs(tokenAddresses []string) (*PairsResponse, error) {
	// Join addresses with commas, limiting to 30 addresses per request
	if len(tokenAddresses) > 30 {
		tokenAddresses = tokenAddresses[:30]
	}
	addressesString := strings.Join(tokenAddresses, ",")

	client := &http.Client{}
	url := fmt.Sprintf("%s%s", pairsByTokenAddress, addressesString)

	req, err := http.NewRequest(GET, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var pairsResponse PairsResponse
	if err := json.Unmarshal(body, &pairsResponse); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return &pairsResponse, nil
}
