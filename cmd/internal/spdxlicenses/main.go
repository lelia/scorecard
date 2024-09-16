// Copyright 2024 OpenSSF Scorecard Authors
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

/*
Tool to get the latest list of supported licenses from SPDX (https://spdx.org/licenses).
Ref: https://github.com/ossf/scorecard/issues/4031

Intended usage:

1. Use the tool to generate a file (or write to stdout) with a list of supported SPDX licenses.

2. Update `fsfOsiApprovedLicenseMap` in checks/raw/license.go with the content of the file (or stdout).

3. Commit the changes to checks/raw/license.go and file a PR updating the SPDX license list.
*/
package main

import "github.com/ossf/scorecard/v5/cmd/internal/spdxlicenses/app"

func main() {
	app.Execute()
}
