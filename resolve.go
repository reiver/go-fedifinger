package fedifinger

import (
	"github.com/reiver/go-erorr"
	"github.com/reiver/go-fediverseid"
	"github.com/reiver/go-webfinger"
)

const (
	errEmptyAcctURI     = erorr.Error("fedifinger: empty acct-uri")
	errEmptyActivityURL = erorr.Error("fedifinger: empty activity-url")
	errEmptyHost        = erorr.Error("fedifinger: empty host")
	errMissingLink      = erorr.Error("fedifinger: missing link")
)

// Resolve resolves a Fediverse-ID to an HTTPS URL using WebFinger.
//
// For example:
//
//	url, err := fedifinger.Resolve("@reiver@mastodon.social")
func Resolve(fediverseID string) (string, error) {

	var fid fediverseid.FediverseID
	{
		var err error
		fid, err = fediverseid.ParseFediverseIDString(fediverseID)
		if nil != err {
			var nada string
			return nada, erorr.Errorf("fedifinger: problem parsing fediverse-id: %w", err)
		}
	}

	var host string
	{
		var found bool
		host, found = fid.Host()
		if !found {
			var nada string
			return nada, errEmptyHost
		}
		if "" == host {
			var nada string
			return nada, errEmptyHost
		}
	}

	var accturi string = fid.AcctURI()
	if "" == accturi {
		var nada string
		return nada, errEmptyAcctURI
	}

	var activityURL string
	{
		var webFingerResponse webfinger.Response
		{
			err := webfinger.Get(&webFingerResponse, host, accturi)
			if nil != err {
				var nada string
				return nada, erorr.Errorf("fedifinger: problem making WebFinger request to ???: %w", err)
			}
		}

		var link *webfinger.Link
		{
			loop: for _, wfrLink := range webFingerResponse.Links {
				if isActivityLink(&wfrLink) {
					link = &wfrLink
					break loop
				}
			}
		}
		if nil == link {
			var nada string
			return nada, errMissingLink
		}

		var found bool
		activityURL, found = link.HRef.Get()
		if !found {
			var nada string
			return nada, errEmptyActivityURL
		}
		if "" == activityURL {
			var nada string
			return nada, errEmptyActivityURL
		}
	}

	return activityURL, nil
}
