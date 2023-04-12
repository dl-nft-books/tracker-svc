// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package marketplace

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IMarketplaceBaseTokenParams is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceBaseTokenParams struct {
	TokenContract    common.Address
	IsDisabled       bool
	PricePerOneToken *big.Int
	TokenName        string
}

// IMarketplaceBuyParams is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceBuyParams struct {
	PaymentDetails IMarketplacePaymentDetails
	TokenContract  common.Address
	FutureTokenId  *big.Int
	EndTimestamp   *big.Int
	TokenURI       string
}

// IMarketplaceDetailedTokenParams is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceDetailedTokenParams struct {
	TokenContract common.Address
	TokenParams   IMarketplaceTokenParams
	TokenName     string
	TokenSymbol   string
}

// IMarketplacePaymentDetails is an auto generated low-level Go binding around an user-defined struct.
type IMarketplacePaymentDetails struct {
	PaymentTokenAddress common.Address
	PaymentTokenPrice   *big.Int
	Discount            *big.Int
	NftTokenId          *big.Int
}

// IMarketplaceSig is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceSig struct {
	R [32]byte
	S [32]byte
	V uint8
}

// IMarketplaceTokenParams is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceTokenParams struct {
	PricePerOneToken     *big.Int
	MinNFTFloorPrice     *big.Int
	VoucherTokensAmount  *big.Int
	VoucherTokenContract common.Address
	FundsRecipient       common.Address
	IsNFTBuyable         bool
	IsDisabled           bool
}

// MarketplaceMetaData contains all meta data concerning the Marketplace contract.
var MarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newBaseTokenContractsURI\",\"type\":\"string\"}],\"name\":\"BaseTokenContractsURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PaidTokensWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minNFTFloorPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voucherTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voucherTokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundsRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNFTBuyable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structIMarketplace.TokenParams\",\"name\":\"tokenParams\",\"type\":\"tuple\"}],\"name\":\"TokenContractDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minNFTFloorPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voucherTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voucherTokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundsRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNFTBuyable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structIMarketplace.TokenParams\",\"name\":\"tokenParams\",\"type\":\"tuple\"}],\"name\":\"TokenContractParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintedTokenPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paidTokensAmount\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"paymentTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIMarketplace.PaymentDetails\",\"name\":\"paymentDetails\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"futureTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIMarketplace.BuyParams\",\"name\":\"buyParams\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"enumIMarketplace.PaymentType\",\"name\":\"paymentType\",\"type\":\"uint8\"}],\"name\":\"TokenSuccessfullyPurchased\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseTokenContractsURI_\",\"type\":\"string\"}],\"name\":\"__Marketplace_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minNFTFloorPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voucherTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voucherTokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundsRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNFTBuyable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"}],\"internalType\":\"structIMarketplace.TokenParams\",\"name\":\"tokenParams_\",\"type\":\"tuple\"}],\"name\":\"addToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenProxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseTokenContractsURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"paymentTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIMarketplace.PaymentDetails\",\"name\":\"paymentDetails\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"futureTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"internalType\":\"structIMarketplace.BuyParams\",\"name\":\"buyParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structIMarketplace.Sig\",\"name\":\"sig_\",\"type\":\"tuple\"}],\"name\":\"buyTokenWithERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"paymentTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIMarketplace.PaymentDetails\",\"name\":\"paymentDetails\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"futureTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"internalType\":\"structIMarketplace.BuyParams\",\"name\":\"buyParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structIMarketplace.Sig\",\"name\":\"sig_\",\"type\":\"tuple\"}],\"name\":\"buyTokenWithETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"paymentTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIMarketplace.PaymentDetails\",\"name\":\"paymentDetails\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"futureTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"internalType\":\"structIMarketplace.BuyParams\",\"name\":\"buyParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structIMarketplace.Sig\",\"name\":\"sig_\",\"type\":\"tuple\"}],\"name\":\"buyTokenWithNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"paymentTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIMarketplace.PaymentDetails\",\"name\":\"paymentDetails\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"futureTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"internalType\":\"structIMarketplace.BuyParams\",\"name\":\"buyParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structIMarketplace.Sig\",\"name\":\"sig_\",\"type\":\"tuple\"}],\"name\":\"buyTokenWithVoucher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveTokenContractsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokenContracts_\",\"type\":\"address[]\"}],\"name\":\"getBaseTokenParams\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"}],\"internalType\":\"structIMarketplace.BaseTokenParams[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getBaseTokenParamsPart\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"}],\"internalType\":\"structIMarketplace.BaseTokenParams[]\",\"name\":\"tokenParams_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokenContracts_\",\"type\":\"address[]\"}],\"name\":\"getDetailedTokenParams\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minNFTFloorPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voucherTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voucherTokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundsRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNFTBuyable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"}],\"internalType\":\"structIMarketplace.TokenParams\",\"name\":\"tokenParams\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"internalType\":\"structIMarketplace.DetailedTokenParams[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getDetailedTokenParamsPart\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minNFTFloorPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voucherTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voucherTokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundsRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNFTBuyable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"}],\"internalType\":\"structIMarketplace.TokenParams\",\"name\":\"tokenParams\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"internalType\":\"structIMarketplace.DetailedTokenParams[]\",\"name\":\"tokenParams_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenContractsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getTokenContractsPart\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContract_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"userAddr_\",\"type\":\"address\"}],\"name\":\"getUserTokenIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIDs_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseTokenContractsURI_\",\"type\":\"string\"}],\"name\":\"setBaseTokenContractsURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContract_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minNFTFloorPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voucherTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voucherTokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundsRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNFTBuyable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"}],\"internalType\":\"structIMarketplace.TokenParams\",\"name\":\"newTokenParams_\",\"type\":\"tuple\"}],\"name\":\"updateAllParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient_\",\"type\":\"address\"}],\"name\":\"withdrawCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketplaceMetaData.ABI instead.
var MarketplaceABI = MarketplaceMetaData.ABI

