// Copyright 2024 OpenSSF Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//nolint:stylecheck
package hasNoGitHubWorkflowPermissionUnknown

import (
	"embed"
	"fmt"

	"github.com/ossf/scorecard/v5/checker"
	"github.com/ossf/scorecard/v5/finding"
	"github.com/ossf/scorecard/v5/probes/internal/utils/permissions"
	"github.com/ossf/scorecard/v5/probes/internal/utils/uerror"
)

//go:embed *.yml
var fs embed.FS

const Probe = "hasNoGitHubWorkflowPermissionUnknown"

func Run(raw *checker.RawResults) ([]finding.Finding, string, error) {
	if raw == nil {
		return nil, "", fmt.Errorf("%w: raw", uerror.ErrNil)
	}

	results := raw.TokenPermissionsResults
	var findings []finding.Finding

	if results.NumTokens == 0 {
		f, err := finding.NewWith(fs, Probe,
			"No token permissions found",
			nil, finding.OutcomeNotApplicable)
		if err != nil {
			return nil, Probe, fmt.Errorf("create finding: %w", err)
		}
		findings = append(findings, *f)
		return findings, Probe, nil
	}

	for _, r := range results.TokenPermissions {
		if r.Type != checker.PermissionLevelUnknown {
			continue
		}

		// Create finding
		f, err := permissions.CreateFalseFinding(r, Probe, fs, raw.Metadata.Metadata)
		if err != nil {
			return nil, Probe, fmt.Errorf("create finding: %w", err)
		}
		findings = append(findings, *f)
	}

	if len(findings) == 0 {
		f, err := finding.NewWith(fs, Probe,
			"no workflows with unknown permissions",
			nil, finding.OutcomeTrue)
		if err != nil {
			return nil, Probe, fmt.Errorf("create finding: %w", err)
		}
		findings = append(findings, *f)
	}
	return findings, Probe, nil
}
