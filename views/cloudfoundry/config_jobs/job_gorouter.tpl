    {{with .CloudFoundryJobs}}
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">Jobs-{{.gorouter.JobName}}</h2>
		</div>
		<div class="form-horizontal">
		  	<div class="panel-body">
			  	<div class="form-group">
			    	<label for="{{.gorouter.JobName}}_name" class="col-sm-2 control-label">Name</label>
				    <div class="col-sm-10">
				      	<input type="text" class="form-control" id="{{.gorouter.JobName}}_name" placeholder="Name" name="{{.gorouter.JobName}}_name" value = "{{.gorouter.Name}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.gorouter.JobName}}_resourcesPool" class="col-sm-2 control-label">ResourcesPool</label>
				    <div class="col-sm-10">
				      	<input type="hidden" id="{{.gorouter.JobName}}_resourcesPool" name="{{.gorouter.JobName}}_resourcesPool" value = "{{.gorouter.ResourcesPool}}">
				    	<select class="form-control" id="{{.gorouter.JobName}}_resourcesPool" name="{{.gorouter.JobName}}_resourcesPool_select" >
						{{$gorouter := .gorouter.JobName}}
						{{range $.ResourcesPools}}<option id="{{$gorouter}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
						</select>
						<script type="text/javascript">
						$("#{{.gorouter.JobName}}_{{.gorouter.ResourcesPool}}").attr("selected",true);  
						</script>
					</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.gorouter.JobName}}_instances" class="col-sm-2 control-label">Instances</label>
				    <div class="col-sm-10">
				      	<input type="number" class="form-control" id="{{.gorouter.JobName}}_instances" placeholder="Instances" name="{{.gorouter.JobName}}_instances" value = "{{.gorouter.Instances}}" required >
				    </div>
			  	</div>
			</div>
		</div>
	</div>
	{{end}}