package visit

import "strings"

type (
	VisitOS      string
	VisitBrowser string
)

const (
	VisitOSMac     VisitOS = "Mac"
	VisitOSiPhone  VisitOS = "iPhone"
	VisitOSLinux   VisitOS = "Linux"
	VisitOSWindows VisitOS = "Windows"
	VisitOSAndroid VisitOS = "Android"
	VisitOSOther   VisitOS = "Other"
)

const (
	VisitBrowserChrome  VisitBrowser = "Chrome"
	VisitBrowserFirefox VisitBrowser = "Firefox"
	VisitBrowserSafari  VisitBrowser = "Safari"
	VisitBrowserOther   VisitBrowser = "Other"
)

type UserAgent struct {
	OS      VisitOS
	Browser VisitBrowser
}

func ParseUserAgent(ua string) *UserAgent {
	os := getOsName(ua)
	browser := getBrowser(ua)
	return &UserAgent{
		OS:      os,
		Browser: browser,
	}
}

func getOsName(ua string) VisitOS {
	macIdx := strings.Index(ua, "Mac")
	if strings.Index(ua, "iPhone") > 0 && macIdx > 0 {
		return VisitOSiPhone
	}
	if strings.Index(ua, "Macintosh") > 0 && macIdx > 0 {
		return VisitOSMac
	}

	linuxIdx := strings.Index(ua, "Linux")
	androidIdx := strings.Index(ua, "Android")

	if linuxIdx > 0 && androidIdx == -1 {
		return VisitOSLinux
	}
	if linuxIdx > 0 && androidIdx > 0 {
		return VisitOSAndroid
	}

	if strings.Index(ua, "Windows") > 0 {
		return VisitOSWindows
	}

	return VisitOSOther
}

func getBrowser(ua string) VisitBrowser {
	if strings.Index(ua, "Firefox") > 0 {
		return VisitBrowserFirefox
	}
	chromeIdx := strings.Index(ua, "Chrome")
	safariIdx := strings.Index(ua, "Safari")
	if safariIdx > 0 && chromeIdx < 0 {
		return VisitBrowserSafari
	}
	if chromeIdx > 0 {
		return VisitBrowserChrome
	}
	return VisitBrowserOther
}
