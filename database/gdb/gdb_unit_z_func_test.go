// Copyright 2019 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gdb

import (
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func Test_Func_doQuoteWord(t *testing.T) {
	gtest.Case(t, func() {
		array := map[string]string{
			"user":                   "`user`",
			"user u":                 "user u",
			"user_detail":            "`user_detail`",
			"user,user_detail":       "user,user_detail",
			"user u, user_detail ut": "user u, user_detail ut",
			"u.id asc":               "u.id asc",
			"u.id asc, ut.uid desc":  "u.id asc, ut.uid desc",
		}
		for k, v := range array {
			gtest.Assert(doQuoteWord(k, "`", "`"), v)
		}
	})
}

func Test_Func_doQuoteString(t *testing.T) {
	gtest.Case(t, func() {
		// "user", "user u", "user,user_detail", "user u, user_detail ut", "u.id asc".
		array := map[string]string{
			"user":                             "`user`",
			"user u":                           "`user` u",
			"user,user_detail":                 "`user`,`user_detail`",
			"user u, user_detail ut":           "`user` u,`user_detail` ut",
			"u.id asc":                         "`u`.`id` asc",
			"u.id asc, ut.uid desc":            "`u`.`id` asc,`ut`.`uid` desc",
			"user.user u, user.user_detail ut": "`user`.`user` u,`user`.`user_detail` ut",
		}
		for k, v := range array {
			gtest.Assert(doQuoteString(k, "`", "`"), v)
		}
	})
}

func Test_Func_addTablePrefix(t *testing.T) {
	gtest.Case(t, func() {
		prefix := ""
		array := map[string]string{
			"user":                         "`user`",
			"user u":                       "`user` u",
			"user as u":                    "`user` as u",
			"user,user_detail":             "`user`,`user_detail`",
			"user u, user_detail ut":       "`user` u,`user_detail` ut",
			"user as u, user_detail as ut": "`user` as u,`user_detail` as ut",
		}
		for k, v := range array {
			gtest.Assert(doHandleTableName(k, prefix, "`", "`"), v)
		}
	})
	gtest.Case(t, func() {
		prefix := "gf_"
		array := map[string]string{
			"user":                         "`gf_user`",
			"user u":                       "`gf_user` u",
			"user as u":                    "`gf_user` as u",
			"user,user_detail":             "`gf_user`,`gf_user_detail`",
			"user u, user_detail ut":       "`gf_user` u,`gf_user_detail` ut",
			"user as u, user_detail as ut": "`gf_user` as u,`gf_user_detail` as ut",
		}
		for k, v := range array {
			gtest.Assert(doHandleTableName(k, prefix, "`", "`"), v)
		}
	})
}
