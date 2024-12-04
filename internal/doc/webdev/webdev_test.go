package webdev

import (
	"testing"
)

func TestInitWebDav(t *testing.T) {
	url := "https://toi.teracloud.jp/dav/"
	username := "hxzhouh"
	password := "mVsRKYTW72SvGoPu"
	webDav := InitWebDav(url, username, password, "")

	err := webDav.SyncFile()
	if err != nil {

	}
}
