package version

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	goVer "github.com/hashicorp/go-version"
	"github.com/sirupsen/logrus"

	"github.com/openshift/rosa/pkg/cache"
	"github.com/openshift/rosa/pkg/clients"
)

//go:generate mockgen -source=retriever.go -package=version -destination=./retriever_mock.go
type Retriever interface {
	RetrieveLatestVersionFromMirror() (*goVer.Version, error)
	RetrievePossibleVersionsFromCache() ([]string, bool)
	RetrievePossibleVersionsFromMirror() ([]string, error)
}

func NewRetriever(spec RetrieverSpec) Retriever {
	return &retriever{
		logger: spec.Logger,
		client: spec.Client,
		cache:  spec.Cache,
	}
}

type RetrieverSpec struct {
	Logger *logrus.Logger
	Client clients.HTTPClient
	Cache  cache.RosaCacheService
}

var _ Retriever = &retriever{}

type retriever struct {
	logger *logrus.Logger
	client clients.HTTPClient
	cache  cache.RosaCacheService
}

func (r retriever) RetrieveLatestVersionFromMirror() (*goVer.Version, error) {
	possibleVersions, err := r.RetrievePossibleVersionsFromMirror()
	if err != nil {
		return nil, fmt.Errorf("there was a problem retrieving possible versions from mirror: %v", err)
	}
	if len(possibleVersions) == 0 {
		return nil, fmt.Errorf("no versions available in mirror %s", baseReleasesFolder)
	}
	latestVersion, err := goVer.NewVersion(possibleVersions[0])
	if err != nil {
		return nil, fmt.Errorf("there was a problem retrieving latest version: %v", err)
	}
	for _, ver := range possibleVersions[1:] {
		curVersion, err := goVer.NewVersion(ver)
		if err != nil {
			continue
		}
		if curVersion.GreaterThan(latestVersion) {
			latestVersion = curVersion
		}
	}
	return latestVersion, nil
}

func (r retriever) RetrievePossibleVersionsFromCache() ([]string, bool) {
	cachedVersions, hasCachedVersions := r.cache.Get(cache.VersionCacheKey)
	if !hasCachedVersions {
		return []string{}, false
	}

	possibleVersions, hasExtracted, _ := cache.ConvertToStringSlice(cachedVersions)
	if !hasExtracted {
		return []string{}, false
	}
	return possibleVersions, true
}

func (r retriever) RetrievePossibleVersionsFromMirror() ([]string, error) {
	var possibleVersions []string

	possibleVersions, gotPossibleVersionsFromCache := r.RetrievePossibleVersionsFromCache()
	if gotPossibleVersionsFromCache {
		return possibleVersions, nil
	}

	resp, err := r.client.Get(baseReleasesFolder)
	if err != nil {
		return []string{}, fmt.Errorf("error setting up request for latest released rosa cli: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < http.StatusOK || resp.StatusCode > 299 {
		return []string{},
			fmt.Errorf("error while requesting latest released rosa cli: %d %s", resp.StatusCode, resp.Status)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return []string{}, fmt.Errorf("error parsing response body: %v", err)
	}
	doc.Find(".file").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(j int, ss *goquery.Selection) {
			if ver, ok := ss.Attr("href"); ok {
				ver = strings.TrimSpace(ver)
				ver = strings.TrimRight(ver, "/")
				if ver != "latest" {
					possibleVersions = append(possibleVersions, ver)
				}
			}
		})
	})
	if err := r.cache.Set(cache.VersionCacheKey, possibleVersions); err != nil {
		r.logger.Debugf("Failed to set possible versions in cache : %v", err)
	}
	r.logger.Debugf("Versions available for download: %v", possibleVersions)
	return possibleVersions, nil
}
