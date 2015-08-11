<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">PaaS ResourcesPools</h2>
		</div>
  		<div class="panel-body">
			<form class="form-horizontal" method="post" action="cloudfoundry" id="resourceform">
				<input type = "hidden" name="model" value = "{{.Model}}">
				<input type = "hidden" id="poollength" value = "{{.Pools}}">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      	<button type="submit" class="btn btn-default " data-loading-text="Saving..." >Save</button>
						<button type="button" class="btn btn-default " id="add_resources_pool" >More>></button>
				    </div>
			  	</div>
				{{with .CloudFoundry}}
				{{range $index,$element := .ResourcesPools}}
			    <div class="panel panel-default" id="deletePanel{{$index}}">
					<div class="panel-heading" >
						<h2 class="panel-title">ResourcesPools<button name="delete" class="btn btn-default" type="button" onclick="deletePanel('deletePanel{{$index}}')">Delete</button></h2>
					</div>
					<div class="form-horizontal">
					  	<div class="panel-body">
						  	<div class="form-group">
						    	<label for="name" class="col-sm-2 control-label">Name</label>
							    <div class="col-sm-10">
									<input type="hidden" name ="id" value ="{{.Id}}">
							      	<input type="text" class="form-control" id="name" placeholder="Name" name="name" value = "{{.Name}}" required>
							    </div>
						  	</div>
				
						  	<div class="form-group">
						    	<label for="instanceType" class="col-sm-2 control-label">InstanceType</label>
							    <div class="col-sm-10">
							      	<input type="text" class="form-control" id="instanceType" placeholder="InstanceType" name="instanceType" value = "{{.InstanceType}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="availabilityZone" class="col-sm-2 control-label">AvailabilityZone</label>
							    <div class="col-sm-10">
							      	<input type="text" class="form-control" id="availabilityZone" placeholder="AvailabilityZone" name="availabilityZone" value = "{{.AvailabilityZone}}" >
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="size" class="col-sm-2 control-label">Size</label>
							    <div class="col-sm-10">
							      	<input type="number" class="form-control" id="size" placeholder="Size" name="size" value = "{{.Size}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="defaultNetWork" class="col-sm-2 control-label">DefaultNetWork</label>
							    <div class="col-sm-10">
							      	<input type="text" class="form-control" id="defaultNetWork" placeholder="DefaultNetWork" name="defaultNetWork" value = "{{.DefaultNetWork}}" required readonly>
							    </div>
						  	</div>
						</div>
					</div>
				</div>
				{{end}}
				{{end}}
			</form>
  		</div>
	</div>
</div>
<script src="/static/js/resourcespool.js"></script>
