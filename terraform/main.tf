terraform {
  backend "http" {
    address        = "http://localhost:3000/state/123"
    lock_address   = "http://localhost:3000/state/123"
    unlock_address = "http://localhost:3000/state/123"
  }
}

terraform {
  required_providers {
    docker = {
      source = "kreuzwerker/docker"
      version = "~> 3.0.1"
    }
  }
}

provider "docker" {}

resource "docker_image" "nginx" {
  name         = "nginx:latest"
  keep_locally = false
}

resource "docker_container" "nginx" {
  image = docker_image.nginx.image_id
  name  = "tutorial"
  ports {
    internal = 80
    external = 8001
  }
}
