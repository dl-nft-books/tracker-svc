package handlers

import (
	"fmt"
	"github.com/dl-nft-books/tracker-svc/internal/service/api/requests"
	"github.com/dl-nft-books/tracker-svc/resources"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
	"strconv"
	"strings"
)

func GetStatisticsByBook(w http.ResponseWriter, r *http.Request) {
	var (
		response resources.StatisticsByBookResponse
	)
	request, err := requests.NewGetStatisticsByBookRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	statistics, err := DB(r).KeyValue().New().Select([]string{
		// chain_pie_chart
		fmt.Sprintf("stats-book-%d-chain_id-%", request.BookId),

		// date graph
		fmt.Sprintf("stats-book-%d-amount-date-%", request.BookId),

		// token histogram
		fmt.Sprintf("stats-book-%d-token_symbol-%%-price_token", request.BookId),
		fmt.Sprintf("stats-book-%d-token_symbol-%%-price_usd", request.BookId),
		fmt.Sprintf("stats-book-%d-price_usd", request.BookId),
	}, []string{})
	if err != nil {
		Log(r).WithError(err).Error("failed to get statistics")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	fmt.Println(statistics)
	for _, stats := range statistics {
		switch {
		case strings.HasPrefix(stats.Key, fmt.Sprintf("stats-book-%d-token_symbol", request.BookId)):
			tokenSymbol := getTokenSymbol(stats.Key)
			if strings.Index(stats.Key, "usd") != -1 {
				response = setUsdPriceByBook(response, tokenSymbol, cast.ToFloat64(stats.Value))
				break
			}
			response = setTokenPriceByBook(response, tokenSymbol, cast.ToFloat64(stats.Value))
			break
		case strings.HasPrefix(stats.Key, fmt.Sprintf("stats-book-%d-price_usd", request.BookId)):
			response.Data.Attributes.TokensHistogram.Attributes.Total = cast.ToFloat64(stats.Value)
			break
		case strings.HasPrefix(stats.Key, fmt.Sprintf("stats-book-%d-amount-date-%", request.BookId)):
			date := getDate(stats.Key)
			response.Data.Attributes.DateGraph = append(response.Data.Attributes.DateGraph, resources.DateGraph{
				Key: resources.Key{Type: resources.DATE_GRAPH},
				Attributes: resources.DateGraphAttributes{
					Amount: cast.ToInt64(stats.Value),
					Date:   date,
				},
			})
			break
		case strings.HasPrefix(stats.Key, fmt.Sprintf("stats-book-%d-chain_id-%", request.BookId)):
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
	nftPayments, err := DB(r).NftPayments().Page(request.NFT, request.Sort).Select()
	for _, nftPayment := range nftPayments {
		response.Data.Attributes.NftList = append(response.Data.Attributes.NftList, resources.NftListItem{
			Key: resources.Key{
				ID:   strconv.Itoa(int(nftPayment.NftId)),
				Type: resources.NFT_LIST,
			},
			Attributes: resources.NftListItemAttributes{
				Address: nftPayment.NftAddress,
			},
		})
	}
	ape.Render(w, response)
}

func setTokenPriceByBook(response resources.StatisticsByBookResponse, tokenSymbol string, price float64) resources.StatisticsByBookResponse {
	for _, token := range response.Data.Attributes.TokensHistogram.Attributes.Tokens {
		if token.Attributes.Name == tokenSymbol {
			token.Attributes.NativeCurrency = price
			return response
		}
	}
	response.Data.Attributes.TokensHistogram.Attributes.Tokens = append(response.Data.Attributes.TokensHistogram.Attributes.Tokens, resources.PricePieChartTokens{
		Key: resources.Key{Type: resources.TOKENS},
		Attributes: resources.PricePieChartTokensAttributes{
			Name:           tokenSymbol,
			NativeCurrency: price,
		},
	})
	return response
}
func setUsdPriceByBook(response resources.StatisticsByBookResponse, tokenSymbol string, price float64) resources.StatisticsByBookResponse {
	for _, token := range response.Data.Attributes.TokensHistogram.Attributes.Tokens {
		if token.Attributes.Name == tokenSymbol {
			token.Attributes.Usd = price
			return response
		}
	}
	response.Data.Attributes.TokensHistogram.Attributes.Tokens = append(response.Data.Attributes.TokensHistogram.Attributes.Tokens, resources.PricePieChartTokens{
		Key: resources.Key{Type: resources.TOKENS},
		Attributes: resources.PricePieChartTokensAttributes{
			Name: tokenSymbol,
			Usd:  price,
		},
	})
	return response
}

func getDate(key string) string {
	// Find the first and second occurrence of "-" in the key
	firstIndex := strings.Index(key, "date-") + len("date-")
	return key[firstIndex:]
}
