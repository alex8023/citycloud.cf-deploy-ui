    {{with .CloudFoundryJobs}}
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">Jobs-{{.haproxy.JobName}}</h2>
		</div>
		<div class="form-horizontal">
		  	<div class="panel-body">
			  	<div class="form-group">
			    	<label for="{{.haproxy.JobName}}_name" class="col-sm-3 control-label">Name</label>
				    <div class="col-sm-7">
				      	<input type="text" class="form-control" id="{{.haproxy.JobName}}_name" placeholder="Name" name="{{.haproxy.JobName}}_name" value = "{{.haproxy.Name}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.haproxy.JobName}}_resourcesPool" class="col-sm-3 control-label">ResourcesPool</label>
				    <div class="col-sm-7">
				      	<input type="hidden" id="{{.haproxy.JobName}}_resourcesPool" name="{{.haproxy.JobName}}_resourcesPool" value = "{{.haproxy.ResourcesPool}}">
				    	<select class="form-control" id="{{.haproxy.JobName}}_resourcesPool" name="{{.haproxy.JobName}}_resourcesPool_select" >
						{{$haproxy := .haproxy.JobName}}
						{{range $.ResourcesPools}}<option id="{{$haproxy}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
						</select>
						<script type="text/javascript">
						$("#{{.haproxy.JobName}}_{{.haproxy.ResourcesPool}}").attr("selected",true);  
						</script>
					</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.haproxy.JobName}}_instances" class="col-sm-3 control-label">Instances</label>
				    <div class="col-sm-7">
				      	<input type="number" class="form-control" id="{{.haproxy.JobName}}_instances" placeholder="Instances" name="{{.haproxy.JobName}}_instances" value = "{{.haproxy.Instances}}" required readonly>
				    </div>
			  	</div>
			</div>
		</div>
	</div>
	{{end}}