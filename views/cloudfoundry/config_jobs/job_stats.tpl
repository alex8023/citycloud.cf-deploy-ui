    {{with .CloudFoundry}}
	{{with .CloudFoundryJobs}}
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">{{i18n $.Lang "Jobs"}}-{{.stats.JobName}}</h2>
		</div>
		<div class="form-horizontal">
		  	<div class="panel-body">
			  	<div class="form-group">
			    	<label for="{{.stats.JobName}}_name" class="col-sm-3 control-label">{{i18n $.Lang "Name"}}</label>
				    <div class="col-sm-7">
				      	<input type="text" class="form-control" id="{{.stats.JobName}}_name" placeholder="{{i18n $.Lang "Name"}}" name="{{.stats.JobName}}_name" value = "{{.stats.Name}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.stats.JobName}}_resourcesPool" class="col-sm-3 control-label">{{i18n $.Lang "ResourcesPool"}}</label>
				    <div class="col-sm-7">
				      	<input type="hidden" id="{{.stats.JobName}}_resourcesPool" name="{{.stats.JobName}}_resourcesPool" value = "{{.stats.ResourcesPool}}">
				    	<select class="form-control" id="{{.stats.JobName}}_resourcesPool" name="{{.stats.JobName}}_resourcesPool_select" >
						{{$stats := .stats.JobName}}
						{{range $.CloudFoundry.ResourcesPools}}<option id="{{$stats}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
						</select>
						<script type="text/javascript">
						$("#{{.stats.JobName}}_{{.stats.ResourcesPool}}").attr("selected",true);  
						</script>
					</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="{{.stats.JobName}}_instances" class="col-sm-3 control-label">{{i18n $.Lang "Instances"}}</label>
				    <div class="col-sm-7">
				      	<input type="number" class="form-control" id="{{.stats.JobName}}_instances" placeholder="{{i18n $.Lang "Instances"}}" name="{{.stats.JobName}}_instances" value = "{{.stats.Instances}}" required readonly>
				    </div>
			  	</div>
			</div>
		</div>
	</div>
	{{end}}{{end}}