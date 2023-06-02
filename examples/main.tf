terraform {
  required_providers {
    json-formatter = {
      source = "TheNicholi/json-formatter"
    }
  }
}

data "json-formatter_format" "example" {
    json = "{\"a\":\"b\",\"myObj\":{\"prop1\":1}}"
    indent = "    "
}

output "example_output" {
    value = data.json-formatter_format.example.result
}