package retrieval_success_rate

import (
	"context"
	"github.com/FotiosBistas/retrieval-success-rate/pkg"
	"github.com/ipfs/go-cid"
	log "github.com/sirupsen/logrus"
)

func startProvidingEstimator(h *pkg.Host, cid cid.Cid) {
	ctx := context.Background()

	log.Info("Start providing content estimator for %s", cid.String())
	//TODO provide Estimator is missing from dht
	err := h.dht.ProvideEstimator(ctx, cid)
	log.Info("Done providing content estimator for cid %s", cid.String())
	if err != nil {
		log.Errorf("error providing cid: %s", cid.String())
	}
}
