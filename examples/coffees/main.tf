terraform {
  required_providers {
    awsgeneric = {
      source = "hashicorp.com/edu/hashicups"
    }
  }
}

provider "awsgeneric" {
  host     = "http://localhost:19090"
  username = "education"
  password = "test123"
}

data "awsgeneric_coffees" "edu" {}

output "edu_coffees" {
  value = data.awsgeneric_coffees.edu
}