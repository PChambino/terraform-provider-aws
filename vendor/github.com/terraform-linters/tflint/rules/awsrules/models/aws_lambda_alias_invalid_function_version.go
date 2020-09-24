// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsLambdaAliasInvalidFunctionVersionRule checks the pattern is valid
type AwsLambdaAliasInvalidFunctionVersionRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsLambdaAliasInvalidFunctionVersionRule returns new rule with default attributes
func NewAwsLambdaAliasInvalidFunctionVersionRule() *AwsLambdaAliasInvalidFunctionVersionRule {
	return &AwsLambdaAliasInvalidFunctionVersionRule{
		resourceType:  "aws_lambda_alias",
		attributeName: "function_version",
		max:           1024,
		min:           1,
		pattern:       regexp.MustCompile(`^(\$LATEST|[0-9]+)$`),
	}
}

// Name returns the rule name
func (r *AwsLambdaAliasInvalidFunctionVersionRule) Name() string {
	return "aws_lambda_alias_invalid_function_version"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaAliasInvalidFunctionVersionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaAliasInvalidFunctionVersionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaAliasInvalidFunctionVersionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaAliasInvalidFunctionVersionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"function_version must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"function_version must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^(\$LATEST|[0-9]+)$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}