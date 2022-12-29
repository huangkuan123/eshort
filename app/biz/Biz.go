package biz

import (
	"eshort/app/biz/strategy"
	"eshort/pkg/config"
)

var BIZ Business

func SetBusiness() Business {
	s := config.GetString("eshort.app_type")
	if s == "shorturl" {
		BIZ = strategy.ShortURL{}
	}
	if s == "passphrase" {
		BIZ = strategy.Passphrate{}
	}
	return BIZ
}
