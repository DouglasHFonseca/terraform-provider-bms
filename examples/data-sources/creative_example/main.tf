terraform {
  required_providers {
    bms = {
      source = "hashicorp.com/edu/bms"
    }
  }
}

provider "bms" {
  account_id = "385963540035"
}

data "bms_creatives" "example" {
  filters = {
    name = "terraform"
  }
}

output "example_creatives" {
  value = data.bms_creatives.example
}
