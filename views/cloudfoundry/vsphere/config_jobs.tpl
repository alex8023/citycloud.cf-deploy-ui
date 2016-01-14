<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">PaaS {{i18n $.Lang "Jobs"}}</h2>
		</div>
  		<div class="panel-body">
			<form class="form-horizontal" method="post" action="vspherecloudfoundry">
				<input type = "hidden" name="model" value = "{{.Model}}">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      <button type="submit" class="btn btn-default " data-loading-text="Saving..."><span class="glyphicon glyphicon-floppy-save"></span> {{i18n $.Lang "Save"}}</button>
						<a class="btn btn-default " href="vspherecloudfoundry"><span class="glyphicon glyphicon-step-backward"></span> {{i18n $.Lang "Back"}}</a>
				    </div>
			  	</div>
					{{template "cloudfoundry/vsphere/config_jobs/job_haproxy.tpl" .}}
					{{template "cloudfoundry/vsphere/config_jobs/job_gorouter.tpl" .}}
					{{template "cloudfoundry/vsphere/config_jobs/job_dea_next.tpl" .}}
					{{template "cloudfoundry/vsphere/config_jobs/job_postgres.tpl" .}}
					{{template "cloudfoundry/vsphere/config_jobs/job_nfs.tpl" .}}
					{{template "cloudfoundry/vsphere/config_jobs/job_nats.tpl" .}}
					{{template "cloudfoundry/vsphere/config_jobs/job_syslog_aggregator.tpl" .}}
					{{template "cloudfoundry/vsphere/config_jobs/job_etcd.tpl" .}}
					{{template "cloudfoundry/vsphere/config_jobs/job_loggregator.tpl" .}}
					{{template "cloudfoundry/vsphere/config_jobs/job_loggregator_traffic.tpl" .}}
					{{template "cloudfoundry/vsphere/config_jobs/job_uaa.tpl" .}}
					{{template "cloudfoundry/vsphere/config_jobs/job_cloud_controller_ng.tpl" .}}
					{{template "cloudfoundry/vsphere/config_jobs/job_cloud_controller_worker.tpl" .}}
					{{template "cloudfoundry/vsphere/config_jobs/job_cloud_controller_clock.tpl" .}}
					{{template "cloudfoundry/vsphere/config_jobs/job_hm9000.tpl" .}}
					{{template "cloudfoundry/vsphere/config_jobs/job_stats.tpl" .}}
			</form>
  		</div>
	</div>
</div>
