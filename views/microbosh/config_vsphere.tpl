<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">VsphereMicroBOSH Deployment</h2>
		</div>
  		<div class="panel-body">
			<form class="form-horizontal" method="post" action="vspheremicrobosh" enctype="multipart/form-data">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      <button type="submit" class="btn btn-default " data-loading-text="Saving..."><span class="glyphicon glyphicon-floppy-save"></span> Save</button>
					<a class="btn btn-default " href="vspheremicrobosh"><span class="glyphicon glyphicon-step-backward"></span> Back</a>
				    </div>
			  	</div>
				{{with .VsphereMicroBOSH}}
			  	<div class="form-group">
			    	<label for="name" class="col-sm-3 control-label">Deployment Name</label>
				    <div class="col-sm-7">
						<input type="hidden" name="vsphereMicroId" value = "{{.Id}}">
				      	<input type="text" class="form-control" id="name" placeholder="Deployment Name" name="name" value = "{{.Name}}" required>
				    </div>
			  	</div>
				{{with .VsphereNetWork}}
			  	<div class="form-group">
			    	<label for="vip" class="col-sm-3 control-label">IP</label>
				    <div class="col-sm-7">
						<input type="hidden" name="vsphereNetworkId" value = "{{.Id}}">
				      	<input type="text" class="form-control" id="vip" placeholder="IP" name="ip" value = "{{.Ip}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="netMask" class="col-sm-3 control-label">NetMask</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="netMask" placeholder="NetMask" name="netMask" value = "{{.NetMask}}" required>
			    	</div>
			  	</div>
				<div class="form-group">
			    	<label for="gateWay" class="col-sm-3 control-label">GateWay</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="gateWay" placeholder="GateWay" name="gateWay" value = "{{.GateWay}}" required>
			    	</div>
			  	</div>
				<div class="form-group">
			    	<label for="dns" class="col-sm-3 control-label">Dns</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="dns" placeholder="Dns" name="dns" value = "{{.Dns}}" required>
			    	</div>
			  	</div>
				<div class="form-group">
			    	<label for="vlanName" class="col-sm-3 control-label">VlanName</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="vlanName" placeholder="VlanName" name="vlanName" value = "{{.VlanName}}" required>
			    	</div>
			  	</div>
				{{end}}
				{{with .VsphereResource}}
			  	<div class="form-group">
			    	<label for="persistentDisk" class="col-sm-3 control-label">PersistentDisk</label>
				    <div class="col-sm-7">
						<input type="hidden" name="vsphereResourceId" value = "{{.Id}}">
				      	<input type="text" class="form-control" id="persistentDisk" placeholder="16384" name="persistentDisk" value = "{{.PersistentDisk}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="ram" class="col-sm-3 control-label">Ram</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="ram" placeholder="Ram" name="ram" value = "{{.Ram}}" required>
			    	</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="disk" class="col-sm-3 control-label">Disk</label>
			    	<div class="col-sm-7">
			      	<input type="text" class="form-control" id="disk" placeholder="Disk" name="disk" value = "{{.Disk}}" required>
			    	</div>
			  	</div>
				<div class="form-group">
			    	<label for="cpu" class="col-sm-3 control-label">Cpu</label>
			    	<div class="col-sm-7">
			      	<input type="text" class="form-control" id="cpu" placeholder="Cpu" name="cpu" value = "{{.Cpu}}" required>
			    	</div>
			  	</div>
				{{end}}
				{{with .VsphereVcenter}}
			  	<div class="form-group">
			    	<label for="userName" class="col-sm-3 control-label">UserName</label>
			    	<div class="col-sm-7">
					<input type="hidden" name="vsphereVcenterId" value = "{{.Id}}">
			      	<input type="text" class="form-control" id="userName" placeholder="UserName" name="userName" value = "{{.UserName}}" required>
			    	</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="passwd" class="col-sm-3 control-label">Passwd</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="passwd" placeholder=" Passwd" name="passwd" value = "{{.Passwd}}" required>
			    	</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="host" class="col-sm-3 control-label">Host</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="host" placeholder="Host" name="host" value = "{{.Host}}" required>
			    	</div>
			  	</div>
				<div class="form-group">
				    <label for="dataCenterName" class="col-sm-3 control-label">DataCenterName</label>
				    <div class="col-sm-7">
				      	<input type="text" class="form-control" id="dataCenterName" placeholder="DataCenterName" name="dataCenterName" value = "{{.DataCenterName}}" required>
				    </div>
				</div>
			  	<div class="form-group">
			    	<label for="vmFolder" class="col-sm-3 control-label">VmFolder</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="vmFolder" placeholder="VmFolder" name="vmFolder" value = "{{.VmFolder}}" required>
			    	</div>
			  	</div>
				<div class="form-group">
			    	<label for="templateFolder" class="col-sm-3 control-label">TemplateFolder</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="templateFolder" placeholder="TemplateFolder" name="templateFolder" value = "{{.TemplateFolder}}" required>
			    	</div>
			  	</div>
				<div class="form-group">
			    	<label for="diskPath" class="col-sm-3 control-label">DiskPath</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="diskPath" placeholder="DiskPath" name="diskPath" value = "{{.DiskPath}}" required>
			    	</div>
			  	</div>
				<div class="form-group">
			    	<label for="dataStore" class="col-sm-3 control-label">DataStore</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="dataStore" placeholder="DataStore" name="dataStore" value = "{{.DataStore}}" required>
			    	</div>
			  	</div>
				<div class="form-group">
			    	<label for="clustersName" class="col-sm-3 control-label">ClustersName</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="clustersName" placeholder="ClustersName" name="clustersName" value = "{{.ClustersName}}" required>
			    	</div>
			  	</div>
				{{end}}	
				{{end}}
			</form>
  		</div>
	</div>
</div>
