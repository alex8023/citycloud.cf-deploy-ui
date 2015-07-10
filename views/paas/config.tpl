<input type="hidden" id="navfocus" value = "{{.NavBarFocus}}">
<div style="height:20px"></div>
<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">MicroBOSH Deployment</h2>
		</div>
  		<div class="panel-body">
			<form class="form-horizontal" method="post" action="microbosh">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      <button type="submit" class="btn btn-default ">Save</button>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="name" class="col-sm-2 control-label">Deployment Name</label>
				    <div class="col-sm-10">
				      <input type="text" class="form-control" id="name" placeholder="Deployment Name" name="name">
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="vip" class="col-sm-2 control-label">Vip</label>
				    <div class="col-sm-10">
				      	<input type="text" class="form-control" id="vip" placeholder="vip" name="vip">
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="net_id" class="col-sm-2 control-label">NetId</label>
			    	<div class="col-sm-10">
			      		<input type="text" class="form-control" id="net_id" placeholder="NetId" name="net_id">
			    	</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="persistent_disk" class="col-sm-2 control-label">PersistentDiskSize</label>
				    <div class="col-sm-10">
				      	<input type="text" class="form-control" id="persistent_disk" placeholder="16384" name="persistent_disk">
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="instance_type" class="col-sm-2 control-label">InstanceFlavor</label>
			    	<div class="col-sm-10">
			      		<input type="text" class="form-control" id="instance_type" placeholder="InstanceFlavor" name="instance_type">
			    	</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="availability_zone" class="col-sm-2 control-label">AvailabilityZone</label>
			    	<div class="col-sm-10">
			      		<input type="text" class="form-control" id="availability_zone" placeholder="AvailabilityZone" name="availability_zone">
			    	</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="auth_url" class="col-sm-2 control-label">OpenStackAuthUrl</label>
			    	<div class="col-sm-10">
			      		<input type="text" class="form-control" id="auth_url" placeholder="AuthUrl" name="auth_url">
			    	</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="username" class="col-sm-2 control-label">OpenStackUserName</label>
			    	<div class="col-sm-10">
			      		<input type="text" class="form-control" id="username" placeholder="OpenStack username" name="username">
			    	</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="api_key" class="col-sm-2 control-label">OpenStackApiKey</label>
			    	<div class="col-sm-10">
			      		<input type="text" class="form-control" id="api_key" placeholder="api_key" name="api_key">
			    	</div>
			  	</div>
				<div class="form-group">
				    <label for="tenant" class="col-sm-2 control-label">OpenStackTenant</label>
				    <div class="col-sm-10">
				      	<input type="text" class="form-control" id="tenant" placeholder="Tenant" name="tenant">
				    </div>
				</div>
			  	<div class="form-group">
			    	<label for="default_key_name" class="col-sm-2 control-label">OpenStackSecretKey</label>
			    	<div class="col-sm-10">
			      		<input type="text" class="form-control" id="default_key_name" placeholder="SecretKey" name="default_key_name">
			    	</div>
			  	</div>
				<div class="form-group">
				    <label for="private_key"  class="col-sm-2 control-label">OpenStackPrivateKey</label>
					<div class="col-sm-10">
				    	<input type="file" id="private_key">
				    	<p class="help-block">Please upload your PrivateKey file about SecretKey</p>
					</div>
				</div>

			  	<div class="form-group">
			    	<label for="cci_ebs_url" class="col-sm-2 control-label">NBSAuthUrl</label>
			    	<div class="col-sm-10">
			      		<input type="text" class="form-control" id="cci_ebs_url" placeholder="NBS AuthUrl" name="cci_ebs_url">
			    	</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="cci_ebs_accesskey" class="col-sm-2 control-label">NBSAccesskey</label>
			    	<div class="col-sm-10">
			      		<input type="text" class="form-control" id="cci_ebs_accesskey" placeholder="NBS Accesskey" name="cci_ebs_accesskey">
			    	</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="cci_ebs_secretkey" class="col-sm-2 control-label">NBSSecretkey</label>
			    	<div class="col-sm-10">
			      		<input type="text" class="form-control" id="cci_ebs_secretkey" placeholder="NBS Secretkey" name="cci_ebs_secretkey">
			    	</div>
			  	</div>
			</form>
  		</div>
	</div>
</div>
