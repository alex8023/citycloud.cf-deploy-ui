    {{with .CloudFoundry}}
	{{with .CloudFoundryJobs}}
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">{{i18n $.Lang "Jobs"}}-{{.postgres.JobName}}</h2>
		</div>
		<div class="form-horizontal">
		  	<div class="panel-body">
			  	<div class="form-group">
			    	<label for="{{.postgres.JobName}}_name" class="col-sm-3 control-label">{{i18n $.Lang "Name"}}</label>
				    <div class="col-sm-7">
				      	<input type="text" class="form-control" id="{{.postgres.JobName}}_name" placeholder="{{i18n $.Lang "Name"}}" name="{{.postgres.JobName}}_name" value = "{{.postgres.Name}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.postgres.JobName}}_resourcesPool" class="col-sm-3 control-label">{{i18n $.Lang "ResourcesPool"}}</label>
				    <div class="col-sm-7">
				      	<input type="hidden" id="{{.postgres.JobName}}_resourcesPool" name="{{.postgres.JobName}}_resourcesPool" value = "{{.postgres.ResourcesPool}}">
				    	<select class="form-control" id="{{.postgres.JobName}}_resourcesPool" name="{{.postgres.JobName}}_resourcesPool_select" >
						{{$postgres := .postgres.JobName}}
						{{range $.CloudFoundry.ResourcesPools}}<option id="{{$postgres}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
						</select>
						<script type="text/javascript">
						$("#{{.postgres.JobName}}_{{.postgres.ResourcesPool}}").attr("selected",true);  
						</script>
					</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.postgres.JobName}}_instances" class="col-sm-3 control-label">{{i18n $.Lang "Instances"}}</label>
				    <div class="col-sm-7">
				      	<input type="number" class="form-control" id="{{.postgres.JobName}}_instances" placeholder="{{i18n $.Lang "Instances"}}" name="{{.postgres.JobName}}_instances" value = "{{.postgres.Instances}}" required readonly>
				    </div>
			  	</div>
			</div>
		</div>
	</div>
	{{end}}{{end}}