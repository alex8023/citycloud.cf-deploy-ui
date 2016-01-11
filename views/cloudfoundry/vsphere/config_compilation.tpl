<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">PaaS Compilation</h2>
		</div>
  		<div class="panel-body">
			<form class="form-horizontal" method="post" action="vspherecloudfoundry">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      <button type="submit" class="btn btn-default " data-loading-text="Saving..."><span class="glyphicon glyphicon-floppy-save"></span> Save</button>
					<a class="btn btn-default " href="vspherecloudfoundry"><span class="glyphicon glyphicon-step-backward"></span> Back</a>
				    </div>
			  	</div>
				{{with .CloudFoundry}}
				{{with .VsphereCompilation}}
				<input type = "hidden" name = "id" value = "{{.Id}}">
			  	<div class="form-group">
			    	<label for="workers" class="col-sm-3 control-label">Workers</label>
				    <div class="col-sm-7">
				      <input type="number" class="form-control" id="workers" placeholder="Workers" name="workers" value = "{{.Workers}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="defaultNetWork" class="col-sm-3 control-label">DefaultNetWork</label>
				    <div class="col-sm-7">
				      <input type="text" class="form-control" id="defaultNetWork" placeholder="DefaultNetWork" name="defaultNetWork" value = "{{.DefaultNetWork}}" readonly required>
				    </div>
			  	</div>
				<div class="form-group">
			    	<label for="" class="col-sm-3 control-label">FlavorPool</label>
				    <div class="col-sm-7">
				    	<select class="form-control" id="compilation_flavor_pool" name="vid" >
						{{range $.VsphereResource}}<option id="compilation_flavor_pool_{{.Id}}" value="{{.Id}}">Ram({{.Ram}}M)-CPU({{.Cpu}})-Disk({{.Disk}}M)</option>{{end}}
						</select>
						<script type="text/javascript">
						$("#compilation_flavor_pool_{{.Vid}}").attr("selected",true);  
						</script>
					</div>
			  	</div>
				{{end}}
				{{end}}
				<input type = "hidden" name="model" value = "{{.Model}}">
			</form>
  		</div>
	</div>
</div>
