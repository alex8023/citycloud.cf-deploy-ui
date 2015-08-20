    {{with .CloudFoundryJobs}}
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">Jobs-{{.hm9000.JobName}}</h2>
		</div>
		<div class="form-horizontal">
		  	<div class="panel-body">
			  	<div class="form-group">
			    	<label for="{{.hm9000.JobName}}_name" class="col-sm-2 control-label">Name</label>
				    <div class="col-sm-10">
				      	<input type="text" class="form-control" id="{{.hm9000.JobName}}_name" placeholder="Name" name="{{.hm9000.JobName}}_name" value = "{{.hm9000.Name}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.hm9000.JobName}}_resourcesPool" class="col-sm-2 control-label">ResourcesPool</label>
				    <div class="col-sm-10">
				      	<input type="hidden" id="{{.hm9000.JobName}}_resourcesPool" name="{{.hm9000.JobName}}_resourcesPool" value = "{{.hm9000.ResourcesPool}}">
				    	<select class="form-control" id="{{.hm9000.JobName}}_resourcesPool" name="{{.hm9000.JobName}}_resourcesPool_select" >
						{{$hm9000 := .hm9000.JobName}}
						{{range $.ResourcesPools}}<option id="{{$hm9000}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
						</select>
						<script type="text/javascript">
						$("#{{.hm9000.JobName}}_{{.hm9000.ResourcesPool}}").attr("selected",true);  
						</script>
					</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.hm9000.JobName}}_instances" class="col-sm-2 control-label">Instances</label>
				    <div class="col-sm-10">
				      	<input type="number" class="form-control" id="{{.hm9000.JobName}}_instances" placeholder="Instances" name="{{.hm9000.JobName}}_instances" value = "{{.hm9000.Instances}}" required readonly>
				    </div>
			  	</div>
			</div>
		</div>
	</div>
	{{end}}