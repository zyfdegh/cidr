package cidr

import (
	"errors"
	"strconv"
	"strings"
)

var (
	// ErrEmptyCIDR is returned when the input CIDR in invalid
	ErrEmptyCIDR = errors.New("empty CIDR")
)

// CIDR 地址范围
func GetRange(cidr string) (fromAddr, toAddr string, err error) {
	if len(cidr) == 0 {
		return
	}
	fromAddr, toAddr = getCidrIpRange(cidr)
	return
}

//// copied from "duhf_think"
//// 《使用Go语言计算网络IP地址的CIDR》 https://studygolang.com/articles/3775
func getCidrIpRange(cidr string) (string, string) {
	ip := strings.Split(cidr, "/")[0]
	ipSegs := strings.Split(ip, ".")
	maskLen, _ := strconv.Atoi(strings.Split(cidr, "/")[1])
	seg2MinIp, seg2MaxIp := getIpSeg2Range(ipSegs, maskLen)
	seg3MinIp, seg3MaxIp := getIpSeg3Range(ipSegs, maskLen)
	seg4MinIp, seg4MaxIp := getIpSeg4Range(ipSegs, maskLen)
	ipPrefix := ipSegs[0] + "."

	return ipPrefix + strconv.Itoa(seg2MinIp) + "." + strconv.Itoa(seg3MinIp) + "." + strconv.Itoa(seg4MinIp),
		ipPrefix + strconv.Itoa(seg2MaxIp) + "." + strconv.Itoa(seg3MaxIp) + "." + strconv.Itoa(seg4MaxIp)
}

// 得到第二段IP的区间（第一片段.第二片段.第三片段.第四片段）
func getIpSeg2Range(ipSegs []string, maskLen int) (int, int) {
	if maskLen > 16 {
		segIp, _ := strconv.Atoi(ipSegs[1])
		return segIp, segIp
	}
	ipSeg, _ := strconv.Atoi(ipSegs[1])
	return getIpSegRange(uint8(ipSeg), uint8(16-maskLen))
}

//得到第三段IP的区间（第一片段.第二片段.第三片段.第四片段）
func getIpSeg3Range(ipSegs []string, maskLen int) (int, int) {
	if maskLen > 24 {
		segIp, _ := strconv.Atoi(ipSegs[2])
		return segIp, segIp
	}
	ipSeg, _ := strconv.Atoi(ipSegs[2])
	return getIpSegRange(uint8(ipSeg), uint8(24-maskLen))
}

//得到第四段IP的区间（第一片段.第二片段.第三片段.第四片段）
func getIpSeg4Range(ipSegs []string, maskLen int) (int, int) {
	ipSeg, _ := strconv.Atoi(ipSegs[3])
	segMinIp, segMaxIp := getIpSegRange(uint8(ipSeg), uint8(32-maskLen))
	if segMaxIp > segMinIp {
		return segMinIp + 1, segMaxIp
	}
	return segMinIp, segMaxIp
}

//根据用户输入的基础IP地址和CIDR掩码计算一个IP片段的区间
func getIpSegRange(userSegIp, offset uint8) (int, int) {
	var ipSegMax uint8 = 255
	netSegIp := ipSegMax << offset
	segMinIp := netSegIp & userSegIp
	segMaxIp := userSegIp&(255<<offset) | ^(255 << offset)
	return int(segMinIp), int(segMaxIp)
}
