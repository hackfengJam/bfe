// Copyright (c) 2019 The BFE Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mod_geo

import (
	"fmt"
	"testing"
)

func TestConfModGeoCase1(t *testing.T) {
	config, err := ConfLoad("./test_data/mod_geo/mod_geo.conf", "")
	if err != nil {
		msg := fmt.Sprintf("confModGeoLoad():err=%s", err.Error())
		t.Error(msg)
		return
	}

	if config.Basic.GeoDBPath != GeoDBDefaultPath {
		t.Error("GeoDBPath should be mod_geo/geo.db")
	}
}

func TestConfModGeoCase2(t *testing.T) {
	config, err := ConfLoad("./test_data/mod_geo/mod_geo1.conf", "")
	if err != nil {
		msg := fmt.Sprintf("confModGeoLoad():err=%s", err.Error())
		t.Error(msg)
		return
	}

	// GeoDBPath is empty, default use mod_geo/geo.db
	if config.Basic.GeoDBPath != GeoDBDefaultPath {
		t.Error("GeoDBPath should be mod_geo/geo.db")
	}
}
