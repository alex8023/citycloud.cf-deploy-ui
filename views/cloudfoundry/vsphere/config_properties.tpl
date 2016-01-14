<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">PaaS {{i18n $.Lang "Properties"}}</h2>
		</div>
  		<div class="panel-body">
			<form class="form-horizontal" method="post" action="vspherecloudfoundry">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      <button type="submit" class="btn btn-default " data-loading-text="Saving..."><span class="glyphicon glyphicon-floppy-save"></span> {{i18n $.Lang "Save"}}</button>
					<a class="btn btn-default " href="vspherecloudfoundry"><span class="glyphicon glyphicon-step-backward"></span> {{i18n $.Lang "Back"}}</a>
				    </div>
			  	</div>
				{{with .CloudFoundry}}
				{{with .CloudFoundryProperties}}
			  	<div class="form-group">
			    	<label for="name" class="col-sm-3 control-label">{{i18n $.Lang "Deployment Name"}}</label>
				    <div class="col-sm-7">
						<input type="hidden" name="id" value="{{.Id}}">
				      	<input type="text" class="form-control" id="name" placeholder="{{i18n $.Lang "Deployment Name"}}" name="name" value = "{{.Name}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="uuid" class="col-sm-3 control-label">{{i18n $.Lang "Director UUID"}}</label>
				    <div class="col-sm-7">
				      <input type="text" class="form-control" id="uuid" placeholder="{{i18n $.Lang "Director UUID"}}" name="uuid" value = "{{.Uuid}}" required>
				    </div>
			  	</div>
				<div class="form-group">
			    	<label for="systemDomain" class="col-sm-3 control-label">{{i18n $.Lang "SystemDomain"}}</label>
				    <div class="col-sm-7">
				      <input type="text" class="form-control" id="systemDomain" placeholder="{{i18n $.Lang "SystemDomain"}}" name="systemDomain" value = "{{.SystemDomain}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="systemDomainOrg" class="col-sm-3 control-label">{{i18n $.Lang "SystemOrg"}}</label>
				    <div class="col-sm-7">
				      <input type="text" class="form-control" id="systemDomainOrg" placeholder="{{i18n $.Lang "SystemOrg"}}" name="systemDomainOrg" value = "{{.SystemDomainOrg}}" required>
				    </div>
			  	</div>
				{{end}}{{end}}
				<input type = "hidden" name="model" value = "{{.Model}}">
			</form>
  		</div>
	</div>
</div>
