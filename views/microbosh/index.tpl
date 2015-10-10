<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">MicroBOSH Deployment</h2>
		</div>
		<div class="form-horizontal">
	  		<div class="panel-body">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      	<button class="btn btn-default " id = "config-microbosh"><span class="glyphicon glyphicon-edit"></span> Config</button>
					  	<button class="btn btn-default " id = "deploy-microbosh"><span class="glyphicon glyphicon-cog"></span> Deploy</button>
					<!--<button class="btn btn-primary ">primary</button>
						<button class="btn btn-warning ">warning</button>
						<button class="btn btn-danger ">danger</button>
						<button class="btn btn-link ">link</button>-->
				    </div>
					<script type="text/javascript">
						$('#config-microbosh').on('click', function(){
				    		window.location.href = "microbosh?action=config";
				  		})
						$('#deploy-microbosh').on('click',function(){
							window.location.href = "microbosh?action=deploy";
						})
					</script>
				</div>
				{{with .MicroBOSH}}
				<div class="form-group">
					<label class="col-sm-3 control-label">Deployment Name</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.Name}}</p>
					</div>
				</div>
				{{with .NetWork}}
				<div class="form-group">
					<label class="col-sm-3 control-label">Vip</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.Vip}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-3 control-label">NetId</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.NetId}}</p>
					</div>
				</div>
				{{end}}
				{{with .Resources}}
				<div class="form-group">
					<label class="col-sm-3 control-label">PersistentDiskSize</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.PersistentDisk}}</p>
					</div>
				</div>
				<div class="form-group">
					<label  class="col-sm-3 control-label">InstanceType</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.InstanceType}}</p>
					</div>
				</div>
				<div class="form-group">
					<label  class="col-sm-3 control-label">AvailabilityZone</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.AvailabilityZone}}</p>
					</div>
				</div>
				{{end}}
				{{with .CloudProperties}}
				<div class="form-group">
					<label  class="col-sm-3 control-label">AuthUrl</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.AuthUrl}}</p>
					</div>
				</div>
				<div class="form-group">
					<label  class="col-sm-3 control-label">UserName</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.UserName}}</p>
					</div>
				</div>
				<div class="form-group">
					<label  class="col-sm-3 control-label">ApiKey</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.ApiKey}}</p>
					</div>
				</div>
				<div class="form-group">
					<label  class="col-sm-3 control-label">Tenant</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.Tenant}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-3 control-label">SecretKey</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.DefaultKeyName}}</p>
					</div>
				</div>
				<div class="form-group">
					<label for="private_key"  class="col-sm-3 control-label">PrivateKey</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.PrivateKey}}</p>
					</div>
				</div>
				{{if ne $.IaaSVersion $.DefaultVersion}}
				<div class="form-group">
					<label  class="col-sm-3 control-label">NBSAuthUrl</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.CciEbsUrl}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-3 control-label">NBSAccesskey</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.CciEbsAccesskey}}</p>
					</div>
				</div>
				<div class="form-group">
					<label  class="col-sm-3 control-label">NBSSecretkey</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.CciEbsSecretkey}}</p>
					</div>
				</div>
				{{end}}
				{{end}}
				{{end}}
			</div>
		</div>
	</div>
</div>
