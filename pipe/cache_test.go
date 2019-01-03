package pipe

import (
	"testing"

	"github.com/uber/storagetapper/test"
)

func TestPipeCacheBasic(t *testing.T) {
	pcfg := cfg.Pipe
	p, err := CacheGet("local", &pcfg, nil)
	test.CheckFail(err, t)
	test.Assert(t, p.Type() == "local", "should be local type")
	_, err = CacheGet("local", &pcfg, nil)
	test.CheckFail(err, t)
	test.Assert(t, len(cache) == 1, "second local CacheGet shouldn't create an entry")

	p, err = CacheGet("file", &pcfg, nil)
	test.CheckFail(err, t)
	test.Assert(t, len(cache) == 2, "an entry should be created for first file pipe")
	test.Assert(t, p.Type() == "file", "should be file type")
	_, err = CacheGet("file", &pcfg, nil)
	test.CheckFail(err, t)
	test.Assert(t, len(cache) == 2, "second file CacheGet shouldn't create an entry")

	pcfg1 := cfg.Pipe
	pcfg1.BaseDir = "/tmp/somedir"

	p, err = CacheGet("file", &pcfg1, nil)
	test.CheckFail(err, t)
	test.Assert(t, p.Type() == "file", "should be file type")
	test.Assert(t, len(cache) == 3, "different config, should create an entry")
	test.Assert(t, p.Config().BaseDir == "/tmp/somedir", "check that config is ours")

	p, err = CacheGet("file", &pcfg, nil)
	test.CheckFail(err, t)
	test.Assert(t, p.Type() == "file", "should be file type")
	test.Assert(t, p.Config().BaseDir != "/tmp/somedir", "check that we can still get pipe with original config")

	CacheDestroy()

	test.Assert(t, len(cache) == 0, "destroyed")
}
