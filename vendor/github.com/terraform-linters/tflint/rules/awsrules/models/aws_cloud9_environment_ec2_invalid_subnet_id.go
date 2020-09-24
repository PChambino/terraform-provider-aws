// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsCloud9EnvironmentEc2InvalidSubnetIDRule checks the pattern is valid
type AwsCloud9EnvironmentEc2InvalidSubnetIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCloud9EnvironmentEc2InvalidSubnetIDRule returns new rule with default attributes
func NewAwsCloud9EnvironmentEc2InvalidSubnetIDRule() *AwsCloud9EnvironmentEc2InvalidSubnetIDRule {
	return &AwsCloud9EnvironmentEc2InvalidSubnetIDRule{
		resourceType:  "aws_cloud9_environment_ec2",
		attributeName: "subnet_id",
		max:           30,
		min:           5,
	}
}

// Name returns the rule name
func (r *AwsCloud9EnvironmentEc2InvalidSubnetIDRule) Name() string {
	return "aws_cloud9_environment_ec2_invalid_subnet_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloud9EnvironmentEc2InvalidSubnetIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloud9EnvironmentEc2InvalidSubnetIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloud9EnvironmentEc2InvalidSubnetIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloud9EnvironmentEc2InvalidSubnetIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"subnet_id must be 30 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"subnet_id must be 5 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}