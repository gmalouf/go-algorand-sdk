# Changelog

All notable changes to this project will be documented in this file.

## v2.9.0 - 2025-02-27

<!-- Release notes generated using configuration in .github/release.yml at sdk-release-branchtest -->
### What's Changed

Enhanced support for /v2/blockheaders compatible with [Indexer 3.7.x](https://github.com/algorand/indexer/releases/tag/v3.7.2).

Support for /v2/block/{round}?header-only=true in upcoming go-algorand 4.0.2 release.

#### Enhancements

* build(deps): bump golang.org/x/crypto from 0.29.0 to 0.31.0 by @dependabot in https://github.com/algorand/go-algorand-sdk/pull/675
* tests: cross-sdk heartbeat tests by @algorandskiy in https://github.com/algorand/go-algorand-sdk/pull/676
* consensus: update for minor name change in go-algorand by @cce in https://github.com/algorand/go-algorand-sdk/pull/679
* API: Support for header-only flag on /v2/block algod endpoint. by @gmalouf in https://github.com/algorand/go-algorand-sdk/pull/684

#### Other

* Regenerate code with the latest specification file (9a6c0845) by @github-actions in https://github.com/algorand/go-algorand-sdk/pull/677
* Regenerate code with the latest specification file (5bff0845) by @github-actions in https://github.com/algorand/go-algorand-sdk/pull/680

### New Contributors

* @cce made their first contribution in https://github.com/algorand/go-algorand-sdk/pull/679

**Full Changelog**: https://github.com/gmalouf/go-algorand-sdk/compare/v2.8.0...v2.9.0
