package tflint

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/terraform/configs"
)

// Runner acts as a client for each plugin to query the host process about the Terraform configurations.
type Runner interface {
	// WalkResourceAttributes visits attributes with the passed function.
	// You must pass a resource type as the first argument and an attribute name as the second argument.
	WalkResourceAttributes(string, string, func(*hcl.Attribute) error) error

	// WalkResourceBlocks visits blocks with the passed function.
	// You must pass a resource type as the first argument and a block type as the second argument.
	// This API currently does not support labeled blocks.
	WalkResourceBlocks(string, string, func(*hcl.Block) error) error

	// WalkResources visits resources with the passed function.
	// You must pass a resource type as the first argument.
	WalkResources(string, func(*configs.Resource) error) error

	// WalkModuleCalls visits module calls with the passed function.
	WalkModuleCalls(func(*configs.ModuleCall) error) error

	// Backend returns the backend configuration, if any.
	Backend() (*configs.Backend, error)

	// Config returns the Terraform configuration.
	// This object contains almost all accessible data structures from plugins.
	Config() (*configs.Config, error)

	// EvaluateExpr evaluates the passed expression and reflects the result in ret.
	// Since this function returns an application error, it is expected to use the EnsureNoError
	// to determine whether to continue processing.
	EvaluateExpr(expr hcl.Expression, ret interface{}) error

	// EmitIssue sends an issue with an expression to TFLint. You need to pass the message of the issue and the expression.
	EmitIssueOnExpr(rule Rule, message string, expr hcl.Expression) error

	// EmitIssue sends an issue to TFLint. You need to pass the message of the issue and the range.
	// You should use EmitIssueOnExpr if you want to emit an issue for an expression.
	// This API provides a lower level interface.
	EmitIssue(rule Rule, message string, location hcl.Range) error

	// EnsureNoError is a helper for error handling. Depending on the type of error generated by EvaluateExpr,
	// determine whether to exit, skip, or continue. If it is continued, the passed function will be executed.
	EnsureNoError(error, func() error) error
}

// Rule is the interface that the plugin's rules should satisfy.
type Rule interface {
	// Name will be displayed with a message of an issue and will be the identifier used to control
	// the behavior of this rule in the configuration file etc.
	// Therefore, it is expected that this will not duplicate the rule names provided by other plugins.
	Name() string

	// Enabled indicates whether the rule is enabled by default.
	Enabled() bool

	// Severity indicates the severity of the rule.
	Severity() string

	// Link allows you to add a reference link to the rule.
	Link() string

	// Check is the entrypoint of the rule. You can fetch Terraform configurations and send issues via Runner.
	Check(Runner) error
}