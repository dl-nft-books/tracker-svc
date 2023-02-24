# Change Log

All notable changes to this project will be documented in this file.

## [1.2.0] - 2022-02-23

### Added
- Support for NFT exchange: add floor price to book

### Fixed
- Rollback of book updates after losing connection in the tracker

## [1.1.0] - 2022-01-27

### Added
- Support for various networks
- Voucher update event

### Fixed
- Duplicate payments

## [1.0.0] - 2022-12-09
 
### Added
- Full system refactoring.

### Changed
- Instead of a runner that iterates through blocks in each contract batch, we added a separate tracker+consumer pair that uses both rpc (for reading previously skipped events) and web socket (for listening events) which drastically increases speed of catching and processing events. 
