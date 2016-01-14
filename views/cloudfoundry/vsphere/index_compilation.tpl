{{with .CloudFoundry}}
			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">PaaS {{i18n $.Lang "Compilation"}}</h2>
					</div>
					<div class="form-horizontal">
				  		<div class="panel-body">
							<div class="form-group">
							    <div class="col-sm-offset-2 col-sm-10">
							      	<button class="btn btn-default " id = "config-Compilation"><span class="glyphicon glyphicon-edit"></span> {{i18n $.Lang "Config"}}</button>
							    </div>
								<script type="text/javascript">
									$('#config-Compilation').on('click', function(){
							    		window.location.href = "vspherecloudfoundry?action=config&model=Compilation";
							  		})
								</script>
							</div>
							{{with .VsphereCompilation}}
							<div class="form-group">
								<label class="col-sm-3 control-label">{{i18n $.Lang "Workers"}}</label>
								<div class="col-sm-5">
									<p class="form-control-static">{{.Workers}}</p>
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
									<p class="form-control-static">{{with .VsphereResource}}Ram({{.Ram}}M)-CPU({{.Cpu}})-Disk({{.Disk}}M){{end}}</p>
								</div>
							</div>
							{{end}}
						</div>
					</div>
				</div>
{{end}}