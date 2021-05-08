terraform {
  required_providers {
    codename = {
      version = ">= 0.1.0"
      source  = "github.com/lucasepe/codename"
    }
  }
}

provider "codename" {}

resource "codename" "example1" {
  snakefy      = true
  token_length = 4
}

resource "codename" "example2" {
  prefix = "it->"
}

output "codename1" {
  value = codename.example1.id
}

output "codename2" {
  value = codename.example2.id
}
