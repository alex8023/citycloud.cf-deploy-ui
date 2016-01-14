    {{with .CloudFoundry.CloudFoundryJobs}}
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">{{i18n $.Lang "Jobs"}}-{{.syslog_aggregator.JobName}}</h2>
		</div>
		<div class="form-horizontal">
		  	<div class="panel-body">
			  	<div class="form-group">
			    	<label for="{{.syslog_aggregator.JobName}}_name" class="col-sm-3 control-label">{{i18n $.Lang "Name"}}</label>
				    <div class="col-sm-7">
				      	<input type="text" class="form-control" id="{{.syslog_aggregator.JobName}}_name" placeholder="{{i18n $.Lang "Name"}}" name="{{.syslog_aggregator.JobName}}_name" value = "{{.syslog_aggregator.Name}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.syslog_aggregator.JobName}}_resourcesPool" class="col-sm-3 control-label">{{i18n $.Lang "ResourcesPool"}}</label>
				    <div class="col-sm-7">
				      	<input type="hidden" id="{{.syslog_aggregator.JobName}}_resourcesPool" name="{{.syslog_aggregator.JobName}}_resourcesPool" value = "{{.syslog_aggregator.ResourcesPool}}">
				    	<select class="form-control" id="{{.syslog_aggregator.JobName}}_resourcesPool" name="{{.syslog_aggregator.JobName}}_resourcesPool_select" >
						{{$syslog_aggregator := .syslog_aggregator.JobName}}
						{{range $.CloudFoundry.VsphereResourcesPools}}<option id="{{$syslog_aggregator}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
						</select>
						<script type="text/javascript">
						$("#{{.syslog_aggregator.JobName}}_{{.syslog_aggregator.ResourcesPool}}").attr("selected",true);  
						</script>
					</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.syslog_aggregator.JobName}}_instances" class="col-sm-3 control-label">{{i18n $.Lang "Instances"}}</label>
				    <div class="col-sm-7">
				      	<input type="number" class="form-control" id="{{.syslog_aggregator.JobName}}_instances" placeholder="{{i18n $.Lang "Instances"}}" name="{{.syslog_aggregator.JobName}}_instances" value = "{{.syslog_aggregator.Instances}}" required readonly>
				    </div>
			  	</div>
			</div>
		</div>
	</div>
	{{end}}