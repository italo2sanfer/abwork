variable "project_id" {
  type    = string
  default = "infinite-maxim-417114"
}

variable "regiao" {
  type = string
  default = "us-west4"
}

variable "nome" {
  type = string
  default = "vm-webserver"
}

variable "tipo_maquina" {
  type = string
  default = "f1-micro"
}

variable "zona" {
  type = string
  default = "us-west4-a"
}

variable "imagem" {
  type = string
  default = "debian-cloud/debian-12"
}

variable "nome_fw" {
  type = string
  default = "webserver-firewall"
}

variable "portas" {
  type = list
  default = ["80"]
}
