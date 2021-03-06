		{{with .CloudFoundry}}
			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">PaaS {{i18n $.Lang "NetWorks"}}</h2>
					</div>
					<div class="form-horizontal">
				  		<div class="panel-body">
							<div class="form-group">
							    <div class="col-sm-offset-2 col-sm-10">
							      	<button class="btn btn-default " id = "config-NetWorks"><span class="glyphicon glyphicon-edit"></span> {{i18n $.Lang "Config"}}</button>
							    </div>
								<script type="text/javascript">
									$('#config-NetWorks').on('click', function(){
							    		window.location.href = "cloudfoundry?action=config&model=NetWorks";
							  		})
								</script>
							</div>
							{{with .NetWorks}}
							{{with .private}}
							<div class="panel panel-default">
								<div class="panel-heading" >
									<h2 class="panel-title">{{i18n $.Lang "NetWorks-Private"}}</h2>
								</div>
								<div class="form-horizontal">
								  	<div class="panel-body">
										<div class="form-group">
											<label class="col-sm-3 control-label">{{i18n $.Lang "Name"}}</label>
											<div class="col-sm-5">
												<p class="form-control-static">{{.Name}}</p>
											</div>
										</div>
										<div class="form-group">
											<label class="col-sm-3 control-label">{{i18n $.Lang "NetType"}}</label>
											<div class="col-sm-5">
												<p class="form-control-static">{{.NetType}}</p>
											</div>
										</div>
										<div class="form-group">
											<label class="col-sm-3 control-label">{{i18n $.Lang "NetId"}}</label>
											<div class="col-sm-5">
												<p class="form-control-static">{{.NetId}}</p>
											</div>
										</div>
										<div class="form-group">
											<label class="col-sm-3 control-label">{{i18n $.Lang "Cidr"}}</label>
											<div class="col-sm-5">
												<p class="form-control-static">{{.Cidr}}</p>
											</div>
										</div>
										<div class="form-group">
											<label class="col-sm-3 control-label">{{i18n $.Lang "Dns"}}</label>
											<div class="col-sm-5">
												<p class="form-control-static">{{.Dns}}</p>
											</div>
										</div>
										<div class="form-group">
											<label class="col-sm-3 control-label">{{i18n $.Lang "ReservedIps"}}</label>
											<div class="col-sm-5">
												<p class="form-control-static">{{.ReservedIp}}</p>
											</div>
										</div>
										<div class="form-group">
											<label class="col-sm-3 control-label">{{i18n $.Lang "StaticIps"}}</label>
											<div class="col-sm-5">
												<p class="form-control-static">{{.StaticIp}}</p>
											</div>
										</div>
									</div>
								</div>
							</div>
							{{end}}
							{{with .public}}
							<div class="panel panel-default">
								<div class="panel-heading" >
									<h2 class="panel-title">{{i18n $.Lang "NetWorks-Public"}}</h2>
								</div>
								<div class="form-horizontal">
								  	<div class="panel-body">
										<div class="form-group">
											<label class="col-sm-3 control-label">{{i18n $.Lang "Name"}}</label>
											<div class="col-sm-5">
												<p class="form-control-static">{{.Name}}</p>
											</div>
										</div>
										<div class="form-group">
											<label class="col-sm-3 control-label">{{i18n $.Lang "NetType"}}</label>
											<div class="col-sm-5">
												<p class="form-control-static">{{.NetType}}</p>
											</div>
										</div>
										<div class="form-group">
											<label class="col-sm-3 control-label">{{i18n $.Lang "NetId"}}</label>
											<div class="col-sm-5">
												<p class="form-control-static">{{.NetId}}</p>
											</div>
										</div>
										<div class="form-group">
											<label class="col-sm-3 control-label">{{i18n $.Lang "Cidr"}}</label>
											<div class="col-sm-5">
												<p class="form-control-static">{{.Cidr}}</p>
											</div>
										</div>
										<div class="form-group">
											<label class="col-sm-3 control-label">{{i18n $.Lang "StaticIps"}}</label>
											<div class="col-sm-5">
												<p class="form-control-static">{{.StaticIp}}</p>
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
		{{end}}
