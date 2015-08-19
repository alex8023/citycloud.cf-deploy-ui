package entity

import (
	"github.com/citycloud/citycloud.cf-deploy-ui/utils"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Testing with Ginkgo", func() {
	It("cloud foundry template", func() {

		var cft = InitCloudFoundryV2()
		_, err := cft.CreateCloudFoundryYaml(CloudFoundryTemplateV2, "/home/ubuntu/temp/cloudfoundryv2.yml")

		if err != nil {
			GinkgoT().Error(err)
		}
	})
	It("meta", func() {

		var cft = InitCloudFoundryV2()
		_, err := cft.CreateCloudFoundryYaml(CloudFoundryMetaTemplateV2+`{{template "meta" .}}`, "/home/ubuntu/temp/metav2.yml")

		if err != nil {
			GinkgoT().Error(err)
		}
	})
	It("update", func() {

		var cft = InitCloudFoundryV2()
		_, err := cft.CreateCloudFoundryYaml(CloudFoundryUpdateTempalte+`{{template "update" .}}`, "/home/ubuntu/temp/update.yml")

		if err != nil {
			GinkgoT().Error(err)
		}
	})
	It("compilation", func() {

		var cft = InitCloudFoundryV2()
		_, err := cft.CreateCloudFoundryYaml(CloudFoundryCompilationTemplate+`{{template "compilation" .}}`, "/home/ubuntu/temp/compilation.yml")

		if err != nil {
			GinkgoT().Error(err)
		}
	})
	It("net works", func() {

		var cft = InitCloudFoundryV2()
		_, err := cft.CreateCloudFoundryYaml(CloudFoundryNetworksTemplateV2+`{{template "networks" .}}`, "/home/ubuntu/temp/networksv2.yml")

		if err != nil {
			GinkgoT().Error(err)
		}
	})
	It("net works v3", func() {

		var cft = InitCloudFoundryV3()
		_, err := cft.CreateCloudFoundryYaml(CloudFoundryNetworksTemplateV3+`{{template "networks" .}}`, "/home/ubuntu/temp/networksv3.yml")

		if err != nil {
			GinkgoT().Error(err)
		}
	})
	It("resource pool", func() {

		var cft = InitCloudFoundryV2()
		_, err := cft.CreateCloudFoundryYaml(CloudFoundryResourcePoolTemplate+`{{template "resourcepool" .}}`, "/home/ubuntu/temp/resourcepool.yml")

		if err != nil {
			GinkgoT().Error(err)
		}
	})
	It("jobs", func() {

		var cft = InitCloudFoundryV2()
		_, err := cft.CreateCloudFoundryYaml(CloudFoundryJobsTemplate+`{{template "jobs" .}}`, "/home/ubuntu/temp/jobs.yml")

		if err != nil {
			GinkgoT().Error(err)
		}
	})
	It("properties", func() {

		var cft = InitCloudFoundryV2()
		_, err := cft.CreateCloudFoundryYaml(CloudFoundryPropertiesTemplate+`{{template "properties" .}}`, "/home/ubuntu/temp/properties.yml")

		if err != nil {
			GinkgoT().Error(err)
		}
	})
	It("cloud foundry v3 template", func() {

		var cft = InitCloudFoundryV3()
		_, err := cft.CreateCloudFoundryYaml(CloudFoundryTemplateV3, "/home/ubuntu/temp/cloudfoundryv3.yml")

		if err != nil {
			GinkgoT().Error(err)
		}
	})
})

