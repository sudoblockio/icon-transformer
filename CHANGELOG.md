# Changelog

## [0.3.6](https://github.com/sudoblockio/icon-transformer/compare/v0.3.5...v0.3.6) (2025-07-25)


### Bug Fixes

* ci targets ([1c74cef](https://github.com/sudoblockio/icon-transformer/commit/1c74cef2acfe63d289fbcfd9f3e6242373b2e0e3))

## [0.3.5](https://github.com/sudoblockio/icon-transformer/compare/v0.3.4...v0.3.5) (2025-03-24)


### Bug Fixes

* add redis sentinel pass option ([71de015](https://github.com/sudoblockio/icon-transformer/commit/71de015e040b8aa063f396fe376f3bb0f21ecf71))
* don't die when contracts remove their name ([ab93bfe](https://github.com/sudoblockio/icon-transformer/commit/ab93bfe47a1ec921e67bbd359ba3e84d61d25b99))
* handle token transfer processing errors more gracefully when contracts are abis are changed ([0a94034](https://github.com/sudoblockio/icon-transformer/commit/0a94034bd25c2b6e8307902c287cd135cc412a0e))
* issue with more workers in recovery ([7f7ea8c](https://github.com/sudoblockio/icon-transformer/commit/7f7ea8cdc4ea79bc8dda54cccf90d9e9f9580998))
* skip in routines ([0a5c338](https://github.com/sudoblockio/icon-transformer/commit/0a5c3388509ab3a6fe0c027e646b15b9ec5f7879))
* update regex for primary key ([5e7f4f4](https://github.com/sudoblockio/icon-transformer/commit/5e7f4f4380494d7e4096465d6f9435f55e482516))
* update unit test ([5de1012](https://github.com/sudoblockio/icon-transformer/commit/5de10129b32ffdffde9efe402e4ee1763639af85))

## [0.3.4](https://github.com/sudoblockio/icon-transformer/compare/v0.3.3...v0.3.4) (2023-03-20)


### Bug Fixes

* add decimals to token address balance [#74](https://github.com/sudoblockio/icon-transformer/issues/74) ([2e3ce95](https://github.com/sudoblockio/icon-transformer/commit/2e3ce954f4f6a3c34286db46378afc2cc7a08913))

## [0.3.3](https://github.com/sudoblockio/icon-transformer/compare/v0.3.2...v0.3.3) (2023-03-01)


### Bug Fixes

* add transaction_index to token_transfers to allow sorting [#71](https://github.com/sudoblockio/icon-transformer/issues/71) ([70907cc](https://github.com/sudoblockio/icon-transformer/commit/70907cc73d596a2cd73ff3e7ddd53fa7051267d4))
* add tx index to token_transfers_by_addres ([b6fbf1e](https://github.com/sudoblockio/icon-transformer/commit/b6fbf1e5cbe8e7d67606e31cd1b85c16505cc41f))

## [0.3.2](https://github.com/sudoblockio/icon-transformer/compare/v0.3.1...v0.3.2) (2023-02-07)


### Bug Fixes

* add appropriate indexes for speeding up queries for *_by_address tables [#66](https://github.com/sudoblockio/icon-transformer/issues/66) ([0b55c08](https://github.com/sudoblockio/icon-transformer/commit/0b55c089d83a090c0f3934d54fb6ad56d4175414))

## [0.3.1](https://github.com/sudoblockio/icon-transformer/compare/v0.3.0...v0.3.1) (2022-12-20)


### Bug Fixes

* add count routine to cron ([3c5262b](https://github.com/sudoblockio/icon-transformer/commit/3c5262bc005f714a8e1e8f69dc52f8fad77f4e48))

## [0.3.0](https://github.com/sudoblockio/icon-transformer/compare/v0.2.0...v0.3.0) (2022-11-18)


### Features

* add deduplication preprocessor to crud loader ([447ee0b](https://github.com/sudoblockio/icon-transformer/commit/447ee0b53a74e43a270c17ad60dce985f6cbeb64))
* add metrics to addresses transformer ([05c4b28](https://github.com/sudoblockio/icon-transformer/commit/05c4b28154444234963fa72feb1cfe42a9fc63b1))


### Bug Fixes

* add addr on tx create score transformer ([1b829ef](https://github.com/sudoblockio/icon-transformer/commit/1b829eff65dd86ae21e672b2b94e116554e1ea54))
* add contract metadata to addresses table ([ec638e6](https://github.com/sudoblockio/icon-transformer/commit/ec638e614d0b4a862142035c4a5f0e4cca837e79))
* add token contract deduper to reduce deadlock issues ([9a3df67](https://github.com/sudoblockio/icon-transformer/commit/9a3df674f6109c8c54668d43ba2eecab8eaf81cd))
* adding recovery for key=token_transfer_count_by_token_contract_ ([be3ac3a](https://github.com/sudoblockio/icon-transformer/commit/be3ac3a9b52979d636ce434dcdf02a52c409b203))
* deadlock errors on upsert with metric to count errors ([e8e2a26](https://github.com/sudoblockio/icon-transformer/commit/e8e2a26be46669e94ce47471e2809e96f4a6526e))
* deadlock on find missing routine ([28006c5](https://github.com/sudoblockio/icon-transformer/commit/28006c59cbabd2ca7267338b36c2721d843beb86))
* dropped records from non-string primary keys in removeDuplicatePrimaryKeys func ([45f94d4](https://github.com/sudoblockio/icon-transformer/commit/45f94d46ecb57770a662a0c6ead15683f1126da8))
* load address types for non-mainnet networks ([563957f](https://github.com/sudoblockio/icon-transformer/commit/563957f4d9572fa286a295bea59e018d2dee9d8b))
* log counts not showing up properly [#58](https://github.com/sudoblockio/icon-transformer/issues/58) ([463b1c8](https://github.com/sudoblockio/icon-transformer/commit/463b1c85d53e17bcf116cf270b66021b165454b2))
* returning nil for addreses with no balance / failed txs ([53f806f](https://github.com/sudoblockio/icon-transformer/commit/53f806fbfab7516bf41afa55f835dca28951a2e7))
* soft error handling for address count routine that was overwriting with wrong values ([31c1564](https://github.com/sudoblockio/icon-transformer/commit/31c1564b0edeb4762842d05f25fca250c8267a34))

## [0.2.0](https://github.com/sudoblockio/icon-transformer/compare/v0.1.2...v0.2.0) (2022-08-17)


### Features

* add batching to token_transfers_by_address ([49708e3](https://github.com/sudoblockio/icon-transformer/commit/49708e36ee583c5b3783b00dc38660dcbc302a55))
* add irc3 token transfer and nft id to table ([9ce08fb](https://github.com/sudoblockio/icon-transformer/commit/9ce08fba00d9db9833a4bb54d46b7531e3b3fa77))
* add irc31 TransferSingle method for token transfers ([701c2a3](https://github.com/sudoblockio/icon-transformer/commit/701c2a3534fa4a5754060a3e74f4361f7ba8fbd3))
* add new generic crud base and rebuild token_transfers with batch ([e3f2428](https://github.com/sudoblockio/icon-transformer/commit/e3f2428b9fdd4bac922d7e31bba93925a0ba7b4e))
* add transaction_type classification ([a317e84](https://github.com/sudoblockio/icon-transformer/commit/a317e845a427695a81b067847fedf32f9416fecb))


### Bug Fixes

* error in tx count crud ([0b541f4](https://github.com/sudoblockio/icon-transformer/commit/0b541f46d61a2dea4db6a5ade3e5c6f6b2f4e5e5))
* transaction counts in recovery ([66e3dc8](https://github.com/sudoblockio/icon-transformer/commit/66e3dc873f1d687204a7b89995432b60064364cf))
* transaction type updater which was not comparing the right hash before ([a17fcb0](https://github.com/sudoblockio/icon-transformer/commit/a17fcb06440bbcbca6b9b62b305493ec53f7199b))
