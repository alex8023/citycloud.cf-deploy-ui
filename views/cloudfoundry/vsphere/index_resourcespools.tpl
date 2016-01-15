{{with .CloudFoundry}}
			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">PaaS {{i18n $.Lang "ResourcesPools"}}</h2>
					</div>
					<div class="form-horizontal">
				  		<div class="panel-body">
							<div class="form-group">
							    <div class="col-sm-offset-2 col-sm-10">
							      	<button class="btn btn-default " id = "config-ResourcesPools"><span class="glyphicon glyphicon-edit"></span> {{i18n $.Lang "Config"}}</button>
									<button class="btn btn-default " id = "config-flavor"><span class="glyphicon glyphicon-edit"></span> {{i18n $.Lang "Config-Flavor"}}</button>
							    </div>
								<script type="text/javascript">
									$('#config-ResourcesPools').on('click', function(){
							    		window.location.href = "vspherecloudfoundry?action=config&model=ResourcesPools";
							  		})
									$('#config-flavor').on('click', function(){
							    		window.location.href = "vsphereresource?action=list";
							  		})
								</script>
							</div>
								{{range .VsphereResourcesPools}}
							    <div class="panel panel-default">
									<div class="panel-heading" >
										<h2 class="panel-title">{{i18n $.Lang "ResourcesPool"}}</h2>
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
												<label class="col-sm-3 control-label">{{i18n $.Lang "Size"}}</label>
												<div class="col-sm-5">
													<p class="form-control-static">{{.Size}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-3 control-label">{{i18n $.Lang "DefaultNetWork"}}</label>
												<div class="col-sm-5">
													<p class="form-control-static">{{.DefaultNetWork}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-3 control-label">{{i18n $.Lang "Instance-Flavor"}}</label>
												<div class="col-sm-5">
													<p class="form-control-static">{{with .VsphereResource}}{{i18n $.Lang "Ram"}}({{.Ram}}M)-{{i18n $.Lang "Cpu"}}({{.Cpu}})-{{i18n $.Lang "Disk"}}({{.Disk}}M){{end}}</p>
												</div>
											</div>
										</div>
									</div>
								</div>
								{{end}}
						</div>
					</div>
				</div>
{{end}}