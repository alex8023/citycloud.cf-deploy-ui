    {{with .CloudFoundryJobs}}
	<div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">Jobs-{{.etcd.JobName}}</h2>
		</div>
		<div class="form-horizontal">
		  	<div class="panel-body">
			  	<div class="form-group">
			    	<label for="{{.etcd.JobName}}_name" class="col-sm-3 control-label">Name</label>
				    <div class="col-sm-7">
				      	<input type="text" class="form-control" id="{{.etcd.JobName}}_name" placeholder="Name" name="{{.etcd.JobName}}_name" value = "{{.etcd.Name}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.etcd.JobName}}_resourcesPool" class="col-sm-3 control-label">ResourcesPool</label>
				    <div class="col-sm-7">
				      	<input type="hidden" id="{{.etcd.JobName}}_resourcesPool" name="{{.etcd.JobName}}_resourcesPool" value = "{{.etcd.ResourcesPool}}">
				    	<select class="form-control" id="{{.etcd.JobName}}_resourcesPool" name="{{.etcd.JobName}}_resourcesPool_select" >
						{{$etcd := .etcd.JobName}}
						{{range $.VsphereResourcesPools}}<option id="{{$etcd}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
						</select>
						<script type="text/javascript">
						$("#{{.etcd.JobName}}_{{.etcd.ResourcesPool}}").attr("selected",true);  
						</script>
					</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.etcd.JobName}}_instances" class="col-sm-3 control-label">Instances</label>
				    <div class="col-sm-7">
				      	<input type="number" class="form-control" id="{{.etcd.JobName}}_instances" placeholder="Instances" name="{{.etcd.JobName}}_instances" value = "{{.etcd.Instances}}" required readonly>
				    </div>
			  	</div>
			</div>
		</div>
	</div>
	{{end}}