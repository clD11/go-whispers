provider "aws" {
  region                      = var.region
  skip_credentials_validation = var.is_dev
  skip_metadata_api_check     = var.is_dev
  skip_requesting_account_id  = var.is_dev
}