package utils

import (
	"net"
)

type IPFactory struct {
	//Ip 池
	ippool map[string]string
	//被使用的Ip,key=ip,value=Jobs
	usedip map[string]string

	jobMaps map[string][]string

	ips []string
}

func NewIPFactory() (iPFactory IPFactory) {
	return IPFactory{}
}

//初始化IPFactory
func (iPFactory IPFactory) InitIPPool(start, end string) {
	ips := parseStaticIP(start, end)
	ippool := make(map[string]string)
	for _, ip := range ips {
		ippool[ip] = ip
	}
	iPFactory.ippool = ippool
	iPFactory.ips = ips

	iPFactory.usedip = make(map[string]string)

	iPFactory.jobMaps = make(map[string][]string)

}

//获取使用ip，传入job名,使用数量
func (iPFactory IPFactory) AssignationIP2Job(jobName string, size int) []string {
	//	jobMap := iPFactory.jobMaps

	return nil
}

//将给定的起始ip和结束ip之间的ip返回
func parseStaticIP(start, end string) []string {
	//转换为4位ip
	startIp := net.ParseIP(start).To4()
	endIp := net.ParseIP(end).To4()

	ips := make([]string, 0)
	ips = append(ips, startIp.String())
	for startIp[3] < endIp[3] {
		startIp[3] = byte(int(startIp[3]) + 1)
		ips = append(ips, startIp.String())
	}

	return ips
}

//验证指定ip是否在指定的cidr范围内
func CheckCIDRContainsIP(cidr string, ip string) bool {
	_, ipnet, err := net.ParseCIDR(cidr)

	if err != nil {
		return false
	}

	pip := net.ParseIP(ip)

	if pip == nil {
		return false
	}
	return ipnet.Contains(pip)
}
