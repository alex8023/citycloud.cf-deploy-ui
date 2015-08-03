package entity

import (
	"github.com/citycloud/citycloud.cf-deploy-ui/utils"
	"testing"
)

func TestCloudFoundryTemplate(t *testing.T) {
	var cft = InitCloudFoundryV2()
	_, err := cft.CreateCloudFoundryYaml(CloudFoundryTemplateV2, "/home/ubuntu/temp/cloudfoundryv2.yml")

	if err != nil {
		t.Error(err)
	}
}

func TestMeta(t *testing.T) {
	var cft = InitCloudFoundryV2()
	_, err := cft.CreateCloudFoundryYaml(CloudFoundryMetaTemplateV2+`{{template "meta" .}}`, "/home/ubuntu/temp/metav2.yml")

	if err != nil {
		t.Error(err)
	}
}

func TestUpdate(t *testing.T) {
	var cft = InitCloudFoundryV2()
	_, err := cft.CreateCloudFoundryYaml(CloudFoundryUpdateTempalte+`{{template "update" .}}`, "/home/ubuntu/temp/update.yml")

	if err != nil {
		t.Error(err)
	}
}

func TestCompilation(t *testing.T) {
	var cft = InitCloudFoundryV2()
	_, err := cft.CreateCloudFoundryYaml(CloudFoundryCompilationTemplate+`{{template "compilation" .}}`, "/home/ubuntu/temp/compilation.yml")

	if err != nil {
		t.Error(err)
	}
}

func TestNetWorks(t *testing.T) {
	var cft = InitCloudFoundryV2()
	_, err := cft.CreateCloudFoundryYaml(CloudFoundryNetworksTemplateV2+`{{template "networks" .}}`, "/home/ubuntu/temp/networksv2.yml")

	if err != nil {
		t.Error(err)
	}
}

func TestResourcePool(t *testing.T) {
	var cft = InitCloudFoundryV2()
	_, err := cft.CreateCloudFoundryYaml(CloudFoundryResourcePoolTemplate+`{{template "resourcepool" .}}`, "/home/ubuntu/temp/resourcepool.yml")

	if err != nil {
		t.Error(err)
	}
}

func TestJobs(t *testing.T) {
	var cft = InitCloudFoundryV2()
	_, err := cft.CreateCloudFoundryYaml(CloudFoundryJobsTemplate+`{{template "jobs" .}}`, "/home/ubuntu/temp/jobs.yml")

	if err != nil {
		t.Error(err)
	}
}
func TestProperties(t *testing.T) {
	var cft = InitCloudFoundryV2()
	_, err := cft.CreateCloudFoundryYaml(CloudFoundryPropertiesTemplate+`{{template "properties" .}}`, "/home/ubuntu/temp/properties.yml")

	if err != nil {
		t.Error(err)
	}
}

func TestCloudFoundryV3Template(t *testing.T) {
	var cloudFoundryProperties = NewCloudFoundryProperties("cf-release",
		"57cfc863-786d-4495-bb97-86d2f650a038", "192.168.133.102", "ccipaas.net", "cci")
	var cft = NewCloudFoundryTemplate(CloudFoundry{CloudFoundryProperties: cloudFoundryProperties})
	_, err := cft.CreateCloudFoundryV3Yaml("/home/ubuntu/temp/cloudfoundryv3.yml")

	if err != nil {
		t.Error(err)
	}
}

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

	var networks = NewNetWorks("cf1",
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
	return cft
}
