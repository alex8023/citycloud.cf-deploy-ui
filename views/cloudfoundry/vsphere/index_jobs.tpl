{{with .CloudFoundry}}
			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">PaaS Jobs</h2>
					</div>
					<div class="form-horizontal">
				  		<div class="panel-body">
							<div class="form-group">
							    <div class="col-sm-offset-2 col-sm-10">
							      	<button class="btn btn-default " id = "config-CloudFoundryJobs"><span class="glyphicon glyphicon-edit"></span> Config</button>
							    </div>
								<script type="text/javascript">
									$('#config-CloudFoundryJobs').on('click', function(){
							    		window.location.href = "vspherecloudfoundry?action=config&model=CloudFoundryJobs";
							  		})
								</script>
							</div>
							{{with .CloudFoundryJobs}}
								{{template "cloudfoundry/index_jobs/job_haproxy.tpl" .}}
								{{template "cloudfoundry/index_jobs/job_gorouter.tpl" .}}
								{{template "cloudfoundry/index_jobs/job_dea_next.tpl" .}}
								{{template "cloudfoundry/index_jobs/job_postgres.tpl" .}}
								{{template "cloudfoundry/index_jobs/job_nfs.tpl" .}}
								{{template "cloudfoundry/index_jobs/job_nats.tpl" .}}
								{{template "cloudfoundry/index_jobs/job_syslog_aggregator.tpl" .}}
								{{template "cloudfoundry/index_jobs/job_etcd.tpl" .}}
								{{template "cloudfoundry/index_jobs/job_loggregator.tpl" .}}
								{{template "cloudfoundry/index_jobs/job_loggregator_traffic.tpl" .}}
								{{template "cloudfoundry/index_jobs/job_uaa.tpl" .}}
								{{template "cloudfoundry/index_jobs/job_cloud_controller_ng.tpl" .}}
								{{template "cloudfoundry/index_jobs/job_cloud_controller_worker.tpl" .}}
								{{template "cloudfoundry/index_jobs/job_cloud_controller_clock.tpl" .}}
								{{template "cloudfoundry/index_jobs/job_hm9000.tpl" .}}
								{{template "cloudfoundry/index_jobs/job_stats.tpl" .}}
							{{end}}
							
						</div>
					</div>
				</div>
{{end}}