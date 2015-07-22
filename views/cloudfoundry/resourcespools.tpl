<input type="hidden" id="navfocus" value = "{{.NavBarFocus}}">
<div style="height:20px"></div>
<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">PaaS ResourcesPools</h2>
		</div>
  		<div class="panel-body">
			<form class="form-horizontal" method="post" action="cloudfoundry">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      	<button type="submit" class="btn btn-default " data-loading-text="Saving...">Save</button>
						<button type="button" class="btn btn-default " id="add_resources_pool" >More>></button>
				    </div>
			  	</div>
				{{with .CloudFoundry}}{{$index := 0}}
				{{range .ResourcesPools}}{{$index := $index + 1}}
			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">ResourcesPools</h2>
					</div>
					<div class="form-horizontal">
					  	<div class="panel-body">
						  	<div class="form-group">
						    	<label for="name" class="col-sm-2 control-label">Name@@{{$index}}</label>
							    <div class="col-sm-10">
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
							      	<input type="text" class="form-control" id="availabilityZone" placeholder="AvailabilityZone" name="availabilityZone" value = "{{.AvailabilityZone}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="size" class="col-sm-2 control-label">Size</label>
							    <div class="col-sm-10">
							      	<input type="text" class="form-control" id="size" placeholder="Size" name="size" value = "{{.Size}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="defaultNetWork" class="col-sm-2 control-label">DefaultNetWork</label>
							    <div class="col-sm-10">
							      	<input type="text" class="form-control" id="defaultNetWork" placeholder="DefaultNetWork" name="defaultNetWork" value = "{{.DefaultNetWork}}" required>
							    </div>
						  	</div>
						</div>
					</div>
				</div>
				{{end}}
				{{end}}
				<input type = "hidden" name="model" value = "{{.Model}}">
			</form>
  		</div>
	</div>
</div>
<script src="/static/js/resourcespool.js"></script>
