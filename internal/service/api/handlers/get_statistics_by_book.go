package handlers

import (
	"github.com/dl-nft-books/tracker-svc/internal/service/api/requests"
	"github.com/dl-nft-books/tracker-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"
	"strconv"
)

func GetStatisticsByBook(w http.ResponseWriter, r *http.Request) {
	var response resources.StatisticsByBookResponse
	response.Data.Type = resources.STATISTICS
	request, err := requests.NewGetStatisticsByBookRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	bookStats, err := DB(r).Statistics().BookStatisticsQ.New().FilterByBookId(request.BookId).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get book statistics")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	if bookStats == nil {
		Log(r).WithFields(logan.F{"book_id": request.BookId}).Error("book statistics not found")
		// return empty statistics
		ape.Render(w, response)
		return
	}
	response.Data.Attributes.TokensHistogram.Attributes.Total = bookStats.UsdPrice
	// Token Histogram
	tokenStats, err := DB(r).Statistics().TokenStatisticsQ.New().
		FilterByBookId(request.BookId).Select()
	if err != nil {
		Log(r).WithError(err).Error("failed to get token statistics")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	for _, token := range tokenStats {
		response.Data.Attributes.TokensHistogram.Attributes.Tokens =
			append(response.Data.Attributes.TokensHistogram.Attributes.Tokens, resources.PricePieChartTokens{
				Key: resources.Key{},
				Attributes: resources.PricePieChartTokensAttributes{
					Name:           token.TokenSymbol,
					NativeCurrency: token.TokenPrice,
					Usd:            token.UsdPrice,
				},
			})
	}

	// Date Graph chart
	dateStats, err := DB(r).Statistics().DateStatisticsQ.New().FilterByBookId(request.BookId).Select()
	if err != nil {
		Log(r).WithError(err).Error("failed to get date statistics")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	for _, date := range dateStats {
		response.Data.Attributes.DateGraph =
			append(response.Data.Attributes.DateGraph, resources.DateGraph{
				Key: resources.Key{},
				Attributes: resources.DateGraphAttributes{
					Amount: date.Amount,
					Date:   date.Date,
				},
			})
	}

	// Chain pie chart
	chainStats, err := DB(r).Statistics().ChainStatisticsQ.New().FilterByBookId(request.BookId).Select()
	if err != nil {
		Log(r).WithError(err).Error("failed to get chain statistics")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	for _, chain := range chainStats {
		response.Data.Attributes.ChainPieChart.Attributes.Chains =
			append(response.Data.Attributes.ChainPieChart.Attributes.Chains, resources.ChainPieChartChains{
				Key: resources.Key{},
				Attributes: resources.ChainPieChartChainsAttributes{
					Amount:  chain.Amount,
					ChainId: chain.ChainId,
				},
			})
	}

	nftPayments, err := DB(r).Payments().Page(pgdb.OffsetPageParams{
		Limit:      request.Limit,
		Order:      "desc",
		PageNumber: 0,
	}, "price_token").FilterByType(resources.NFT).FilterByBookId(request.BookId).Select()
	if err != nil {
		Log(r).WithError(err).Error("failed to get nft payments")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	for _, nftPayment := range nftPayments {
		response.Data.Attributes.NftList = append(response.Data.Attributes.NftList, resources.NftListItem{
			Key: resources.Key{
				ID:   strconv.Itoa(int(nftPayment.NftId)),
				Type: resources.NFT_LIST,
			},
			Attributes: resources.NftListItemAttributes{
				Address: nftPayment.TokenAddress,
				ChainId: nftPayment.ChainId,
			},
		})
	}
	ape.Render(w, response)
}
