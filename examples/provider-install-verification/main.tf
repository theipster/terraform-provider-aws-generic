terraform {
  required_providers {
    awsgeneric = {
      source = "hashicorp.com/edu/hashicups"
    }
  }
}

provider "awsgeneric" {}

data "awsgeneric_coffees" "example" {}
