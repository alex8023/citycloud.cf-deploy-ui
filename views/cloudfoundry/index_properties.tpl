{{with .CloudFoundry}}
 				<div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">PaaS {{i18n $.Lang "Properties"}}</h2>
					</div>
					<div class="form-horizontal">
				  		<div class="panel-body">
							<div class="form-group">
							    <div class="col-sm-offset-2 col-sm-10">
							      	<button class="btn btn-default " id = "config-CloudFoundryProperties"><span class="glyphicon glyphicon-edit"></span> {{i18n $.Lang "Config"}}</button>
							    </div>
								<script type="text/javascript">
									$('#config-CloudFoundryProperties').on('click', function(){
							    		window.location.href = "cloudfoundry?action=config&model=CloudFoundryProperties";
							  		})
								</script>
							</div>
							{{with .CloudFoundryProperties}}
							<div class="form-group">
								<label class="col-sm-3 control-label">{{i18n $.Lang "Deployment Name"}}</label>
								<div class="col-sm-5">
									<p class="form-control-static">{{.Name}}</p>
								</div>
							</div>
							<div class="form-group">
								<label class="col-sm-3 control-label">{{i18n $.Lang "Director UUID"}}</label>
								<div class="col-sm-5">
									<p class="form-control-static">{{.Uuid}}</p>
								</div>
							</div>
							{{if eq $.IaaSVersion $.DefaultVersion}}
							{{else}}
							<div class="form-group">
								<label class="col-sm-3 control-label">{{i18n $.Lang "FloatingIp"}}</label>
								<div class="col-sm-5">
									<p class="form-control-static">{{.FloatingIp}}</p>
								</div>
							</div>
							{{end}}
							<div class="form-group">
								<label class="col-sm-3 control-label">{{i18n $.Lang "SystemDomain"}}</label>
								<div class="col-sm-5">
									<p class="form-control-static">{{.SystemDomain}}</p>
								</div>
							</div>
							<div class="form-group">
								<label class="col-sm-3 control-label">{{i18n $.Lang "SystemOrg"}}</label>
								<div class="col-sm-5">
									<p class="form-control-static">{{.SystemDomainOrg}}</p>
								</div>
							</div>
							{{end}}
						</div>
					</div>
				</div>
{{end}}