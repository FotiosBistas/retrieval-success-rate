package pkg

import (
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

	new_prefix := cid.Prefix{
		Version:  1,
		Codec:    cid.Raw,
		MhType:   mh.SHA2_256,
		MhLength: -1,
	}

	new_cid, err := new_prefix.Sum(content)
	if err != nil {
		return cid.Cid{}, errors.Wrap(err, " summing the content")
	}
	log.Infof("New cid %s", new_cid.String())
	return new_cid, nil
}
