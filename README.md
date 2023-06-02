# Terraform provider for formatting JSON strings

This plugin will format / beautify a JSON string.

## Usage

### Arguments

|Name|Required|Description|
|:---|:---|:---|
|json|Yes|A JSON string.|
|indent|No|The character(s) to use for an indent. Has a default value of two spaces.|

### Example

```hcl
data "json-formatter_format_json" "example" {
    json = "{\"a\":\"b\",\"myObj\":{\"prop1\":1}}"
    indent = "    "
}

output "example_output" {
    value = data.json-formatter_format_json.example.result
}
```

This will output:

```sh
Outputs:

test_output = <<EOT
{
    "a": "b",
    "myObj": {
        "prop1": 1
    }
}
EOT
```