// Marketplace is an auto generated Go binding around an Ethereum contract.
type Marketplace struct {
	MarketplaceCaller     // Read-only binding to the contract
	MarketplaceTransactor // Write-only binding to the contract
	MarketplaceFilterer   // Log filterer for contract events
}

// MarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketplaceSession struct {
	Contract     *Marketplace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketplaceCallerSession struct {
	Contract *MarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketplaceTransactorSession struct {
	Contract     *MarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketplaceRaw struct {
	Contract *Marketplace // Generic contract binding to access the raw methods on
}

// MarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketplaceCallerRaw struct {
	Contract *MarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketplaceTransactorRaw struct {
	Contract *MarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplace creates a new instance of Marketplace, bound to a specific deployed contract.
func NewMarketplace(address common.Address, backend bind.ContractBackend) (*Marketplace, error) {
	contract, err := bindMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marketplace{MarketplaceCaller: MarketplaceCaller{contract: contract}, MarketplaceTransactor: MarketplaceTransactor{contract: contract}, MarketplaceFilterer: MarketplaceFilterer{contract: contract}}, nil
}

// NewMarketplaceCaller creates a new read-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*MarketplaceCaller, error) {
	contract, err := bindMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCaller{contract: contract}, nil
}

// NewMarketplaceTransactor creates a new write-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketplaceTransactor, error) {
	contract, err := bindMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTransactor{contract: contract}, nil
}

// NewMarketplaceFilterer creates a new log filterer instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketplaceFilterer, error) {
	contract, err := bindMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketplaceFilterer{contract: contract}, nil
}

// bindMarketplace binds a generic wrapper to an already deployed contract.
func bindMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketplaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.MarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transact(opts, method, params...)
}

// BaseTokenContractsURI is a free data retrieval call binding the contract method 0x72148c1c.
//
// Solidity: function baseTokenContractsURI() view returns(string)
func (_Marketplace *MarketplaceCaller) BaseTokenContractsURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "baseTokenContractsURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseTokenContractsURI is a free data retrieval call binding the contract method 0x72148c1c.
//
// Solidity: function baseTokenContractsURI() view returns(string)
func (_Marketplace *MarketplaceSession) BaseTokenContractsURI() (string, error) {
	return _Marketplace.Contract.BaseTokenContractsURI(&_Marketplace.CallOpts)
}

// BaseTokenContractsURI is a free data retrieval call binding the contract method 0x72148c1c.
//
// Solidity: function baseTokenContractsURI() view returns(string)
func (_Marketplace *MarketplaceCallerSession) BaseTokenContractsURI() (string, error) {
	return _Marketplace.Contract.BaseTokenContractsURI(&_Marketplace.CallOpts)
}

// GetActiveTokenContractsCount is a free data retrieval call binding the contract method 0x206a3da6.
//
// Solidity: function getActiveTokenContractsCount() view returns(uint256)
func (_Marketplace *MarketplaceCaller) GetActiveTokenContractsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getActiveTokenContractsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveTokenContractsCount is a free data retrieval call binding the contract method 0x206a3da6.
//
// Solidity: function getActiveTokenContractsCount() view returns(uint256)
func (_Marketplace *MarketplaceSession) GetActiveTokenContractsCount() (*big.Int, error) {
	return _Marketplace.Contract.GetActiveTokenContractsCount(&_Marketplace.CallOpts)
}

// GetActiveTokenContractsCount is a free data retrieval call binding the contract method 0x206a3da6.
//
// Solidity: function getActiveTokenContractsCount() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) GetActiveTokenContractsCount() (*big.Int, error) {
	return _Marketplace.Contract.GetActiveTokenContractsCount(&_Marketplace.CallOpts)
}

// GetBaseTokenParams is a free data retrieval call binding the contract method 0x20acff14.
//
// Solidity: function getBaseTokenParams(address[] tokenContracts_) view returns((address,bool,uint256,string)[])
func (_Marketplace *MarketplaceCaller) GetBaseTokenParams(opts *bind.CallOpts, tokenContracts_ []common.Address) ([]IMarketplaceBaseTokenParams, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getBaseTokenParams", tokenContracts_)

	if err != nil {
		return *new([]IMarketplaceBaseTokenParams), err
	}

	out0 := *abi.ConvertType(out[0], new([]IMarketplaceBaseTokenParams)).(*[]IMarketplaceBaseTokenParams)

	return out0, err

}

// GetBaseTokenParams is a free data retrieval call binding the contract method 0x20acff14.
//
// Solidity: function getBaseTokenParams(address[] tokenContracts_) view returns((address,bool,uint256,string)[])
func (_Marketplace *MarketplaceSession) GetBaseTokenParams(tokenContracts_ []common.Address) ([]IMarketplaceBaseTokenParams, error) {
	return _Marketplace.Contract.GetBaseTokenParams(&_Marketplace.CallOpts, tokenContracts_)
}

// GetBaseTokenParams is a free data retrieval call binding the contract method 0x20acff14.
//
// Solidity: function getBaseTokenParams(address[] tokenContracts_) view returns((address,bool,uint256,string)[])
func (_Marketplace *MarketplaceCallerSession) GetBaseTokenParams(tokenContracts_ []common.Address) ([]IMarketplaceBaseTokenParams, error) {
	return _Marketplace.Contract.GetBaseTokenParams(&_Marketplace.CallOpts, tokenContracts_)
}

// GetBaseTokenParamsPart is a free data retrieval call binding the contract method 0x27a99cf7.
//
// Solidity: function getBaseTokenParamsPart(uint256 offset_, uint256 limit_) view returns((address,bool,uint256,string)[] tokenParams_)
func (_Marketplace *MarketplaceCaller) GetBaseTokenParamsPart(opts *bind.CallOpts, offset_ *big.Int, limit_ *big.Int) ([]IMarketplaceBaseTokenParams, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getBaseTokenParamsPart", offset_, limit_)

	if err != nil {
		return *new([]IMarketplaceBaseTokenParams), err
	}

	out0 := *abi.ConvertType(out[0], new([]IMarketplaceBaseTokenParams)).(*[]IMarketplaceBaseTokenParams)

	return out0, err

}

