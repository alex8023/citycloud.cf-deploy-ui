<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">VsphereMicroBOSH {{i18n $.Lang "Deployment"}}</h2>
		</div>
		<div class="form-horizontal">
	  		<div class="panel-body">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      	<button class="btn btn-default " id = "config-microbosh"><span class="glyphicon glyphicon-edit"></span> {{i18n $.Lang "Config"}}</button>
					  	<button class="btn btn-default " id = "deploy-microbosh"><span class="glyphicon glyphicon-cog"></span> {{i18n $.Lang "Deploy"}}</button>
				    </div>
					<script type="text/javascript">
						$('#config-microbosh').on('click', function(){
				    		window.location.href = "vspheremicrobosh?action=config";
				  		})
						$('#deploy-microbosh').on('click',function(){
							window.location.href = "vspheremicrobosh?action=deploy";
						})
					</script>
				</div>
				{{with .VsphereMicroBOSH}}
				<div class="form-group">
					<label class="col-sm-3 control-label">{{i18n $.Lang "Deployment Name"}}</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.Name}}</p>
					</div>
				</div>
				{{with .VsphereNetWork}}
				<div class="form-group">
					<label class="col-sm-3 control-label">{{i18n $.Lang "IP"}}</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.Ip}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-3 control-label">{{i18n $.Lang "NetMask"}}</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.NetMask}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-3 control-label">{{i18n $.Lang "GateWay"}}</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.GateWay}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-3 control-label">{{i18n $.Lang "Dns"}}</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.Dns}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-3 control-label">{{i18n $.Lang "VlanName"}}</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.VlanName}}</p>
					</div>
				</div>
				{{end}}
				{{with .VsphereResource}}
				<div class="form-group">
					<label class="col-sm-3 control-label">{{i18n $.Lang "PersistentDisk"}}</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.PersistentDisk}}</p>
					</div>
				</div>
				<div class="form-group">
					<label  class="col-sm-3 control-label">{{i18n $.Lang "Ram"}}</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.Ram}}</p>
					</div>
				</div>
				<div class="form-group">
					<label  class="col-sm-3 control-label">{{i18n $.Lang "Disk"}}</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.Disk}}</p>
					</div>
				</div>
				<div class="form-group">
					<label  class="col-sm-3 control-label">{{i18n $.Lang "Cpu"}}</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.Cpu}}</p>
					</div>
				</div>
				{{end}}
				{{with .VsphereVcenter}}
				<div class="form-group">
					<label  class="col-sm-3 control-label">{{i18n $.Lang "UserName"}}</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.UserName}}</p>
					</div>
				</div>
				<div class="form-group">
					<label  class="col-sm-3 control-label">{{i18n $.Lang "Passwd"}}</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.Passwd}}</p>
					</div>
				</div>
				<div class="form-group">
					<label  class="col-sm-3 control-label">{{i18n $.Lang "Host"}}</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.Host}}</p>
					</div>
				</div>
				<div class="form-group">
					<label  class="col-sm-3 control-label">{{i18n $.Lang "DataCenterName"}}</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.DataCenterName}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-3 control-label">{{i18n $.Lang "VmFolder"}}</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.VmFolder}}</p>
					</div>
				</div>
				<div class="form-group">
					<label for="private_key"  class="col-sm-3 control-label">{{i18n $.Lang "TemplateFolder"}}</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.TemplateFolder}}</p>
					</div>
				</div>
				<div class="form-group">
					<label for="private_key"  class="col-sm-3 control-label">{{i18n $.Lang "DiskPath"}}</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.DiskPath}}</p>
					</div>
				</div>
				<div class="form-group">
					<label for="private_key"  class="col-sm-3 control-label">{{i18n $.Lang "DataStore"}}</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.DataStore}}</p>
					</div>
				</div>
				<div class="form-group">
					<label for="private_key"  class="col-sm-3 control-label">{{i18n $.Lang "ClustersName"}}</label>
					<div class="col-sm-7">
						<p class="form-control-static">{{.ClustersName}}</p>
					</div>
				</div>
				{{end}}
				{{end}}
			</div>
		</div>
	</div>
</div>
