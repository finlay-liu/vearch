// Copyright 2019 The Vearch Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package entity_test

import (
	"encoding/json"
	"testing"

	"github.com/vearch/vearch/internal/util/assert"
	"github.com/vearch/vearch/proto/entity"
)

func TestSpaceString(t *testing.T) {
	space := &entity.Space{}
	str := `{"engine": "caprice"}`
	if err := json.Unmarshal([]byte(str), &space); err != nil {
		t.Errorf(err.Error())
	}
}

func TestEngineSpaceString(t *testing.T) {
	space := &entity.Space{}
	str := `{"engine": {"nprobe":10, "name":"gamma", "index_size":1000}}`
	if err := json.Unmarshal([]byte(str), &space); err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, space.Engine.IndexSize, int64(1000), "unmarshal string to engine err")
}
