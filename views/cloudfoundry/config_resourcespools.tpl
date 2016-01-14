<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">PaaS {{i18n $.Lang "ResourcesPools"}}</h2>
		</div>
  		<div class="panel-body">
			<form class="form-horizontal" method="post" action="cloudfoundry" id="resourceform">
				<input type = "hidden" name="model" value = "{{.Model}}">
				<input type = "hidden" id="poollength" value = "{{.Pools}}">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      	<button type="submit" class="btn btn-default " data-loading-text="Saving..." ><span class="glyphicon glyphicon-floppy-save"></span> {{i18n $.Lang "Save"}}</button>
						<button type="button" class="btn btn-default " id="add_resources_pool" ><span class="glyphicon glyphicon-plus"></span> {{i18n $.Lang "More"}}>></button>
						<a class="btn btn-default " href="cloudfoundry">Back</a>
				    </div>
			  	</div>
				{{with .CloudFoundry}}
				{{range $index,$element := .ResourcesPools}}
			    <div class="panel panel-default" id="deletePanel{{$index}}">
					<div class="panel-heading" >
						<h2 class="panel-title">{{i18n $.Lang "ResourcesPool"}}<button name="delete" class="btn btn-default" type="button" onclick="deletePanel('deletePanel{{$index}}')">{{i18n $.Lang "Delete"}}</button></h2>
					</div>
					<div class="form-horizontal">
					  	<div class="panel-body">
						  	<div class="form-group">
						    	<label for="name" class="col-sm-3 control-label">{{i18n $.Lang "Name"}}</label>
							    <div class="col-sm-7">
									<input type="hidden" name ="id" value ="{{.Id}}">
							      	<input type="text" class="form-control" id="name" placeholder="{{i18n $.Lang "Name"}}" name="name" value = "{{.Name}}" required>
							    </div>
						  	</div>
				
						  	<div class="form-group">
						    	<label for="instanceType" class="col-sm-3 control-label">{{i18n $.Lang "InstanceType"}}</label>
							    <div class="col-sm-7">
							      	<input type="text" class="form-control" id="instanceType" placeholder="{{i18n $.Lang "InstanceType"}}" name="instanceType" value = "{{.InstanceType}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="availabilityZone" class="col-sm-3 control-label">{{i18n $.Lang "AvailabilityZone"}}</label>
							    <div class="col-sm-7">
							      	<input type="text" class="form-control" id="availabilityZone" placeholder="{{i18n $.Lang "AvailabilityZone"}}" name="availabilityZone" value = "{{.AvailabilityZone}}" >
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="size" class="col-sm-3 control-label">{{i18n $.Lang "Size"}}</label>
							    <div class="col-sm-7">
							      	<input type="number" class="form-control" id="size" placeholder="{{i18n $.Lang "Size"}}" name="size" value = "{{.Size}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="defaultNetWork" class="col-sm-3 control-label">{{i18n $.Lang "DefaultNetWork"}}</label>
							    <div class="col-sm-7">
							      	<input type="text" class="form-control" id="defaultNetWork" placeholder="{{i18n $.Lang "DefaultNetWork"}}" name="defaultNetWork" value = "{{.DefaultNetWork}}" required readonly>
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
