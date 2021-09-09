package verus

import (
	"github.com/btcsuite/btcd/chaincfg"
)

const (
	sighashMask                 = 0x1f
	blake2BSighash              = "ZcashSigHash"
	prevoutsHashPersonalization = "ZcashPrevoutHash"
	sequenceHashPersonalization = "ZcashSequencHash"
	outputsHashPersonalization  = "ZcashOutputsHash"

	versionOverwinter        int32  = 3
	versionOverwinterGroupID uint32 = 0x3C48270
	versionSapling                  = 4
	versionSaplingGroupID           = 0x892f2085
)

// Params signifies the chain specific parameters of the Verus network.
type Params struct {
	// TODO: We do not actually need to embed the entire chaincfg params object.
	*chaincfg.Params

	P2SHPrefix  []byte
	P2PKHPrefix []byte
	// TODO: add ID and Quantum prefix fields

	Upgrades []ParamsUpgrade
}

// ParamsUpgrade ...
type ParamsUpgrade struct {
	ActivationHeight uint32
	BranchID         []byte
}

var (
	witnessMarkerBytes = []byte{0x00, 0x01}

	// MainNetParams defines the mainnet configuration.
	MainNetParams = Params{
		Params: &chaincfg.MainNetParams,

		P2PKHPrefix: []byte{0x3C},
		P2SHPrefix:  []byte{0x55},
		// TODO: add ID {0x66} and and Quantum {0x3A} prefix fields

		Upgrades: []ParamsUpgrade{
			{0, []byte{0x19, 0x1B, 0xA8, 0x5B}},
			{227520, []byte{0xBB, 0x09, 0xB8, 0x76}},
		},
	}

	// TestNet3Params defines the testnet configuration.
	TestNet3Params = Params{
		Params: &chaincfg.TestNet3Params,

		P2PKHPrefix: []byte{0x3C},
		P2SHPrefix:  []byte{0x55},
		// TODO: add ID {0x66} and and Quantum {0x3A} prefix fields
		Upgrades: []ParamsUpgrade{
			{0, []byte{0xBB, 0x09, 0xB8, 0x76}},
		},
	}

	// RegressionNetParams defines a devet/regnet configuration.
	RegressionNetParams = Params{
		Params: &chaincfg.RegressionNetParams,

		P2PKHPrefix: []byte{0x3C},
		P2SHPrefix:  []byte{0x55},
		// TODO: add ID {0x66} and and Quantum {0x3A} prefix fields
		Upgrades: []ParamsUpgrade{
			{0, []byte{0xBB, 0x09, 0xB8, 0x76}},
		},
	}
)
