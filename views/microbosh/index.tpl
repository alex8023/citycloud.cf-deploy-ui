<input type="hidden" id="navfocus" value = "{{.NavBarFocus}}">
<div style="height:20px"></div>
<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">MicroBOSH Deployment</h2>
		</div>
		<div class="form-horizontal">
	  		<div class="panel-body">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      	<button class="btn btn-default " id = "config-microbosh">Config</button>
					  	<button class="btn btn-default " id = "deploy-microbosh">Deploy</button>
					<!--<button class="btn btn-primary ">primary</button>
						<button class="btn btn-warning ">warning</button>
						<button class="btn btn-danger ">danger</button>
						<button class="btn btn-link ">link</button>-->
				    </div>
					<script>
						$('#config-microbosh').on('click', function(){
				    		window.location.href = "microbosh?action=config";
				  		})
					</script>
				</div>
				{{with .MicroBOSH}}
				<div class="form-group">
					<label class="col-sm-2 control-label">Deployment Name</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.Name}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-2 control-label">Vip</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{with .Network}}{{.Vip}}{{end}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-2 control-label">NetId</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{with .Network}}{{.NetId}}{{end}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-2 control-label">PersistentDiskSize</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{with .Resources}}{{.PersistentDisk}}{{end}}</p>
					</div>
				</div>
				<div class="form-group">
					<label  class="col-sm-2 control-label">InstanceFlavor</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{with .Resources}}{{.InstanceType}}{{end}}</p>
					</div>
				</div>
				<div class="form-group">
					<label  class="col-sm-2 control-label">AvailabilityZone</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{with .Resources}}{{.AvailabilityZone}}{{end}}</p>
					</div>
				</div>
				<div class="form-group">
					<label  class="col-sm-2 control-label">AuthUrl</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{with .CloudProperties}}{{.AuthUrl}}</p>
					</div>
				</div>
				<div class="form-group">
					<label  class="col-sm-2 control-label">UserName</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.UserName}}</p>
					</div>
				</div>
				<div class="form-group">
					<label  class="col-sm-2 control-label">ApiKey</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.ApiKey}}</p>
					</div>
				</div>
				<div class="form-group">
					<label  class="col-sm-2 control-label">Tenant</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.Tenant}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-2 control-label">SecretKey</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.DefaultKeyName}}</p>
					</div>
				</div>
				<div class="form-group">
					<label for="private_key"  class="col-sm-2 control-label">PrivateKey</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.PrivateKey}}</p>
					</div>
				</div>
				
				<div class="form-group">
					<label  class="col-sm-2 control-label">NBSAuthUrl</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.CciEbsUrl}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-2 control-label">NBSAccesskey</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.CciEbsAccesskey}}</p>
					</div>
				</div>
				<div class="form-group">
					<label  class="col-sm-2 control-label">NBSSecretkey</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.CciEbsSecretkey}}{{end}}</p>
					</div>
				</div>
				{{end}}
			</div>
		</div>
	</div>
</div>
