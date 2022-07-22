package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/drand/drand/chain"
	"github.com/drand/drand/key"
	"github.com/drand/kyber"
)

type StubInfo struct {
	PublicKey   string        `json:"public_key"`
	Period      time.Duration `json:"period"`
	GenesisTime int64         `json:"genesis_time"`
	GroupHash   []byte        `json:"group_hash"`
}

func main() {
	decodeString, err := hex.DecodeString("8990e7a9aaed2ffed73dbd7092123d6f289930540d7651336225dc172e51b2ce")
	if err != nil {
		panic(err)
	}

	stubInfo := StubInfo{
		PublicKey:   "868f005eb8e6e4ca0a47c8a77ceaa5309a47978a7c71bc5cce96366b5d7a569937c529eeda66c7293784a9402801af31",
		Period:      30,
		GenesisTime: 1595431050,
		GroupHash:   decodeString,
	}

	jsonData, err := json.Marshal(stubInfo)
	if err != nil {
		panic(err)
	}

	info, err := chain.InfoFromJSON(bytes.NewReader(jsonData))
	if err != nil {
		panic(err)
	}
	_ = info
	fmt.Printf("Round # %s", string(jsonData))

	sig, _ := hex.DecodeString("90957ebc0719f8bfb67640aff8ca219bf9f2c5240e60a8711c968d93370d38f87b38ed234a8c63863eb81f234efce55b047478848c0de025527b3d3476dfe860632c1b799550de50a6b9540463e9fb66c8016b89c04a9f52dabdc988e69463c1")
	psig, _ := hex.DecodeString("859504eade86790ad09b2b3474d5e09d1718b549ef7107d7bbd18f5e221765ce8252d7db02664c1f6b20f40c6e8e138704d2acfeb6c5abcc14c77e3a842b2f84515e7366248ca37b1460d23b4f98493c246fbb02851f2a43a710c968a349f8d6")

	Verify(367, sig, psig, info.PublicKey)
	random := chain.RandomnessFromSignature(sig)

	fmt.Println(hex.EncodeToString(random))
}

func Verify(round uint64, sig, prevSig []byte, key kyber.Point) bool {
	err := VerifyBeacon(round, prevSig, sig, key)
	if err != nil {
		panic(err)
	}

	fmt.Println("verified")
	return true
}

func VerifyBeacon(round uint64, prevSig, sig []byte, pubkey kyber.Point) error {
	msg := DigestMessage(round, prevSig)

	return key.Scheme.VerifyRecovered(pubkey, msg, sig)
}

func DigestMessage(currRound uint64, prevSig []byte) []byte {
	h := sha256.New()

	_, _ = h.Write(prevSig)
	_, _ = h.Write(RoundToBytes(currRound))
	return h.Sum(nil)
}

func RoundToBytes(r uint64) []byte {
	var buff bytes.Buffer
	_ = binary.Write(&buff, binary.BigEndian, r)
	return buff.Bytes()
}
