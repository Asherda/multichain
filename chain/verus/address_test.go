package verus_test

import (
	"bytes"
	"math/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil"
	"github.com/renproject/id"
	"github.com/renproject/multichain/api/address"
	"github.com/renproject/multichain/chain/verus"
)

var _ = Describe("Verus Address", func() {
	Context("address", func() {
		addrEncodeDecoder := verus.NewAddressEncodeDecoder(&verus.RegressionNetParams)

		It("addr pub key hash", func() {
			pk := id.NewPrivKey()
			wif, err := btcutil.NewWIF((*btcec.PrivateKey)(pk), verus.RegressionNetParams.Params, true)
			Expect(err).NotTo(HaveOccurred())
			addrPubKeyHash, err := verus.NewAddressPubKeyHash(btcutil.Hash160(wif.PrivKey.PubKey().SerializeUncompressed()), &verus.RegressionNetParams)
			Expect(err).NotTo(HaveOccurred())
			addr := address.Address(addrPubKeyHash.EncodeAddress())

			decodedRawAddr, err := addrEncodeDecoder.DecodeAddress(addr)
			Expect(err).NotTo(HaveOccurred())
			encodedAddr, err := addrEncodeDecoder.EncodeAddress(decodedRawAddr)
			Expect(err).NotTo(HaveOccurred())
			Expect(encodedAddr).To(Equal(addr))
		})

		It("addr script hash", func() {
			script := make([]byte, rand.Intn(100))
			rand.Read(script)
			addrScriptHash, err := verus.NewAddressScriptHash(script, &verus.RegressionNetParams)
			Expect(err).NotTo(HaveOccurred())
			addr := address.Address(addrScriptHash.EncodeAddress())

			decodedRawAddr, err := addrEncodeDecoder.DecodeAddress(addr)
			Expect(err).NotTo(HaveOccurred())
			encodedAddr, err := addrEncodeDecoder.EncodeAddress(decodedRawAddr)
			Expect(err).NotTo(HaveOccurred())
			Expect(encodedAddr).To(Equal(addr))
		})
	})

	Context("AddressEncodeDecoder", func() {
		It("should give an error when decoding address on different network", func() {
			params := []verus.Params{
				verus.MainNetParams,
				verus.TestNet3Params,
				verus.RegressionNetParams,
			}

			for i, param := range params {
				// Generate a P2PKH address with the params
				pk := id.NewPrivKey()
				wif, err := btcutil.NewWIF((*btcec.PrivateKey)(pk), param.Params, true)
				Expect(err).NotTo(HaveOccurred())
				addrPubKeyHash, err := verus.NewAddressPubKeyHash(btcutil.Hash160(wif.PrivKey.PubKey().SerializeUncompressed()), &param)
				Expect(err).NotTo(HaveOccurred())
				p2pkhAddr := address.Address(addrPubKeyHash.EncodeAddress())

				// Generate a P2SH address with the params
				script := make([]byte, rand.Intn(100))
				rand.Read(script)
				addrScriptHash, err := verus.NewAddressScriptHash(script, &param)
				Expect(err).NotTo(HaveOccurred())
				p2shAddr := address.Address(addrScriptHash.EncodeAddress())

				// Try decode the address using decoders with different network params
				for j := range params {
					addrEncodeDecoder := verus.NewAddressEncodeDecoder(&params[j])
					_, err := addrEncodeDecoder.DecodeAddress(p2pkhAddr)
					// Check the prefix in the params instead of comparing the network directly
					// because testnet and regression network has the same prefix.
					if bytes.Equal(params[i].P2PKHPrefix, params[j].P2PKHPrefix) {
						Expect(err).NotTo(HaveOccurred())
					} else {
						Expect(err).To(HaveOccurred())
					}

					_, err = addrEncodeDecoder.DecodeAddress(p2shAddr)
					if bytes.Equal(params[i].P2PKHPrefix, params[j].P2PKHPrefix) {
						Expect(err).NotTo(HaveOccurred())
					} else {
						Expect(err).To(HaveOccurred())
					}
				}
			}
		})
	})
})
