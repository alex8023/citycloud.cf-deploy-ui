package entity_test

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/citycloud/citycloud.cf-deploy-ui/entity"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Testing OpsMonitor CRUD", func() {
	var result = `{
	    "value": {
	        "properties": {
	            "logging": {
	                "max_log_file_size": ""
	            }
	        },
	        "job": {
	            "name": "postgres",
	            "release": "",
	            "template": "postgres",
	            "version": "5",
	            "sha1": "0228430151570ea1e6cbb0c26c79b5ceaa83ef4d",
	            "blobstore_id": "f26d465e-55fb-4efa-9b6f-3c03e3c06bba",
	            "templates": [
	                {
	                    "name": "postgres",
	                    "version": "5",
	                    "sha1": "0228430151570ea1e6cbb0c26c79b5ceaa83ef4d",
	                    "blobstore_id": "f26d465e-55fb-4efa-9b6f-3c03e3c06bba"
	                }
	            ]
	        },
	        "packages": {
	            "common": {
	                "name": "common",
	                "version": "6.1",
	                "sha1": "41bbc068713cc4a2de2433a3b8d8f03ce399c442",
	                "blobstore_id": "17e30dc6-434b-4d5a-5cda-6203395d0333"
	            },
	            "postgres": {
	                "name": "postgres",
	                "version": "5.1",
	                "sha1": "92a88bb74c02c9fe8459590ef695dbc6f4d413f0",
	                "blobstore_id": "5780870b-09f3-46d8-6bb2-b96f52ed3529"
	            }
	        },
	        "configuration_hash": " ",
	        "networks": {
	            "cf12": {
	                "cloud_properties": {
	                    "net_id": "53f22403-b386-4158-a30f-07f0b8cc2ea7"
	                },
	                "default": [
	                    "dns",
	                    "gateway"
	                ],
	                "dns": [
	                    "192.168.138.11",
	                    "10.10.245.59"
	                ],
	                "dns_record_name": "0.postgres.cf12.cf-release.microbosh",
	                "ip": "192.168.138.22",
	                "netmask": "255.255.255.0"
	            }
	        },
	        "resource_pool": {
	            "cloud_properties": {
	                "availability_zone": "bigdata",
	                "instance_type": "flavor_737"
	            },
	            "name": "normal",
	            "stemcell": {
	                "name": "bosh-openstack-kvm-ubuntu-lucid-go_agent",
	                "version": "2719"
	            }
	        },
	        "deployment": "cf-release",
	        "index": 0,
	        "persistent_disk": 30720,
	        "persistent_disk_pool": {
	            "cloud_properties": {},
	            "disk_size": 30720,
	            "name": "e629ad1c-48ea-44ca-82fe-02a023c50819"
	        },
	        "rendered_templates_archive": {
	            "sha1": "36f90aaa68ef2ad08e3e8a3b2e2d01490e8ad02a",
	            "blobstore_id": "a43cbd05-4917-47dc-968d-f316063d078b"
	        },
	        "agent_id": "6b3733ee-7b6f-4d65-9b68-d3d08ddba267",
	        "bosh_protocol": "1",
	        "job_state": "running",
	        "vitals": {
	            "cpu": {
	                "sys": "0.0",
	                "user": "0.1",
	                "wait": "0.0"
	            },
	            "disk": {
	                "ephemeral": {
	                    "inode_percent": "0",
	                    "percent": "1"
	                },
	                "persistent": {
	                    "inode_percent": "0",
	                    "percent": "1"
	                },
	                "system": {
	                    "inode_percent": "32",
	                    "percent": "40"
	                }
	            },
	            "load": [
	                "0.02",
	                "0.08",
	                "0.12"
	            ],
	            "mem": {
	                "kb": "164720",
	                "percent": "8"
	            },
	            "swap": {
	                "kb": "0",
	                "percent": "0"
	            }
	        },
	        "vm": {
	            "name": "vm-447d7acb-4a60-4932-b8fc-d2dc155eeccc"
	        },
	        "ntp": {
	            "offset": "0.134190",
	            "timestamp": "22 Feb 11:15:07"
	        }
	    }
	}`
	var nats entity.NatsResult
	BeforeEach(func() {
		err := json.Unmarshal([]byte(result), &nats)
		fmt.Println(err)
	})

	It("Testing Insert OpsMonitor", func() {
		monitor := entity.Monitor{}
		monitor.AgentId = nats.Value.AgentId
		monitor.Index = nats.Value.Index
		monitor.JobName = nats.Value.Job.Name
		monitor.JobState = nats.Value.JobState
		monitor.Value = result
		monitor.Save()
		monitor.Delete()
	})

	It("Testing Load OpsMonitor", func() {
		sq, err := orm.NewOrm().Raw("insert into monitor (updated) values(?)", time.Now()).Exec()
		Expect(err).NotTo(HaveOccurred())
		id, _ := sq.LastInsertId()
		sq, err = orm.NewOrm().Raw("delete from monitor where id = ?", id).Exec()
		Expect(err).NotTo(HaveOccurred())
		results, _ := sq.RowsAffected()
		Expect(results).To(Equal(int64(1)))
	})
})
