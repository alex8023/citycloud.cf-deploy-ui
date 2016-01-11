{{with .CloudFoundry}}
			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">PaaS NetWorks</h2>
					</div>
					<div class="form-horizontal">
				  		<div class="panel-body">
							<div class="form-group">
							    <div class="col-sm-offset-2 col-sm-10">
							      	<button class="btn btn-default " id = "config-NetWorks"><span class="glyphicon glyphicon-edit"></span> Config</button>
							    </div>
								<script type="text/javascript">
									$('#config-NetWorks').on('click', function(){
							    		window.location.href = "vspherecloudfoundry?action=config&model=NetWorks";
							  		})
								</script>
							</div>
							{{with .NetWorks}}{{with .private}}
							<div class="form-group">
								<label class="col-sm-3 control-label">Name</label>
								<div class="col-sm-5">
									<p class="form-control-static">{{.Name}}</p>
								</div>
							</div>
							<div class="form-group">
								<label class="col-sm-3 control-label">VlanName</label>
								<div class="col-sm-5">
									<p class="form-control-static">{{.NetId}}</p>
								</div>
							</div>
							<div class="form-group">
								<label class="col-sm-3 control-label">Cidr</label>
								<div class="col-sm-5">
									<p class="form-control-static">{{.Cidr}}</p>
								</div>
							</div>
							<div class="form-group">
								<label class="col-sm-3 control-label">GateWay</label>
								<div class="col-sm-5">
									<p class="form-control-static">{{.NetType}}</p>
								</div>
							</div>
							<div class="form-group">
								<label class="col-sm-3 control-label">Dns</label>
								<div class="col-sm-5">
									<p class="form-control-static">{{.Dns}}</p>
								</div>
							</div>
							<div class="form-group">
								<label class="col-sm-3 control-label">ReservedIp</label>
								<div class="col-sm-5">
									<p class="form-control-static">{{.ReservedIp}}</p>
								</div>
							</div>
							<div class="form-group">
								<label class="col-sm-3 control-label">StaticIp</label>
								<div class="col-sm-5">
									<p class="form-control-static">{{.StaticIp}}</p>
								</div>
							</div>
							{{end}}{{end}}
						</div>
					</div>
				</div>
{{end}}
