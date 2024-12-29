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
	req, err := http.NewRequest(GET, string(latestTokenProfiles), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	body, err := sendGetRequest(req)
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
	req, err := http.NewRequest(GET, string(latestBoostedTokens), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	body, err := sendGetRequest(req)
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
	req, err := http.NewRequest(GET, string(tokensWithMostActiveBoosts), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	body, err := sendGetRequest(req)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var profiles []BoostedTokenProfile
	if err := json.Unmarshal(body, &profiles); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return profiles, nil
}

func CheckOrdersPaid(chainId string, tokenAddress string) (*OrderPaidResponse, error) {
	url := fmt.Sprintf("%s%s/%s", ordersPaidForToken, chainId, tokenAddress)

	req, err := http.NewRequest(GET, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	body, err := sendGetRequest(req)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var orderPaidResponse OrderPaidResponse
	if err := json.Unmarshal(body, &orderPaidResponse); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return &orderPaidResponse, nil
}

func GetPairsByChainAndPairId(chainId string, pairId string) (*PairsResponse, error) {
	url := fmt.Sprintf("%s%s/%s", pairsByChainAndPairId, chainId, pairId)

	req, err := http.NewRequest(GET, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	body, err := sendGetRequest(req)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var pairsResponse PairsResponse
	if err := json.Unmarshal(body, &pairsResponse); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return &pairsResponse, nil
}

// GetTokenPairs fetches pair information for the given token addresses
func GetTokenPairs(tokenAddresses []string) (*PairsResponse, error) {
	// Join addresses with commas, limiting to 30 addresses per request
	if len(tokenAddresses) > 30 {
		tokenAddresses = tokenAddresses[:30]
	}
	addressesString := strings.Join(tokenAddresses, ",")

	url := fmt.Sprintf("%s%s", pairsByTokenAddress, addressesString)

	req, err := http.NewRequest(GET, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	body, err := sendGetRequest(req)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var pairsResponse PairsResponse
	if err := json.Unmarshal(body, &pairsResponse); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return &pairsResponse, nil
}

func SearchPairs(searchText string) (*PairsResponse, error) {
	url := fmt.Sprintf("%s?q=%s", searchPairs, searchText)

	req, err := http.NewRequest(GET, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	body, err := sendGetRequest(req)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var pairsResponse PairsResponse
	if err := json.Unmarshal(body, &pairsResponse); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return &pairsResponse, nil
}

func sendGetRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
