terraform {
  required_providers {
    bms = {
      source = "hashicorp.com/edu/bms"
    }
  }
  required_version = ">= 1.1.0"
}

provider "bms" {
  account_id = "385963540035"
}

resource "bms_ads_creative" "creative_1" {
  name   = "creating a creative1 by terraform provider bms"
  domain = "example.com"
  tags   = ["resource", "terraform", "bms"]

  banner {
    width   = 300
    height  = 250
    snippet = "<a href\"https://example.com/\"><img src=\"https://via.placeholder.com/782x90\"></a>"
  }
}

resource "bms_ads_creative" "creative_2" {
  name   = "creating a creative2 by terraform provider bms"
  domain = "example.com"
  tags   = ["resource", "terraform", "bms"]

  banner {
    width   = 300
    height  = 250
    snippet = "<a href\"https://teste.com/\"><img src=\"https://via.placeholder.com/782x90\"></a>"
  }
}

output "bms_ads_creative" {
  value = bms_ads_creative.creative_1
}

resource "bms_ads_creative_group" "creative_group_1" {
  name   = "creating a creative group by terraform provider bms"
  domain = "example.com"
  tags   = ["resource", "group", "bms"]

  spec {
    banner {
      width  = 300
      height = 250
    }
  }

  creative {
    creative_id = bms_ads_creative.creative_1.creative_id
    weight      = 1
  }

  creative {
    creative_id = bms_ads_creative.creative_2.creative_id
    weight      = 2
  }
}

output "bms_ads_creative_group" {
  value = bms_ads_creative_group.creative_group_1
}

resource "bms_ads_ad" "ad_1" {
  domain = "example.com"
  name   = "creating a ad by terraform provider bms"
  tags   = ["resource", "terraform", "bms"]

  spec {
    banner {
      width  = 300
      height = 250
    }
  }

  rule {
    name              = "creating a rule by terraform provider bms"
    creative_group_id = bms_ads_creative_group.creative_group_1.creative_group_id

    schedule_condition {
      type = "schedule"
    }
  }
}