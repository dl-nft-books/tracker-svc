package combiners

import (
	"context"
	"github.com/dl-nft-books/tracker-svc/internal/config"
	"github.com/dl-nft-books/tracker-svc/internal/data/etherdata"
	"github.com/dl-nft-books/tracker-svc/internal/service/runners/consumers"
	"github.com/dl-nft-books/tracker-svc/internal/service/runners/trackers"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/logan/v3"
)

type TokenCombiner struct {
	tracker            *trackers.MarketPlaceTracker
	consumer           *consumers.MarketPlaceConsumer
	logger             *logan.Entry
	ctx                context.Context
	marketplaceAddress common.Address
}

func NewTokenCombiner(cfg config.Config, ctx context.Context, deployedToken trackers.DeployedToken) *TokenCombiner {
	return &TokenCombiner{
		tracker:            trackers.NewMarketPlaceTracker(cfg, ctx, deployedToken),
		consumer:           consumers.NewTokenConsumer(cfg, ctx, deployedToken),
		logger:             cfg.Log(),
		ctx:                ctx,
		marketplaceAddress: deployedToken.Address,
	}
}

func (c *TokenCombiner) ProduceAndConsumeMintEvents() {
	// Running tracker (producer) and consumer with a combinersChannel joining them
	c.logger.Infof("Initializing mint event consumer and producer for %s", c.marketplaceAddress)
	go func() {
		ch := make(chan etherdata.SuccessfulMintEvent)
		go c.tracker.TrackMintEvents(ch)
		go c.consumer.ConsumeMintEvents(ch)
	}()
}

func (c *TokenCombiner) ProduceAndConsumeMintByNftEvents() {
	// Running tracker (producer) and consumer with a combinersChannel joining them
	c.logger.Infof("Initializing mint by NFT event consumer and producer for %s", c.marketplaceAddress)
	go func() {
		ch := make(chan etherdata.SuccessfullyMintedByNftEvent)
		go c.tracker.TrackMintByNftEvents(ch)
		go c.consumer.ConsumeMintByNftEvents(ch)
	}()
}

func (c *TokenCombiner) ProduceAndConsumeAllEvents() {
	c.logger.Infof("Initializing all possible consumers and producers for %s", c.marketplaceAddress)
	c.ProduceAndConsumeMintEvents()
	c.ProduceAndConsumeMintByNftEvents()
}
