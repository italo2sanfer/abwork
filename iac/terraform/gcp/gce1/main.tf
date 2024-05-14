# Configures the project
provider "google" {
  credentials = "${file("key-terraform.json")}"
  project     = "${var.project_id}"
  region      = "${var.region}"
}

# Create a VM with Google Compute Engine
resource "google_compute_instance" "webserver" {
  name          = "${var.name}"
  machine_type  = "${var.machine_type}"
  zone          = "${var.zone}"

  boot_disk {
    initialize_params {
      image = "${var.image}"
    }
  }

  # Install Web Apache Server
  metadata_startup_script = "sudo apt-get update; sudo apt-get install apache2 -y; echo 'Testing' > /var/www/html/index.html"

  # Able network VM and a public IP
  network_interface {
    network = "default"
    access_config {

    }
  }
}

# Create the Firewall for VM
resource "google_compute_firewall" "webfirewall" {
  name        = "${var.fw_name}"
  network     = "default"
  source_tags = []

  allow {
    protocol  = "tcp"
    ports     = "${var.ports}"
  }
}