// GetBaseTokenParamsPart is a free data retrieval call binding the contract method 0x27a99cf7.
//
// Solidity: function getBaseTokenParamsPart(uint256 offset_, uint256 limit_) view returns((address,bool,uint256,string)[] tokenParams_)
func (_Marketplace *MarketplaceSession) GetBaseTokenParamsPart(offset_ *big.Int, limit_ *big.Int) ([]IMarketplaceBaseTokenParams, error) {
	return _Marketplace.Contract.GetBaseTokenParamsPart(&_Marketplace.CallOpts, offset_, limit_)
}

// GetBaseTokenParamsPart is a free data retrieval call binding the contract method 0x27a99cf7.
//
// Solidity: function getBaseTokenParamsPart(uint256 offset_, uint256 limit_) view returns((address,bool,uint256,string)[] tokenParams_)
func (_Marketplace *MarketplaceCallerSession) GetBaseTokenParamsPart(offset_ *big.Int, limit_ *big.Int) ([]IMarketplaceBaseTokenParams, error) {
	return _Marketplace.Contract.GetBaseTokenParamsPart(&_Marketplace.CallOpts, offset_, limit_)
}

// GetDetailedTokenParams is a free data retrieval call binding the contract method 0xcc1270cb.
//
// Solidity: function getDetailedTokenParams(address[] tokenContracts_) view returns((address,(uint256,uint256,uint256,address,address,bool,bool),string,string)[])
func (_Marketplace *MarketplaceCaller) GetDetailedTokenParams(opts *bind.CallOpts, tokenContracts_ []common.Address) ([]IMarketplaceDetailedTokenParams, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getDetailedTokenParams", tokenContracts_)

	if err != nil {
		return *new([]IMarketplaceDetailedTokenParams), err
	}

	out0 := *abi.ConvertType(out[0], new([]IMarketplaceDetailedTokenParams)).(*[]IMarketplaceDetailedTokenParams)

	return out0, err

}

// GetDetailedTokenParams is a free data retrieval call binding the contract method 0xcc1270cb.
//
// Solidity: function getDetailedTokenParams(address[] tokenContracts_) view returns((address,(uint256,uint256,uint256,address,address,bool,bool),string,string)[])
func (_Marketplace *MarketplaceSession) GetDetailedTokenParams(tokenContracts_ []common.Address) ([]IMarketplaceDetailedTokenParams, error) {
	return _Marketplace.Contract.GetDetailedTokenParams(&_Marketplace.CallOpts, tokenContracts_)
}

// GetDetailedTokenParams is a free data retrieval call binding the contract method 0xcc1270cb.
//
// Solidity: function getDetailedTokenParams(address[] tokenContracts_) view returns((address,(uint256,uint256,uint256,address,address,bool,bool),string,string)[])
func (_Marketplace *MarketplaceCallerSession) GetDetailedTokenParams(tokenContracts_ []common.Address) ([]IMarketplaceDetailedTokenParams, error) {
	return _Marketplace.Contract.GetDetailedTokenParams(&_Marketplace.CallOpts, tokenContracts_)
}

// GetDetailedTokenParamsPart is a free data retrieval call binding the contract method 0xce60e378.
//
// Solidity: function getDetailedTokenParamsPart(uint256 offset_, uint256 limit_) view returns((address,(uint256,uint256,uint256,address,address,bool,bool),string,string)[] tokenParams_)
func (_Marketplace *MarketplaceCaller) GetDetailedTokenParamsPart(opts *bind.CallOpts, offset_ *big.Int, limit_ *big.Int) ([]IMarketplaceDetailedTokenParams, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getDetailedTokenParamsPart", offset_, limit_)

	if err != nil {
		return *new([]IMarketplaceDetailedTokenParams), err
	}

	out0 := *abi.ConvertType(out[0], new([]IMarketplaceDetailedTokenParams)).(*[]IMarketplaceDetailedTokenParams)

	return out0, err

}

// GetDetailedTokenParamsPart is a free data retrieval call binding the contract method 0xce60e378.
//
// Solidity: function getDetailedTokenParamsPart(uint256 offset_, uint256 limit_) view returns((address,(uint256,uint256,uint256,address,address,bool,bool),string,string)[] tokenParams_)
func (_Marketplace *MarketplaceSession) GetDetailedTokenParamsPart(offset_ *big.Int, limit_ *big.Int) ([]IMarketplaceDetailedTokenParams, error) {
	return _Marketplace.Contract.GetDetailedTokenParamsPart(&_Marketplace.CallOpts, offset_, limit_)
}

// GetDetailedTokenParamsPart is a free data retrieval call binding the contract method 0xce60e378.
//
// Solidity: function getDetailedTokenParamsPart(uint256 offset_, uint256 limit_) view returns((address,(uint256,uint256,uint256,address,address,bool,bool),string,string)[] tokenParams_)
func (_Marketplace *MarketplaceCallerSession) GetDetailedTokenParamsPart(offset_ *big.Int, limit_ *big.Int) ([]IMarketplaceDetailedTokenParams, error) {
	return _Marketplace.Contract.GetDetailedTokenParamsPart(&_Marketplace.CallOpts, offset_, limit_)
}

// GetTokenContractsCount is a free data retrieval call binding the contract method 0x11b2bf29.
//
// Solidity: function getTokenContractsCount() view returns(uint256)
func (_Marketplace *MarketplaceCaller) GetTokenContractsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getTokenContractsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenContractsCount is a free data retrieval call binding the contract method 0x11b2bf29.
//
// Solidity: function getTokenContractsCount() view returns(uint256)
func (_Marketplace *MarketplaceSession) GetTokenContractsCount() (*big.Int, error) {
	return _Marketplace.Contract.GetTokenContractsCount(&_Marketplace.CallOpts)
}

