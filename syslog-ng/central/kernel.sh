#!/usr/bin/env bash

sysctl -w net.core.rmem_max=524288000
sysctl -w net.core.netdev_max_backlog=10000

sysctl -w net.ipv4.udp_mem='262144 327680 393216'
echo 536870912 > /proc/sys/net/core/rmem_max