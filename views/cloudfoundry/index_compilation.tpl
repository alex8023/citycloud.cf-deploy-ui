{{with .CloudFoundry}}
			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">PaaS Compilation</h2>
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
							{{with .Compilation}}
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
								<label class="col-sm-3 control-label">Workers</label>
								<div class="col-sm-5">
									<p class="form-control-static">{{.Workers}}</p>
								</div>
							</div>
							<div class="form-group">
								<label class="col-sm-3 control-label">DefaultNetWork</label>
								<div class="col-sm-5">
									<p class="form-control-static">{{.DefaultNetWork}}</p>
								</div>
							</div>
							{{end}}
						</div>
					</div>
				</div>
{{end}}