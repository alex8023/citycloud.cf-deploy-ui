	<div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">Jobs-{{.JobName}}</h2>
		</div>
		<div class="form-horizontal">
		  	<div class="panel-body">
			  	<div class="form-group">
			    	<label for="{{.JobName}}_name" class="col-sm-2 control-label">Name</label>
				    <div class="col-sm-10">
				      	<input type="text" class="form-control" id="{{.JobName}}_name" placeholder="Name" name="{{.JobName}}_name" value = "{{.Name}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.JobName}}_resourcesPool" class="col-sm-2 control-label">ResourcesPool</label>
				    <div class="col-sm-10">
				      	<input type="hidden" id="{{.JobName}}_resourcesPool" name="{{.JobName}}_resourcesPool" value = "{{.ResourcesPool}}">
				    	<select class="form-control" id="{{.JobName}}_resourcesPool" name="{{.JobName}}_resourcesPool_select" >
						{{$JobName := .JobName}}
						{{range $.ResourcesPools}}<option id="{{$JobName}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
						</select>
						<script type="text/javascript">
						if ($("#{{.JobName}}_{{.ResourcesPool}}")){
							$("#{{.JobName}}_{{.ResourcesPool}}").attr("selected",true);  
						}
						</script>
					</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.JobName}}_instances" class="col-sm-2 control-label">Instances</label>
				    <div class="col-sm-10">
				      	<input type="number" class="form-control" id="{{.JobName}}_instances" placeholder="Instances" name="{{.JobName}}_instances" value = "{{.Instances}}" required readonly>
				    </div>
			  	</div>
			</div>
		</div>
	</div>