package types

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/drand/drand/chain"
	"github.com/drand/drand/key"
	"github.com/drand/kyber"
)

func (p *ProvenRandomness) Verify(c *ChainInfo) error {
	groupHashBytes, err := hex.DecodeString("8990e7a9aaed2ffed73dbd7092123d6f289930540d7651336225dc172e51b2ce")
	if err != nil {
		return err
	}

	groupHashBase64 := base64.StdEncoding.EncodeToString(groupHashBytes)
	chainInfoJson := fmt.Sprintf(`{
		"public_key": "%s",
		"period": %d,
		"genesis_time": %d,
		"group_hash": "%s"
	}`, c.PublicKey, c.Period, c.GenesisTime, groupHashBase64)

	info, err := chain.InfoFromJSON(
		bytes.NewReader([]byte(chainInfoJson)))
	if err != nil {
		return err
	}

	sig, err := hex.DecodeString(p.PreviousSignature)
	if err != nil {
		return err
	}

	psig, err := hex.DecodeString(p.Signature)
	if err != nil {
		return err
	}

	err = p.verifyBeacon(p.Round, sig, psig, info.PublicKey)
	if err != nil {
		return err
	}

	random := chain.RandomnessFromSignature(sig)
	if hex.EncodeToString(random) != p.Randomness {
		return errors.New("The recovered randomness does not match the one sent")
	}

	return nil
}

func (p *ProvenRandomness) verifyBeacon(round uint64, prevSig, sig []byte, pubkey kyber.Point) error {
	msg := p.digestMessage(round, prevSig)

	return key.Scheme.VerifyRecovered(pubkey, msg, sig)
}

func (p *ProvenRandomness) digestMessage(currRound uint64, prevSig []byte) []byte {
	h := sha256.New()

	_, _ = h.Write(prevSig)
	_, _ = h.Write(p.roundToBytes(currRound))
	return h.Sum(nil)
}

func (p *ProvenRandomness) roundToBytes(r uint64) []byte {
	var buff bytes.Buffer
	_ = binary.Write(&buff, binary.BigEndian, r)
	return buff.Bytes()
}
