package handlers

import (
	"github.com/dl-nft-books/tracker-svc/internal/service/api/requests"
	"github.com/dl-nft-books/tracker-svc/resources"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
	"strconv"
	"strings"
)

func GetStatistics(w http.ResponseWriter, r *http.Request) {
	var response resources.StatisticsResponse
	response.Data.Type = resources.STATISTICS
	request, err := requests.NewGetStatisticsRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	statistics, err := DB(r).KeyValue().New().Select([]string{
		// amount_pie_chart
		"stats-book-%-amount",
		"stats-book-amount-total",

		// price_pie_chart
		"stats-token_symbol-%-price_token",
		"stats-token_symbol-%-price_usd",
		"stats-price_usd",

		// chain_pie_chart
		"stats-chain_id-%",
	}, []string{})
	if err != nil {
		Log(r).WithError(err).Error("failed to get statistics")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	for _, stats := range statistics {
		switch {
		case strings.HasPrefix(stats.Key, "stats-book-"):
			id := getBookID(stats.Key)
			if id != 0 {
				book, err := Booker(r).GetBookById(id)
				if err != nil {
					Log(r).WithError(err).Error("failed to get book")
					ape.RenderErr(w, problems.InternalError())
					return
				}
				response.Data.Attributes.AmountPieChart.Attributes.Books =
					append(response.Data.Attributes.AmountPieChart.Attributes.Books, resources.AmountPieChartBooks{
						Key: resources.Key{Type: resources.AMOUNT_PIE_CHART},
						Attributes: resources.AmountPieChartBooksAttributes{
							Amount:  cast.ToInt64(stats.Value),
							Address: book.Data.Attributes.Networks[0].Attributes.ContractAddress,
						},
					})
				break
			}
			response.Data.Attributes.AmountPieChart.Attributes.Total = cast.ToInt64(stats.Value)
		case strings.HasPrefix(stats.Key, "stats-token_symbol-"):
			tokenSymbol := getTokenSymbol(stats.Key)
			if strings.Index(stats.Key, "usd") != -1 {
				response = setUsdPrice(response, tokenSymbol, cast.ToFloat64(stats.Value))
				response.Data.Attributes.TokenList = append(response.Data.Attributes.TokenList, tokenSymbol)
				break
			}
			response = setTokenPrice(response, tokenSymbol, cast.ToFloat64(stats.Value))
			break
		case strings.HasPrefix(stats.Key, "stats-price_usd"):
			response.Data.Attributes.PricePieChart.Attributes.Total = cast.ToFloat64(stats.Value)
			break
		case strings.HasPrefix(stats.Key, "stats-chain_id-"):
			id := getChainID(stats.Key)
			if id != 0 {
				if err != nil {
					Log(r).WithError(err).Error("failed to get book")
					ape.RenderErr(w, problems.InternalError())
					return
				}
				response.Data.Attributes.ChainPieChart.Attributes.Chains = append(
					response.Data.Attributes.ChainPieChart.Attributes.Chains,
					resources.ChainPieChartChains{
						Key: resources.Key{Type: resources.CHAIN_PIE_CHART},
						Attributes: resources.ChainPieChartChainsAttributes{
							Amount:  cast.ToInt64(stats.Value),
							ChainId: id,
						},
					})
				break
			}
			response.Data.Attributes.ChainPieChart.Attributes.Total = cast.ToInt64(stats.Value)
			break
		}
	}
	nftPayments, err := DB(r).Payments().Page(request.NFT, request.Sort).FilterByType(int8(resources.NFT)).Select()
	for _, nftPayment := range nftPayments {
		response.Data.Attributes.NftList = append(response.Data.Attributes.NftList, resources.NftListItem{
			Key: resources.Key{
				ID:   strconv.Itoa(int(nftPayment.NftId)),
				Type: resources.NFT_LIST,
			},
			Attributes: resources.NftListItemAttributes{
				Address: nftPayment.TokenAddress,
			},
		})
	}
	ape.Render(w, response)
}

func getBookID(key string) int64 {
	// Find the first and second occurrence of "-" in the key
	firstIndex := strings.Index(key, "book-") + len("book-")
	secondIndex := strings.Index(key, "-amount") - 1
	bookID, _ := strconv.Atoi(key[firstIndex:secondIndex])
	return int64(bookID)
}

func getChainID(key string) int64 {
	// Find the first and second occurrence of "-" in the key
	firstIndex := strings.Index(key, "chain_id-") + len("chain_id-")
	chainID, _ := strconv.Atoi(key[firstIndex:])
	return int64(chainID)
}

func getTokenSymbol(key string) string {
	// Find the first and second occurrence of "-" in the key
	firstIndex := strings.Index(key, "-token_symbol-") + len("-token_symbol-")
	secondIndex := strings.Index(key, "-price")
	return key[firstIndex:secondIndex]
}

func setTokenPrice(response resources.StatisticsResponse, tokenSymbol string, price float64) resources.StatisticsResponse {
	for i, token := range response.Data.Attributes.PricePieChart.Attributes.Tokens {
		if token.Attributes.Name == tokenSymbol {
			response.Data.Attributes.PricePieChart.Attributes.Tokens[i].Attributes.NativeCurrency = price
			return response
		}
	}
	response.Data.Attributes.PricePieChart.Attributes.Tokens = append(response.Data.Attributes.PricePieChart.Attributes.Tokens, resources.PricePieChartTokens{
		Key: resources.Key{Type: resources.TOKENS},
		Attributes: resources.PricePieChartTokensAttributes{
			Name:           tokenSymbol,
			NativeCurrency: price,
		},
	})
	return response
}
func setUsdPrice(response resources.StatisticsResponse, tokenSymbol string, price float64) resources.StatisticsResponse {
	for i, token := range response.Data.Attributes.PricePieChart.Attributes.Tokens {
		if token.Attributes.Name == tokenSymbol {
			response.Data.Attributes.PricePieChart.Attributes.Tokens[i].Attributes.Usd = price
			return response
		}
	}
	response.Data.Attributes.PricePieChart.Attributes.Tokens = append(response.Data.Attributes.PricePieChart.Attributes.Tokens, resources.PricePieChartTokens{
		Key: resources.Key{Type: resources.PRICE_PIE_CHART},
		Attributes: resources.PricePieChartTokensAttributes{
			Name: tokenSymbol,
			Usd:  price,
		},
	})
	return response
}
