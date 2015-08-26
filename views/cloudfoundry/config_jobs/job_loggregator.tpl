    {{with .CloudFoundryJobs}}
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">Jobs-{{.loggregator.JobName}}</h2>
		</div>
		<div class="form-horizontal">
		  	<div class="panel-body">
			  	<div class="form-group">
			    	<label for="{{.loggregator.JobName}}_name" class="col-sm-3 control-label">Name</label>
				    <div class="col-sm-7">
				      	<input type="text" class="form-control" id="{{.loggregator.JobName}}_name" placeholder="Name" name="{{.loggregator.JobName}}_name" value = "{{.loggregator.Name}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.loggregator.JobName}}_resourcesPool" class="col-sm-3 control-label">ResourcesPool</label>
				    <div class="col-sm-7">
				      	<input type="hidden" id="{{.loggregator.JobName}}_resourcesPool" name="{{.loggregator.JobName}}_resourcesPool" value = "{{.loggregator.ResourcesPool}}">
				    	<select class="form-control" id="{{.loggregator.JobName}}_resourcesPool" name="{{.loggregator.JobName}}_resourcesPool_select" >
						{{$loggregator := .loggregator.JobName}}
						{{range $.ResourcesPools}}<option id="{{$loggregator}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
						</select>
						<script type="text/javascript">
						$("#{{.loggregator.JobName}}_{{.loggregator.ResourcesPool}}").attr("selected",true);  
						</script>
					</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.loggregator.JobName}}_instances" class="col-sm-3 control-label">Instances</label>
				    <div class="col-sm-7">
				      	<input type="number" class="form-control" id="{{.loggregator.JobName}}_instances" placeholder="Instances" name="{{.loggregator.JobName}}_instances" value = "{{.loggregator.Instances}}" required readonly>
				    </div>
			  	</div>
			</div>
		</div>
	</div>
	{{end}}
				