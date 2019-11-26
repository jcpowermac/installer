locals {
  target_group_ip_assoc = [
    for pair in setproduct(var.target_group_arns, flatten([var.ip_address])) : {
      target_group_arn = pair[0]
      target_id        = pair[1]
    }
  ]
}

resource "null_resource" "ip_addresses_found" {
  triggers = {
    ip_address = var.ip_address
  }
}


resource "aws_lb_target_group_attachment" "bootstrap" {
  depends_on = [
    null_resource.ip_addresses_found
  ]

  for_each = {
    for tgia in local.target_group_ip_assoc : "${tgia.target_group_arn}.${tgia.target_id}" => tgia
  }

  target_group_arn = each.value.target_group_arn
  target_id        = each.value.target_id

  availability_zone = var.availability_zone
}