// GetTokenContractsCount is a free data retrieval call binding the contract method 0x11b2bf29.
//
// Solidity: function getTokenContractsCount() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) GetTokenContractsCount() (*big.Int, error) {
	return _Marketplace.Contract.GetTokenContractsCount(&_Marketplace.CallOpts)
}

// GetTokenContractsPart is a free data retrieval call binding the contract method 0xcc7c925a.
//
// Solidity: function getTokenContractsPart(uint256 offset_, uint256 limit_) view returns(address[])
func (_Marketplace *MarketplaceCaller) GetTokenContractsPart(opts *bind.CallOpts, offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getTokenContractsPart", offset_, limit_)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenContractsPart is a free data retrieval call binding the contract method 0xcc7c925a.
//
// Solidity: function getTokenContractsPart(uint256 offset_, uint256 limit_) view returns(address[])
func (_Marketplace *MarketplaceSession) GetTokenContractsPart(offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	return _Marketplace.Contract.GetTokenContractsPart(&_Marketplace.CallOpts, offset_, limit_)
}

// GetTokenContractsPart is a free data retrieval call binding the contract method 0xcc7c925a.
//
// Solidity: function getTokenContractsPart(uint256 offset_, uint256 limit_) view returns(address[])
func (_Marketplace *MarketplaceCallerSession) GetTokenContractsPart(offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	return _Marketplace.Contract.GetTokenContractsPart(&_Marketplace.CallOpts, offset_, limit_)
}

// GetUserTokenIDs is a free data retrieval call binding the contract method 0xe420e11f.
//
// Solidity: function getUserTokenIDs(address tokenContract_, address userAddr_) view returns(uint256[] tokenIDs_)
func (_Marketplace *MarketplaceCaller) GetUserTokenIDs(opts *bind.CallOpts, tokenContract_ common.Address, userAddr_ common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getUserTokenIDs", tokenContract_, userAddr_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserTokenIDs is a free data retrieval call binding the contract method 0xe420e11f.
//
// Solidity: function getUserTokenIDs(address tokenContract_, address userAddr_) view returns(uint256[] tokenIDs_)
func (_Marketplace *MarketplaceSession) GetUserTokenIDs(tokenContract_ common.Address, userAddr_ common.Address) ([]*big.Int, error) {
	return _Marketplace.Contract.GetUserTokenIDs(&_Marketplace.CallOpts, tokenContract_, userAddr_)
}

// GetUserTokenIDs is a free data retrieval call binding the contract method 0xe420e11f.
//
// Solidity: function getUserTokenIDs(address tokenContract_, address userAddr_) view returns(uint256[] tokenIDs_)
func (_Marketplace *MarketplaceCallerSession) GetUserTokenIDs(tokenContract_ common.Address, userAddr_ common.Address) ([]*big.Int, error) {
	return _Marketplace.Contract.GetUserTokenIDs(&_Marketplace.CallOpts, tokenContract_, userAddr_)
}

// MarketplaceInit is a paid mutator transaction binding the contract method 0xe709d541.
//
// Solidity: function __Marketplace_init(string baseTokenContractsURI_) returns()
func (_Marketplace *MarketplaceTransactor) MarketplaceInit(opts *bind.TransactOpts, baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "__Marketplace_init", baseTokenContractsURI_)
}

// MarketplaceInit is a paid mutator transaction binding the contract method 0xe709d541.
//
// Solidity: function __Marketplace_init(string baseTokenContractsURI_) returns()
func (_Marketplace *MarketplaceSession) MarketplaceInit(baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceInit(&_Marketplace.TransactOpts, baseTokenContractsURI_)
}

// MarketplaceInit is a paid mutator transaction binding the contract method 0xe709d541.
//
// Solidity: function __Marketplace_init(string baseTokenContractsURI_) returns()
func (_Marketplace *MarketplaceTransactorSession) MarketplaceInit(baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceInit(&_Marketplace.TransactOpts, baseTokenContractsURI_)
}

// AddToken is a paid mutator transaction binding the contract method 0x3a65069c.
//
// Solidity: function addToken(string name_, string symbol_, (uint256,uint256,uint256,address,address,bool,bool) tokenParams_) returns(address tokenProxy)
func (_Marketplace *MarketplaceTransactor) AddToken(opts *bind.TransactOpts, name_ string, symbol_ string, tokenParams_ IMarketplaceTokenParams) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "addToken", name_, symbol_, tokenParams_)
}

// AddToken is a paid mutator transaction binding the contract method 0x3a65069c.
//
// Solidity: function addToken(string name_, string symbol_, (uint256,uint256,uint256,address,address,bool,bool) tokenParams_) returns(address tokenProxy)
func (_Marketplace *MarketplaceSession) AddToken(name_ string, symbol_ string, tokenParams_ IMarketplaceTokenParams) (*types.Transaction, error) {
	return _Marketplace.Contract.AddToken(&_Marketplace.TransactOpts, name_, symbol_, tokenParams_)
}

// AddToken is a paid mutator transaction binding the contract method 0x3a65069c.
//
// Solidity: function addToken(string name_, string symbol_, (uint256,uint256,uint256,address,address,bool,bool) tokenParams_) returns(address tokenProxy)
func (_Marketplace *MarketplaceTransactorSession) AddToken(name_ string, symbol_ string, tokenParams_ IMarketplaceTokenParams) (*types.Transaction, error) {
	return _Marketplace.Contract.AddToken(&_Marketplace.TransactOpts, name_, symbol_, tokenParams_)
}

// BuyTokenWithERC20 is a paid mutator transaction binding the contract method 0x46334dc1.
//
// Solidity: function buyTokenWithERC20(((address,uint256,uint256,uint256),address,uint256,uint256,string) buyParams_, (bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceTransactor) BuyTokenWithERC20(opts *bind.TransactOpts, buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "buyTokenWithERC20", buyParams_, sig_)
}

// BuyTokenWithERC20 is a paid mutator transaction binding the contract method 0x46334dc1.
//
// Solidity: function buyTokenWithERC20(((address,uint256,uint256,uint256),address,uint256,uint256,string) buyParams_, (bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceSession) BuyTokenWithERC20(buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenWithERC20(&_Marketplace.TransactOpts, buyParams_, sig_)
}

