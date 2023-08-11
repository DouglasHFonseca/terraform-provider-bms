terraform {
  required_providers {
    bms = {
      source  = "hashicorp.com/edu/bms"
    }
  }
  required_version = ">= 1.1.0"
}

provider "bms" {
  account_id = "385963540035"
}

resource "bms_creative" "creative_1" {
  name = "updating a creative by terraform provider bms"
  domain = "example.com"
  tags = ["resource", "terraform", "bms"]
  banner = {
    width = 300
    height = 250
    snippet = "<a href\"https://example.com/\"><img src=\"https://via.placeholder.com/782x90\"></a>"
  }
}

output "bms_creative" {
    value = bms_creative.creative_1
}

