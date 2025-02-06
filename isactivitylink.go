package fedifinger

import (
	"github.com/reiver/go-opt"
	"github.com/reiver/go-webfinger"
)

var activityLinkRel  opt.Optional[string] = opt.Something("self")
var activityLinkType opt.Optional[string] = opt.Something("application/activity+json")

func isActivityLink(link *webfinger.Link) bool {
	if nil == link {
		return false
	}

	if activityLinkRel != link.Rel {
		return false
	}
	if activityLinkType != link.Type {
		return false
	}

	return true
}
