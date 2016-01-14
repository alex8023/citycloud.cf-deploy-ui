    {{with .CloudFoundry.CloudFoundryJobs}}
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">{{i18n $.Lang "Jobs"}}-{{.dea_next.JobName}}</h2>
		</div>
		<div class="form-horizontal">
		  	<div class="panel-body">
			  	<div class="form-group">
			    	<label for="{{.dea_next.JobName}}_name" class="col-sm-3 control-label">{{i18n $.Lang "Name"}}</label>
				    <div class="col-sm-7">
				      	<input type="text" class="form-control" id="{{.dea_next.JobName}}_name" placeholder="{{i18n $.Lang "Jobs"}}" name="{{.dea_next.JobName}}_name" value = "{{.dea_next.Name}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.dea_next.JobName}}_resourcesPool" class="col-sm-3 control-label">{{i18n $.Lang "ResourcesPool"}}</label>
				    <div class="col-sm-7">
				      	<input type="hidden" id="{{.dea_next.JobName}}_resourcesPool" name="{{.dea_next.JobName}}_resourcesPool" value = "{{.dea_next.ResourcesPool}}">
				    	<select class="form-control" id="{{.dea_next.JobName}}_resourcesPool" name="{{.dea_next.JobName}}_resourcesPool_select" >
						{{$dea_next := .dea_next.JobName}}
						{{range $.CloudFoundry.VsphereResourcesPools}}<option id="{{$dea_next}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
						</select>
						<script type="text/javascript">
						$("#{{.dea_next.JobName}}_{{.dea_next.ResourcesPool}}").attr("selected",true);  
						</script>
					</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.dea_next.JobName}}_instances" class="col-sm-3 control-label">{{i18n $.Lang "Instances"}}</label>
				    <div class="col-sm-7">
				      	<input type="number" class="form-control" id="{{.dea_next.JobName}}_instances" placeholder="{{i18n $.Lang "Instances"}}" name="{{.dea_next.JobName}}_instances" value = "{{.dea_next.Instances}}" required >
				    </div>
			  	</div>
			</div>
		</div>
	</div>
	{{end}}