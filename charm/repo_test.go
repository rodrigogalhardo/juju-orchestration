package charm_test

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	. "launchpad.net/gocheck"
	"launchpad.net/juju-core/charm"
	"launchpad.net/juju-core/environs/config"
	"launchpad.net/juju-core/log"
	"launchpad.net/juju-core/testing"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type MockStore struct {
	mux          *http.ServeMux
	lis          net.Listener
	bundleBytes  []byte
	bundleSha256 string
	downloads    []*charm.URL
}

func NewMockStore(c *C) *MockStore {
	s := &MockStore{}
	bytes, err := ioutil.ReadFile(testing.Charms.BundlePath(c.MkDir(), "dummy"))
	c.Assert(err, IsNil)
	s.bundleBytes = bytes
	h := sha256.New()
	h.Write(bytes)
	s.bundleSha256 = hex.EncodeToString(h.Sum(nil))
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/charm-info", func(w http.ResponseWriter, r *http.Request) {
		s.ServeInfo(w, r)
	})
	s.mux.HandleFunc("/charm/", func(w http.ResponseWriter, r *http.Request) {
		s.ServeCharm(w, r)
	})
	lis, err := net.Listen("tcp", "127.0.0.1:4444")
	c.Assert(err, IsNil)
	s.lis = lis
	go http.Serve(s.lis, s)
	return s
}

func (s *MockStore) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *MockStore) ServeInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	response := map[string]*charm.InfoResponse{}
	for _, url := range r.Form["charms"] {
		cr := &charm.InfoResponse{}
		response[url] = cr
		curl := charm.MustParseURL(url)
		switch curl.Name {
		case "borken":
			cr.Errors = append(cr.Errors, "badness")
			continue
		case "unwise":
			cr.Warnings = append(cr.Warnings, "foolishness")
			fallthrough
		default:
			if curl.Revision == -1 {
				cr.Revision = 23
			} else {
				cr.Revision = curl.Revision
			}
			cr.Sha256 = s.bundleSha256
		}
	}
	data, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		panic(err)
	}
}

func (s *MockStore) ServeCharm(w http.ResponseWriter, r *http.Request) {
	curl := charm.MustParseURL("cs:" + r.URL.Path[len("/charm/"):])
	s.downloads = append(s.downloads, curl)
	w.Header().Set("Connection", "close")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", strconv.Itoa(len(s.bundleBytes)))
	_, err := w.Write(s.bundleBytes)
	if err != nil {
		panic(err)
	}
}

type StoreSuite struct {
	server      *MockStore
	store       charm.Repository
	oldJujuHome string
}

var _ = Suite(&StoreSuite{})

func (s *StoreSuite) SetUpSuite(c *C) {
	s.server = NewMockStore(c)
}

func (s *StoreSuite) SetUpTest(c *C) {
	s.oldJujuHome = config.SetJujuHome(c.MkDir())
	s.store = charm.NewStore("http://127.0.0.1:4444")
	s.server.downloads = nil
}

func (s *StoreSuite) TearDownSuite(c *C) {
	config.SetJujuHome(s.oldJujuHome)
	s.server.lis.Close()
}

func (s *StoreSuite) TestError(c *C) {
	curl := charm.MustParseURL("cs:series/borken")
	expect := `charm info errors for "cs:series/borken": badness`
	_, err := s.store.Latest(curl)
	c.Assert(err, ErrorMatches, expect)
	_, err = s.store.Get(curl)
	c.Assert(err, ErrorMatches, expect)
}

func (s *StoreSuite) TestWarning(c *C) {
	defer log.SetTarget(log.SetTarget(c))
	curl := charm.MustParseURL("cs:series/unwise")
	expect := `.* WARNING charm: charm store reports for "cs:series/unwise": foolishness` + "\n"
	r, err := s.store.Latest(curl)
	c.Assert(r, Equals, 23)
	c.Assert(err, IsNil)
	c.Assert(c.GetTestLog(), Matches, expect)
	ch, err := s.store.Get(curl)
	c.Assert(ch, NotNil)
	c.Assert(err, IsNil)
	c.Assert(c.GetTestLog(), Matches, expect+expect)
}

func (s *StoreSuite) TestLatest(c *C) {
	for _, str := range []string{
		"cs:series/blah",
		"cs:series/blah-2",
		"cs:series/blah-99",
	} {
		r, err := s.store.Latest(charm.MustParseURL(str))
		c.Assert(r, Equals, 23)
		c.Assert(err, IsNil)
	}
}

