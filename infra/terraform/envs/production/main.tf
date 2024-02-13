terraform {
  required_version = "~> 1.6.1"

  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

variable "do_token" {}

provider "digitalocean" {
  token = var.do_token
}

module "droplet" {
  source = "../../modules/droplet"
  name = "production-droplet"
  tags = [ "env:production", "managed_by:terraform" ]
}
