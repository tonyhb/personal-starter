package types

import (
	"time"
)

type Features struct {
	MaxUsers               int
	MaxPages               int
	MinScanFreq            time.Duration
	ScanHistory            int // how long we keep scans for, in days
	HasSlackNotification   bool
	HasWebhookNotification bool
	HasGrouping            bool
	HasBranding            bool
	HasTechnology          bool
	HasContentChangeAPI    bool
	HasPhoneSupport        bool
}

var (
	Free = Features{
		MaxUsers:    1,
		MaxPages:    3,
		MinScanFreq: time.Hour * 8,
		ScanHistory: 365,
	}
	Starter = Features{
		MaxUsers:    1,
		MaxPages:    15,
		MinScanFreq: time.Hour * 8,
		ScanHistory: 365,
	}
	Pro = Features{
		MaxUsers:               5,
		MaxPages:               45,
		MinScanFreq:            time.Hour * 8,
		ScanHistory:            365,
		HasContentChangeAPI:    true,
		HasSlackNotification:   true,
		HasWebhookNotification: true,
		HasGrouping:            true,
	}
	Agency = Features{
		MaxUsers:               100,
		MaxPages:               200,
		MinScanFreq:            time.Hour * 2,
		HasSlackNotification:   true,
		HasWebhookNotification: true,
		HasGrouping:            true,
		HasBranding:            true,
		HasTechnology:          true,
		HasContentChangeAPI:    true,
		HasPhoneSupport:        true,
	}
)
