# Change Log

All notable changes to this project will be documented in this file.

## [1.1.0] - 2022-02-23

### Added
- Support for NFT exchange: add floor price to book
- Support for various networks
- Voucher update event

### Fixed
- Rollback of book updates after losing connection in the tracker
- Duplicate payments

## [1.0.0] - 2022-12-09
 
### Added
- Full system refactoring.
- Service connector.

### Changed
- Instead of a runner that iterates through blocks in each contract batch, we added a separate tracker+consumer pair that uses both rpc (for reading previously skipped events) and web socket (for listening events) which drastically increases speed of catching and processing events. 
- Changed interaction with other services from direct database transactions to calling connectors of other microservices. 

## [1.0.0-rc.0] - 2022-10-11 

### Added
- First relativitely stable version of a tracker service that tracks events from the factory and token contracts and updates other services' databases accordingly.
- Implemented deploy token listener using solely rpc by iterating through blocks with a configurable step. 
- Implemented token successful mint event and update event listeners (without web socket).   

[1.1.0]: https://gitlab.com/tokend/nft-books/contract-tracker/compare/1.0.0...1.1.0
[1.0.0]: https://gitlab.com/tokend/nft-books/contract-tracker/compare/1.0.0-rc.0...1.0.0
[1.0.0-rc.0]: https://gitlab.com/tokend/nft-books/contract-tracker/tags/1.0.0-rc.0