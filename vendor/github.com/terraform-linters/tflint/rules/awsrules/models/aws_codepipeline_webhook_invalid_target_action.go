// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsCodepipelineWebhookInvalidTargetActionRule checks the pattern is valid
type AwsCodepipelineWebhookInvalidTargetActionRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCodepipelineWebhookInvalidTargetActionRule returns new rule with default attributes
func NewAwsCodepipelineWebhookInvalidTargetActionRule() *AwsCodepipelineWebhookInvalidTargetActionRule {
	return &AwsCodepipelineWebhookInvalidTargetActionRule{
		resourceType:  "aws_codepipeline_webhook",
		attributeName: "target_action",
		max:           100,
		min:           1,
		pattern:       regexp.MustCompile(`^[A-Za-z0-9.@\-_]+$`),
	}
}

// Name returns the rule name
func (r *AwsCodepipelineWebhookInvalidTargetActionRule) Name() string {
	return "aws_codepipeline_webhook_invalid_target_action"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodepipelineWebhookInvalidTargetActionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodepipelineWebhookInvalidTargetActionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodepipelineWebhookInvalidTargetActionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodepipelineWebhookInvalidTargetActionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"target_action must be 100 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"target_action must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[A-Za-z0-9.@\-_]+$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}