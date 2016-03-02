$(document).ready(function(){
	var result = 
	{
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
	}
	
	
	setInterval(function(){
		result.value.vitals.cpu.sys = Math.random()+"";
		result.value.vitals.cpu.user = Math.random()+"";
		result.value.vitals.mem.kb = Math.random()*1000000+"";
		result.value.vitals.disk.persistent.percent = Math.random()*100+"";
		result.value.vitals.disk.system.percent = Math.random()*100+"";
		result.value.vitals.load[0] = Math.random()*2+"";
		result.value.vitals.load[1] = Math.random()*2+"";
		result.value.vitals.load[2] = Math.random()*2+"";
	},1000);
	
	
	
		Highcharts.setOptions({
            global: {
                useUTC: false
            }
        });

        $('#cpu-monitor').highcharts({
            chart: {
                type: 'spline',
                animation: false, 
                marginRight: 10,
                events: {
                    load: function () {
                        var system_cpu = this.series[0];
						var user_cpu =  this.series[1];
                        setInterval(function () {
							var sysy = Number(result.value.vitals.cpu.sys),x = (new Date()).getTime();
							var usery = Number(result.value.vitals.cpu.user);
                            system_cpu.addPoint([x, sysy], true, true);
							user_cpu.addPoint([x,usery],true,true);
                        }, 5000);
                    }
                }
            },
            title: {
                text: 'CPU Monitor',
				x: -20
            },
            xAxis: {
                type: 'datetime',
                tickPixelInterval: 150
            },
            yAxis: {
                title: {
                    text: 'CPU used percent(%)'
                },
                plotLines: [{
                    value: 0,
                    width: 1,
                    color: '#808080'
                }]
            },
            tooltip: {
                formatter: function () {
                    return '<b>' + this.series.name + '</b><br/>' +
                        Highcharts.dateFormat('%Y-%m-%d %H:%M:%S', this.x) + '<br/>' +
                        Highcharts.numberFormat(this.y, 2);
                }
            },
            legend: {
                layout: 'vertical',
            	align: 'right',
            	verticalAlign: 'top',
            	borderWidth: 0
            },
            exporting: {
                enabled: false
            },
            series: [{
                name: 'sys used',
                data: (function () {
                    var data = [],
                        time = (new Date()).getTime(),
                        i;

                    for (i = -20; i <= 0; i += 1) {
                        data.push({
                            x: time + i * 5000,
                            y: 0
                        });
                    }
                    return data;
                }())
            },{
                name: 'user used',
                data: (function () {
                    // generate an array of random data
                    var data = [],
                        time = (new Date()).getTime(),
                        i;

                    for (i = -20; i <= 0; i += 1) {
                        data.push({
                            x: time + i * 5000,
                            y: 0
                        });
                    }
                    return data;
                }())
            }],
			credits:{
				enabled:false
			}
        });
        $('#mem-monitor').highcharts({
            chart: {
                type: 'spline',
                animation: false, 
                marginRight: 10,
                events: {
                    load: function () {
                        var memused = this.series[0];
                        setInterval(function () {
                            var x = (new Date()).getTime(), 
                                y = Number(result.value.vitals.mem.kb)/1024;
                            memused.addPoint([x, y], true, true);
                        }, 5000);
                    }
                }
            },
            title: {
                text: 'Memory Monitor',
				x: -20
            },
            xAxis: {
                type: 'datetime',
                tickPixelInterval: 150
            },
            yAxis: {
                title: {
                    text: 'Memory used (MB)'
                },
                plotLines: [{
                    value: 0,
                    width: 1,
                    color: '#808080'
                }]
            },
            tooltip: {
                formatter: function () {
                    return '<b>' + this.series.name + '</b><br/>' +
                        Highcharts.dateFormat('%Y-%m-%d %H:%M:%S', this.x) + '<br/>' +
                        Highcharts.numberFormat(this.y, 2);
                }
            },
            legend: {
                layout: 'vertical',
            	align: 'right',
            	verticalAlign: 'top',
            	borderWidth: 0
            },
            exporting: {
                enabled: false
            },
            series: [{
                name: 'Memory used',
                data: (function () {
                    var data = [],
                        time = (new Date()).getTime(),
                        i;

                    for (i = -20; i <= 0; i += 1) {
                        data.push({
                            x: time + i * 5000,
                            y: 0
                        });
                    }
                    return data;
                }())
            }],
			credits:{
				enabled:false
			}
        });
        $('#disk-monitor').highcharts({
            chart: {
                type: 'spline',
                animation: false, 
                marginRight: 10,
                events: {
                    load: function () {
                        var persistent_disk_used = this.series[0];
						var system_disk_used =  this.series[1];
                        setInterval(function () {
                            var x = (new Date()).getTime(), 
                                persistent_disk_used_y = Number(result.value.vitals.disk.persistent.percent),
								system_disk_used_y = Number(result.value.vitals.disk.system.percent);
                            persistent_disk_used.addPoint([x, persistent_disk_used_y], true, true);
							system_disk_used.addPoint([x,system_disk_used_y],true,true);
                        }, 5000);
                    }
                }
            },
            title: {
                text: 'Disk Monitor',
				x: -20
            },
            xAxis: {
                type: 'datetime',
                tickPixelInterval: 150
            },
            yAxis: {
                title: {
                    text: 'Disk used percent (%)'
                },
                plotLines: [{
                    value: 0,
                    width: 1,
                    color: '#808080'
                }],
				max:100
            },
            tooltip: {
                formatter: function () {
                    return '<b>' + this.series.name + '</b><br/>' +
                        Highcharts.dateFormat('%Y-%m-%d %H:%M:%S', this.x) + '<br/>' +
                        Highcharts.numberFormat(this.y, 0);
                }
            },
            legend: {
                layout: 'vertical',
            	align: 'right',
            	verticalAlign: 'top',
            	borderWidth: 0
            },
            exporting: {
                enabled: false
            },
            series: [{
                name: 'Disk Persistent',
                data: (function () {
                    var data = [],
                        time = (new Date()).getTime(),
                        i;

                    for (i = -20; i <= 0; i += 1) {
                        data.push({
                            x: time + i * 5000,
                            y: 0
                        });
                    }
                    return data;
                }())
            },{
                name: 'Disk System',
                data: (function () {
                    var data = [],
                        time = (new Date()).getTime(),
                        i;

                    for (i = -20; i <= 0; i += 1) {
                        data.push({
                            x: time + i * 5000,
                            y: 0
                        });
                    }
                    return data;
                }())
            }],
			credits:{
				enabled:false
			}
        });
        $('#load-monitor').highcharts({
            chart: {
                type: 'spline',
                animation: false, 
                marginRight: 10,
                events: {
                    load: function () {
                        var one_min = this.series[0];
						var five_min =  this.series[1];
						var fifteen_min =  this.series[2];
                        setInterval(function () {
                            var x = (new Date()).getTime(), 
                                one_min_y = Number(result.value.vitals.load[0]),
								five_min_y = Number(result.value.vitals.load[1]),
								fifteen_min_y = Number(result.value.vitals.load[2]);
                            one_min.addPoint([x, one_min_y], true, true);
							five_min.addPoint([x,five_min_y],true,true);
							fifteen_min.addPoint([x,fifteen_min_y],true,true);
                        }, 5000);
                    }
                }
            },
            title: {
                text: 'Load Average Monitor',
				x: -20
            },
            xAxis: {
                type: 'datetime',
                tickPixelInterval: 150
            },
            yAxis: {
                title: {
                    text: 'Disk used percent (%)'
                },
                plotLines: [{
                    value: 0,
                    width: 1,
                    color: '#808080'
                }]
            },
            tooltip: {
                formatter: function () {
                    return '<b>' + this.series.name + '</b><br/>' +
                        Highcharts.dateFormat('%Y-%m-%d %H:%M:%S', this.x) + '<br/>' +
                        Highcharts.numberFormat(this.y, 2);
                }
            },
            legend: {
                layout: 'vertical',
            	align: 'right',
            	verticalAlign: 'top',
            	borderWidth: 0
            },
            exporting: {
                enabled: false
            },
            series: [{
                name: '1 min',
                data: (function () {
                    var data = [],
                        time = (new Date()).getTime(),
                        i;

                    for (i = -20; i <= 0; i += 1) {
                        data.push({
                            x: time + i * 5000,
                            y: 0
                        });
                    }
                    return data;
                }())
            },{
                name: '5 min',
                data: (function () {
                    var data = [],
                        time = (new Date()).getTime(),
                        i;

                    for (i = -20; i <= 0; i += 1) {
                        data.push({
                            x: time + i * 5000,
                            y: 0
                        });
                    }
                    return data;
                }())
            },{
                name: '15 min',
                data: (function () {
                    var data = [],
                        time = (new Date()).getTime(),
                        i;

                    for (i = -20; i <= 0; i += 1) {
                        data.push({
                            x: time + i * 5000,
                            y: 0
                        });
                    }
                    return data;
                }())
            }],
			credits:{
				enabled:false
			}
        });
});