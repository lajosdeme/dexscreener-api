# DexScreener Go API

A Go library for interacting with the [DexScreener API](https://docs.dexscreener.com/api/reference). This library provides a simple and intuitive way to fetch token profiles, pair information, and search across DexScreener's supported DEXes.

## Installation

```bash
go get github.com/lajosdeme/dexscreener-api
```

## Quick Start

```go
package main

import (
    "fmt"
    api "github.com/lajosdeme/dexscreener-api"
)

func main() {
    // Get latest token profiles
    profiles, err := api.GetLatestTokenProfiles()
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Found %d token profiles\n", len(profiles))
}
```

## Features

- Full coverage of DexScreener API endpoints
- Proper error handling
- Easy to use interface
- Fully type-safe responses

## API Methods

### Get Latest Token Profiles
Fetches the most recently updated token profiles.

```go
profiles, err := GetLatestTokenProfiles() ([]TokenProfile, error)
```

### Get Latest Boosted Token Profiles
Retrieves the most recently boosted token profiles.

```go
profiles, err := GetLatestBoostedTokenProfiles() ([]BoostedTokenProfile, error)
```

### Get Tokens with Most Active Boosts
Fetches tokens sorted by number of active boosts.

```go
profiles, err := GetTokensWithMostActiveBoosts() ([]BoostedTokenProfile, error)
```

### Check Orders Paid
Checks if orders are paid for a specific token on a chain.

```go
response, err := CheckOrdersPaid(chainId string, tokenAddress string) (*OrderPaidResponse, error)
```

### Get Pairs by Chain and Pair ID
Retrieves specific trading pairs by chain ID and pair ID.

```go
pairs, err := GetPairsByChainAndPairId(chainId string, pairId string) (*PairsResponse, error)
```

### Get Token Pairs
Fetches pair information for multiple token addresses (up to 30).

```go
pairs, err := GetTokenPairs(tokenAddresses []string) (*PairsResponse, error)
```

### Search Pairs
Searches for pairs across all supported DEXes.

```go
pairs, err := SearchPairs(searchText string) (*PairsResponse, error)
```

## Example Usage

### Searching for Pairs

```go
package main

import (
    "fmt"
    api "github.com/lajosdeme/dexscreener-api"
)

func main() {
    // Search for USDC pairs
    pairs, err := api.SearchPairs("SOL/USDC")
    if err != nil {
        panic(err)
    }
    
    // Print pair information
    for _, pair := range pairs.Pairs {
        fmt.Printf("Chain: %s\n", pair.ChainId)
        fmt.Printf("DEX: %s\n", pair.DexId)
        fmt.Printf("Price USD: %s\n", pair.PriceUsd)
        fmt.Printf("Liquidity USD: %.2f\n", pair.Liquidity.USD)
        fmt.Printf("---\n")
    }
}
```

### Getting Token Pairs

```go
package main

import (
    "fmt"
    api "github.com/lajosdeme/dexscreener-api"
)

func main() {
    // Token addresses to check
    addresses := []string{
        "JUPyiwrYJFskUPiHa7hkeR8VUtAeFoSYbKedZNsDvCN",
        "7GCihgDB8fe6KNjn2MYtkzZcRjQy3t9GHdC8uHYmW2hr",
    }
    
    // Get pair information
    response, err := api.GetTokenPairs(addresses)
    if err != nil {
        panic(err)
    }
    
    // Print pair information
    for _, pair := range response.Pairs {
        fmt.Printf("Base Token: %s (%s)\n", pair.BaseToken.Name, pair.BaseToken.Symbol)
        fmt.Printf("Quote Token: %s (%s)\n", pair.QuoteToken.Name, pair.QuoteToken.Symbol)
        fmt.Printf("Price: $%s\n", pair.PriceUsd)
        fmt.Printf("---\n")
    }
}
```

## Rate Limiting

DexScreener API has rate limits. This library does not automatically handle rate limiting, but it does provide clear error messages when you hit rate limits. It's recommended to implement your own rate limiting strategy based on your needs.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This library is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Official Documentation

For more details about the API endpoints and responses, please refer to the [official DexScreener API documentation](https://docs.dexscreener.com/api/reference).