// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsStoragegatewayCachedIscsiVolumeInvalidSourceVolumeArnRule checks the pattern is valid
type AwsStoragegatewayCachedIscsiVolumeInvalidSourceVolumeArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsStoragegatewayCachedIscsiVolumeInvalidSourceVolumeArnRule returns new rule with default attributes
func NewAwsStoragegatewayCachedIscsiVolumeInvalidSourceVolumeArnRule() *AwsStoragegatewayCachedIscsiVolumeInvalidSourceVolumeArnRule {
	return &AwsStoragegatewayCachedIscsiVolumeInvalidSourceVolumeArnRule{
		resourceType:  "aws_storagegateway_cached_iscsi_volume",
		attributeName: "source_volume_arn",
		max:           500,
		min:           50,
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidSourceVolumeArnRule) Name() string {
	return "aws_storagegateway_cached_iscsi_volume_invalid_source_volume_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidSourceVolumeArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidSourceVolumeArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidSourceVolumeArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidSourceVolumeArnRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"source_volume_arn must be 500 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"source_volume_arn must be 50 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}