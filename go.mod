module github.com/renproject/multichain

go 1.14

require (
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/btcsuite/btcutil v1.0.2
	github.com/cosmos/cosmos-sdk v0.39.1
	github.com/dchest/blake2b v1.0.0
	github.com/ethereum/go-ethereum v1.10.6
	github.com/filecoin-project/go-address v0.0.5
	github.com/filecoin-project/go-jsonrpc v0.1.4-0.20210217175800-45ea43ac2bec
	github.com/filecoin-project/go-state-types v0.1.1-0.20210506134452-99b279731c48
	github.com/filecoin-project/lotus v1.10.0
	github.com/ipfs/go-cid v0.0.7
	github.com/minio/blake2b-simd v0.0.0-20160723061019-3f5f724cb5b1
	github.com/multiformats/go-varint v0.0.6
	github.com/near/borsh-go v0.3.0
	github.com/onsi/ginkgo v1.14.0
	github.com/onsi/gomega v1.10.1
	github.com/renproject/id v0.4.2
	github.com/renproject/pack v0.2.5
	github.com/renproject/solana-ffi v0.1.2
	github.com/renproject/surge v1.2.6
	github.com/tendermint/tendermint v0.33.8
	github.com/terra-project/core v0.4.0-rc.4
	github.com/tyler-smith/go-bip39 v1.1.0
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2
)

replace github.com/cosmos/ledger-cosmos-go => github.com/terra-project/ledger-terra-go v0.11.1-terra

replace github.com/CosmWasm/go-cosmwasm => github.com/terra-project/go-cosmwasm v0.10.1-terra

replace github.com/filecoin-project/filecoin-ffi => ./chain/filecoin/filecoin-ffi

replace github.com/renproject/solana-ffi => ./chain/solana/solana-ffi

replace github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