func (s *StoreSuite) assertCached(c *C, curl *charm.URL) {
	s.server.downloads = nil
	ch, err := s.store.Get(curl)
	c.Assert(err, IsNil)
	c.Assert(ch, NotNil)
	c.Assert(s.server.downloads, IsNil)
}

func (s *StoreSuite) TestGetCacheImplicitRevision(c *C) {
	base := "cs:series/blah"
	curl := charm.MustParseURL(base)
	revCurl := charm.MustParseURL(base + "-23")
	ch, err := s.store.Get(curl)
	c.Assert(err, IsNil)
	c.Assert(ch, NotNil)
	c.Assert(s.server.downloads, DeepEquals, []*charm.URL{revCurl})
	s.assertCached(c, curl)
	s.assertCached(c, revCurl)
}

func (s *StoreSuite) TestGetCacheExplicitRevision(c *C) {
	base := "cs:series/blah-12"
	curl := charm.MustParseURL(base)
	ch, err := s.store.Get(curl)
	c.Assert(err, IsNil)
	c.Assert(ch, NotNil)
	c.Assert(s.server.downloads, DeepEquals, []*charm.URL{curl})
	s.assertCached(c, curl)
}

func (s *StoreSuite) TestGetBadCache(c *C) {
	c.Assert(os.Mkdir(config.JujuHomePath("cache"), 0777), IsNil)
	base := "cs:series/blah"
	curl := charm.MustParseURL(base)
	revCurl := charm.MustParseURL(base + "-23")
	name := charm.Quote(revCurl.String()) + ".charm"
	err := ioutil.WriteFile(config.JujuHomePath("cache", name), nil, 0666)
	c.Assert(err, IsNil)
	ch, err := s.store.Get(curl)
	c.Assert(err, IsNil)
	c.Assert(ch, NotNil)
	c.Assert(s.server.downloads, DeepEquals, []*charm.URL{revCurl})
	s.assertCached(c, curl)
	s.assertCached(c, revCurl)
}

type LocalRepoSuite struct {
	testing.LoggingSuite
	repo       *charm.LocalRepository
	seriesPath string
}

var _ = Suite(&LocalRepoSuite{})

func (s *LocalRepoSuite) SetUpTest(c *C) {
	s.LoggingSuite.SetUpTest(c)
	root := c.MkDir()
	s.repo = &charm.LocalRepository{root}
	s.seriesPath = filepath.Join(root, "series")
	c.Assert(os.Mkdir(s.seriesPath, 0777), IsNil)
}

func (s *LocalRepoSuite) addBundle(name string) string {
	return testing.Charms.BundlePath(s.seriesPath, name)
}

func (s *LocalRepoSuite) addDir(name string) string {
	return testing.Charms.ClonedDirPath(s.seriesPath, name)
}

func (s *LocalRepoSuite) TestMissingCharm(c *C) {
	_, err := s.repo.Latest(charm.MustParseURL("local:series/zebra"))
	c.Assert(err, ErrorMatches, `no charms found matching "local:series/zebra" in `+s.repo.Path)
	_, err = s.repo.Get(charm.MustParseURL("local:series/zebra"))
	c.Assert(err, ErrorMatches, `no charms found matching "local:series/zebra" in `+s.repo.Path)
	_, err = s.repo.Latest(charm.MustParseURL("local:badseries/zebra"))
	c.Assert(err, ErrorMatches, `no charms found matching "local:badseries/zebra" in `+s.repo.Path)
	_, err = s.repo.Get(charm.MustParseURL("local:badseries/zebra"))
	c.Assert(err, ErrorMatches, `no charms found matching "local:badseries/zebra" in `+s.repo.Path)
}

func (s *LocalRepoSuite) TestMissingRepo(c *C) {
	c.Assert(os.RemoveAll(s.repo.Path), IsNil)
	_, err := s.repo.Latest(charm.MustParseURL("local:series/zebra"))
	c.Assert(err, ErrorMatches, `no repository found at ".*"`)
	_, err = s.repo.Get(charm.MustParseURL("local:series/zebra"))
	c.Assert(err, ErrorMatches, `no repository found at ".*"`)
	c.Assert(ioutil.WriteFile(s.repo.Path, nil, 0666), IsNil)
	_, err = s.repo.Latest(charm.MustParseURL("local:series/zebra"))
	c.Assert(err, ErrorMatches, `no repository found at ".*"`)
	_, err = s.repo.Get(charm.MustParseURL("local:series/zebra"))
	c.Assert(err, ErrorMatches, `no repository found at ".*"`)
}

