// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsRoute53HealthCheckInvalidSearchStringRule checks the pattern is valid
type AwsRoute53HealthCheckInvalidSearchStringRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsRoute53HealthCheckInvalidSearchStringRule returns new rule with default attributes
func NewAwsRoute53HealthCheckInvalidSearchStringRule() *AwsRoute53HealthCheckInvalidSearchStringRule {
	return &AwsRoute53HealthCheckInvalidSearchStringRule{
		resourceType:  "aws_route53_health_check",
		attributeName: "search_string",
		max:           255,
	}
}

// Name returns the rule name
func (r *AwsRoute53HealthCheckInvalidSearchStringRule) Name() string {
	return "aws_route53_health_check_invalid_search_string"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53HealthCheckInvalidSearchStringRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53HealthCheckInvalidSearchStringRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53HealthCheckInvalidSearchStringRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53HealthCheckInvalidSearchStringRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"search_string must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}