<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">PaaS {{i18n $.Lang "ResourcesPools"}}</h2>
		</div>
  		<div class="panel-body">
			<form class="form-horizontal" method="post" action="vspherecloudfoundry" id="resourceform">
				<input type = "hidden" name="model" value = "{{.Model}}">
				<input type = "hidden" id="poollength" value = "{{.Pools}}">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      	<button type="submit" class="btn btn-default " data-loading-text="Saving..." ><span class="glyphicon glyphicon-floppy-save"></span> {{i18n $.Lang "Save"}}</button>
						<button type="button" class="btn btn-default " id="add_resources_pool" ><span class="glyphicon glyphicon-plus"></span> {{i18n $.Lang "More"}}>></button>
						<a class="btn btn-default " href="vspherecloudfoundry"> {{i18n $.Lang "Back"}}</a>
				    </div>
			  	</div>
				{{with .CloudFoundry}}
				{{range $index,$element := .VsphereResourcesPools}}
			    <div class="panel panel-default" id="deletePanel{{$index}}">
					<div class="panel-heading" >
						<h2 class="panel-title">ResourcesPools<button name="delete" class="btn btn-default" type="button" onclick="deletePanel('deletePanel{{$index}}')">Delete</button></h2>
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
							
							<div class="form-group">
						    	<label for="" class="col-sm-3 control-label">{{i18n $.Lang "Instance-Flavor"}}</label>
							    <div class="col-sm-7">
							    	<select class="form-control" id="compilation_flavor_pool" name="vid" >
									{{$vsphereResourcePoolId := .Id}}
									{{range $.VsphereResource}}<option id="compilation_flavor_pool_{{$vsphereResourcePoolId}}_{{.Id}}" value="{{.Id}}">Ram({{.Ram}}M)-CPU({{.Cpu}})-Disk({{.Disk}}M)</option>{{end}}
									</select>
									<script type="text/javascript">
									$("#compilation_flavor_pool_{{.Id}}_{{.Vid}}").attr("selected",true);  
									</script>
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
