package main

import (
	"errors"
	"fmt"
	"net"
	neturl "net/url"
	"strings"
)

func ValidUrl(urlStr string) error {
	u, err := neturl.Parse(urlStr)
	if err != nil {
		return err
	}

	if u.Hostname() == "" {
		return errors.New("host name is null")
	}

	if _, has := allowPortsMap[u.Port()]; !has {
		return errors.New(fmt.Sprintf("allowed ports are %s, but "+u.Port(), allowPorts[1:]))
	}

	if isLocalDomain(u.Hostname()) {
		return errors.New("not allow local domain")
	}

	hostips, _ := net.LookupIP(u.Hostname())
	for _, hostip := range hostips {
		if isPrivateIP(hostip) {
			return errors.New("not allow private ip")
		}
	}

	return nil
}

func isLocalDomain(hostname string) bool {
	parts := strings.Split(hostname, ".")
	if len(parts) == 0 {
		return false
	}
	topPart := strings.ToLower(strings.TrimSpace(parts[len(parts)-1]))
	return topPart == "local" || topPart == "localhost"
}

func isPrivateIP(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() {
		return true
	}

	for _, block := range privateIPBlocks {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}

// Ref from https://stackoverflow.com/questions/41240761
var privateIPBlocks []*net.IPNet
var allowPorts = []string{"", "80", "443"}
var allowPortsMap map[string]struct{}

func init() {
	for _, cidr := range []string{
		"127.0.0.0/8",    // IPv4 loopback
		"10.0.0.0/8",     // RFC1918
		"172.16.0.0/12",  // RFC1918
		"192.168.0.0/16", // RFC1918
		"169.254.0.0/16", // RFC3927 link-local
		"::1/128",        // IPv6 loopback
		"fe80::/10",      // IPv6 link-local
		"fc00::/7",       // IPv6 unique local addr
	} {
		_, block, err := net.ParseCIDR(cidr)
		if err != nil {
			panic(fmt.Errorf("parse error on %q: %v", cidr, err))
		}
		privateIPBlocks = append(privateIPBlocks, block)
	}

	allowPortsMap = make(map[string]struct{}, len(allowPorts))
	for _, port := range allowPorts {
		allowPortsMap[port] = struct{}{}
	}
}
