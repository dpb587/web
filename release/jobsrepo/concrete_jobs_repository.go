package jobsrepo

import (
	"encoding/json"

	boshlog "github.com/cloudfoundry/bosh-agent/logger"
	bosherr "github.com/cloudfoundry/bosh-agent/errors"
	bpreljob "github.com/cppforlife/bosh-provisioner/release/job"

	bhrelver "github.com/cppforlife/bosh-hub/release/relver"
	bhrelsrepo "github.com/cppforlife/bosh-hub/release/releasesrepo"
)

type JobsRepository interface {
	FindAll(bhrelsrepo.ReleaseVersionRec) ([]bpreljob.Job, error)
}

type CJRepository struct {
	relVerFactory bhrelver.Factory
	logger boshlog.Logger
}

type relVerRecKey struct {
	Source     string
	VersionRaw string
}

func NewConcreteJobsRepository(
	relVerFactory bhrelver.Factory,
	logger boshlog.Logger,
) CJRepository {
	return CJRepository{
		relVerFactory: relVerFactory,
		logger: logger,
	}
}

func (r CJRepository) FindAll(relVerRec bhrelsrepo.ReleaseVersionRec) ([]bpreljob.Job, error) {
	var relJobs []bpreljob.Job

	relVer, err := r.relVerFactory.Find(relVerRec.Source, relVerRec.VersionRaw)
	if err != nil {
		return nil, err
	}

	contents, err := relVer.Read("jobs.v1.yml")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &relJobs)
	if err != nil {
		return nil, bosherr.WrapError(err, "Unmarshaling release jobs")
	}

	return relJobs, nil
}
