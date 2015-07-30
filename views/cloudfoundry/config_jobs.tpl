<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">PaaS Jobs</h2>
		</div>
  		<div class="panel-body">
			<form class="form-horizontal" method="post" action="cloudfoundry">
				<input type = "hidden" name="model" value = "{{.Model}}">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      <button type="submit" class="btn btn-default " data-loading-text="Saving...">Save</button>
				    </div>
			  	</div>
				{{with .CloudFoundry}}
				{{$ResourcesPools := .ResourcesPools}}
				{{with .CloudFoundryJobs}}
			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">Jobs-{{.haproxy.JobName}}</h2>
					</div>
					<div class="form-horizontal">
					  	<div class="panel-body">
						  	<div class="form-group">
						    	<label for="{{.haproxy.JobName}}_name" class="col-sm-2 control-label">Name</label>
							    <div class="col-sm-10">
							      	<input type="text" class="form-control" id="{{.haproxy.JobName}}_name" placeholder="Name" name="{{.haproxy.JobName}}_name" value = "{{.haproxy.Name}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.haproxy.JobName}}_resourcesPool" class="col-sm-2 control-label">ResourcesPool</label>
							    <div class="col-sm-10">
							      	<input type="hidden" id="{{.haproxy.JobName}}_resourcesPool" name="{{.haproxy.JobName}}_resourcesPool" value = "{{.haproxy.ResourcesPool}}">
							    	<select class="form-control" id="{{.haproxy.JobName}}_resourcesPool" name="{{.haproxy.JobName}}_resourcesPool_select" >
									{{$haproxy := .haproxy.JobName}}
									{{range $ResourcesPools}}<option id="{{$haproxy}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
									</select>
									<script type="text/javascript">
									$("#{{.haproxy.JobName}}_{{.haproxy.ResourcesPool}}").attr("selected",true);  
									</script>
								</div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.haproxy.JobName}}_instances" class="col-sm-2 control-label">Instances</label>
							    <div class="col-sm-10">
							      	<input type="number" class="form-control" id="{{.haproxy.JobName}}_instances" placeholder="Instances" name="{{.haproxy.JobName}}_instances" value = "{{.haproxy.Instances}}" required readonly>
							    </div>
						  	</div>
						</div>
					</div>
				</div>

			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">Jobs-{{.gorouter.JobName}}</h2>
					</div>
					<div class="form-horizontal">
					  	<div class="panel-body">
						  	<div class="form-group">
						    	<label for="{{.gorouter.JobName}}_name" class="col-sm-2 control-label">Name</label>
							    <div class="col-sm-10">
							      	<input type="text" class="form-control" id="{{.gorouter.JobName}}_name" placeholder="Name" name="{{.gorouter.JobName}}_name" value = "{{.gorouter.Name}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.gorouter.JobName}}_resourcesPool" class="col-sm-2 control-label">ResourcesPool</label>
							    <div class="col-sm-10">
							      	<input type="hidden" id="{{.gorouter.JobName}}_resourcesPool" name="{{.gorouter.JobName}}_resourcesPool" value = "{{.gorouter.ResourcesPool}}">
							    	<select class="form-control" id="{{.gorouter.JobName}}_resourcesPool" name="{{.gorouter.JobName}}_resourcesPool_select" >
									{{$gorouter := .gorouter.JobName}}
									{{range $ResourcesPools}}<option id="{{$gorouter}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
									</select>
									<script type="text/javascript">
									$("#{{.gorouter.JobName}}_{{.gorouter.ResourcesPool}}").attr("selected",true);  
									</script>
								</div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.gorouter.JobName}}_instances" class="col-sm-2 control-label">Instances</label>
							    <div class="col-sm-10">
							      	<input type="number" class="form-control" id="{{.gorouter.JobName}}_instances" placeholder="Instances" name="{{.gorouter.JobName}}_instances" value = "{{.gorouter.Instances}}" required >
							    </div>
						  	</div>
						</div>
					</div>
				</div>

			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">Jobs-{{.postgres.JobName}}</h2>
					</div>
					<div class="form-horizontal">
					  	<div class="panel-body">
						  	<div class="form-group">
						    	<label for="{{.postgres.JobName}}_name" class="col-sm-2 control-label">Name</label>
							    <div class="col-sm-10">
							      	<input type="text" class="form-control" id="{{.postgres.JobName}}_name" placeholder="Name" name="{{.postgres.JobName}}_name" value = "{{.postgres.Name}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.postgres.JobName}}_resourcesPool" class="col-sm-2 control-label">ResourcesPool</label>
							    <div class="col-sm-10">
							      	<input type="hidden" id="{{.postgres.JobName}}_resourcesPool" name="{{.postgres.JobName}}_resourcesPool" value = "{{.postgres.ResourcesPool}}">
							    	<select class="form-control" id="{{.postgres.JobName}}_resourcesPool" name="{{.postgres.JobName}}_resourcesPool_select" >
									{{$postgres := .postgres.JobName}}
									{{range $ResourcesPools}}<option id="{{$postgres}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
									</select>
									<script type="text/javascript">
									$("#{{.postgres.JobName}}_{{.postgres.ResourcesPool}}").attr("selected",true);  
									</script>
								</div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.postgres.JobName}}_instances" class="col-sm-2 control-label">Instances</label>
							    <div class="col-sm-10">
							      	<input type="number" class="form-control" id="{{.postgres.JobName}}_instances" placeholder="Instances" name="{{.postgres.JobName}}_instances" value = "{{.postgres.Instances}}" required readonly>
							    </div>
						  	</div>
						</div>
					</div>
				</div>

			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">Jobs-{{.nfs.JobName}}</h2>
					</div>
					<div class="form-horizontal">
					  	<div class="panel-body">
						  	<div class="form-group">
						    	<label for="{{.nfs.JobName}}_name" class="col-sm-2 control-label">Name</label>
							    <div class="col-sm-10">
							      	<input type="text" class="form-control" id="{{.nfs.JobName}}_name" placeholder="Name" name="{{.nfs.JobName}}_name" value = "{{.nfs.Name}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.nfs.JobName}}_resourcesPool" class="col-sm-2 control-label">ResourcesPool</label>
							    <div class="col-sm-10">
							      	<input type="hidden" id="{{.nfs.JobName}}_resourcesPool" name="{{.nfs.JobName}}_resourcesPool" value = "{{.nfs.ResourcesPool}}">
							    	<select class="form-control" id="{{.nfs.JobName}}_resourcesPool" name="{{.nfs.JobName}}_resourcesPool_select" >
									{{$nfs := .nfs.JobName}}
									{{range $ResourcesPools}}<option id="{{$nfs}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
									</select>
									<script type="text/javascript">
									$("#{{.nfs.JobName}}_{{.nfs.ResourcesPool}}").attr("selected",true);  
									</script>
								</div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.nfs.JobName}}_instances" class="col-sm-2 control-label">Instances</label>
							    <div class="col-sm-10">
							      	<input type="number" class="form-control" id="{{.nfs.JobName}}_instances" placeholder="Instances" name="{{.nfs.JobName}}_instances" value = "{{.nfs.Instances}}" required readonly>
							    </div>
						  	</div>
						</div>
					</div>
				</div>

			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">Jobs-{{.nats.JobName}}</h2>
					</div>
					<div class="form-horizontal">
					  	<div class="panel-body">
						  	<div class="form-group">
						    	<label for="{{.nats.JobName}}_name" class="col-sm-2 control-label">Name</label>
							    <div class="col-sm-10">
							      	<input type="text" class="form-control" id="{{.nats.JobName}}_name" placeholder="Name" name="{{.nats.JobName}}_name" value = "{{.nats.Name}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.nats.JobName}}_resourcesPool" class="col-sm-2 control-label">ResourcesPool</label>
							    <div class="col-sm-10">
							      	<input type="hidden" id="{{.nats.JobName}}_resourcesPool" name="{{.nats.JobName}}_resourcesPool" value = "{{.nats.ResourcesPool}}">
							    	<select class="form-control" id="{{.nats.JobName}}_resourcesPool" name="{{.nats.JobName}}_resourcesPool_select" >
									{{$nats := .nats.JobName}}
									{{range $ResourcesPools}}<option id="{{$nats}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
									</select>
									<script type="text/javascript">
									$("#{{.nats.JobName}}_{{.nats.ResourcesPool}}").attr("selected",true);  
									</script>
								</div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.nats.JobName}}_instances" class="col-sm-2 control-label">Instances</label>
							    <div class="col-sm-10">
							      	<input type="number" class="form-control" id="{{.nats.JobName}}_instances" placeholder="Instances" name="{{.nats.JobName}}_instances" value = "{{.nats.Instances}}" required readonly>
							    </div>
						  	</div>
						</div>
					</div>
				</div>

			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">Jobs-{{.syslog_aggregator.JobName}}</h2>
					</div>
					<div class="form-horizontal">
					  	<div class="panel-body">
						  	<div class="form-group">
						    	<label for="{{.syslog_aggregator.JobName}}_name" class="col-sm-2 control-label">Name</label>
							    <div class="col-sm-10">
							      	<input type="text" class="form-control" id="{{.syslog_aggregator.JobName}}_name" placeholder="Name" name="{{.syslog_aggregator.JobName}}_name" value = "{{.syslog_aggregator.Name}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.syslog_aggregator.JobName}}_resourcesPool" class="col-sm-2 control-label">ResourcesPool</label>
							    <div class="col-sm-10">
							      	<input type="hidden" id="{{.syslog_aggregator.JobName}}_resourcesPool" name="{{.syslog_aggregator.JobName}}_resourcesPool" value = "{{.syslog_aggregator.ResourcesPool}}">
							    	<select class="form-control" id="{{.syslog_aggregator.JobName}}_resourcesPool" name="{{.syslog_aggregator.JobName}}_resourcesPool_select" >
									{{$syslog_aggregator := .syslog_aggregator.JobName}}
									{{range $ResourcesPools}}<option id="{{$syslog_aggregator}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
									</select>
									<script type="text/javascript">
									$("#{{.syslog_aggregator.JobName}}_{{.syslog_aggregator.ResourcesPool}}").attr("selected",true);  
									</script>
								</div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.syslog_aggregator.JobName}}_instances" class="col-sm-2 control-label">Instances</label>
							    <div class="col-sm-10">
							      	<input type="number" class="form-control" id="{{.syslog_aggregator.JobName}}_instances" placeholder="Instances" name="{{.syslog_aggregator.JobName}}_instances" value = "{{.syslog_aggregator.Instances}}" required readonly>
							    </div>
						  	</div>
						</div>
					</div>
				</div>

			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">Jobs-{{.etcd.JobName}}</h2>
					</div>
					<div class="form-horizontal">
					  	<div class="panel-body">
						  	<div class="form-group">
						    	<label for="{{.etcd.JobName}}_name" class="col-sm-2 control-label">Name</label>
							    <div class="col-sm-10">
							      	<input type="text" class="form-control" id="{{.etcd.JobName}}_name" placeholder="Name" name="{{.etcd.JobName}}_name" value = "{{.etcd.Name}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.etcd.JobName}}_resourcesPool" class="col-sm-2 control-label">ResourcesPool</label>
							    <div class="col-sm-10">
							      	<input type="hidden" id="{{.etcd.JobName}}_resourcesPool" name="{{.etcd.JobName}}_resourcesPool" value = "{{.etcd.ResourcesPool}}">
							    	<select class="form-control" id="{{.etcd.JobName}}_resourcesPool" name="{{.etcd.JobName}}_resourcesPool_select" >
									{{$etcd := .etcd.JobName}}
									{{range $ResourcesPools}}<option id="{{$etcd}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
									</select>
									<script type="text/javascript">
									$("#{{.etcd.JobName}}_{{.etcd.ResourcesPool}}").attr("selected",true);  
									</script>
								</div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.etcd.JobName}}_instances" class="col-sm-2 control-label">Instances</label>
							    <div class="col-sm-10">
							      	<input type="number" class="form-control" id="{{.etcd.JobName}}_instances" placeholder="Instances" name="{{.etcd.JobName}}_instances" value = "{{.etcd.Instances}}" required readonly>
							    </div>
						  	</div>
						</div>
					</div>
				</div>

			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">Jobs-{{.loggregator.JobName}}</h2>
					</div>
					<div class="form-horizontal">
					  	<div class="panel-body">
						  	<div class="form-group">
						    	<label for="{{.loggregator.JobName}}_name" class="col-sm-2 control-label">Name</label>
							    <div class="col-sm-10">
							      	<input type="text" class="form-control" id="{{.loggregator.JobName}}_name" placeholder="Name" name="{{.loggregator.JobName}}_name" value = "{{.loggregator.Name}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.loggregator.JobName}}_resourcesPool" class="col-sm-2 control-label">ResourcesPool</label>
							    <div class="col-sm-10">
							      	<input type="hidden" id="{{.loggregator.JobName}}_resourcesPool" name="{{.loggregator.JobName}}_resourcesPool" value = "{{.loggregator.ResourcesPool}}">
							    	<select class="form-control" id="{{.loggregator.JobName}}_resourcesPool" name="{{.loggregator.JobName}}_resourcesPool_select" >
									{{$loggregator := .loggregator.JobName}}
									{{range $ResourcesPools}}<option id="{{$loggregator}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
									</select>
									<script type="text/javascript">
									$("#{{.loggregator.JobName}}_{{.loggregator.ResourcesPool}}").attr("selected",true);  
									</script>
								</div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.loggregator.JobName}}_instances" class="col-sm-2 control-label">Instances</label>
							    <div class="col-sm-10">
							      	<input type="number" class="form-control" id="{{.loggregator.JobName}}_instances" placeholder="Instances" name="{{.etcd.JobName}}_instances" value = "{{.loggregator.Instances}}" required readonly>
							    </div>
						  	</div>
						</div>
					</div>
				</div>

			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">Jobs-{{.uaa.JobName}}</h2>
					</div>
					<div class="form-horizontal">
					  	<div class="panel-body">
						  	<div class="form-group">
						    	<label for="{{.uaa.JobName}}_name" class="col-sm-2 control-label">Name</label>
							    <div class="col-sm-10">
							      	<input type="text" class="form-control" id="{{.uaa.JobName}}_name" placeholder="Name" name="{{.uaa.JobName}}_name" value = "{{.uaa.Name}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.uaa.JobName}}_resourcesPool" class="col-sm-2 control-label">ResourcesPool</label>
							    <div class="col-sm-10">
							      	<input type="hidden" id="{{.uaa.JobName}}_resourcesPool" name="{{.uaa.JobName}}_resourcesPool" value = "{{.uaa.ResourcesPool}}">
							    	<select class="form-control" id="{{.uaa.JobName}}_resourcesPool" name="{{.uaa.JobName}}_resourcesPool_select" >
									{{$uaa := .uaa.JobName}}
									{{range $ResourcesPools}}<option id="{{$uaa}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
									</select>
									<script type="text/javascript">
									$("#{{.uaa.JobName}}_{{.uaa.ResourcesPool}}").attr("selected",true);  
									</script>
								</div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.uaa.JobName}}_instances" class="col-sm-2 control-label">Instances</label>
							    <div class="col-sm-10">
							      	<input type="number" class="form-control" id="{{.uaa.JobName}}_instances" placeholder="Instances" name="{{.etcd.JobName}}_instances" value = "{{.uaa.Instances}}" required readonly>
							    </div>
						  	</div>
						</div>
					</div>
				</div>

			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">Jobs-{{.cloud_controller_ng.JobName}}</h2>
					</div>
					<div class="form-horizontal">
					  	<div class="panel-body">
						  	<div class="form-group">
						    	<label for="{{.cloud_controller_ng.JobName}}_name" class="col-sm-2 control-label">Name</label>
							    <div class="col-sm-10">
							      	<input type="text" class="form-control" id="{{.cloud_controller_ng.JobName}}_name" placeholder="Name" name="{{.cloud_controller_ng.JobName}}_name" value = "{{.cloud_controller_ng.Name}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.cloud_controller_ng.JobName}}_resourcesPool" class="col-sm-2 control-label">ResourcesPool</label>
							    <div class="col-sm-10">
							      	<input type="hidden" id="{{.cloud_controller_ng.JobName}}_resourcesPool" name="{{.cloud_controller_ng.JobName}}_resourcesPool" value = "{{.cloud_controller_ng.ResourcesPool}}">
							    	<select class="form-control" id="{{.cloud_controller_ng.JobName}}_resourcesPool" name="{{.cloud_controller_ng.JobName}}_resourcesPool_select" >
									{{$cloud_controller_ng := .cloud_controller_ng.JobName}}
									{{range $ResourcesPools}}<option id="{{$cloud_controller_ng}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
									</select>
									<script type="text/javascript">
									$("#{{.cloud_controller_ng.JobName}}_{{.cloud_controller_ng.ResourcesPool}}").attr("selected",true);  
									</script>
								</div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.cloud_controller_ng.JobName}}_instances" class="col-sm-2 control-label">Instances</label>
							    <div class="col-sm-10">
							      	<input type="number" class="form-control" id="{{.cloud_controller_ng.JobName}}_instances" placeholder="Instances" name="{{.etcd.JobName}}_instances" value = "{{.cloud_controller_ng.Instances}}" required readonly>
							    </div>
						  	</div>
						</div>
					</div>
				</div>

			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">Jobs-{{.cloud_controller_worker.JobName}}</h2>
					</div>
					<div class="form-horizontal">
					  	<div class="panel-body">
						  	<div class="form-group">
						    	<label for="{{.cloud_controller_worker.JobName}}_name" class="col-sm-2 control-label">Name</label>
							    <div class="col-sm-10">
							      	<input type="text" class="form-control" id="{{.cloud_controller_worker.JobName}}_name" placeholder="Name" name="{{.cloud_controller_worker.JobName}}_name" value = "{{.cloud_controller_worker.Name}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.cloud_controller_worker.JobName}}_resourcesPool" class="col-sm-2 control-label">ResourcesPool</label>
							    <div class="col-sm-10">
							      	<input type="hidden" id="{{.cloud_controller_worker.JobName}}_resourcesPool" name="{{.cloud_controller_worker.JobName}}_resourcesPool" value = "{{.cloud_controller_worker.ResourcesPool}}">
							    	<select class="form-control" id="{{.cloud_controller_worker.JobName}}_resourcesPool" name="{{.cloud_controller_worker.JobName}}_resourcesPool_select" >
									{{$cloud_controller_worker := .cloud_controller_worker.JobName}}
									{{range $ResourcesPools}}<option id="{{$cloud_controller_worker}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
									</select>
									<script type="text/javascript">
									$("#{{.cloud_controller_worker.JobName}}_{{.cloud_controller_worker.ResourcesPool}}").attr("selected",true);  
									</script>
								</div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.cloud_controller_worker.JobName}}_instances" class="col-sm-2 control-label">Instances</label>
							    <div class="col-sm-10">
							      	<input type="number" class="form-control" id="{{.cloud_controller_worker.JobName}}_instances" placeholder="Instances" name="{{.etcd.JobName}}_instances" value = "{{.cloud_controller_worker.Instances}}" required readonly>
							    </div>
						  	</div>
						</div>
					</div>
				</div>

			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">Jobs-{{.cloud_controller_clock.JobName}}</h2>
					</div>
					<div class="form-horizontal">
					  	<div class="panel-body">
						  	<div class="form-group">
						    	<label for="{{.cloud_controller_clock.JobName}}_name" class="col-sm-2 control-label">Name</label>
							    <div class="col-sm-10">
							      	<input type="text" class="form-control" id="{{.cloud_controller_clock.JobName}}_name" placeholder="Name" name="{{.cloud_controller_clock.JobName}}_name" value = "{{.cloud_controller_clock.Name}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.cloud_controller_clock.JobName}}_resourcesPool" class="col-sm-2 control-label">ResourcesPool</label>
							    <div class="col-sm-10">
							      	<input type="hidden" id="{{.cloud_controller_clock.JobName}}_resourcesPool" name="{{.cloud_controller_clock.JobName}}_resourcesPool" value = "{{.cloud_controller_clock.ResourcesPool}}">
							    	<select class="form-control" id="{{.cloud_controller_clock.JobName}}_resourcesPool" name="{{.cloud_controller_clock.JobName}}_resourcesPool_select" >
									{{$cloud_controller_clock := .cloud_controller_clock.JobName}}
									{{range $ResourcesPools}}<option id="{{$cloud_controller_clock}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
									</select>
									<script type="text/javascript">
									$("#{{.cloud_controller_clock.JobName}}_{{.cloud_controller_clock.ResourcesPool}}").attr("selected",true);  
									</script>
								</div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.cloud_controller_clock.JobName}}_instances" class="col-sm-2 control-label">Instances</label>
							    <div class="col-sm-10">
							      	<input type="number" class="form-control" id="{{.cloud_controller_clock.JobName}}_instances" placeholder="Instances" name="{{.etcd.JobName}}_instances" value = "{{.cloud_controller_clock.Instances}}" required readonly>
							    </div>
						  	</div>
						</div>
					</div>
				</div>

			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">Jobs-{{.hm9000.JobName}}</h2>
					</div>
					<div class="form-horizontal">
					  	<div class="panel-body">
						  	<div class="form-group">
						    	<label for="{{.hm9000.JobName}}_name" class="col-sm-2 control-label">Name</label>
							    <div class="col-sm-10">
							      	<input type="text" class="form-control" id="{{.hm9000.JobName}}_name" placeholder="Name" name="{{.hm9000.JobName}}_name" value = "{{.hm9000.Name}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.hm9000.JobName}}_resourcesPool" class="col-sm-2 control-label">ResourcesPool</label>
							    <div class="col-sm-10">
							      	<input type="hidden" id="{{.hm9000.JobName}}_resourcesPool" name="{{.hm9000.JobName}}_resourcesPool" value = "{{.hm9000.ResourcesPool}}">
							    	<select class="form-control" id="{{.hm9000.JobName}}_resourcesPool" name="{{.hm9000.JobName}}_resourcesPool_select" >
									{{$hm9000 := .hm9000.JobName}}
									{{range $ResourcesPools}}<option id="{{$hm9000}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
									</select>
									<script type="text/javascript">
									$("#{{.hm9000.JobName}}_{{.hm9000.ResourcesPool}}").attr("selected",true);  
									</script>
								</div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.hm9000.JobName}}_instances" class="col-sm-2 control-label">Instances</label>
							    <div class="col-sm-10">
							      	<input type="number" class="form-control" id="{{.hm9000.JobName}}_instances" placeholder="Instances" name="{{.etcd.JobName}}_instances" value = "{{.hm9000.Instances}}" required readonly>
							    </div>
						  	</div>
						</div>
					</div>
				</div>

			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">Jobs-{{.stats.JobName}}</h2>
					</div>
					<div class="form-horizontal">
					  	<div class="panel-body">
						  	<div class="form-group">
						    	<label for="{{.stats.JobName}}_name" class="col-sm-2 control-label">Name</label>
							    <div class="col-sm-10">
							      	<input type="text" class="form-control" id="{{.stats.JobName}}_name" placeholder="Name" name="{{.stats.JobName}}_name" value = "{{.stats.Name}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.stats.JobName}}_resourcesPool" class="col-sm-2 control-label">ResourcesPool</label>
							    <div class="col-sm-10">
							      	<input type="hidden" id="{{.stats.JobName}}_resourcesPool" name="{{.stats.JobName}}_resourcesPool" value = "{{.stats.ResourcesPool}}">
							    	<select class="form-control" id="{{.stats.JobName}}_resourcesPool" name="{{.stats.JobName}}_resourcesPool_select" >
									{{$stats := .stats.JobName}}
									{{range $ResourcesPools}}<option id="{{$stats}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
									</select>
									<script type="text/javascript">
									$("#{{.stats.JobName}}_{{.stats.ResourcesPool}}").attr("selected",true);  
									</script>
								</div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.stats.JobName}}_instances" class="col-sm-2 control-label">Instances</label>
							    <div class="col-sm-10">
							      	<input type="number" class="form-control" id="{{.stats.JobName}}_instances" placeholder="Instances" name="{{.etcd.JobName}}_instances" value = "{{.stats.Instances}}" required readonly>
							    </div>
						  	</div>
						</div>
					</div>
				</div>

			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">Jobs-{{.dea_next.JobName}}</h2>
					</div>
					<div class="form-horizontal">
					  	<div class="panel-body">
						  	<div class="form-group">
						    	<label for="{{.dea_next.JobName}}_name" class="col-sm-2 control-label">Name</label>
							    <div class="col-sm-10">
							      	<input type="text" class="form-control" id="{{.dea_next.JobName}}_name" placeholder="Name" name="{{.dea_next.JobName}}_name" value = "{{.dea_next.Name}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.dea_next.JobName}}_resourcesPool" class="col-sm-2 control-label">ResourcesPool</label>
							    <div class="col-sm-10">
							      	<input type="hidden" id="{{.dea_next.JobName}}_resourcesPool" name="{{.dea_next.JobName}}_resourcesPool" value = "{{.dea_next.ResourcesPool}}">
							    	<select class="form-control" id="{{.dea_next.JobName}}_resourcesPool" name="{{.dea_next.JobName}}_resourcesPool_select" >
									{{$dea_next := .dea_next.JobName}}
									{{range $ResourcesPools}}<option id="{{$dea_next}}_{{.Name}}" value="{{.Name}}">{{.Name}}</option>{{end}}
									</select>
									<script type="text/javascript">
									$("#{{.dea_next.JobName}}_{{.dea_next.ResourcesPool}}").attr("selected",true);  
									</script>
								</div>
						  	</div>
						  	<div class="form-group">
						    	<label for="{{.dea_next.JobName}}_instances" class="col-sm-2 control-label">Instances</label>
							    <div class="col-sm-10">
							      	<input type="number" class="form-control" id="{{.dea_next.JobName}}_instances" placeholder="Instances" name="{{.etcd.JobName}}_instances" value = "{{.dea_next.Instances}}" required >
							    </div>
						  	</div>
						</div>
					</div>
				</div>
		
				{{end}}
				{{end}}
			</form>
  		</div>
	</div>
</div>
