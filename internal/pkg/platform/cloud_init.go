package platform

import "fmt"

func TailscaleCloudInit(authKey, tid, controlServer string) string {
	return fmt.Sprintf("#!/bin/bash\necho 'net.ipv4.ip_forward = 1' | sudo tee -a /etc/sysctl.conf\necho 'net.ipv6.conf.all.forwarding = 1' | sudo tee -a /etc/sysctl.conf\nsudo sysctl -p /etc/sysctl.conf\n\ncurl -fsSL https://tailscale.com/install.sh | sh\n\nsudo tailscale up --auth-key=%s --hostname=%s --login-server=%s --advertise-exit-node", authKey, tid, controlServer)
}