// BuyTokenWithERC20 is a paid mutator transaction binding the contract method 0x46334dc1.
//
// Solidity: function buyTokenWithERC20(((address,uint256,uint256,uint256),address,uint256,uint256,string) buyParams_, (bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceTransactorSession) BuyTokenWithERC20(buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenWithERC20(&_Marketplace.TransactOpts, buyParams_, sig_)
}

// BuyTokenWithETH is a paid mutator transaction binding the contract method 0x21871a3d.
//
// Solidity: function buyTokenWithETH(((address,uint256,uint256,uint256),address,uint256,uint256,string) buyParams_, (bytes32,bytes32,uint8) sig_) payable returns()
func (_Marketplace *MarketplaceTransactor) BuyTokenWithETH(opts *bind.TransactOpts, buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "buyTokenWithETH", buyParams_, sig_)
}

// BuyTokenWithETH is a paid mutator transaction binding the contract method 0x21871a3d.
//
// Solidity: function buyTokenWithETH(((address,uint256,uint256,uint256),address,uint256,uint256,string) buyParams_, (bytes32,bytes32,uint8) sig_) payable returns()
func (_Marketplace *MarketplaceSession) BuyTokenWithETH(buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenWithETH(&_Marketplace.TransactOpts, buyParams_, sig_)
}

// BuyTokenWithETH is a paid mutator transaction binding the contract method 0x21871a3d.
//
// Solidity: function buyTokenWithETH(((address,uint256,uint256,uint256),address,uint256,uint256,string) buyParams_, (bytes32,bytes32,uint8) sig_) payable returns()
func (_Marketplace *MarketplaceTransactorSession) BuyTokenWithETH(buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenWithETH(&_Marketplace.TransactOpts, buyParams_, sig_)
}

// BuyTokenWithNFT is a paid mutator transaction binding the contract method 0x29d9621d.
//
// Solidity: function buyTokenWithNFT(((address,uint256,uint256,uint256),address,uint256,uint256,string) buyParams_, (bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceTransactor) BuyTokenWithNFT(opts *bind.TransactOpts, buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "buyTokenWithNFT", buyParams_, sig_)
}

// BuyTokenWithNFT is a paid mutator transaction binding the contract method 0x29d9621d.
//
// Solidity: function buyTokenWithNFT(((address,uint256,uint256,uint256),address,uint256,uint256,string) buyParams_, (bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceSession) BuyTokenWithNFT(buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenWithNFT(&_Marketplace.TransactOpts, buyParams_, sig_)
}

// BuyTokenWithNFT is a paid mutator transaction binding the contract method 0x29d9621d.
//
// Solidity: function buyTokenWithNFT(((address,uint256,uint256,uint256),address,uint256,uint256,string) buyParams_, (bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceTransactorSession) BuyTokenWithNFT(buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenWithNFT(&_Marketplace.TransactOpts, buyParams_, sig_)
}

// BuyTokenWithVoucher is a paid mutator transaction binding the contract method 0x34e8efe4.
//
// Solidity: function buyTokenWithVoucher(((address,uint256,uint256,uint256),address,uint256,uint256,string) buyParams_, (bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceTransactor) BuyTokenWithVoucher(opts *bind.TransactOpts, buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "buyTokenWithVoucher", buyParams_, sig_)
}

// BuyTokenWithVoucher is a paid mutator transaction binding the contract method 0x34e8efe4.
//
// Solidity: function buyTokenWithVoucher(((address,uint256,uint256,uint256),address,uint256,uint256,string) buyParams_, (bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceSession) BuyTokenWithVoucher(buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenWithVoucher(&_Marketplace.TransactOpts, buyParams_, sig_)
}

// BuyTokenWithVoucher is a paid mutator transaction binding the contract method 0x34e8efe4.
//
// Solidity: function buyTokenWithVoucher(((address,uint256,uint256,uint256),address,uint256,uint256,string) buyParams_, (bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceTransactorSession) BuyTokenWithVoucher(buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenWithVoucher(&_Marketplace.TransactOpts, buyParams_, sig_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Marketplace *MarketplaceTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Marketplace *MarketplaceSession) Pause() (*types.Transaction, error) {
	return _Marketplace.Contract.Pause(&_Marketplace.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Marketplace *MarketplaceTransactorSession) Pause() (*types.Transaction, error) {
	return _Marketplace.Contract.Pause(&_Marketplace.TransactOpts)
}

// SetBaseTokenContractsURI is a paid mutator transaction binding the contract method 0x30dcd8aa.
//
// Solidity: function setBaseTokenContractsURI(string baseTokenContractsURI_) returns()
func (_Marketplace *MarketplaceTransactor) SetBaseTokenContractsURI(opts *bind.TransactOpts, baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setBaseTokenContractsURI", baseTokenContractsURI_)
}

// SetBaseTokenContractsURI is a paid mutator transaction binding the contract method 0x30dcd8aa.
//
// Solidity: function setBaseTokenContractsURI(string baseTokenContractsURI_) returns()
func (_Marketplace *MarketplaceSession) SetBaseTokenContractsURI(baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _Marketplace.Contract.SetBaseTokenContractsURI(&_Marketplace.TransactOpts, baseTokenContractsURI_)
}

