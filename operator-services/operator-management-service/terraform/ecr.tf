resource "aws_ecr_repository" "operator-services-images" {
  name = "operator-services-images"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

}
