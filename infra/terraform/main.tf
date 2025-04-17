provider "aws" {
  region = "us-west-2"
}

resource "aws_s3_bucket" "class_materials" {
  bucket = "skillshare-class-materials"
  acl    = "private"
}