// SetBaseTokenContractsURI is a paid mutator transaction binding the contract method 0x30dcd8aa.
//
// Solidity: function setBaseTokenContractsURI(string baseTokenContractsURI_) returns()
func (_Marketplace *MarketplaceTransactorSession) SetBaseTokenContractsURI(baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _Marketplace.Contract.SetBaseTokenContractsURI(&_Marketplace.TransactOpts, baseTokenContractsURI_)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Marketplace *MarketplaceTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Marketplace *MarketplaceSession) Unpause() (*types.Transaction, error) {
	return _Marketplace.Contract.Unpause(&_Marketplace.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Marketplace *MarketplaceTransactorSession) Unpause() (*types.Transaction, error) {
	return _Marketplace.Contract.Unpause(&_Marketplace.TransactOpts)
}

// UpdateAllParams is a paid mutator transaction binding the contract method 0x34fb02bf.
//
// Solidity: function updateAllParams(address tokenContract_, string name_, string symbol_, (uint256,uint256,uint256,address,address,bool,bool) newTokenParams_) returns()
func (_Marketplace *MarketplaceTransactor) UpdateAllParams(opts *bind.TransactOpts, tokenContract_ common.Address, name_ string, symbol_ string, newTokenParams_ IMarketplaceTokenParams) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "updateAllParams", tokenContract_, name_, symbol_, newTokenParams_)
}

// UpdateAllParams is a paid mutator transaction binding the contract method 0x34fb02bf.
//
// Solidity: function updateAllParams(address tokenContract_, string name_, string symbol_, (uint256,uint256,uint256,address,address,bool,bool) newTokenParams_) returns()
func (_Marketplace *MarketplaceSession) UpdateAllParams(tokenContract_ common.Address, name_ string, symbol_ string, newTokenParams_ IMarketplaceTokenParams) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateAllParams(&_Marketplace.TransactOpts, tokenContract_, name_, symbol_, newTokenParams_)
}

// UpdateAllParams is a paid mutator transaction binding the contract method 0x34fb02bf.
//
// Solidity: function updateAllParams(address tokenContract_, string name_, string symbol_, (uint256,uint256,uint256,address,address,bool,bool) newTokenParams_) returns()
func (_Marketplace *MarketplaceTransactorSession) UpdateAllParams(tokenContract_ common.Address, name_ string, symbol_ string, newTokenParams_ IMarketplaceTokenParams) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateAllParams(&_Marketplace.TransactOpts, tokenContract_, name_, symbol_, newTokenParams_)
}

// WithdrawCurrency is a paid mutator transaction binding the contract method 0x8407ecc4.
//
// Solidity: function withdrawCurrency(address tokenAddr_, address recipient_) returns()
func (_Marketplace *MarketplaceTransactor) WithdrawCurrency(opts *bind.TransactOpts, tokenAddr_ common.Address, recipient_ common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "withdrawCurrency", tokenAddr_, recipient_)
}

// WithdrawCurrency is a paid mutator transaction binding the contract method 0x8407ecc4.
//
// Solidity: function withdrawCurrency(address tokenAddr_, address recipient_) returns()
func (_Marketplace *MarketplaceSession) WithdrawCurrency(tokenAddr_ common.Address, recipient_ common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.WithdrawCurrency(&_Marketplace.TransactOpts, tokenAddr_, recipient_)
}

// WithdrawCurrency is a paid mutator transaction binding the contract method 0x8407ecc4.
//
// Solidity: function withdrawCurrency(address tokenAddr_, address recipient_) returns()
func (_Marketplace *MarketplaceTransactorSession) WithdrawCurrency(tokenAddr_ common.Address, recipient_ common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.WithdrawCurrency(&_Marketplace.TransactOpts, tokenAddr_, recipient_)
}