func InitCloudFoundryV2() CloudFoundryTemplate {
	var cloudProperties = CloudProperties{
		AuthUrl:         "http://192.168.128.2:5000/v2.0",
		UserName:        "paas",
		ApiKey:          "paas123",
		Tenant:          "Project_paas",
		DefaultKeyName:  "test_microbosh",
		PrivateKey:      "~/bosh-workspace/key/test_microbosh.private",
		CciEbsUrl:       "http://192.168.128.5:8080/EBS",
		CciEbsAccesskey: "7fd292943f3a46c491fbe4152cf3c8f0",
		CciEbsSecretkey: "8e100b9430ab429397352359b16f01b0",
	}

	var cloudFoundryProperties = NewCloudFoundryProperties("cf-release",
		"57cfc863-786d-4495-bb97-86d2f650a038", "192.168.133.102", "ccipaas.net", "cci")

	cloudFoundryProperties.CloudProperties = cloudProperties

	var compilation = NewCompilation("flavor_91", "zone2", 6, "cf1")

	var networks = NewNetWorks("cf1", "private",
		"manual",
		"8bb21e6e-dc6a-409c-82d0-a110fb3c9fe1",
		"192.168.129.0/24",
		"10.10.170.2",
		"192.168.133.108",
		"192.168.129.1 - 192.168.129.99",
		"192.168.129.100 - 192.168.129.126")

	var mapNetWorks = make(map[string]NetWorks)
	mapNetWorks["private"] = networks
	mapNetWorks["public"] = NewFloatingNetWork(cloudFoundryProperties.FloatingIp)

	var resourcesPool = NewResourcesPools("normal",
		"flavor_91",
		"zone2",
		"cf1",
		4)
	var mapResourcesPools = make([]ResourcesPools, 0)
	mapResourcesPools = append(mapResourcesPools, resourcesPool)

	cloudFoundryJobsMap := map[string]CloudFoundryJobs{
		utils.Job_Haproxy: NewCloudFoundryJobs(
			"haproxy",
			utils.Job_Haproxy,
			"normal",
			1,
			[]string{""}),
		utils.Job_Gorouter: NewCloudFoundryJobs(
			"gorouter",
			utils.Job_Gorouter,
			"normal",
			2,
			[]string{""}),
		utils.Job_Postgres: NewCloudFoundryJobs(
			"postgres",
			utils.Job_Postgres,
			"normal",
			1,
			[]string{""}),
		utils.Job_NFS: NewCloudFoundryJobs(
			"nfs",
			utils.Job_NFS,
			"normal",
			1,
			[]string{""}),
		utils.Job_NATS: NewCloudFoundryJobs(
			"nats",
			utils.Job_NATS,
			"normal",
			1,
			[]string{""}),
		utils.Job_Syslog: NewCloudFoundryJobs(
			"syslog_aggregator",
			utils.Job_Syslog,
			"normal",
			1,
			[]string{""}),
		utils.Job_Etcd: NewCloudFoundryJobs(
			"etcd",
			utils.Job_Etcd,
			"normal",
			1,
			[]string{""}),
		utils.Job_Loggregator: NewCloudFoundryJobs(
			"loggregator",
			utils.Job_Loggregator,
			"normal",
			1,
			[]string{""}),
		utils.Job_Loggregator_Traffic: NewCloudFoundryJobs(
			"loggregator_traffic",
			utils.Job_Loggregator_Traffic,
			"normal",
			1,
			[]string{""}),
		utils.Job_Uaa: NewCloudFoundryJobs(
			"uaa", utils.Job_Uaa,
			"normal",
			1,
			[]string{""}),
		utils.Job_Cloud_Controller_NG: NewCloudFoundryJobs(
			"cloud_controller_ng",
			utils.Job_Cloud_Controller_NG,
			"normal",
			1,
			[]string{""}),
		utils.Job_Cloud_Controller_Worker: NewCloudFoundryJobs(
			"cloud_controller_worker",
			utils.Job_Cloud_Controller_Worker,
			"normal",
			1,
			[]string{""}),
		utils.Job_Cloud_Controller_Clock: NewCloudFoundryJobs(
			"cloud_controller_clock",
			utils.Job_Cloud_Controller_Clock,
			"normal",
			1,
			[]string{""}),
		utils.Job_Hm9000: NewCloudFoundryJobs(
			"hm9000",
			utils.Job_Hm9000,
			"normal",
			1,
			[]string{""}),
		utils.Job_Stats: NewCloudFoundryJobs("stats",
			utils.Job_Stats,
			"normal",
			1,
			[]string{""}),
		utils.Job_Dea_Next: NewCloudFoundryJobs(
			"dea_next",
			utils.Job_Dea_Next,
			"normal",
			1,
			[]string{""}),
	}

	ipfactory := utils.NewIPFactory()
	ipArr := utils.SpliteIPs(networks.StaticIp)
	success, _ := ipfactory.InitFactory(ipArr[0], ipArr[1])

	for k, v := range cloudFoundryJobsMap {
		if success {
			assignaIP, _ := ipfactory.AssignaIP2Job(k, v.Instances)
			v.StaticIp = assignaIP
		}
		cloudFoundryJobsMap[k] = v
	}

	for index, value := range mapResourcesPools {
		poolsize := 0
		for _, v := range cloudFoundryJobsMap {
			if value.Name == v.ResourcesPool {
				poolsize = poolsize + 1
			}
		}
		value.Size = poolsize
		mapResourcesPools[index] = value
	}

	var cft = NewCloudFoundryTemplate(CloudFoundry{
		CloudFoundryProperties: cloudFoundryProperties,
		Compilation:            compilation,
		NetWorks:               mapNetWorks,
		ResourcesPools:         mapResourcesPools,
		CloudFoundryJobs:       cloudFoundryJobsMap,
	})
	cft.CloudFoundry.Release = NewRelease("cf", "170")
	cft.CloudFoundry.Stemcells = NewStemcell("", "")
	return cft
}

