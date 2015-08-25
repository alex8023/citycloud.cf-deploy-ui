    {{with .CloudFoundryJobs}}
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">Jobs-{{.nfs.JobName}}</h2>
		</div>
		<div class="form-horizontal">
		  	<div class="panel-body">
			  	<div class="form-group">
			    	<label for="{{.nfs.JobName}}_name" class="col-sm-2 control-label">Name</label>
				    <div class="col-sm-10">
				      	<input type="text" class="form-control" id="{{.nfs.JobName}}_name" placeholder="Name" name="{{.nfs.JobName}}_name" value = "{{.nfs.Name}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.nfs.JobName}}_resourcesPool" class="col-sm-2 control-label">ResourcesPool</label>
				    <div class="col-sm-10">
				      	<input type="hidden" id="{{.nfs.JobName}}_resourcesPool" name="{{.nfs.JobName}}_resourcesPool" value = "{{.nfs.ResourcesPool}}">
				    	<select class="form-control" id="{{.nfs.JobName}}_resourcesPool" name="{{.nfs.JobName}}_resourcesPool_select" >
						{{$nfs := .nfs.JobName}}
						{{range $.ResourcesPools}}<option id="{{$nfs}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
						</select>
						<script type="text/javascript">
						$("#{{.nfs.JobName}}_{{.nfs.ResourcesPool}}").attr("selected",true);  
						</script>
					</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.nfs.JobName}}_instances" class="col-sm-2 control-label">Instances</label>
				    <div class="col-sm-10">
				      	<input type="number" class="form-control" id="{{.nfs.JobName}}_instances" placeholder="Instances" name="{{.nfs.JobName}}_instances" value = "{{.nfs.Instances}}" required readonly>
				    </div>
			  	</div>
			</div>
		</div>
	</div>
	{{end}}