// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/utils/cryptography/draft-EIP712Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/cryptography/ECDSAUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/IERC20MetadataUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/utils/SafeERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/IERC721Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/IERC721EnumerableUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC721/utils/ERC721Holder.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "@dlsl/dev-modules/contracts-registry/AbstractDependant.sol";
import "@dlsl/dev-modules/libs/decimals/DecimalsConverter.sol";
import "@dlsl/dev-modules/libs/arrays/Paginator.sol";
import "@dlsl/dev-modules/utils/Globals.sol";

import "./interfaces/IMarketplace.sol";
import "./interfaces/IRoleManager.sol";
import "./interfaces/IContractsRegistry.sol";
import "./interfaces/ITokenFactory.sol";
import "./interfaces/tokens/IERC721MintableToken.sol";

contract Marketplace is IMarketplace, ERC721Holder, AbstractDependant, EIP712Upgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;
    using Paginator for EnumerableSet.AddressSet;
    using DecimalsConverter for uint256;
    using SafeERC20Upgradeable for IERC20MetadataUpgradeable;

    bytes32 internal constant _BUY_TYPEHASH =
        keccak256(
            "Buy(address tokenContract,uint256 futureTokenId,address paymentTokenAddress,uint256 paymentTokenPrice,uint256 discount,uint256 endTimestamp,bytes32 tokenURI)"
        );

    IRoleManager private _roleManager;
    ITokenFactory private _tokenFactory;

    EnumerableSet.AddressSet internal _tokenContracts;
    mapping(address => TokenParams) internal _tokenParams;

    string public baseTokenContractsURI;

    modifier onlyMarketplaceManager() {
        _onlyMarketplaceManager();
        _;
    }

    function __Marketplace_init(
        string memory baseTokenContractsURI_
    ) external override initializer {
        __EIP712_init("Marketplace", "1");

        baseTokenContractsURI = baseTokenContractsURI_;
        // __ReentrancyGuard_init();
    }

    function setDependencies(
        address contractsRegistry_,
        bytes calldata
    ) external override dependant {
        IContractsRegistry registry_ = IContractsRegistry(contractsRegistry_);

        _roleManager = IRoleManager(registry_.getRoleManagerContract());
        _tokenFactory = ITokenFactory(registry_.getTokenFactoryContract());
    }

    function addToken(
        string memory name_,
        string memory symbol_,
        TokenParams memory tokenParams_
    ) external onlyMarketplaceManager returns (address tokenProxy_) {
        require(
            bytes(name_).length > 0 && bytes(symbol_).length > 0,
            "Marketplace: Token name or symbol is empty."
        );
        tokenProxy_ = _tokenFactory.deployToken(name_, symbol_);

        _tokenParams[tokenProxy_] = tokenParams_;

        _tokenContracts.add(tokenProxy_);

        emit TokenContractDeployed(tokenProxy_, name_, symbol_, tokenParams_);
    }

    function updateAllParams(
        address tokenContract_,
        string memory name_,
        string memory symbol_,
        TokenParams memory newTokenParams_
    ) external override onlyMarketplaceManager {
        require(
            _tokenContracts.contains(tokenContract_),
            "Marketplace: Token contract not found."
        );

        require(
            bytes(name_).length > 0 && bytes(symbol_).length > 0,
            "Marketplace: Token name or symbol is empty."
        );

        _tokenParams[tokenContract_] = newTokenParams_;

        IERC721MintableToken(tokenContract_).updateTokenParams(name_, symbol_);

        emit TokenContractParamsUpdated(tokenContract_, name_, symbol_, newTokenParams_);
    }

    // TODO: when not paused?; nonReentrant?
    function buyToken(
        address tokenContract_,
        uint256 futureTokenId_,
        address paymentTokenAddress_,
        uint256 paymentTokenPrice_,
        uint256 discount_,
        uint256 endTimestamp_,
        string memory tokenURI_,
        bytes32 r_,
        bytes32 s_,
        uint8 v_
    ) external payable {
        _verifySignature(
            tokenContract_,
            futureTokenId_,
            paymentTokenAddress_,
            paymentTokenPrice_,
            discount_,
            endTimestamp_,
            tokenURI_,
            r_,
            s_,
            v_
        );

        uint256 amountToPay_;

        if (paymentTokenPrice_ != 0 || paymentTokenAddress_ != address(0)) {
            if (paymentTokenAddress_ == address(0)) {
                amountToPay_ = _payWithETH(tokenContract_, paymentTokenPrice_, discount_);
            } else {
                amountToPay_ = _payWithERC20(
                    tokenContract_,
                    IERC20MetadataUpgradeable(paymentTokenAddress_),
                    paymentTokenPrice_,
                    discount_
                );
            }
        }

        TokenParams storage _currentTokenParams = _tokenParams[tokenContract_];

        _mintToken(tokenContract_, futureTokenId_, tokenURI_);
        MintedTokenInfo memory mintedTokenInfo = MintedTokenInfo(
            futureTokenId_,
            _currentTokenParams.pricePerOneToken,
            tokenURI_
        );

        emit SuccessfullyMinted(
            tokenContract_,
            msg.sender,
            mintedTokenInfo,
            paymentTokenAddress_,
            amountToPay_,
            paymentTokenPrice_,
            discount_
        );
    }

    function buyTokenByNFT(
        address tokenContract_,
        uint256 futureTokenId_,
        address nftAddress_,
        uint256 nftFloorPrice_,
        uint256 tokenId_,
        uint256 endTimestamp_,
        string memory tokenURI_,
        bytes32 r_,
        bytes32 s_,
        uint8 v_
    ) external override {
        TokenParams storage _currentTokenParams = _tokenParams[tokenContract_];

        require(
            _currentTokenParams.isNFTBuyable,
            "Marketplace: This token cannot be purchased with NFT."
        );

        _verifySignature(
            tokenContract_,
            futureTokenId_,
            nftAddress_,
            nftFloorPrice_,
            0, // Discount is zero for NFT by NFT option
            endTimestamp_,
            tokenURI_,
            r_,
            s_,
            v_
        );

        _payWithNFT(tokenContract_, IERC721Upgradeable(nftAddress_), nftFloorPrice_, tokenId_);

        _mintToken(tokenContract_, futureTokenId_, tokenURI_);

        emit SuccessfullyMintedByNFT(
            tokenContract_,
            msg.sender,
            MintedTokenInfo(futureTokenId_, _currentTokenParams.minNFTFloorPrice, tokenURI_),
            nftAddress_,
            tokenId_,
            nftFloorPrice_
        );
    }

    function setBaseTokenContractsURI(
        string memory baseTokenContractsURI_
    ) external override onlyMarketplaceManager {
        baseTokenContractsURI = baseTokenContractsURI_;

        emit BaseTokenContractsURIUpdated(baseTokenContractsURI_);
    }

    function _payWithERC20(
        address tokenContract_,
        IERC20MetadataUpgradeable tokenAddr_,
        uint256 tokenPrice_,
        uint256 discount_
    ) internal returns (uint256) {
        require(msg.value == 0, "Marketplace: Currency amount must be a zero.");

        TokenParams storage _currentTokenParams = _tokenParams[tokenContract_];

        uint256 amountToPay_ = tokenPrice_ != 0
            ? _getAmountAfterDiscount(
                (_currentTokenParams.pricePerOneToken * DECIMAL) / tokenPrice_,
                discount_
            )
            : _currentTokenParams.voucherTokensAmount;

        tokenAddr_.safeTransferFrom(
            msg.sender,
            _currentTokenParams.fundsRecipient == address(0)
                ? address(this)
                : _currentTokenParams.fundsRecipient,
            amountToPay_.from18(tokenAddr_.decimals())
        );

        return amountToPay_;
    }

    function _payWithETH(
        address tokenContract_,
        uint256 ethPrice_,
        uint256 discount_
    ) internal returns (uint256) {
        TokenParams storage _currentTokenParams = _tokenParams[tokenContract_];

        uint256 amountToPay_ = _getAmountAfterDiscount(
            (_currentTokenParams.pricePerOneToken * DECIMAL) / ethPrice_,
            discount_
        );

        require(msg.value >= amountToPay_, "Marketplace: Invalid currency amount.");

        if (
            _currentTokenParams.fundsRecipient != address(0) &&
            _currentTokenParams.fundsRecipient != address(this)
        ) {
            (bool success_, ) = _currentTokenParams.fundsRecipient.call{value: amountToPay_}("");
            require(success_, "Marketplace: Failed to send currency to recipient.");
        }

        uint256 extraCurrencyAmount_ = msg.value - amountToPay_;

        if (extraCurrencyAmount_ > 0) {
            (bool success_, ) = msg.sender.call{value: extraCurrencyAmount_}("");
            require(success_, "Marketplace: Failed to return currency.");
        }

        return amountToPay_;
    }

    function _payWithNFT(
        address tokenContract_,
        IERC721Upgradeable nft_,
        uint256 nftFloorPrice_,
        uint256 tokenId_
    ) internal {
        TokenParams storage _currentTokenParams = _tokenParams[tokenContract_];

        require(
            nftFloorPrice_ >= _currentTokenParams.minNFTFloorPrice,
            "Marketplace: NFT floor price is less than the minimal."
        );
        require(
            IERC721Upgradeable(nft_).ownerOf(tokenId_) == msg.sender,
            "Marketplace: Sender is not the owner."
        );

        nft_.safeTransferFrom(
            msg.sender,
            _currentTokenParams.fundsRecipient == address(0)
                ? address(this)
                : _currentTokenParams.fundsRecipient,
            tokenId_
        );
    }

    function _mintToken(
        address tokenContract_,
        uint256 mintTokenId_,
        string memory tokenURI_
    ) internal {
        IERC721MintableToken(tokenContract_).mint(msg.sender, mintTokenId_, tokenURI_);
    }

    function getUserTokenIDs(
        address tokenContract_,
        address userAddr_
    ) external view override returns (uint256[] memory tokenIDs_) {
        uint256 _tokensCount = IERC721Upgradeable(tokenContract_).balanceOf(userAddr_);

        tokenIDs_ = new uint256[](_tokensCount);

        for (uint256 i; i < _tokensCount; i++) {
            tokenIDs_[i] = IERC721EnumerableUpgradeable(tokenContract_).tokenOfOwnerByIndex(
                userAddr_,
                i
            );
        }
    }

    function getTokenParams(
        address tokenContract_
    ) external view override returns (TokenParams memory) {
        return _tokenParams[tokenContract_];
    }

    function getTokenContractsCount() external view override returns (uint256) {
        return _tokenContracts.length();
    }

    function getTokenContractsPart(
        uint256 offset_,
        uint256 limit_
    ) external view override returns (address[] memory) {
        return _tokenContracts.part(offset_, limit_);
    }

    function _verifySignature(
        address tokenContract_,
        uint256 futureTokenId_,
        address paymentTokenAddress_,
        uint256 paymentTokenPrice_,
        uint256 discount_,
        uint256 endTimestamp_,
        string memory tokenURI_,
        bytes32 r_,
        bytes32 s_,
        uint8 v_
    ) internal view {
        bytes32 structHash_ = keccak256(
            abi.encode(
                _BUY_TYPEHASH,
                tokenContract_,
                futureTokenId_,
                paymentTokenAddress_,
                paymentTokenPrice_,
                discount_,
                endTimestamp_,
                keccak256(abi.encodePacked(tokenURI_))
            )
        );

        address signer_ = ECDSAUpgradeable.recover(_hashTypedDataV4(structHash_), v_, r_, s_);

        require(_roleManager.isAdmin(signer_), "Marketplace: Invalid signature.");
        require(block.timestamp <= endTimestamp_, "Marketplace: Signature expired.");
    }

    function _onlyMarketplaceManager() internal view {
        require(
            _roleManager.isMarketplaceManager(msg.sender),
            "Marketplace: Caller is not a marketplace manager."
        );
    }

    function _getAmountAfterDiscount(
        uint256 amount_,
        uint256 discount_
    ) internal pure returns (uint256) {
        return (amount_ * (PERCENTAGE_100 - discount_)) / PERCENTAGE_100;
    }
}