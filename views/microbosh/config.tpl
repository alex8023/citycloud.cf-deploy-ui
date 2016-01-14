<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">MicroBOSH {{i18n $.Lang "Deployment"}}</h2>
		</div>
  		<div class="panel-body">
			<form class="form-horizontal" method="post" action="microbosh" enctype="multipart/form-data">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      <button type="submit" class="btn btn-default " data-loading-text="Saving..."><span class="glyphicon glyphicon-floppy-save"></span> {{i18n $.Lang "Save"}}</button>
					<a class="btn btn-default " href="microbosh"><span class="glyphicon glyphicon-step-backward"></span> {{i18n $.Lang "Back"}}</a>
				    </div>
			  	</div>
				{{with .MicroBOSH}}
			  	<div class="form-group">
			    	<label for="name" class="col-sm-3 control-label">{{i18n $.Lang "Deployment Name"}}</label>
				    <div class="col-sm-7">
						<input type="hidden" name="id" value = "{{.Id}}">
				      	<input type="text" class="form-control" id="name" placeholder="{{i18n $.Lang "Deployment Name"}}" name="name" value = "{{.Name}}" required>
				    </div>
			  	</div>
				{{with .NetWork}}
			  	<div class="form-group">
			    	<label for="vip" class="col-sm-3 control-label">{{i18n $.Lang "Vip"}}</label>
				    <div class="col-sm-7">
						<input type="hidden" name="networkId" value = "{{.Id}}">
				      	<input type="text" class="form-control" id="vip" placeholder="{{i18n $.Lang "Vip"}}" name="vip" value = "{{.Vip}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="net_id" class="col-sm-3 control-label">{{i18n $.Lang "NetId"}}</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="net_id" placeholder="{{i18n $.Lang "NetId"}}" name="net_id" value = "{{.NetId}}" required>
			    	</div>
			  	</div>
				{{end}}
				{{with .Resources}}
			  	<div class="form-group">
			    	<label for="persistent_disk" class="col-sm-3 control-label">{{i18n $.Lang "PersistentDisk"}}</label>
				    <div class="col-sm-7">
						<input type="hidden" name="resourcesId" value = "{{.Id}}">
				      	<input type="text" class="form-control" id="persistent_disk" placeholder="16384" name="persistent_disk" value = "{{.PersistentDisk}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="instance_type" class="col-sm-3 control-label">{{i18n $.Lang "InstanceType"}}</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="instance_type" placeholder="{{i18n $.Lang "InstanceType"}}" name="instance_type" value = "{{.InstanceType}}" required>
			    	</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="availability_zone" class="col-sm-3 control-label">{{i18n $.Lang "AvailabilityZone"}}</label>
			    	<div class="col-sm-7">
			      	<input type="text" class="form-control" id="availability_zone" placeholder="{{i18n $.Lang "AvailabilityZone"}}" name="availability_zone" value = "{{.AvailabilityZone}}" required>
			    	</div>
			  	</div>
				{{end}}
				{{with .CloudProperties}}
			  	<div class="form-group">
			    	<label for="auth_url" class="col-sm-3 control-label">{{i18n $.Lang "AuthUrl"}}</label>
			    	<div class="col-sm-7">
					<input type="hidden" name="cloudPropertiesId" value = "{{.Id}}">
			      	<input type="text" class="form-control" id="auth_url" placeholder="{{i18n $.Lang "AuthUrl"}}" name="auth_url" value = "{{.AuthUrl}}" required>
			    	</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="username" class="col-sm-3 control-label">{{i18n $.Lang "UserName"}}</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="username" placeholder="{{i18n $.Lang "UserName"}}" name="username" value = "{{.UserName}}" required>
			    	</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="api_key" class="col-sm-3 control-label">{{i18n $.Lang "ApiKey"}}</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="api_key" placeholder="{{i18n $.Lang "ApiKey"}}" name="api_key" value = "{{.ApiKey}}" required>
			    	</div>
			  	</div>
				<div class="form-group">
				    <label for="tenant" class="col-sm-3 control-label">{{i18n $.Lang "Tenant"}}</label>
				    <div class="col-sm-7">
				      	<input type="text" class="form-control" id="tenant" placeholder="{{i18n $.Lang "Tenant"}}" name="tenant" value = "{{.Tenant}}" required>
				    </div>
				</div>
			  	<div class="form-group">
			    	<label for="default_key_name" class="col-sm-3 control-label">{{i18n $.Lang "SecretKey"}}</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="default_key_name" placeholder="{{i18n $.Lang "SecretKey"}}" name="default_key_name" value = "{{.DefaultKeyName}}" required>
			    	</div>
			  	</div>
				<div class="form-group">
				    <label for="private_key"  class="col-sm-3 control-label">{{i18n $.Lang "PrivateKey"}}</label>
					<div class="col-sm-7">
				    	<input type="file" id="private_key" name = "private_key">
				    	<p class="help-block">{{i18n $.Lang "Please upload your PrivateKey file about SecretKey"}}</p>
					</div>
				</div>
					{{if ne $.IaaSVersion $.DefaultVersion}}
					  	<div class="form-group">
						    	<label for="cci_ebs_url" class="col-sm-3 control-label">{{i18n $.Lang "NBSAuthUrl1"}}</label>
						    	<div class="col-sm-7">
						      		<input type="text" class="form-control" id="cci_ebs_url" placeholder="{{i18n $.Lang "NBSAuthUrl1"}}" name="cci_ebs_url" value = "{{.CciEbsUrl}}" required>
						    	</div>
					  	</div>
			
					  	<div class="form-group">
						    	<label for="cci_ebs_accesskey" class="col-sm-3 control-label">{{i18n $.Lang "NBSAccesskey"}}</label>
						    	<div class="col-sm-7">
						      		<input type="text" class="form-control" id="cci_ebs_accesskey" placeholder="{{i18n $.Lang "NBSAccesskey"}}" name="cci_ebs_accesskey" value = "{{.CciEbsAccesskey}}" required>
						    	</div>
					  	</div>
		
					  	<div class="form-group">
						    	<label for="cci_ebs_secretkey" class="col-sm-3 control-label">{{i18n $.Lang "NBSSecretkey"}}</label>
						    	<div class="col-sm-7">
						      		<input type="text" class="form-control" id="cci_ebs_secretkey" placeholder="{{i18n $.Lang "NBSSecretkey"}}" name="cci_ebs_secretkey" value = "{{.CciEbsSecretkey}}" required>
						    	</div>
					  	</div>
					{{end}}	
					{{end}}	
				{{end}}
			</form>
  		</div>
	</div>
</div>
