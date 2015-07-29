<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">PaaS PaaSProperties</h2>
		</div>
  		<div class="panel-body">
			<form class="form-horizontal" method="post" action="cloudfoundry">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      <button type="submit" class="btn btn-default " data-loading-text="Saving...">Save</button>
				    </div>
			  	</div>
				{{with .CloudFoundry}}
				{{with .CloudFoundryProperties}}
			  	<div class="form-group">
			    	<label for="name" class="col-sm-2 control-label">Deployment Name</label>
				    <div class="col-sm-10">
				      <input type="text" class="form-control" id="name" placeholder="Deployment Name" name="name" value = "{{.Name}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="uuid" class="col-sm-2 control-label">Director UUID</label>
				    <div class="col-sm-10">
				      <input type="text" class="form-control" id="uuid" placeholder="Director UUID" name="uuid" value = "{{.Uuid}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="floatingIp" class="col-sm-2 control-label">FloatingIp</label>
				    <div class="col-sm-10">
				      <input type="text" class="form-control" id="floatingIp" placeholder="FloatingIp" name="floatingIp" value = "{{.FloatingIp}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="systemDomain" class="col-sm-2 control-label">SystemDomain</label>
				    <div class="col-sm-10">
				      <input type="text" class="form-control" id="systemDomain" placeholder="SystemDomain" name="systemDomain" value = "{{.SystemDomain}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="systemDomainOrg" class="col-sm-2 control-label">SystemOrg</label>
				    <div class="col-sm-10">
				      <input type="text" class="form-control" id="systemDomainOrg" placeholder="SystemOrg" name="systemDomainOrg" value = "{{.SystemDomainOrg}}" required>
				    </div>
			  	</div>
				{{end}}
				{{end}}
				<input type = "hidden" name="model" value = "{{.Model}}">
			</form>
  		</div>
	</div>
</div>