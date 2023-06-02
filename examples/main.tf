terraform {
  required_providers {
    json-formatter = {
      source = "TheNicholi/json-formatter"
    }
  }
}

data "json-formatter_format_json" "example" {
    json = "{\"a\":\"b\",\"myObj\":{\"prop1\":1}}"
    indent = "    "
}

output "example_output" {
    value = data.json-formatter_format_json.example.result
}