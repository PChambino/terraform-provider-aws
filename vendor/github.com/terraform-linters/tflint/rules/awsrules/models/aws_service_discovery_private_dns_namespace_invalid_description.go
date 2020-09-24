// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsServiceDiscoveryPrivateDNSNamespaceInvalidDescriptionRule checks the pattern is valid
type AwsServiceDiscoveryPrivateDNSNamespaceInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsServiceDiscoveryPrivateDNSNamespaceInvalidDescriptionRule returns new rule with default attributes
func NewAwsServiceDiscoveryPrivateDNSNamespaceInvalidDescriptionRule() *AwsServiceDiscoveryPrivateDNSNamespaceInvalidDescriptionRule {
	return &AwsServiceDiscoveryPrivateDNSNamespaceInvalidDescriptionRule{
		resourceType:  "aws_service_discovery_private_dns_namespace",
		attributeName: "description",
		max:           1024,
	}
}

// Name returns the rule name
func (r *AwsServiceDiscoveryPrivateDNSNamespaceInvalidDescriptionRule) Name() string {
	return "aws_service_discovery_private_dns_namespace_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsServiceDiscoveryPrivateDNSNamespaceInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsServiceDiscoveryPrivateDNSNamespaceInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsServiceDiscoveryPrivateDNSNamespaceInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsServiceDiscoveryPrivateDNSNamespaceInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}