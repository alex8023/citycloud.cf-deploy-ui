 {{with .CloudFoundry}}
 {{with .CloudFoundryJobs}}
	<div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">{{i18n $.Lang "Jobs"}}-{{.cloud_controller_clock.JobName}}</h2>
		</div>
		<div class="form-horizontal">
		  	<div class="panel-body">
			  	<div class="form-group">
			    	<label for="{{.cloud_controller_clock.JobName}}_name" class="col-sm-3 control-label">{{i18n $.Lang "Name"}}</label>
				    <div class="col-sm-7">
				      	<input type="text" class="form-control" id="{{.cloud_controller_clock.JobName}}_name" placeholder="{{i18n $.Lang "Name"}}" name="{{.cloud_controller_clock.JobName}}_name" value = "{{.cloud_controller_clock.Name}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.cloud_controller_clock.JobName}}_resourcesPool" class="col-sm-3 control-label">{{i18n $.Lang "ResourcesPool"}}</label>
				    <div class="col-sm-7">
				      	<input type="hidden" id="{{.cloud_controller_clock.JobName}}_resourcesPool" name="{{.cloud_controller_clock.JobName}}_resourcesPool" value = "{{.cloud_controller_clock.ResourcesPool}}">
				    	<select class="form-control" id="{{.cloud_controller_clock.JobName}}_resourcesPool" name="{{.cloud_controller_clock.JobName}}_resourcesPool_select" >
						{{$cloud_controller_clock := .cloud_controller_clock.JobName}}
						{{range $.CloudFoundry.ResourcesPools}}<option id="{{$cloud_controller_clock}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
						</select>
						<script type="text/javascript">
						$("#{{.cloud_controller_clock.JobName}}_{{.cloud_controller_clock.ResourcesPool}}").attr("selected",true);  
						</script>
					</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.cloud_controller_clock.JobName}}_instances" class="col-sm-3 control-label">{{i18n $.Lang "Instances"}}</label>
				    <div class="col-sm-7">
				      	<input type="number" class="form-control" id="{{.cloud_controller_clock.JobName}}_instances" placeholder="{{i18n $.Lang "Instances"}}" name="{{.cloud_controller_clock.JobName}}_instances" value = "{{.cloud_controller_clock.Instances}}" required readonly>
				    </div>
			  	</div>
			</div>
		</div>
	</div>
	{{end}}	{{end}}