<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">MicroBOSH Deployment</h2>
		</div>
  		<div class="panel-body">
			<form class="form-horizontal" method="post" action="microbosh" enctype="multipart/form-data">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      <button type="submit" class="btn btn-default " data-loading-text="Saving...">Save</button>
					<a class="btn btn-default " href="microbosh">Back</a>
				    </div>
			  	</div>
				{{with .MicroBOSH}}
			  	<div class="form-group">
			    	<label for="name" class="col-sm-2 control-label">Deployment Name</label>
				    <div class="col-sm-10">
						<input type="hidden" name="id" value = "{{.Id}}">
				      	<input type="text" class="form-control" id="name" placeholder="Deployment Name" name="name" value = "{{.Name}}" required>
				    </div>
			  	</div>
				{{with .NetWork}}
			  	<div class="form-group">
			    	<label for="vip" class="col-sm-2 control-label">Vip</label>
				    <div class="col-sm-10">
						<input type="hidden" name="networkId" value = "{{.Id}}">
				      	<input type="text" class="form-control" id="vip" placeholder="Vip" name="vip" value = "{{.Vip}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="net_id" class="col-sm-2 control-label">NetId</label>
			    	<div class="col-sm-10">
			      		<input type="text" class="form-control" id="net_id" placeholder="NetId" name="net_id" value = "{{.NetId}}" required>
			    	</div>
			  	</div>
				{{end}}
				{{with .Resources}}
			  	<div class="form-group">
			    	<label for="persistent_disk" class="col-sm-2 control-label">PersistentDiskSize</label>
				    <div class="col-sm-10">
						<input type="hidden" name="resourcesId" value = "{{.Id}}">
				      	<input type="text" class="form-control" id="persistent_disk" placeholder="16384" name="persistent_disk" value = "{{.PersistentDisk}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="instance_type" class="col-sm-2 control-label">InstanceFlavor</label>
			    	<div class="col-sm-10">
			      		<input type="text" class="form-control" id="instance_type" placeholder="InstanceFlavor" name="instance_type" value = "{{.InstanceType}}" required>
			    	</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="availability_zone" class="col-sm-2 control-label">AvailabilityZone</label>
			    	<div class="col-sm-10">
			      	<input type="text" class="form-control" id="availability_zone" placeholder="AvailabilityZone" name="availability_zone" value = "{{.AvailabilityZone}}" required>
			    	</div>
			  	</div>
				{{end}}
				{{with .CloudProperties}}
			  	<div class="form-group">
			    	<label for="auth_url" class="col-sm-2 control-label">AuthUrl</label>
			    	<div class="col-sm-10">
					<input type="hidden" name="cloudPropertiesId" value = "{{.Id}}">
			      	<input type="text" class="form-control" id="auth_url" placeholder="AuthUrl" name="auth_url" value = "{{.AuthUrl}}" required>
			    	</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="username" class="col-sm-2 control-label">UserName</label>
			    	<div class="col-sm-10">
			      		<input type="text" class="form-control" id="username" placeholder=" UserName" name="username" value = "{{.UserName}}" required>
			    	</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="api_key" class="col-sm-2 control-label">ApiKey</label>
			    	<div class="col-sm-10">
			      		<input type="text" class="form-control" id="api_key" placeholder="ApiKey" name="api_key" value = "{{.ApiKey}}" required>
			    	</div>
			  	</div>
				<div class="form-group">
				    <label for="tenant" class="col-sm-2 control-label">Tenant</label>
				    <div class="col-sm-10">
				      	<input type="text" class="form-control" id="tenant" placeholder="Tenant" name="tenant" value = "{{.Tenant}}" required>
				    </div>
				</div>
			  	<div class="form-group">
			    	<label for="default_key_name" class="col-sm-2 control-label">SecretKey</label>
			    	<div class="col-sm-10">
			      		<input type="text" class="form-control" id="default_key_name" placeholder="SecretKey" name="default_key_name" value = "{{.DefaultKeyName}}" required>
			    	</div>
			  	</div>
				<div class="form-group">
				    <label for="private_key"  class="col-sm-2 control-label">PrivateKey</label>
					<div class="col-sm-10">
				    	<input type="file" id="private_key" name = "private_key">
				    	<p class="help-block">Please upload your PrivateKey file about SecretKey</p>
					</div>
				</div>
					{{if ne $.IaaSVersion $.DefaultVersion}}
					  	<div class="form-group">
						    	<label for="cci_ebs_url" class="col-sm-2 control-label">NBSAuthUrl1</label>
						    	<div class="col-sm-10">
						      		<input type="text" class="form-control" id="cci_ebs_url" placeholder="NBSAuthUrl" name="cci_ebs_url" value = "{{.CciEbsUrl}}" required>
						    	</div>
					  	</div>
			
					  	<div class="form-group">
						    	<label for="cci_ebs_accesskey" class="col-sm-2 control-label">NBSAccesskey</label>
						    	<div class="col-sm-10">
						      		<input type="text" class="form-control" id="cci_ebs_accesskey" placeholder="NBSAccesskey" name="cci_ebs_accesskey" value = "{{.CciEbsAccesskey}}" required>
						    	</div>
					  	</div>
		
					  	<div class="form-group">
						    	<label for="cci_ebs_secretkey" class="col-sm-2 control-label">NBSSecretkey</label>
						    	<div class="col-sm-10">
						      		<input type="text" class="form-control" id="cci_ebs_secretkey" placeholder="NBSSecretkey" name="cci_ebs_secretkey" value = "{{.CciEbsSecretkey}}" required>
						    	</div>
					  	</div>
					{{end}}	
					{{end}}	
				{{end}}
			</form>
  		</div>
	</div>
</div>
