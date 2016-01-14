    {{with .CloudFoundry}}
    {{with .CloudFoundryJobs}}
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">{{i18n $.Lang "Jobs"}}-{{.nats.JobName}}</h2>
		</div>
		<div class="form-horizontal">
		  	<div class="panel-body">
			  	<div class="form-group">
			    	<label for="{{.nats.JobName}}_name" class="col-sm-3 control-label">{{i18n $.Lang "Name"}}</label>
				    <div class="col-sm-7">
				      	<input type="text" class="form-control" id="{{.nats.JobName}}_name" placeholder="{{i18n $.Lang "Name"}}" name="{{.nats.JobName}}_name" value = "{{.nats.Name}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.nats.JobName}}_resourcesPool" class="col-sm-3 control-label">{{i18n $.Lang "ResourcesPool"}}</label>
				    <div class="col-sm-7">
				      	<input type="hidden" id="{{.nats.JobName}}_resourcesPool" name="{{.nats.JobName}}_resourcesPool" value = "{{.nats.ResourcesPool}}">
				    	<select class="form-control" id="{{.nats.JobName}}_resourcesPool" name="{{.nats.JobName}}_resourcesPool_select" >
						{{$nats := .nats.JobName}}
						{{range $.CloudFoundry.ResourcesPools}}<option id="{{$nats}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
						</select>
						<script type="text/javascript">
						$("#{{.nats.JobName}}_{{.nats.ResourcesPool}}").attr("selected",true);  
						</script>
					</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.nats.JobName}}_instances" class="col-sm-3 control-label">{{i18n $.Lang "Instances"}}</label>
				    <div class="col-sm-7">
				      	<input type="number" class="form-control" id="{{.nats.JobName}}_instances" placeholder="{{i18n $.Lang "Instances"}}" name="{{.nats.JobName}}_instances" value = "{{.nats.Instances}}" required readonly>
				    </div>
			  	</div>
			</div>
		</div>
	</div>
	{{end}}{{end}}