terraform {
  required_providers {
    awsgeneric = {
      source = "hashicorp.com/edu/hashicups"
    }
  }
  required_version = ">= 1.8.0"
}

provider "awsgeneric" {
  username = "education"
  password = "test123"
  host     = "http://localhost:19090"
}

output "total_price" {
  value = provider::awsgeneric::compute_tax(5.00, 0.085)
}
