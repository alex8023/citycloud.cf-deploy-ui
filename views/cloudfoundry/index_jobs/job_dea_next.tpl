{{with .CloudFoundry.CloudFoundryJobs.dea_next}}
 <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">{{i18n $.Lang "Jobs"}}-{{.JobName}}</h2>
		</div>
		<div class="form-horizontal">
	  		<div class="panel-body">
				<div class="form-group">
					<label class="col-sm-3 control-label">{{i18n $.Lang "Name"}}</label>
					<div class="col-sm-5">
						<p class="form-control-static">{{.Name}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-3 control-label">{{i18n $.Lang "ResourcesPool"}}</label>
					<div class="col-sm-5">
						<p class="form-control-static">{{.ResourcesPool}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-3 control-label">{{i18n $.Lang "Instances"}}</label>
					<div class="col-sm-5">
						<p class="form-control-static">{{.Instances}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-3 control-label">{{i18n $.Lang "StaticIp"}}</label>
					<div class="col-sm-5">
						<p class="form-control-static">{{.StaticIp}}</p>
					</div>
				</div>
			</div>
		</div>
	</div>
{{end}}