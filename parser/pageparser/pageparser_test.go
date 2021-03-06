// Copyright 2018 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pageparser

import (
	"strings"
	"testing"
)

func BenchmarkParse(b *testing.B) {
	start := `
	

---
title: "Front Matters"
description: "It really does"
---

This is some summary. This is some summary. This is some summary. This is some summary.

 <!--more-->


`
	input := []byte(start + strings.Repeat(strings.Repeat("this is text", 30)+"{{< myshortcode >}}This is some inner content.{{< /myshortcode >}}", 10))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := parseBytes(input); err != nil {
			b.Fatal(err)
		}
	}
}
