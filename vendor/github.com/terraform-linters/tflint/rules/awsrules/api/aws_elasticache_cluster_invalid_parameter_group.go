// This file generated by `generator/main.go`. DO NOT EDIT

package api

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsElastiCacheClusterInvalidParameterGroupRule checks whether attribute value actually exists
type AwsElastiCacheClusterInvalidParameterGroupRule struct {
	resourceType  string
	attributeName string
	data          map[string]bool
	dataPrepared  bool
}

// NewAwsElastiCacheClusterInvalidParameterGroupRule returns new rule with default attributes
func NewAwsElastiCacheClusterInvalidParameterGroupRule() *AwsElastiCacheClusterInvalidParameterGroupRule {
	return &AwsElastiCacheClusterInvalidParameterGroupRule{
		resourceType:  "aws_elasticache_cluster",
		attributeName: "parameter_group_name",
		data:          map[string]bool{},
		dataPrepared:  false,
	}
}

// Name returns the rule name
func (r *AwsElastiCacheClusterInvalidParameterGroupRule) Name() string {
	return "aws_elasticache_cluster_invalid_parameter_group"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElastiCacheClusterInvalidParameterGroupRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElastiCacheClusterInvalidParameterGroupRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElastiCacheClusterInvalidParameterGroupRule) Link() string {
	return ""
}

// Check checks whether the attributes are included in the list retrieved by DescribeCacheParameterGroups
func (r *AwsElastiCacheClusterInvalidParameterGroupRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		if !r.dataPrepared {
			log.Print("[DEBUG] invoking DescribeCacheParameterGroups")
			var err error
			r.data, err = runner.AwsClient.DescribeCacheParameterGroups()
			if err != nil {
				err := &tflint.Error{
					Code:    tflint.ExternalAPIError,
					Level:   tflint.ErrorLevel,
					Message: "An error occurred while invoking DescribeCacheParameterGroups",
					Cause:   err,
				}
				log.Printf("[ERROR] %s", err)
				return err
			}
			r.dataPrepared = true
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.data[val] {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is invalid parameter group name.`, val),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}