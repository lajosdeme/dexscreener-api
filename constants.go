package dexscreenerapi

type Endpoint string

const base = "https://api.dexscreener.com/"
const latestTokenProfiles Endpoint = base + "token-profiles/latest/v1"
const latestBoostedTokens Endpoint = base + "token-boosts/latest/v1"
const tokensWithMostActiveBoosts Endpoint = base + "token-boosts/top/v1"
const ordersPaidForToken Endpoint = base + "orders/v1/"
const pairsByChainAndPairId Endpoint = base + "latest/dex/pairs/"
const pairsByTokenAddress Endpoint = base + "latest/dex/tokens/"
const searchPairs Endpoint = base + "latest/dex/search"

const GET = "GET"
