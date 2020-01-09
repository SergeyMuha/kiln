package release

type ExtraConstraint interface {
	Satisfies(set ReleaseRequirement) bool
}

type SatisfyingLocalRelease struct {
	LocalRelease
	ExtraConstraint
}

func (r SatisfyingLocalRelease) Satisfies(rr ReleaseRequirement) bool {
	if r.ExtraConstraint == nil {
		r.ExtraConstraint = noConstraint{}
	}
	return r.Name == rr.Name &&
		r.Version == rr.Version &&
		r.ExtraConstraint.Satisfies(rr)
}

type noConstraint struct{}

func (noConstraint) Satisfies(ReleaseRequirement) bool {
	return true
}

func NewLocalBuiltRelease(id ReleaseID, localPath string) SatisfyingLocalRelease {
	return SatisfyingLocalRelease{
		LocalRelease: LocalRelease{ReleaseID: id, LocalPath: localPath},
	}
}

func NewLocalCompiledRelease(id ReleaseID, stemcellOS, stemcellVersion, localPath string) SatisfyingLocalRelease {
	return SatisfyingLocalRelease{
		LocalRelease: LocalRelease{ReleaseID: id, LocalPath: localPath},
		ExtraConstraint: StemcellConstraint{
			StemcellOS:      stemcellOS,
			StemcellVersion: stemcellVersion,
		},
	}
}

type StemcellConstraint struct {
	StemcellOS      string
	StemcellVersion string
}

func (cr StemcellConstraint) Satisfies(rr ReleaseRequirement) bool {
	return cr.StemcellOS == rr.StemcellOS &&
		cr.StemcellVersion == rr.StemcellVersion
}