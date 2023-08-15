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
  name = "creating a creative1 by terraform provider bms"
  domain = "example.com"
  tags = ["resource", "terraform", "bms"]

  banner {
    width = 300
    height = 250
    snippet = "<a href\"https://example.com/\"><img src=\"https://via.placeholder.com/782x90\"></a>"
  }
}

resource "bms_creative" "creative_2" {
  name = "creating a creative2 by terraform provider bms"
  domain = "example.com"
  tags = ["resource", "terraform", "bms"]

  banner {
    width = 300
    height = 250
    snippet = "<a href\"https://teste.com/\"><img src=\"https://via.placeholder.com/782x90\"></a>"
  }
}

output "bms_creative" {
  value = bms_creative.creative_1
}

resource "bms_creative_group" "creative_group_1" {
    name = "creating a creative group by terraform provider bms"
    domain = "example.com"
    tags = [ "resource", "group", "bms"]

    spec {
      banner {
        width = 300
        height = 250
      }
    }

    creatives {
        creative_id = bms_creative.creative_1.creative_id
        weight = 1
      }

    creatives {
      creative_id = bms_creative.creative_2.creative_id
      weight = 2
    }
}

output "bms_creative_group" {
    value = bms_creative_group.creative_group_1
}