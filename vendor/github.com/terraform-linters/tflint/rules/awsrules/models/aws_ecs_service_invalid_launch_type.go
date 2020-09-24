// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsEcsServiceInvalidLaunchTypeRule checks the pattern is valid
type AwsEcsServiceInvalidLaunchTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsEcsServiceInvalidLaunchTypeRule returns new rule with default attributes
func NewAwsEcsServiceInvalidLaunchTypeRule() *AwsEcsServiceInvalidLaunchTypeRule {
	return &AwsEcsServiceInvalidLaunchTypeRule{
		resourceType:  "aws_ecs_service",
		attributeName: "launch_type",
		enum: []string{
			"EC2",
			"FARGATE",
		},
	}
}

// Name returns the rule name
func (r *AwsEcsServiceInvalidLaunchTypeRule) Name() string {
	return "aws_ecs_service_invalid_launch_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEcsServiceInvalidLaunchTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEcsServiceInvalidLaunchTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEcsServiceInvalidLaunchTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEcsServiceInvalidLaunchTypeRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as launch_type`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}