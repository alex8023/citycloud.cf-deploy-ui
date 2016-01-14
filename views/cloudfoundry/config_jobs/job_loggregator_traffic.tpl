     {{with .CloudFoundry}}
	 {{with .CloudFoundryJobs}}
 	<div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">{{i18n $.Lang "Jobs"}}-{{.loggregator_traffic.JobName}}</h2>
		</div>
		<div class="form-horizontal">
		  	<div class="panel-body">
			  	<div class="form-group">
			    	<label for="{{.loggregator_traffic.JobName}}_name" class="col-sm-3 control-label">{{i18n $.Lang "Name"}}</label>
				    <div class="col-sm-7">
				      	<input type="text" class="form-control" id="{{.loggregator_traffic.JobName}}_name" placeholder="{{i18n $.Lang "Name"}}" name="{{.loggregator_traffic.JobName}}_name" value = "{{.loggregator_traffic.Name}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.loggregator_traffic.JobName}}_resourcesPool" class="col-sm-3 control-label">{{i18n $.Lang "ResourcesPool"}}</label>
				    <div class="col-sm-7">
				      	<input type="hidden" id="{{.loggregator_traffic.JobName}}_resourcesPool" name="{{.loggregator_traffic.JobName}}_resourcesPool" value = "{{.loggregator_traffic.ResourcesPool}}">
				    	<select class="form-control" id="{{.loggregator_traffic.JobName}}_resourcesPool" name="{{.loggregator_traffic.JobName}}_resourcesPool_select" >
						{{$loggregator := .loggregator_traffic.JobName}}
						{{range $.CloudFoundry.ResourcesPools}}<option id="{{$loggregator}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
						</select>
						<script type="text/javascript">
						$("#{{.loggregator_traffic.JobName}}_{{.loggregator_traffic.ResourcesPool}}").attr("selected",true);  
						</script>
					</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.loggregator_traffic.JobName}}_instances" class="col-sm-3 control-label">{{i18n $.Lang "Instances"}}</label>
				    <div class="col-sm-7">
				      	<input type="number" class="form-control" id="{{.loggregator_traffic.JobName}}_instances" placeholder="{{i18n $.Lang "Instances"}}" name="{{.loggregator_traffic.JobName}}_instances" value = "{{.loggregator_traffic.Instances}}" required readonly>
				    </div>
			  	</div>
			</div>
		</div>
	</div>
	{{end}}{{end}}