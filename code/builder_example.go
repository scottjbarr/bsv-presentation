package builder

import (
	"io/ioutil"

	"github.com/btcsuite/btcutil"
	"github.com/scottjbarr/bsv/builder"
	"github.com/scottjbarr/bsv/network"
)

func AddInput() {
	build := builder.New(&network.StressTestNetParams)

	// add an input, provide the UTXO details.
	hash := "c762a29a4beb4821ad843590c3f11ffaed38b7eadc74557bdf36da3539921531"
	index := 0
	value := int64(10000000)
	address := "mupiWN44gq3NZmvZuMMyx8KbRwism69Gbw"

	if _, err := build.AddInput(hash, index, value, address); err != nil {
		panic(err)
	}
}

func AddOutputs() {
	build := builder.New(&network.StressTestNetParams)

	// add an output to the recipient
	recipient := "n1kBjpqmH82jgiRnEHLmFMNv77kvugBomm"
	if _, err := build.AddOutput(recipient, 1000); err != nil {
		panic(err)
	}

	// add the change address
	changeAddr := "mq4htwkZSAG9isuVbEvcLaAiNL59p26W64"
	changeValue := value - 1000 - 200
	if _, err := build.AddOutput(changeAddr, changeValue); err != nil {
		panic(err)
	}
}

func OpReturnData() {
	build := builder.New(&network.StressTestNetParams)

	content, err := ioutil.ReadFile("files/bitcoin.md")
	if err != nil {
		panic(err)
	}

	build.AddOpReturn(content)
}

func Sign() {
	build := builder.New(&network.StressTestNetParams)

	// add your inputs and outputs...

	// load the private key for the STN, not real money :)
	secret := "cQDgbH4C7HP3LSJevMSb1dPMCviCPoLwJ28mxnDRJueMSCa72xjm"

	wif, err := btcutil.DecodeWIF(secret)
	if err != nil {
		panic(err)
	}

	// sign the first input
	if err := build.Sign(privKey, 0); err != nil {
		panic(err)
	}
}

func Serialize() {
	build := builder.New(&network.StressTestNetParams)

	// add inputs, outputs, sign the TX etc..

	// get the raw TX
	b, err := build.Serialize()
	if err != nil {
		panic(err)
	}
}
