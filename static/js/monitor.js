$(document).ready(function(){
	
//	result.value.vitals.cpu.sys = Math.random()+"";
//	result.value.vitals.cpu.user = Math.random()+"";
//	result.value.vitals.mem.kb = Math.random()*1000000+"";
//	result.value.vitals.disk.persistent.percent = Math.random()*100+"";
//	result.value.vitals.disk.system.percent = Math.random()*100+"";
//	result.value.vitals.load[0] = Math.random()*2+"";
//	result.value.vitals.load[1] = Math.random()*2+"";
//	result.value.vitals.load[2] = Math.random()*2+"";

	// 开始执行
	var sys_cpu ,user_cpu = "0.0";
	var mem_kb = "0";
	var disk_persistent_per , disk_system_per = "0";
	var load1,load5,load15 = "0.00";
	
	
	var interval = setInterval(function(){
		RequestByAgentId();
	},5000);
	
	setTimeout(function(){initHighchart()},5000);
//	initHighchart();
	//执行结束
	
	//获取健康数据请求
	function RequestByAgentId(){
		var agent_id = $("#agent_id").val();
		$.ajax({
			url:'opsmonitorrest',
			dataType:'json',
			data:{agent_id:agent_id},
			success:function(data){
				if (data.Code =="200"){
					_vitals = data.Data.value.vitals;
					sys_cpu = _vitals.cpu.sys;
					user_cpu = _vitals.cpu.user;
					mem_kb = _vitals.mem.kb;
					disk_persistent_per = _vitals.disk.persistent.percent;
					disk_system_per =_vitals.disk.system.percent;
					load1 = _vitals.load[0];
					load5 = _vitals.load[0];
					load15 = _vitals.load[0];
				}else{
					resetResult();
					$("#warning-block-request").attr("class","alert alert-danger alert-dismissible");
					$("#warning-block-request-message").html(data.Data);
					clearInterval(interval);
					
				}
			},
			error:function(data){
				$("#warning-block-request").attr("class","alert alert-danger alert-dismissible")
				$("#warning-block-request-message").html("Request Err")
			}
		});
	}
		
	function resetResult(){
		sys_cpu = "0.0";
		user_cpu = "0.0";
		mem_kb = "0";
		disk_persistent_per ="0"; 
		disk_system_per = "0";
		load1 ="0.00";
		load5 ="0.00"; 
		load15 = "0.00";
	}
	
	function demoResult(){
		sys_cpu = Math.random()+"";
		user_cpu = Math.random()+"";
		mem_kb = Math.random()*1000000+"";
		disk_persistent_per = Math.random()*100+"";
		disk_system_per = Math.random()*100+"";
		load1 =Math.random()*2+"";
		load5 = Math.random()*2+"";
		load15 = Math.random()*2+"";
	}
	//初始化图表，并开始加载数据
	function initHighchart(){
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
						var user_used_cpu =  this.series[1];
                        setInterval(function () {
							var sysy = Number(sys_cpu),x = (new Date()).getTime();
							var usery = Number(user_cpu);
                            system_cpu.addPoint([x, sysy], true, true);
							user_used_cpu.addPoint([x,usery],true,true);
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
                    text: 'CPU Used Percent(%)'
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
                                y = Number(mem_kb)/1024;
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
                    text: 'Memory Used (MB)'
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
                                persistent_disk_used_y = Number(disk_persistent_per),
								system_disk_used_y = Number(disk_system_per);
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
                                one_min_y = Number(load1),
								five_min_y = Number(load5),
								fifteen_min_y = Number(load15);
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
                    text: 'Load Average'
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
	}
});
