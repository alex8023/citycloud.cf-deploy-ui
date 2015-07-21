<input type="hidden" id="navfocus" value = "{{.NavBarFocus}}">
<div style="height:20px"></div>
<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">CloudFoundry CloudFoundryProperties</h2>
		</div>
		<div class="form-horizontal">
	  		<div class="panel-body">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      	<button class="btn btn-default " id = "config-CloudFoundryProperties">Config</button>
				    </div>
					<script type="text/javascript">
						$('#config-CloudFoundryProperties').on('click', function(){
				    		window.location.href = "cloudfoundry?action=config&model=CloudFoundryProperties";
				  		})
					</script>
				</div>
				{{with .CloudFoundry}}
				{{with .CloudFoundryProperties}}
				<div class="form-group">
					<label class="col-sm-2 control-label">Deployment Name</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.Name}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-2 control-label">Director UUID</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.Uuid}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-2 control-label">FloatingIp</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.FloatingIp}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-2 control-label">SystemDomain</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.SystemDomain}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-2 control-label">SystemOrg</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.SystemDomainOrg}}</p>
					</div>
				</div>
				{{end}}
				{{end}}
			</div>
		</div>
	</div>
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">CloudFoundry NetWorks</h2>
		</div>
		<div class="form-horizontal">
	  		<div class="panel-body">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      	<button class="btn btn-default " id = "config-cloudfoundry">Config</button>
				    </div>
					<script type="text/javascript">
						$('#config-NetWorks').on('click', function(){
				    		window.location.href = "cloudfoundry?action=config&model=NetWorks";
				  		})
					</script>
				</div>
				{{with .CloudFoundry}}
				{{with .NetWorks}}
				<div class="form-group">
					<label class="col-sm-2 control-label">Name</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.Name}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-2 control-label">NetType</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.NetType}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-2 control-label">NetId</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.NetId}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-2 control-label">Cidr</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.Cidr}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-2 control-label">Dns</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.Dns}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-2 control-label">ReservedIp</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.ReservedIp}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-2 control-label">StaticIp</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.StaticIp}}</p>
					</div>
				</div>
				{{end}}
				{{end}}
			</div>
		</div>
	</div>
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">CloudFoundry Compilation</h2>
		</div>
		<div class="form-horizontal">
	  		<div class="panel-body">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      	<button class="btn btn-default " id = "config-Compilation">Config</button>
				    </div>
					<script type="text/javascript">
						$('#config-Compilation').on('click', function(){
				    		window.location.href = "cloudfoundry?action=config&model=Compilation";
				  		})
					</script>
				</div>
				{{with .CloudFoundry}}
				{{with .Compilation}}
				<div class="form-group">
					<label class="col-sm-2 control-label">InstanceType</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.InstanceType}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-2 control-label">AvailabilityZone</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.AvailabilityZone}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-2 control-label">Workers</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.Workers}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-2 control-label">DefaultNetWork</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.DefaultNetWork}}</p>
					</div>
				</div>
				{{end}}
				{{end}}
			</div>
		</div>
	</div>
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">CloudFoundry ResourcesPool</h2>
		</div>
		<div class="form-horizontal">
	  		<div class="panel-body">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      	<button class="btn btn-default " id = "config-ResourcesPool">Config</button>
				    </div>
					<script type="text/javascript">
						$('#config-ResourcesPool').on('click', function(){
				    		window.location.href = "cloudfoundry?action=config&model=ResourcesPool";
				  		})
					</script>
				</div>
				{{with .CloudFoundry}}
					{{with .ResourcesPool}}
				    <div class="panel panel-default">
						<div class="panel-heading" >
							<h2 class="panel-title">ResourcesPool {{.Name}}</h2>
						</div>
						<div class="form-horizontal">
					  		<div class="panel-body">
								<div class="form-group">
									<label class="col-sm-2 control-label">InstanceType</label>
									<div class="col-sm-10">
										<p class="form-control-static">{{.InstanceType}}</p>
									</div>
								</div>
								<div class="form-group">
									<label class="col-sm-2 control-label">AvailabilityZone</label>
									<div class="col-sm-10">
										<p class="form-control-static">{{.AvailabilityZone}}</p>
									</div>
								</div>
								<div class="form-group">
									<label class="col-sm-2 control-label">Size</label>
									<div class="col-sm-10">
										<p class="form-control-static">{{.Size}}</p>
									</div>
								</div>
								<div class="form-group">
									<label class="col-sm-2 control-label">DefaultNetWork</label>
									<div class="col-sm-10">
										<p class="form-control-static">{{.DefaultNetWork}}</p>
									</div>
								</div>
							</div>
						</div>
					</div>
				    <div class="panel panel-default">
						<div class="panel-heading" >
							<h2 class="panel-title">ResourcesPool {{.Name}}</h2>
						</div>
						<div class="form-horizontal">
					  		<div class="panel-body">
								<div class="form-group">
									<label class="col-sm-2 control-label">InstanceType</label>
									<div class="col-sm-10">
										<p class="form-control-static">{{.InstanceType}}</p>
									</div>
								</div>
								<div class="form-group">
									<label class="col-sm-2 control-label">AvailabilityZone</label>
									<div class="col-sm-10">
										<p class="form-control-static">{{.AvailabilityZone}}</p>
									</div>
								</div>
								<div class="form-group">
									<label class="col-sm-2 control-label">Size</label>
									<div class="col-sm-10">
										<p class="form-control-static">{{.Size}}</p>
									</div>
								</div>
								<div class="form-group">
									<label class="col-sm-2 control-label">DefaultNetWork</label>
									<div class="col-sm-10">
										<p class="form-control-static">{{.DefaultNetWork}}</p>
									</div>
								</div>
							</div>
						</div>
					</div>
					{{end}}
				{{end}}
			</div>
		</div>
	</div>
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">CloudFoundry Jobs</h2>
		</div>
		<div class="form-horizontal">
	  		<div class="panel-body">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      	<button class="btn btn-default " id = "config-Compilation">Config</button>
				    </div>
					<script type="text/javascript">
						$('#config-Compilation').on('click', function(){
				    		window.location.href = "cloudfoundry?action=config&model=Compilation";
				  		})
					</script>
				</div>
				{{with .CloudFoundry}}
				{{with .Compilation}}
				<div class="form-group">
					<label class="col-sm-2 control-label">InstanceType</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.InstanceType}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-2 control-label">AvailabilityZone</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.AvailabilityZone}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-2 control-label">Workers</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.Workers}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-2 control-label">DefaultNetWork</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.DefaultNetWork}}</p>
					</div>
				</div>
				{{end}}
				{{end}}
			</div>
		</div>
	</div>

</div>
