// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsIAMGroupInvalidPathRule checks the pattern is valid
type AwsIAMGroupInvalidPathRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsIAMGroupInvalidPathRule returns new rule with default attributes
func NewAwsIAMGroupInvalidPathRule() *AwsIAMGroupInvalidPathRule {
	return &AwsIAMGroupInvalidPathRule{
		resourceType:  "aws_iam_group",
		attributeName: "path",
		max:           512,
		min:           1,
		pattern:       regexp.MustCompile(`^(\x{002F})|(\x{002F}[\x{0021}-\x{007F}]+\x{002F})$`),
	}
}

// Name returns the rule name
func (r *AwsIAMGroupInvalidPathRule) Name() string {
	return "aws_iam_group_invalid_path"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMGroupInvalidPathRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMGroupInvalidPathRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMGroupInvalidPathRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMGroupInvalidPathRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"path must be 512 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"path must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^(\x{002F})|(\x{002F}[\x{0021}-\x{007F}]+\x{002F})$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}