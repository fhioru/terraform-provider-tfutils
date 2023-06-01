terraform {
  required_providers {
    json-formatter = {
      source = "TheNicholi/json-formatter"
    }
  }
}

data "json-formatter_format" "test" {
    json = "{\"a\":\"b\",\"myObj\":{\"prop1\":1}}"
    indent = "    "
}

output "test_output" {
    value = data.json-formatter_format.test.result
}

resource "local_file" "test" {
  content  = data.json-formatter_format.test.result
  filename = "test.json"
}