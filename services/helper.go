package services

import "net"

func Unique(data []uint32) []uint32 {
	tmp := map[uint32]bool{}
	ret := []uint32{}
	for _, v := range data {
		ok, _ := tmp[v]
		if ok {
			continue
		}
		tmp[v] = true
		ret = append(ret, v)
	}
	return ret
}

func UniqueString(data []string) []string {
	tmp := map[string]bool{}
	ret := []string{}
	for _, v := range data {
		ok, _ := tmp[v]
		if ok {
			continue
		}
		tmp[v] = true
		ret = append(ret, v)
	}
	return ret
}

func VerifyIp(ip string) bool {
	return net.ParseIP(ip) != nil
}
