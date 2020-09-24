// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsDlmLifecyclePolicyInvalidExecutionRoleArnRule checks the pattern is valid
type AwsDlmLifecyclePolicyInvalidExecutionRoleArnRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsDlmLifecyclePolicyInvalidExecutionRoleArnRule returns new rule with default attributes
func NewAwsDlmLifecyclePolicyInvalidExecutionRoleArnRule() *AwsDlmLifecyclePolicyInvalidExecutionRoleArnRule {
	return &AwsDlmLifecyclePolicyInvalidExecutionRoleArnRule{
		resourceType:  "aws_dlm_lifecycle_policy",
		attributeName: "execution_role_arn",
		max:           2048,
		pattern:       regexp.MustCompile(`^arn:aws(-[a-z]{1,3}){0,2}:iam::\d+:role/.*$`),
	}
}

// Name returns the rule name
func (r *AwsDlmLifecyclePolicyInvalidExecutionRoleArnRule) Name() string {
	return "aws_dlm_lifecycle_policy_invalid_execution_role_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDlmLifecyclePolicyInvalidExecutionRoleArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDlmLifecyclePolicyInvalidExecutionRoleArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDlmLifecyclePolicyInvalidExecutionRoleArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDlmLifecyclePolicyInvalidExecutionRoleArnRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"execution_role_arn must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:aws(-[a-z]{1,3}){0,2}:iam::\d+:role/.*$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}