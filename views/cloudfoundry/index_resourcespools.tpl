{{with .CloudFoundry}}
			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">PaaS ResourcesPools</h2>
					</div>
					<div class="form-horizontal">
				  		<div class="panel-body">
							<div class="form-group">
							    <div class="col-sm-offset-2 col-sm-10">
							      	<button class="btn btn-default " id = "config-ResourcesPools"><span class="glyphicon glyphicon-edit"></span> Config</button>
							    </div>
								<script type="text/javascript">
									$('#config-ResourcesPools').on('click', function(){
							    		window.location.href = "cloudfoundry?action=config&model=ResourcesPools";
							  		})
								</script>
							</div>
								{{range .ResourcesPools}}
							    <div class="panel panel-default">
									<div class="panel-heading" >
										<h2 class="panel-title">ResourcesPool</h2>
									</div>
									<div class="form-horizontal">
								  		<div class="panel-body">
											<div class="form-group">
												<label class="col-sm-3 control-label">Name</label>
												<div class="col-sm-5">
													<p class="form-control-static">{{.Name}}</p>
												</div>
											</div>
			
											<div class="form-group">
												<label class="col-sm-3 control-label">InstanceType</label>
												<div class="col-sm-5">
													<p class="form-control-static">{{.InstanceType}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-3 control-label">AvailabilityZone</label>
												<div class="col-sm-5">
													<p class="form-control-static">{{.AvailabilityZone}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-3 control-label">Size</label>
												<div class="col-sm-5">
													<p class="form-control-static">{{.Size}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-3 control-label">DefaultNetWork</label>
												<div class="col-sm-5">
													<p class="form-control-static">{{.DefaultNetWork}}</p>
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