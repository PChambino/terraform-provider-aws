// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsGameliftAliasInvalidDescriptionRule checks the pattern is valid
type AwsGameliftAliasInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsGameliftAliasInvalidDescriptionRule returns new rule with default attributes
func NewAwsGameliftAliasInvalidDescriptionRule() *AwsGameliftAliasInvalidDescriptionRule {
	return &AwsGameliftAliasInvalidDescriptionRule{
		resourceType:  "aws_gamelift_alias",
		attributeName: "description",
		max:           1024,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsGameliftAliasInvalidDescriptionRule) Name() string {
	return "aws_gamelift_alias_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGameliftAliasInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGameliftAliasInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGameliftAliasInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGameliftAliasInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"description must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}