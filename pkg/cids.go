package pkg

import (
	"crypto/sha256"
	"github.com/ipfs/go-cid"
	mh "github.com/multiformats/go-multihash"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"math/rand"
)

func GenerateRandomCid() (cid.Cid, error) {

	content := make([]byte, 1024)

	if _, err := rand.Read(content); err != nil {
		return cid.Cid{}, errors.Wrap(err, "while trying to generate random cid")
	}

	hash := sha256.New()
	hash.Write(content)

	mhash, err := mh.Encode(hash.Sum(nil), mh.SHA2_256)

	if err != nil {
		return cid.Cid{}, errors.Wrap(err, "while generating multihash")
	}

	new_cid := cid.NewCidV0(mhash)
	log.Infof("New cid %s", new_cid.String())
	return new_cid, nil
}
