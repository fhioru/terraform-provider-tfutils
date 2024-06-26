# Terraform provider for utility data sources and functions

> Attribution. This provider is based on the original work published at
> https://github.com/TheNicholi/terraform-provider-json-formatter


## Utility Data Sources

### tfutils_json_format

Formats the provided JSON input to provide a pretty-print type functionality. Only a computer
would be happy to read a 40,000 long line of JSON.

**Usage**

| Name   | Required | Description                                                               |
| :----- | :------- | :------------------------------------------------------------------------ |
| json   | Yes      | A JSON string.                                                            |
| indent | No       | The number of spaces to use for an indent. Defaults to 2.                 |

**Example**

```hcl
data "tfutils_json_format" "example" {
  json = "{\"a\":\"b\",\"myObj\":{\"prop1\":1}}"
  indent = 4
}

output "example_output" {
  value = data.tfutils_json_format.example.result
}
```

**Output**

```console
example_output =
{
    "a": "b",
    "myObj": {
        "prop1": 1
    }
}
```
