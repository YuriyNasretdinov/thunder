// generated by: generate_version_info.sh "1.0.0" "vvv@corp.badoo.com" "" "/Applications/Xcode.app/Contents/Developer/usr/bin/make - " 
// DO NOT EDIT!

package main

import (
	"badoo/_packages/service"
	"badoo/_packages/service/stats"
	"github.com/gogo/protobuf/proto"
)

func init() {
	service.VersionInfo = badoo_service.ResponseVersion{
		Version:        proto.String("1.0.0"),
		Maintainer:     proto.String("vvv@corp.badoo.com"),
		AutoBuildTag:   proto.String(""),
		BuildDate:      proto.String("Thu Sep 15 10:26:36 UTC 2016"),
		BuildHost:      proto.String("nasretdinov.local"),
		BuildGoVersion: proto.String("go version go1.7 darwin/amd64"),
		BuildCommand:   proto.String("/Applications/Xcode.app/Contents/Developer/usr/bin/make - "),
		VcsType:        proto.String("git"),
		VcsBasename:    proto.String("go-proxyd"),
		VcsNum:         proto.String("31"),
		VcsDate:        proto.String("2014-02-17T16:03:45+0400"),
		VcsBranch:      proto.String("master"),
		VcsTag:         proto.String(""),
		VcsTick:        proto.String("31"),
		VcsFullHash:    proto.String("a10437837ce1c22015dbcaef9796644cadb8f0b1"),
		VcsShortHash:   proto.String("a104378"),
		VcsWcModified:  proto.String("1"),
	}
}