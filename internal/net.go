package internal

import "net"

func LookupIp(host string) ([]string, error) {
	if host == "" {
		return nil, nil
	}

	ips, err := net.LookupIP(host)
	if err != nil {
		return nil, err
	}

	var out []string
	for _, ip := range ips {
		out = append(out, ip.String())
	}
	return out, nil
}
