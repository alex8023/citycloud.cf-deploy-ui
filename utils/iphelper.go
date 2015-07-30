package utils

import (
	"fmt"
	"net"
	"strings"
)

type IPFactory struct {
	//Ip 池 key= ip,value = ip
	ippool map[string]string
	//被使用的Ip,key=ip,value=Jobs
	usedip map[string]string

	//key=jobName,value = ips
	jobMaps map[string][]string

	ips []string
}

func NewIPFactory() (iPFactory *IPFactory) {
	return &IPFactory{}
}

//初始化IPFactory
func (iPFactory *IPFactory) InitFactory(start, end string) (bool, error) {
	ips := parseStaticIP(start, end)
	if len(ips) < Job_Instance {
		return false, fmt.Errorf("there is not enough ip for PaaS instances: %s", Job_Instance)
	}

	ippool := make(map[string]string)
	for _, ip := range ips {
		ippool[ip] = ip
	}
	iPFactory.ippool = ippool
	iPFactory.ips = ips
	usedip := make(map[string]string)

	jobMaps := make(map[string][]string)

	jobMaps[Job_Haproxy] = []string{ips[0]}
	usedip[ips[0]] = ips[0]
	jobMaps[Job_Gorouter] = []string{ips[1]}
	usedip[ips[1]] = ips[1]
	jobMaps[Job_Postgres] = []string{ips[2]}
	usedip[ips[2]] = ips[2]
	jobMaps[Job_NFS] = []string{ips[3]}
	usedip[ips[3]] = ips[3]
	jobMaps[Job_NATS] = []string{ips[4]}
	usedip[ips[4]] = ips[4]
	jobMaps[Job_Syslog] = []string{ips[5]}
	usedip[ips[5]] = ips[5]
	jobMaps[Job_Etcd] = []string{ips[6]}
	usedip[ips[6]] = ips[6]
	jobMaps[Job_Loggregator] = []string{ips[7]}
	usedip[ips[7]] = ips[7]
	jobMaps[Job_Loggregator_Traffic] = []string{ips[8]}
	usedip[ips[8]] = ips[8]
	jobMaps[Job_Uaa] = []string{ips[9]}
	usedip[ips[9]] = ips[9]
	jobMaps[Job_Cloud_Controller_NG] = []string{ips[10]}
	usedip[ips[10]] = ips[10]
	jobMaps[Job_Cloud_Controller_Worker] = []string{ips[11]}
	usedip[ips[11]] = ips[11]
	jobMaps[Job_Cloud_Controller_Clock] = []string{ips[12]}
	usedip[ips[12]] = ips[12]
	jobMaps[Job_Hm9000] = []string{ips[13]}
	usedip[ips[13]] = ips[13]
	jobMaps[Job_Stats] = []string{ips[14]}
	usedip[ips[14]] = ips[14]
	jobMaps[Job_Dea_Next] = []string{ips[15]}
	usedip[ips[15]] = ips[15]

	iPFactory.jobMaps = jobMaps
	iPFactory.usedip = usedip
	return true, nil
}

//获取使用ip，参数JobName,规格
func (iPFactory *IPFactory) AssignaIP2Job(jobName string, size int) ([]string, error) {
	jobMap := iPFactory.jobMaps

	if ips, ok := jobMap[jobName]; ok {
		ipslen := len(ips)
		if ipslen == size {
			return ips, nil
		} else if ipslen > size {
			newips := ips[:size]
			jobMap[jobName] = newips
			iPFactory.trancateIp(ips[size:])
			return newips, nil
		} else {
			newips, errs := iPFactory.GetUnUsedIp(size - ipslen)
			if errs == nil {
				ips = append(ips, newips...)
				jobMap[jobName] = ips
				iPFactory.markIP2Used(jobName, ips)
				return ips, nil
			} else {
				return ips, errs
			}
		}
	} else {
		return nil, fmt.Errorf("unknow JobName: %s", jobName)
	}
}

//将给定的起始ip和结束ip之间的ip返回，前24位相同
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

// 获取未使用的ip地址
func (iPFactory *IPFactory) GetUnUsedIp(size int) ([]string, error) {
	//	ipslen := len(iPFactory.ips)
	ips := make([]string, 0)
	i := 0
	for _, ip := range iPFactory.ips {
		if i < size {
			if _, ok := iPFactory.usedip[ip]; !ok {
				ips = append(ips, ip)
				i = i + 1
			}
		}
	}
	if len(ips) != size {
		return nil, fmt.Errorf("there is not enough ip")
	}

	return ips, nil
}

func (iPFactory *IPFactory) trancateIp(ips []string) {
	for _, ip := range ips {
		delete(iPFactory.usedip, ip)
	}
	//	return true, nil
}

func (iPFactory *IPFactory) GetAssignaIP4Job(job string) []string {
	return iPFactory.jobMaps[job]
}

func (iPFactory *IPFactory) markIP2Used(job string, ips []string) {
	for _, ip := range ips {
		iPFactory.usedip[ip] = job
	}
}

func SpliteIPs(s string) []string {
	result := strings.Split(s, "-")
	for index, ip := range result {
		result[index] = strings.TrimSpace(ip)
	}
	return result
}
