#!/bin/bash
set -x

for i in {1..10}; do
        if hostname=$(/bin/vmtoolsd --cmd 'info-get guestinfo.hostname'); then
                break
        fi
        sleep 30s
done
/usr/bin/hostnamectl --transient --static set-hostname ${hostname}

