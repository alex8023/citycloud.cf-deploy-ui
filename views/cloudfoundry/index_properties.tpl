{{with .CloudFoundry}}
 				<div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">PaaS PaaSProperties</h2>
					</div>
					<div class="form-horizontal">
				  		<div class="panel-body">
							<div class="form-group">
							    <div class="col-sm-offset-2 col-sm-10">
							      	<button class="btn btn-default " id = "config-CloudFoundryProperties"><span class="glyphicon glyphicon-edit"></span> Config</button>
							    </div>
								<script type="text/javascript">
									$('#config-CloudFoundryProperties').on('click', function(){
							    		window.location.href = "cloudfoundry?action=config&model=CloudFoundryProperties";
							  		})
								</script>
							</div>
							{{with .CloudFoundryProperties}}
							<div class="form-group">
								<label class="col-sm-3 control-label">Deployment Name</label>
								<div class="col-sm-5">
									<p class="form-control-static">{{.Name}}</p>
								</div>
							</div>
							<div class="form-group">
								<label class="col-sm-3 control-label">Director UUID</label>
								<div class="col-sm-5">
									<p class="form-control-static">{{.Uuid}}</p>
								</div>
							</div>
							{{if eq $.IaaSVersion $.DefaultVersion}}
							{{else}}
							<div class="form-group">
								<label class="col-sm-3 control-label">FloatingIp</label>
								<div class="col-sm-5">
									<p class="form-control-static">{{.FloatingIp}}</p>
								</div>
							</div>
							{{end}}
							<div class="form-group">
								<label class="col-sm-3 control-label">SystemDomain</label>
								<div class="col-sm-5">
									<p class="form-control-static">{{.SystemDomain}}</p>
								</div>
							</div>
							<div class="form-group">
								<label class="col-sm-3 control-label">SystemOrg</label>
								<div class="col-sm-5">
									<p class="form-control-static">{{.SystemDomainOrg}}</p>
								</div>
							</div>
							{{end}}
						</div>
					</div>
				</div>
{{end}}