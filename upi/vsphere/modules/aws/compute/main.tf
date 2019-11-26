locals {
  target_group_arns_length = length(var.target_group_arns)

  target_group_ip_assoc = [
    for pair in setproduct(var.target_group_arns, var.ip_addresses) : {
      target_group_arn = pair[0]
      target_id        = pair[1]
    }
  ]
}


resource "aws_lb_target_group_attachment" "compute" {
  for_each = {
    for tgia in local.target_group_ip_assoc : "${tgia.target_group_arn}.${tgia.target_id}" => tgia
  }

  target_group_arn = each.value.target_group_arn
  target_id        = each.value.target_id

  availability_zone = var.availability_zone
}
