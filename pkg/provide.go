package pkg

import (
	"context"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func StartProvidingEstimator(c context.Context, h *Host) error {

	cid, err := GenerateRandomCid()
	if err != nil {
		return errors.Wrap(err, " trying to generate new cid")
	}
	log.Infof("Start providing content estimator for %s", cid.String())
	err = h.DHT.OptimisticProvide(c, cid)
	log.Infof("Done providing content estimator for cid %s", cid.String())
	if err != nil {
		return errors.Wrap(err, " when providing cid")
	}
	return nil
}
