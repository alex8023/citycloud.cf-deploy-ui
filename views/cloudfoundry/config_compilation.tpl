<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">PaaS {{i18n $.Lang "Compilation"}}</h2>
		</div>
  		<div class="panel-body">
			<form class="form-horizontal" method="post" action="cloudfoundry">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      <button type="submit" class="btn btn-default " data-loading-text="Saving..."><span class="glyphicon glyphicon-floppy-save"></span> {{i18n $.Lang "Save"}}</button>
					<a class="btn btn-default " href="cloudfoundry"><span class="glyphicon glyphicon-step-backward"></span> {{i18n $.Lang "Back"}}</a>
				    </div>
			  	</div>
				{{with .CloudFoundry}}
				{{with .Compilation}}
			  	<div class="form-group">
			    	<label for="instanceType" class="col-sm-3 control-label">{{i18n $.Lang "InstanceType"}}</label>
				    <div class="col-sm-7">
						<input type="hidden" name="id" value="{{.Id}}">
				      <input type="text" class="form-control" id="instanceType" placeholder="{{i18n $.Lang "InstanceType"}}" name="instanceType" value = "{{.InstanceType}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="availabilityZone" class="col-sm-3 control-label">{{i18n $.Lang "AvailabilityZone"}}</label>
				    <div class="col-sm-7">
				      <input type="text" class="form-control" id="availabilityZone" placeholder="{{i18n $.Lang "AvailabilityZone"}}" name="availabilityZone" value = "{{.AvailabilityZone}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="workers" class="col-sm-3 control-label">{{i18n $.Lang "Workers"}}</label>
				    <div class="col-sm-7">
				      <input type="number" class="form-control" id="workers" placeholder="{{i18n $.Lang "Workers"}}" name="workers" value = "{{.Workers}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="defaultNetWork" class="col-sm-3 control-label">{{i18n $.Lang "DefaultNetWork"}}</label>
				    <div class="col-sm-7">
				      <input type="text" class="form-control" id="defaultNetWork" placeholder="{{i18n $.Lang "DefaultNetWork"}}" name="defaultNetWork" value = "{{.DefaultNetWork}}" readonly required>
				    </div>
			  	</div>
				{{end}}
				{{end}}
				<input type = "hidden" name="model" value = "{{.Model}}">
			</form>
  		</div>
	</div>
</div>
