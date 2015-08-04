    {{with .CloudFoundryJobs}}
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">Jobs-{{.uaa.JobName}}</h2>
		</div>
		<div class="form-horizontal">
		  	<div class="panel-body">
			  	<div class="form-group">
			    	<label for="{{.uaa.JobName}}_name" class="col-sm-2 control-label">Name</label>
				    <div class="col-sm-10">
				      	<input type="text" class="form-control" id="{{.uaa.JobName}}_name" placeholder="Name" name="{{.uaa.JobName}}_name" value = "{{.uaa.Name}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.uaa.JobName}}_resourcesPool" class="col-sm-2 control-label">ResourcesPool</label>
				    <div class="col-sm-10">
				      	<input type="hidden" id="{{.uaa.JobName}}_resourcesPool" name="{{.uaa.JobName}}_resourcesPool" value = "{{.uaa.ResourcesPool}}">
				    	<select class="form-control" id="{{.uaa.JobName}}_resourcesPool" name="{{.uaa.JobName}}_resourcesPool_select" >
						{{$uaa := .uaa.JobName}}
						{{range $.ResourcesPools}}<option id="{{$uaa}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
						</select>
						<script type="text/javascript">
						$("#{{.uaa.JobName}}_{{.uaa.ResourcesPool}}").attr("selected",true);  
						</script>
					</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.uaa.JobName}}_instances" class="col-sm-2 control-label">Instances</label>
				    <div class="col-sm-10">
				      	<input type="number" class="form-control" id="{{.uaa.JobName}}_instances" placeholder="Instances" name="{{.etcd.JobName}}_instances" value = "{{.uaa.Instances}}" required readonly>
				    </div>
			  	</div>
			</div>
		</div>
	</div>
	{{end}}