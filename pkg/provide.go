package pkg

import (
	"context"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func startProvidingEstimator(h *Host) error {
	ctx := context.Background()

	cid, err := generate_random_cid()
	if err != nil {
		return errors.Wrap(err, " trying to generate new cid")
	}
	log.Info("Start providing content estimator for %s", cid.String())
	//TODO provide Estimator is missing from dht
	err = h.dht.ProvideEstimator(ctx, cid)
	log.Info("Done providing content estimator for cid %s", cid.String())
	if err != nil {
		return errors.Wrap(err, " when providing cid")
	}
	return nil
}