func (s *LocalRepoSuite) TestMultipleVersions(c *C) {
	curl := charm.MustParseURL("local:series/upgrade")
	s.addDir("upgrade1")
	rev, err := s.repo.Latest(curl)
	c.Assert(err, IsNil)
	c.Assert(rev, Equals, 1)
	ch, err := s.repo.Get(curl)
	c.Assert(err, IsNil)
	c.Assert(ch.Revision(), Equals, 1)

	s.addDir("upgrade2")
	rev, err = s.repo.Latest(curl)
	c.Assert(err, IsNil)
	c.Assert(rev, Equals, 2)
	ch, err = s.repo.Get(curl)
	c.Assert(err, IsNil)
	c.Assert(ch.Revision(), Equals, 2)

	revCurl := curl.WithRevision(1)
	rev, err = s.repo.Latest(revCurl)
	c.Assert(err, IsNil)
	c.Assert(rev, Equals, 2)
	ch, err = s.repo.Get(revCurl)
	c.Assert(err, IsNil)
	c.Assert(ch.Revision(), Equals, 1)

	badRevCurl := curl.WithRevision(33)
	rev, err = s.repo.Latest(badRevCurl)
	c.Assert(err, IsNil)
	c.Assert(rev, Equals, 2)
	ch, err = s.repo.Get(badRevCurl)
	c.Assert(err, ErrorMatches, `no charms found matching "local:series/upgrade-33" in `+s.repo.Path)
}

func (s *LocalRepoSuite) TestBundle(c *C) {
	curl := charm.MustParseURL("local:series/dummy")
	s.addBundle("dummy")

	rev, err := s.repo.Latest(curl)
	c.Assert(err, IsNil)
	c.Assert(rev, Equals, 1)
	ch, err := s.repo.Get(curl)
	c.Assert(err, IsNil)
	c.Assert(ch.Revision(), Equals, 1)
}

func (s *LocalRepoSuite) TestLogsErrors(c *C) {
	err := ioutil.WriteFile(filepath.Join(s.seriesPath, "blah.charm"), nil, 0666)
	c.Assert(err, IsNil)
	err = os.Mkdir(filepath.Join(s.seriesPath, "blah"), 0666)
	c.Assert(err, IsNil)
	samplePath := s.addDir("upgrade2")
	gibberish := []byte("don't parse me by")
	err = ioutil.WriteFile(filepath.Join(samplePath, "metadata.yaml"), gibberish, 0666)
	c.Assert(err, IsNil)

	curl := charm.MustParseURL("local:series/dummy")
	s.addDir("dummy")
	ch, err := s.repo.Get(curl)
	c.Assert(err, IsNil)
	c.Assert(ch.Revision(), Equals, 1)
	c.Assert(c.GetTestLog(), Matches, `
.* WARNING charm: failed to load charm at ".*/series/blah": .*
.* WARNING charm: failed to load charm at ".*/series/blah.charm": .*
.* WARNING charm: failed to load charm at ".*/series/upgrade2": .*
`[1:])
}

func renameSibling(c *C, path, name string) {
	c.Assert(os.Rename(path, filepath.Join(filepath.Dir(path), name)), IsNil)
}

func (s *LocalRepoSuite) TestIgnoresUnpromisingNames(c *C) {
	err := ioutil.WriteFile(filepath.Join(s.seriesPath, "blah.notacharm"), nil, 0666)
	c.Assert(err, IsNil)
	err = os.Mkdir(filepath.Join(s.seriesPath, ".blah"), 0666)
	c.Assert(err, IsNil)
	renameSibling(c, s.addDir("dummy"), ".dummy")
	renameSibling(c, s.addBundle("dummy"), "dummy.notacharm")
	curl := charm.MustParseURL("local:series/dummy")

	_, err = s.repo.Get(curl)
	c.Assert(err, ErrorMatches, `no charms found matching "local:series/dummy" in `+s.repo.Path)
	_, err = s.repo.Latest(curl)
	c.Assert(err, ErrorMatches, `no charms found matching "local:series/dummy" in `+s.repo.Path)
	c.Assert(c.GetTestLog(), Equals, "")
}
