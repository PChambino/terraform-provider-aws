// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsLbListenerInvalidProtocolRule checks the pattern is valid
type AwsLbListenerInvalidProtocolRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsLbListenerInvalidProtocolRule returns new rule with default attributes
func NewAwsLbListenerInvalidProtocolRule() *AwsLbListenerInvalidProtocolRule {
	return &AwsLbListenerInvalidProtocolRule{
		resourceType:  "aws_lb_listener",
		attributeName: "protocol",
		enum: []string{
			"HTTP",
			"HTTPS",
			"TCP",
			"TLS",
			"UDP",
			"TCP_UDP",
		},
	}
}

// Name returns the rule name
func (r *AwsLbListenerInvalidProtocolRule) Name() string {
	return "aws_lb_listener_invalid_protocol"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLbListenerInvalidProtocolRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLbListenerInvalidProtocolRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLbListenerInvalidProtocolRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLbListenerInvalidProtocolRule) Check(runner *tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as protocol`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}