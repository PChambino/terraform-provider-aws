// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsOrganizationsAccountInvalidRoleNameRule checks the pattern is valid
type AwsOrganizationsAccountInvalidRoleNameRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsOrganizationsAccountInvalidRoleNameRule returns new rule with default attributes
func NewAwsOrganizationsAccountInvalidRoleNameRule() *AwsOrganizationsAccountInvalidRoleNameRule {
	return &AwsOrganizationsAccountInvalidRoleNameRule{
		resourceType:  "aws_organizations_account",
		attributeName: "role_name",
		max:           64,
		pattern:       regexp.MustCompile(`^[\w+=,.@-]{1,64}$`),
	}
}

// Name returns the rule name
func (r *AwsOrganizationsAccountInvalidRoleNameRule) Name() string {
	return "aws_organizations_account_invalid_role_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsOrganizationsAccountInvalidRoleNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsOrganizationsAccountInvalidRoleNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsOrganizationsAccountInvalidRoleNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsOrganizationsAccountInvalidRoleNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"role_name must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\w+=,.@-]{1,64}$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}