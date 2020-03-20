locals {
  ignition_encoded = "data:text/plain;charset=utf-8;base64,${base64encode(var.ignition)}"
}

data "ignition_systemd_unit" "hostname-systemd" {
  name    = "vsphere-hostname-vmtoolsd.service"
  content = "[Unit]\nAfter=vmtoolsd.service\n[Service]\nType=oneshot\nExecStart=/usr/bin/hostnamectl --static set-hostname $(/usr/bin/vmtoolsd --cmd 'info-get guestinfo.hostname')\nExecStart=/usr/bin/hostnamectl --transient set-hostname $(/usr/bin/vmtoolsd --cmd 'info-get guestinfo.hostname')\n[Install]\nWantedBy=multi-user.target"

}

data "ignition_config" "ign" {
  //count = var.instance_count

  append {
    source = local.ignition_encoded
  }

  systemd = [
    data.ignition_systemd_unit.hostname-systemd.rendered,
  ]
}

