package main

import (
	"github.com/schumann-it/tflint-ruleset-schumann-it-terraform-ext/rules"
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "terraform-ext",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				rules.NewTerraformVariableAttributeOrderRule(),
				rules.NewTerraformVariableAttributeSeparationRule(),
				rules.NewTerraformVariableAttributesBeforeValidationRulesRule(),
				rules.NewTerraformVariableValidationRuleSeparationRule(),
				rules.NewTerraformVariableValidationRuleAttributeOrderRule(),
			},
		},
	})
}
