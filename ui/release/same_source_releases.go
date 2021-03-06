package release

import (
	"fmt"
	"sort"

	bhrelsrepo "github.com/bosh-io/web/release/releasesrepo"
)

type SameSourceReleases struct {
	Source   Source
	Releases []Release
}

func NewSameSourceReleases(source bhrelsrepo.Source, relVerRecs []bhrelsrepo.ReleaseVersionRec, relName string) SameSourceReleases {
	rels := SameSourceReleases{
		Source: NewSource(source),
	}

	for _, relVerRec := range relVerRecs {
		rel := NewIncompleteRelease(relVerRec, relName)
		rels.Releases = append(rels.Releases, rel)
	}

	sort.Sort(sort.Reverse(ReleaseSorting(rels.Releases)))

	if len(rels.Releases) > 0 {
		rels.Releases[0].IsLatest = true
	}

	return rels
}

func (r SameSourceReleases) FirstXReleases(x int) []Release {
	if len(r.Releases) < x {
		return r.Releases
	}

	return r.Releases[0:x]
}

func (r SameSourceReleases) HasMoreThanXReleases(x int) bool {
	return len(r.Releases) > x
}

func (r SameSourceReleases) AvatarURL() string {
	if len(r.Releases) > 0 {
		return r.Releases[0].AvatarURL()
	}

	return ""
}

func (s SameSourceReleases) ForAPI() []Release { return s.Releases }

func (r SameSourceReleases) AllURL() string { return "/releases" }

func (r SameSourceReleases) URL() string {
	return fmt.Sprintf("/releases/%s?all=1", r.Source)
}
