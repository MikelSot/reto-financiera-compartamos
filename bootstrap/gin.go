package bootstrap

import (
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
	rkgin "github.com/rookie-ninja/rk-gin/v2/boot"
)

func newGinEntry(boot []byte) *rkgin.GinEntry {
	rkentry.BootstrapBuiltInEntryFromYAML(boot)

	entries := rkgin.RegisterGinEntryYAML(boot)

	return entries["service-gin"].(*rkgin.GinEntry)
}
