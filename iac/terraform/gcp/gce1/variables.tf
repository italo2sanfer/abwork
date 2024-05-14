variable "project_id" {
  type    = string
  default = "<project_id_development> change here!"
}

variable "region" {
  type = string
  default = "us-west4"
}

variable "name" {
  type = string
  default = "vm-webserver"
}

variable "machine_type" {
  type = string
  default = "f1-micro"
}

variable "zone" {
  type = string
  default = "us-west4-a"
}

variable "image" {
  type = string
  default = "debian-cloud/debian-12"
}

variable "fw_name" {
  type = string
  default = "webserver-firewall"
}

variable "ports" {
  type = list
  default = ["80"]
}