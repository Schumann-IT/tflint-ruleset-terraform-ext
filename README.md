# TFLint Ruleset Terraform extension
[![Build Status](https://github.com/schumann-it/tflint-ruleset-terraformn-ext/workflows/build/badge.svg?branch=main)](https://github.com/schumann-it/tflint-ruleset-terraformn-ext/actions)

Extends default terraform ruleset with
- validation of variable attributes

## Requirements

- TFLint v0.42+
- Go v1.22

## Rules

|Name| Description                                |Severity|Enabled|Link|
| --- |--------------------------------------------| --- | --- | --- |
|terraform_variable_attribute_order| Validates the order of varibale attributes |ERROR|✔||

## Building the plugin

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```

You can run the built plugin like the following:

```
$ cat << EOS > .tflint.hcl
plugin "terraform-schumann-it-ext" {
  enabled = true
}
EOS
$ tflint
```
