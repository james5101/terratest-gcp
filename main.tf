provider "google" {
  project = "james-app-dev"
  region  = "us-east1"
  zone    = "us-east1-b"
}



resource "google_compute_instance" "vm" {
  name         = "jamesterratest"
  machine_type = "f1-micro"
  zone         = "us-east1-b"

  boot_disk {
    initialize_params {
      image = "ubuntu-1804-lts"
    }
  }

  

  network_interface {
    network = "default"

    access_config {
     
    }
  }

  
}