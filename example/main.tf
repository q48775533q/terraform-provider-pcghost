terraform {
  required_providers {
    pcghost = {
      version = "0.2.0"
      source  = "pcghost.com/com/pcghost"
    }
  }
}


provider "pcghost" {
  address = "http://127.0.0.1:8000"
}

resource "pcghost_pet" "my_pet" {
  name    = "princess"
  species = "cat"
  age     = 4
}
