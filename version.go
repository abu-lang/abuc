// Copyright 2022 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"runtime/debug"
	"time"

	"golang.org/x/mod/module"
	"golang.org/x/mod/semver"
)

// version tries to parse VCS information encoded in the running executable to
// return a Go pseudo-version string. The pseudo-version string is suffixed with
// "+modified" if the VCS information reports an unclean state. On failure the
// value "unknown" or the revision's ID is returned.
func version() string {
	res := "unknown"
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return res
	}
	if semver.IsValid(bi.Main.Version) {
		return bi.Main.Version
	}
	// fallback to using VCS information
	var rev string
	var cmtt *time.Time
	var mod bool
	for _, s := range bi.Settings {
		switch s.Key {
		case "vcs.revision":
			rev = s.Value
		case "vcs.time":
			t, err := time.Parse(time.RFC3339, s.Value)
			if err != nil {
				break
			}
			cmtt = &t
		case "vcs.modified":
			mod = s.Value == "true"
		}
	}
	if rev == "" {
		return res
	} else if cmtt == nil {
		return rev
	}
	res = module.PseudoVersion("v0", previousVersion, *cmtt, rev[:12])
	if mod {
		res += "+modified"
	}
	return res
}

const license = `
Apache License Version 2.0

Copyright 2022 Massimo Comuzzo, Michele Pasqua and Marino Miculan

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.`