func InitCloudFoundryV3() CloudFoundryTemplate {
	var cloudProperties = CloudProperties{
		AuthUrl:        "http://192.168.128.2:5000/v2.0",
		UserName:       "paas",
		ApiKey:         "paas123",
		Tenant:         "Project_paas",
		DefaultKeyName: "test_microbosh",
		PrivateKey:     "~/bosh-workspace/key/test_microbosh.private",
	}

	var cloudFoundryProperties = NewCloudFoundryProperties("cf-release",
		"57cfc863-786d-4495-bb97-86d2f650a038", "10.162.2.28", "ccipaas.net", "cci")

	cloudFoundryProperties.CloudProperties = cloudProperties

	var compilation = NewCompilation("flavor_91", "zone2", 6, "cf1")

	var networks = NewNetWorks("cf1", "private",
		"manual",
		"8bb21e6e-dc6a-409c-82d0-a110fb3c9fe1",
		"192.168.129.0/24",
		"10.10.170.2",
		"192.168.133.108",
		"192.168.129.1 - 192.168.129.99",
		"192.168.129.100 - 192.168.129.126")

	var mapNetWorks = make(map[string]NetWorks)
	mapNetWorks["private"] = networks
	mapNetWorks["public"] = NewNetWorks("publc", "publc",
		"manual",
		"8bb21e6e-dc6a-409c-82d0-a110fb3c9fe1",
		"10.162.2.0/24",
		"",
		"",
		"",
		"10.162.2.28 - 10.162.2.29")

	var resourcesPool = NewResourcesPools("normal",
		"flavor_91",
		"zone2",
		"cf1",
		4)
	var mapResourcesPools = make([]ResourcesPools, 0)
	mapResourcesPools = append(mapResourcesPools, resourcesPool)

	cloudFoundryJobsMap := map[string]CloudFoundryJobs{
		utils.Job_Haproxy: NewCloudFoundryJobs(
			"haproxy",
			utils.Job_Haproxy,
			"normal",
			1,
			[]string{""}),
		utils.Job_Gorouter: NewCloudFoundryJobs(
			"gorouter",
			utils.Job_Gorouter,
			"normal",
			2,
			[]string{""}),
		utils.Job_Postgres: NewCloudFoundryJobs(
			"postgres",
			utils.Job_Postgres,
			"normal",
			1,
			[]string{""}),
		utils.Job_NFS: NewCloudFoundryJobs(
			"nfs",
			utils.Job_NFS,
			"normal",
			1,
			[]string{""}),
		utils.Job_NATS: NewCloudFoundryJobs(
			"nats",
			utils.Job_NATS,
			"normal",
			1,
			[]string{""}),
		utils.Job_Syslog: NewCloudFoundryJobs(
			"syslog_aggregator",
			utils.Job_Syslog,
			"normal",
			1,
			[]string{""}),
		utils.Job_Etcd: NewCloudFoundryJobs(
			"etcd",
			utils.Job_Etcd,
			"normal",
			1,
			[]string{""}),
		utils.Job_Loggregator: NewCloudFoundryJobs(
			"loggregator",
			utils.Job_Loggregator,
			"normal",
			1,
			[]string{""}),
		utils.Job_Loggregator_Traffic: NewCloudFoundryJobs(
			"loggregator_traffic",
			utils.Job_Loggregator_Traffic,
			"normal",
			1,
			[]string{""}),
		utils.Job_Uaa: NewCloudFoundryJobs(
			"uaa", utils.Job_Uaa,
			"normal",
			1,
			[]string{""}),
		utils.Job_Cloud_Controller_NG: NewCloudFoundryJobs(
			"cloud_controller_ng",
			utils.Job_Cloud_Controller_NG,
			"normal",
			1,
			[]string{""}),
		utils.Job_Cloud_Controller_Worker: NewCloudFoundryJobs(
			"cloud_controller_worker",
			utils.Job_Cloud_Controller_Worker,
			"normal",
			1,
			[]string{""}),
		utils.Job_Cloud_Controller_Clock: NewCloudFoundryJobs(
			"cloud_controller_clock",
			utils.Job_Cloud_Controller_Clock,
			"normal",
			1,
			[]string{""}),
		utils.Job_Hm9000: NewCloudFoundryJobs(
			"hm9000",
			utils.Job_Hm9000,
			"normal",
			1,
			[]string{""}),
		utils.Job_Stats: NewCloudFoundryJobs("stats",
			utils.Job_Stats,
			"normal",
			1,
			[]string{""}),
		utils.Job_Dea_Next: NewCloudFoundryJobs(
			"dea_next",
			utils.Job_Dea_Next,
			"normal",
			1,
			[]string{""}),
	}

	ipfactory := utils.NewIPFactory()
	ipArr := utils.SpliteIPs(networks.StaticIp)
	success, _ := ipfactory.InitFactory(ipArr[0], ipArr[1])

	for k, v := range cloudFoundryJobsMap {
		if success {
			assignaIP, _ := ipfactory.AssignaIP2Job(k, v.Instances)
			v.StaticIp = assignaIP
		}
		cloudFoundryJobsMap[k] = v
	}

	for index, value := range mapResourcesPools {
		poolsize := 0
		for _, v := range cloudFoundryJobsMap {
			if value.Name == v.ResourcesPool {
				poolsize = poolsize + 1
			}
		}
		value.Size = poolsize
		mapResourcesPools[index] = value
	}

	var cft = NewCloudFoundryTemplate(CloudFoundry{
		CloudFoundryProperties: cloudFoundryProperties,
		Compilation:            compilation,
		NetWorks:               mapNetWorks,
		ResourcesPools:         mapResourcesPools,
		CloudFoundryJobs:       cloudFoundryJobsMap,
	})

	cft.CloudFoundry.Release = NewRelease("cf", "170")
	cft.CloudFoundry.Stemcells = NewStemcell("", "")
	return cft
}
