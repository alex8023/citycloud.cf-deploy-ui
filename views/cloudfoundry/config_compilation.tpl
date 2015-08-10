<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">PaaS Compilation</h2>
		</div>
  		<div class="panel-body">
			<form class="form-horizontal" method="post" action="cloudfoundry">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      <button type="submit" class="btn btn-default " data-loading-text="Saving...">Save</button>
				    </div>
			  	</div>
				{{with .CloudFoundry}}
				{{with .Compilation}}
			  	<div class="form-group">
			    	<label for="instanceType" class="col-sm-2 control-label">InstanceType</label>
				    <div class="col-sm-10">
						<input type="hidden" name="id" value="{{.Id}}">
				      <input type="text" class="form-control" id="instanceType" placeholder="InstanceType" name="instanceType" value = "{{.InstanceType}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="availabilityZone" class="col-sm-2 control-label">AvailabilityZone</label>
				    <div class="col-sm-10">
				      <input type="text" class="form-control" id="availabilityZone" placeholder="AvailabilityZone" name="availabilityZone" value = "{{.AvailabilityZone}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="workers" class="col-sm-2 control-label">Workers</label>
				    <div class="col-sm-10">
				      <input type="number" class="form-control" id="workers" placeholder="Workers" name="workers" value = "{{.Workers}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="defaultNetWork" class="col-sm-2 control-label">DefaultNetWork</label>
				    <div class="col-sm-10">
				      <input type="text" class="form-control" id="defaultNetWork" placeholder="DefaultNetWork" name="defaultNetWork" value = "{{.DefaultNetWork}}" readonly required>
				    </div>
			  	</div>
				{{end}}
				{{end}}
				<input type = "hidden" name="model" value = "{{.Model}}">
			</form>
  		</div>
	</div>
</div>
