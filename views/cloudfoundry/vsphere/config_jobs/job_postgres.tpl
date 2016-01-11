    {{with .CloudFoundryJobs}}
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">Jobs-{{.postgres.JobName}}</h2>
		</div>
		<div class="form-horizontal">
		  	<div class="panel-body">
			  	<div class="form-group">
			    	<label for="{{.postgres.JobName}}_name" class="col-sm-3 control-label">Name</label>
				    <div class="col-sm-7">
				      	<input type="text" class="form-control" id="{{.postgres.JobName}}_name" placeholder="Name" name="{{.postgres.JobName}}_name" value = "{{.postgres.Name}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.postgres.JobName}}_resourcesPool" class="col-sm-3 control-label">ResourcesPool</label>
				    <div class="col-sm-7">
				      	<input type="hidden" id="{{.postgres.JobName}}_resourcesPool" name="{{.postgres.JobName}}_resourcesPool" value = "{{.postgres.ResourcesPool}}">
				    	<select class="form-control" id="{{.postgres.JobName}}_resourcesPool" name="{{.postgres.JobName}}_resourcesPool_select" >
						{{$postgres := .postgres.JobName}}
						{{range $.VsphereResourcesPools}}<option id="{{$postgres}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
						</select>
						<script type="text/javascript">
						$("#{{.postgres.JobName}}_{{.postgres.ResourcesPool}}").attr("selected",true);  
						</script>
					</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.postgres.JobName}}_instances" class="col-sm-3 control-label">Instances</label>
				    <div class="col-sm-7">
				      	<input type="number" class="form-control" id="{{.postgres.JobName}}_instances" placeholder="Instances" name="{{.postgres.JobName}}_instances" value = "{{.postgres.Instances}}" required readonly>
				    </div>
			  	</div>
			</div>
		</div>
	</div>
	{{end}}