{{with .CloudFoundry}}
			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">PaaS Jobs</h2>
					</div>
					<div class="form-horizontal">
				  		<div class="panel-body">
							<div class="form-group">
							    <div class="col-sm-offset-2 col-sm-10">
							      	<button class="btn btn-default " id = "config-CloudFoundryJobs">Config</button>
							    </div>
								<script type="text/javascript">
									$('#config-CloudFoundryJobs').on('click', function(){
							    		window.location.href = "cloudfoundry?action=config&model=CloudFoundryJobs";
							  		})
								</script>
							</div>
							{{with .CloudFoundryJobs}}
								{{with .haproxy}}
							    <div class="panel panel-default">
									<div class="panel-heading" >
										<h2 class="panel-title">Jobs-{{.JobName}}</h2>
									</div>
									<div class="form-horizontal">
								  		<div class="panel-body">
											<div class="form-group">
												<label class="col-sm-2 control-label">Name</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Name}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">ResourcesPoolName</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.ResourcesPool}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">Instances</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Instances}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">StaticIp</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.StaticIp}}</p>
												</div>
											</div>
										</div>
									</div>
								</div>
								{{end}}
								{{with .gorouter}}
							    <div class="panel panel-default">
									<div class="panel-heading" >
										<h2 class="panel-title">Jobs-{{.JobName}}</h2>
									</div>
									<div class="form-horizontal">
								  		<div class="panel-body">
											<div class="form-group">
												<label class="col-sm-2 control-label">Name</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Name}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">ResourcesPoolName</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.ResourcesPool}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">Instances</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Instances}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">StaticIp</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.StaticIp}}</p>
												</div>
											</div>
										</div>
									</div>
								</div>
								{{end}}
								{{with .postgres}}
							    <div class="panel panel-default">
									<div class="panel-heading" >
										<h2 class="panel-title">Jobs-{{.JobName}}</h2>
									</div>
									<div class="form-horizontal">
								  		<div class="panel-body">
											<div class="form-group">
												<label class="col-sm-2 control-label">Name</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Name}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">ResourcesPoolName</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.ResourcesPool}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">Instances</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Instances}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">StaticIp</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.StaticIp}}</p>
												</div>
											</div>
										</div>
									</div>
								</div>
								{{end}}
								{{with .nfs}}
							    <div class="panel panel-default">
									<div class="panel-heading" >
										<h2 class="panel-title">Jobs-{{.JobName}}</h2>
									</div>
									<div class="form-horizontal">
								  		<div class="panel-body">
											<div class="form-group">
												<label class="col-sm-2 control-label">Name</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Name}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">ResourcesPoolName</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.ResourcesPool}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">Instances</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Instances}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">StaticIp</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.StaticIp}}</p>
												</div>
											</div>
										</div>
									</div>
								</div>
								{{end}}
								{{with .nats}}
							    <div class="panel panel-default">
									<div class="panel-heading" >
										<h2 class="panel-title">Jobs-{{.JobName}}</h2>
									</div>
									<div class="form-horizontal">
								  		<div class="panel-body">
											<div class="form-group">
												<label class="col-sm-2 control-label">Name</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Name}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">ResourcesPoolName</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.ResourcesPool}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">Instances</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Instances}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">StaticIp</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.StaticIp}}</p>
												</div>
											</div>
										</div>
									</div>
								</div>
								{{end}}
								{{with .syslog_aggregator}}
							    <div class="panel panel-default">
									<div class="panel-heading" >
										<h2 class="panel-title">Jobs-{{.JobName}}</h2>
									</div>
									<div class="form-horizontal">
								  		<div class="panel-body">
											<div class="form-group">
												<label class="col-sm-2 control-label">Name</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Name}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">ResourcesPoolName</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.ResourcesPool}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">Instances</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Instances}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">StaticIp</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.StaticIp}}</p>
												</div>
											</div>
										</div>
									</div>
								</div>
								{{end}}
								{{with .etcd}}
							    <div class="panel panel-default">
									<div class="panel-heading" >
										<h2 class="panel-title">Jobs-{{.JobName}}</h2>
									</div>
									<div class="form-horizontal">
								  		<div class="panel-body">
											<div class="form-group">
												<label class="col-sm-2 control-label">Name</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Name}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">ResourcesPoolName</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.ResourcesPool}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">Instances</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Instances}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">StaticIp</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.StaticIp}}</p>
												</div>
											</div>
										</div>
									</div>
								</div>
								{{end}}
								{{with .loggregator}}
							    <div class="panel panel-default">
									<div class="panel-heading" >
										<h2 class="panel-title">Jobs-{{.JobName}}</h2>
									</div>
									<div class="form-horizontal">
								  		<div class="panel-body">
											<div class="form-group">
												<label class="col-sm-2 control-label">Name</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Name}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">ResourcesPoolName</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.ResourcesPool}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">Instances</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Instances}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">StaticIp</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.StaticIp}}</p>
												</div>
											</div>
										</div>
									</div>
								</div>
								{{end}}
								{{with .uaa}}
							    <div class="panel panel-default">
									<div class="panel-heading" >
										<h2 class="panel-title">Jobs-{{.JobName}}</h2>
									</div>
									<div class="form-horizontal">
								  		<div class="panel-body">
											<div class="form-group">
												<label class="col-sm-2 control-label">Name</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Name}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">ResourcesPoolName</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.ResourcesPool}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">Instances</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Instances}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">StaticIp</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.StaticIp}}</p>
												</div>
											</div>
										</div>
									</div>
								</div>
								{{end}}
								{{with .cloud_controller_ng}}
							    <div class="panel panel-default">
									<div class="panel-heading" >
										<h2 class="panel-title">Jobs-{{.JobName}}</h2>
									</div>
									<div class="form-horizontal">
								  		<div class="panel-body">
											<div class="form-group">
												<label class="col-sm-2 control-label">Name</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Name}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">ResourcesPoolName</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.ResourcesPool}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">Instances</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Instances}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">StaticIp</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.StaticIp}}</p>
												</div>
											</div>
										</div>
									</div>
								</div>
								{{end}}
								{{with .cloud_controller_worker}}
							    <div class="panel panel-default">
									<div class="panel-heading" >
										<h2 class="panel-title">Jobs-{{.JobName}}</h2>
									</div>
									<div class="form-horizontal">
								  		<div class="panel-body">
											<div class="form-group">
												<label class="col-sm-2 control-label">Name</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Name}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">ResourcesPoolName</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.ResourcesPool}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">Instances</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Instances}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">StaticIp</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.StaticIp}}</p>
												</div>
											</div>
										</div>
									</div>
								</div>
								{{end}}
								{{with .cloud_controller_clock}}
							    <div class="panel panel-default">
									<div class="panel-heading" >
										<h2 class="panel-title">Jobs-{{.JobName}}</h2>
									</div>
									<div class="form-horizontal">
								  		<div class="panel-body">
											<div class="form-group">
												<label class="col-sm-2 control-label">Name</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Name}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">ResourcesPoolName</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.ResourcesPool}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">Instances</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Instances}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">StaticIp</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.StaticIp}}</p>
												</div>
											</div>
										</div>
									</div>
								</div>
								{{end}}
								{{with .hm9000}}
							    <div class="panel panel-default">
									<div class="panel-heading" >
										<h2 class="panel-title">Jobs-{{.JobName}}</h2>
									</div>
									<div class="form-horizontal">
								  		<div class="panel-body">
											<div class="form-group">
												<label class="col-sm-2 control-label">Name</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Name}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">ResourcesPoolName</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.ResourcesPool}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">Instances</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Instances}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">StaticIp</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.StaticIp}}</p>
												</div>
											</div>
										</div>
									</div>
								</div>
								{{end}}
								{{with .stats}}
							    <div class="panel panel-default">
									<div class="panel-heading" >
										<h2 class="panel-title">Jobs-{{.JobName}}</h2>
									</div>
									<div class="form-horizontal">
								  		<div class="panel-body">
											<div class="form-group">
												<label class="col-sm-2 control-label">Name</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Name}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">ResourcesPoolName</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.ResourcesPool}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">Instances</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Instances}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">StaticIp</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.StaticIp}}</p>
												</div>
											</div>
										</div>
									</div>
								</div>
								{{end}}
								{{with .dea_next}}
							    <div class="panel panel-default">
									<div class="panel-heading" >
										<h2 class="panel-title">Jobs-{{.JobName}}</h2>
									</div>
									<div class="form-horizontal">
								  		<div class="panel-body">
											<div class="form-group">
												<label class="col-sm-2 control-label">Name</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Name}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">ResourcesPoolName</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.ResourcesPool}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">Instances</label>
												<div class="col-sm-10">
													<p class="form-control-static">{{.Instances}}</p>
												</div>
											</div>
											<div class="form-group">
												<label class="col-sm-2 control-label">StaticIp</label>
												<div class="col-sm-10">
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