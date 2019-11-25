locals {
  target_group_arns_length = length(var.target_group_arns)
}

resource "aws_lb_target_group_attachment" "control_plane" {
  count = var.instance_count * local.target_group_arns_length

  target_group_arn  = var.target_group_arns[count.index % local.target_group_arns_length]
  target_id         = var.ip_addresses[floor(count.index / local.target_group_arns_length)]
  availability_zone = var.availability_zone
}

