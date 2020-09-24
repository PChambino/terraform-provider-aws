// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsSsmDocumentInvalidNameRule checks the pattern is valid
type AwsSsmDocumentInvalidNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsSsmDocumentInvalidNameRule returns new rule with default attributes
func NewAwsSsmDocumentInvalidNameRule() *AwsSsmDocumentInvalidNameRule {
	return &AwsSsmDocumentInvalidNameRule{
		resourceType:  "aws_ssm_document",
		attributeName: "name",
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9_\-.]{3,128}$`),
	}
}

// Name returns the rule name
func (r *AwsSsmDocumentInvalidNameRule) Name() string {
	return "aws_ssm_document_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmDocumentInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmDocumentInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmDocumentInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmDocumentInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9_\-.]{3,128}$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}