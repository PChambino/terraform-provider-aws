// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsLambdaFunctionInvalidRoleRule checks the pattern is valid
type AwsLambdaFunctionInvalidRoleRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsLambdaFunctionInvalidRoleRule returns new rule with default attributes
func NewAwsLambdaFunctionInvalidRoleRule() *AwsLambdaFunctionInvalidRoleRule {
	return &AwsLambdaFunctionInvalidRoleRule{
		resourceType:  "aws_lambda_function",
		attributeName: "role",
		pattern:       regexp.MustCompile(`^arn:(aws[a-zA-Z-]*)?:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`),
	}
}

// Name returns the rule name
func (r *AwsLambdaFunctionInvalidRoleRule) Name() string {
	return "aws_lambda_function_invalid_role"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaFunctionInvalidRoleRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaFunctionInvalidRoleRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaFunctionInvalidRoleRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaFunctionInvalidRoleRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:(aws[a-zA-Z-]*)?:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}