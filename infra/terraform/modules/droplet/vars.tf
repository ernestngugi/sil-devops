variable "name" {
  type = string
}

variable "region" {
  type    = string
  default = "fra1"
}

variable "tags" {
  type    = list(string)
  default = []
}

variable "image" {
  type = string
  default = "ubuntu-20-04-x64"
}

variable "size" {
  type = string
  default = "s-1vcpu-1gb"
}
