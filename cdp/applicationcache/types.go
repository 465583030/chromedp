package applicationcache

// AUTOGENERATED. DO NOT EDIT.

import (
	cdp "github.com/knq/chromedp/cdp"
)

// Resource detailed application cache resource information.
type Resource struct {
	URL  string `json:"url"`  // Resource url.
	Size int64  `json:"size"` // Resource size.
	Type string `json:"type"` // Resource type.
}

// ApplicationCache detailed application cache information.
type ApplicationCache struct {
	ManifestURL  string      `json:"manifestURL"`  // Manifest URL.
	Size         float64     `json:"size"`         // Application cache size.
	CreationTime float64     `json:"creationTime"` // Application cache creation time.
	UpdateTime   float64     `json:"updateTime"`   // Application cache update time.
	Resources    []*Resource `json:"resources"`    // Application cache resources.
}

// FrameWithManifest frame identifier - manifest URL pair.
type FrameWithManifest struct {
	FrameID     cdp.FrameID `json:"frameId"`     // Frame identifier.
	ManifestURL string      `json:"manifestURL"` // Manifest URL.
	Status      int64       `json:"status"`      // Application cache status.
}