// MarketplaceBaseTokenContractsURIUpdatedIterator is returned from FilterBaseTokenContractsURIUpdated and is used to iterate over the raw logs and unpacked data for BaseTokenContractsURIUpdated events raised by the Marketplace contract.
type MarketplaceBaseTokenContractsURIUpdatedIterator struct {
	Event *MarketplaceBaseTokenContractsURIUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceBaseTokenContractsURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceBaseTokenContractsURIUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceBaseTokenContractsURIUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceBaseTokenContractsURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceBaseTokenContractsURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceBaseTokenContractsURIUpdated represents a BaseTokenContractsURIUpdated event raised by the Marketplace contract.
type MarketplaceBaseTokenContractsURIUpdated struct {
	NewBaseTokenContractsURI string
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterBaseTokenContractsURIUpdated is a free log retrieval operation binding the contract event 0x35221ce24b65bb230797e03080154779f50580793d0f60505bb8a6d15c507b80.
//
// Solidity: event BaseTokenContractsURIUpdated(string newBaseTokenContractsURI)
func (_Marketplace *MarketplaceFilterer) FilterBaseTokenContractsURIUpdated(opts *bind.FilterOpts) (*MarketplaceBaseTokenContractsURIUpdatedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "BaseTokenContractsURIUpdated")
	if err != nil {
		return nil, err
	}
	return &MarketplaceBaseTokenContractsURIUpdatedIterator{contract: _Marketplace.contract, event: "BaseTokenContractsURIUpdated", logs: logs, sub: sub}, nil
}

// WatchBaseTokenContractsURIUpdated is a free log subscription operation binding the contract event 0x35221ce24b65bb230797e03080154779f50580793d0f60505bb8a6d15c507b80.
//
// Solidity: event BaseTokenContractsURIUpdated(string newBaseTokenContractsURI)
func (_Marketplace *MarketplaceFilterer) WatchBaseTokenContractsURIUpdated(opts *bind.WatchOpts, sink chan<- *MarketplaceBaseTokenContractsURIUpdated) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "BaseTokenContractsURIUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceBaseTokenContractsURIUpdated)
				if err := _Marketplace.contract.UnpackLog(event, "BaseTokenContractsURIUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBaseTokenContractsURIUpdated is a log parse operation binding the contract event 0x35221ce24b65bb230797e03080154779f50580793d0f60505bb8a6d15c507b80.
//
// Solidity: event BaseTokenContractsURIUpdated(string newBaseTokenContractsURI)
func (_Marketplace *MarketplaceFilterer) ParseBaseTokenContractsURIUpdated(log types.Log) (*MarketplaceBaseTokenContractsURIUpdated, error) {
	event := new(MarketplaceBaseTokenContractsURIUpdated)
	if err := _Marketplace.contract.UnpackLog(event, "BaseTokenContractsURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplacePaidTokensWithdrawnIterator is returned from FilterPaidTokensWithdrawn and is used to iterate over the raw logs and unpacked data for PaidTokensWithdrawn events raised by the Marketplace contract.
type MarketplacePaidTokensWithdrawnIterator struct {
	Event *MarketplacePaidTokensWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplacePaidTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplacePaidTokensWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplacePaidTokensWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplacePaidTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplacePaidTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplacePaidTokensWithdrawn represents a PaidTokensWithdrawn event raised by the Marketplace contract.
type MarketplacePaidTokensWithdrawn struct {
	TokenAddr common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPaidTokensWithdrawn is a free log retrieval operation binding the contract event 0xae2408273a2b19d84df6e31c6b33cda24768f5e01118d06cbe3b098e231868f1.
//
// Solidity: event PaidTokensWithdrawn(address indexed tokenAddr, address recipient, uint256 amount)
func (_Marketplace *MarketplaceFilterer) FilterPaidTokensWithdrawn(opts *bind.FilterOpts, tokenAddr []common.Address) (*MarketplacePaidTokensWithdrawnIterator, error) {

	var tokenAddrRule []interface{}
	for _, tokenAddrItem := range tokenAddr {
		tokenAddrRule = append(tokenAddrRule, tokenAddrItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "PaidTokensWithdrawn", tokenAddrRule)
	if err != nil {
		return nil, err
	}
	return &MarketplacePaidTokensWithdrawnIterator{contract: _Marketplace.contract, event: "PaidTokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchPaidTokensWithdrawn is a free log subscription operation binding the contract event 0xae2408273a2b19d84df6e31c6b33cda24768f5e01118d06cbe3b098e231868f1.
//
// Solidity: event PaidTokensWithdrawn(address indexed tokenAddr, address recipient, uint256 amount)
func (_Marketplace *MarketplaceFilterer) WatchPaidTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *MarketplacePaidTokensWithdrawn, tokenAddr []common.Address) (event.Subscription, error) {

	var tokenAddrRule []interface{}
	for _, tokenAddrItem := range tokenAddr {
		tokenAddrRule = append(tokenAddrRule, tokenAddrItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "PaidTokensWithdrawn", tokenAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplacePaidTokensWithdrawn)
				if err := _Marketplace.contract.UnpackLog(event, "PaidTokensWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaidTokensWithdrawn is a log parse operation binding the contract event 0xae2408273a2b19d84df6e31c6b33cda24768f5e01118d06cbe3b098e231868f1.
//
// Solidity: event PaidTokensWithdrawn(address indexed tokenAddr, address recipient, uint256 amount)
func (_Marketplace *MarketplaceFilterer) ParsePaidTokensWithdrawn(log types.Log) (*MarketplacePaidTokensWithdrawn, error) {
	event := new(MarketplacePaidTokensWithdrawn)
	if err := _Marketplace.contract.UnpackLog(event, "PaidTokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceTokenContractDeployedIterator is returned from FilterTokenContractDeployed and is used to iterate over the raw logs and unpacked data for TokenContractDeployed events raised by the Marketplace contract.
type MarketplaceTokenContractDeployedIterator struct {
	Event *MarketplaceTokenContractDeployed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceTokenContractDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceTokenContractDeployed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceTokenContractDeployed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceTokenContractDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceTokenContractDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceTokenContractDeployed represents a TokenContractDeployed event raised by the Marketplace contract.
type MarketplaceTokenContractDeployed struct {
	TokenContract common.Address
	TokenName     string
	TokenSymbol   string
	TokenParams   IMarketplaceTokenParams
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenContractDeployed is a free log retrieval operation binding the contract event 0x781d8b3802c0c6a9e4fb0d1fc008ee5fa955521d5f40eb00a0147560571f4f15.
//
// Solidity: event TokenContractDeployed(address indexed tokenContract, string tokenName, string tokenSymbol, (uint256,uint256,uint256,address,address,bool,bool) tokenParams)
func (_Marketplace *MarketplaceFilterer) FilterTokenContractDeployed(opts *bind.FilterOpts, tokenContract []common.Address) (*MarketplaceTokenContractDeployedIterator, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "TokenContractDeployed", tokenContractRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTokenContractDeployedIterator{contract: _Marketplace.contract, event: "TokenContractDeployed", logs: logs, sub: sub}, nil
}

// WatchTokenContractDeployed is a free log subscription operation binding the contract event 0x781d8b3802c0c6a9e4fb0d1fc008ee5fa955521d5f40eb00a0147560571f4f15.
//
// Solidity: event TokenContractDeployed(address indexed tokenContract, string tokenName, string tokenSymbol, (uint256,uint256,uint256,address,address,bool,bool) tokenParams)
func (_Marketplace *MarketplaceFilterer) WatchTokenContractDeployed(opts *bind.WatchOpts, sink chan<- *MarketplaceTokenContractDeployed, tokenContract []common.Address) (event.Subscription, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "TokenContractDeployed", tokenContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceTokenContractDeployed)
				if err := _Marketplace.contract.UnpackLog(event, "TokenContractDeployed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenContractDeployed is a log parse operation binding the contract event 0x781d8b3802c0c6a9e4fb0d1fc008ee5fa955521d5f40eb00a0147560571f4f15.
//
// Solidity: event TokenContractDeployed(address indexed tokenContract, string tokenName, string tokenSymbol, (uint256,uint256,uint256,address,address,bool,bool) tokenParams)
func (_Marketplace *MarketplaceFilterer) ParseTokenContractDeployed(log types.Log) (*MarketplaceTokenContractDeployed, error) {
	event := new(MarketplaceTokenContractDeployed)
	if err := _Marketplace.contract.UnpackLog(event, "TokenContractDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceTokenContractParamsUpdatedIterator is returned from FilterTokenContractParamsUpdated and is used to iterate over the raw logs and unpacked data for TokenContractParamsUpdated events raised by the Marketplace contract.
type MarketplaceTokenContractParamsUpdatedIterator struct {
	Event *MarketplaceTokenContractParamsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceTokenContractParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceTokenContractParamsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceTokenContractParamsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceTokenContractParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceTokenContractParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceTokenContractParamsUpdated represents a TokenContractParamsUpdated event raised by the Marketplace contract.
type MarketplaceTokenContractParamsUpdated struct {
	TokenContract common.Address
	TokenName     string
	TokenSymbol   string
	TokenParams   IMarketplaceTokenParams
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenContractParamsUpdated is a free log retrieval operation binding the contract event 0x7fe3c496f8848416010aee7f49319e93c40d2414b841001865b1e4db2421f033.
//
// Solidity: event TokenContractParamsUpdated(address indexed tokenContract, string tokenName, string tokenSymbol, (uint256,uint256,uint256,address,address,bool,bool) tokenParams)
func (_Marketplace *MarketplaceFilterer) FilterTokenContractParamsUpdated(opts *bind.FilterOpts, tokenContract []common.Address) (*MarketplaceTokenContractParamsUpdatedIterator, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "TokenContractParamsUpdated", tokenContractRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTokenContractParamsUpdatedIterator{contract: _Marketplace.contract, event: "TokenContractParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenContractParamsUpdated is a free log subscription operation binding the contract event 0x7fe3c496f8848416010aee7f49319e93c40d2414b841001865b1e4db2421f033.
//
// Solidity: event TokenContractParamsUpdated(address indexed tokenContract, string tokenName, string tokenSymbol, (uint256,uint256,uint256,address,address,bool,bool) tokenParams)
func (_Marketplace *MarketplaceFilterer) WatchTokenContractParamsUpdated(opts *bind.WatchOpts, sink chan<- *MarketplaceTokenContractParamsUpdated, tokenContract []common.Address) (event.Subscription, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "TokenContractParamsUpdated", tokenContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceTokenContractParamsUpdated)
				if err := _Marketplace.contract.UnpackLog(event, "TokenContractParamsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenContractParamsUpdated is a log parse operation binding the contract event 0x7fe3c496f8848416010aee7f49319e93c40d2414b841001865b1e4db2421f033.
//
// Solidity: event TokenContractParamsUpdated(address indexed tokenContract, string tokenName, string tokenSymbol, (uint256,uint256,uint256,address,address,bool,bool) tokenParams)
func (_Marketplace *MarketplaceFilterer) ParseTokenContractParamsUpdated(log types.Log) (*MarketplaceTokenContractParamsUpdated, error) {
	event := new(MarketplaceTokenContractParamsUpdated)
	if err := _Marketplace.contract.UnpackLog(event, "TokenContractParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceTokenSuccessfullyPurchasedIterator is returned from FilterTokenSuccessfullyPurchased and is used to iterate over the raw logs and unpacked data for TokenSuccessfullyPurchased events raised by the Marketplace contract.
type MarketplaceTokenSuccessfullyPurchasedIterator struct {
	Event *MarketplaceTokenSuccessfullyPurchased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceTokenSuccessfullyPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceTokenSuccessfullyPurchased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceTokenSuccessfullyPurchased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceTokenSuccessfullyPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceTokenSuccessfullyPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceTokenSuccessfullyPurchased represents a TokenSuccessfullyPurchased event raised by the Marketplace contract.
type MarketplaceTokenSuccessfullyPurchased struct {
	Recipient        common.Address
	MintedTokenPrice *big.Int
	PaidTokensAmount *big.Int
	BuyParams        IMarketplaceBuyParams
	PaymentType      uint8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTokenSuccessfullyPurchased is a free log retrieval operation binding the contract event 0x348fd415ae0c7af831cb6b31e5bbf7487a3302c3347ad7502450d72550a8e33f.
//
// Solidity: event TokenSuccessfullyPurchased(address indexed recipient, uint256 mintedTokenPrice, uint256 paidTokensAmount, ((address,uint256,uint256,uint256),address,uint256,uint256,string) buyParams, uint8 paymentType)
func (_Marketplace *MarketplaceFilterer) FilterTokenSuccessfullyPurchased(opts *bind.FilterOpts, recipient []common.Address) (*MarketplaceTokenSuccessfullyPurchasedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "TokenSuccessfullyPurchased", recipientRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTokenSuccessfullyPurchasedIterator{contract: _Marketplace.contract, event: "TokenSuccessfullyPurchased", logs: logs, sub: sub}, nil
}

// WatchTokenSuccessfullyPurchased is a free log subscription operation binding the contract event 0x348fd415ae0c7af831cb6b31e5bbf7487a3302c3347ad7502450d72550a8e33f.
//
// Solidity: event TokenSuccessfullyPurchased(address indexed recipient, uint256 mintedTokenPrice, uint256 paidTokensAmount, ((address,uint256,uint256,uint256),address,uint256,uint256,string) buyParams, uint8 paymentType)
func (_Marketplace *MarketplaceFilterer) WatchTokenSuccessfullyPurchased(opts *bind.WatchOpts, sink chan<- *MarketplaceTokenSuccessfullyPurchased, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "TokenSuccessfullyPurchased", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceTokenSuccessfullyPurchased)
				if err := _Marketplace.contract.UnpackLog(event, "TokenSuccessfullyPurchased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenSuccessfullyPurchased is a log parse operation binding the contract event 0x348fd415ae0c7af831cb6b31e5bbf7487a3302c3347ad7502450d72550a8e33f.
//
// Solidity: event TokenSuccessfullyPurchased(address indexed recipient, uint256 mintedTokenPrice, uint256 paidTokensAmount, ((address,uint256,uint256,uint256),address,uint256,uint256,string) buyParams, uint8 paymentType)
func (_Marketplace *MarketplaceFilterer) ParseTokenSuccessfullyPurchased(log types.Log) (*MarketplaceTokenSuccessfullyPurchased, error) {
	event := new(MarketplaceTokenSuccessfullyPurchased)
	if err := _Marketplace.contract.UnpackLog(event, "TokenSuccessfullyPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
