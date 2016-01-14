<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">VsphereMicroBOSH {{i18n $.Lang "Deployment"}}</h2>
		</div>
  		<div class="panel-body">
			<form class="form-horizontal" method="post" action="vspheremicrobosh" enctype="multipart/form-data">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      <button type="submit" class="btn btn-default " data-loading-text="Saving..."><span class="glyphicon glyphicon-floppy-save"></span> {{i18n $.Lang "Save"}}</button>
					<a class="btn btn-default " href="vspheremicrobosh"><span class="glyphicon glyphicon-step-backward"></span> {{i18n $.Lang "Back"}}</a>
				    </div>
			  	</div>
				{{with .VsphereMicroBOSH}}
			  	<div class="form-group">
			    	<label for="name" class="col-sm-3 control-label">{{i18n $.Lang "Deployment Name"}}</label>
				    <div class="col-sm-7">
						<input type="hidden" name="vsphereMicroId" value = "{{.Id}}">
				      	<input type="text" class="form-control" id="name" placeholder="{{i18n $.Lang "Deployment Name"}}" name="name" value = "{{.Name}}" required>
				    </div>
			  	</div>
				{{with .VsphereNetWork}}
			  	<div class="form-group">
			    	<label for="vip" class="col-sm-3 control-label">{{i18n $.Lang "IP"}}</label>
				    <div class="col-sm-7">
						<input type="hidden" name="vsphereNetworkId" value = "{{.Id}}">
				      	<input type="text" class="form-control" id="vip" placeholder="{{i18n $.Lang "IP"}}" name="ip" value = "{{.Ip}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="netMask" class="col-sm-3 control-label">{{i18n $.Lang "NetMask"}}</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="netMask" placeholder="{{i18n $.Lang "NetMask"}}" name="netMask" value = "{{.NetMask}}" required>
			    	</div>
			  	</div>
				<div class="form-group">
			    	<label for="gateWay" class="col-sm-3 control-label">{{i18n $.Lang "GateWay"}}</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="gateWay" placeholder="{{i18n $.Lang "GateWay"}}" name="gateWay" value = "{{.GateWay}}" required>
			    	</div>
			  	</div>
				<div class="form-group">
			    	<label for="dns" class="col-sm-3 control-label">{{i18n $.Lang "Dns"}}</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="dns" placeholder="{{i18n $.Lang "Dns"}}" name="dns" value = "{{.Dns}}" required>
			    	</div>
			  	</div>
				<div class="form-group">
			    	<label for="vlanName" class="col-sm-3 control-label">{{i18n $.Lang "VlanName"}}</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="vlanName" placeholder="{{i18n $.Lang "VlanName"}}" name="vlanName" value = "{{.VlanName}}" required>
			    	</div>
			  	</div>
				{{end}}
				{{with .VsphereResource}}
			  	<div class="form-group">
			    	<label for="persistentDisk" class="col-sm-3 control-label">{{i18n $.Lang "PersistentDisk"}}</label>
				    <div class="col-sm-7">
						<input type="hidden" name="vsphereResourceId" value = "{{.Id}}">
				      	<input type="text" class="form-control" id="persistentDisk" placeholder="16384" name="persistentDisk" value = "{{.PersistentDisk}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="ram" class="col-sm-3 control-label">{{i18n $.Lang "Ram"}}</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="ram" placeholder="{{i18n $.Lang "Ram"}}" name="ram" value = "{{.Ram}}" required>
			    	</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="disk" class="col-sm-3 control-label">{{i18n $.Lang "Disk"}}</label>
			    	<div class="col-sm-7">
			      	<input type="text" class="form-control" id="disk" placeholder="{{i18n $.Lang "Disk"}}" name="disk" value = "{{.Disk}}" required>
			    	</div>
			  	</div>
				<div class="form-group">
			    	<label for="cpu" class="col-sm-3 control-label">{{i18n $.Lang "Cpu"}}</label>
			    	<div class="col-sm-7">
			      	<input type="text" class="form-control" id="cpu" placeholder="{{i18n $.Lang "Cpu"}}" name="cpu" value = "{{.Cpu}}" required>
			    	</div>
			  	</div>
				{{end}}
				{{with .VsphereVcenter}}
			  	<div class="form-group">
			    	<label for="userName" class="col-sm-3 control-label">{{i18n $.Lang "UserName"}}</label>
			    	<div class="col-sm-7">
					<input type="hidden" name="vsphereVcenterId" value = "{{.Id}}">
			      	<input type="text" class="form-control" id="userName" placeholder="{{i18n $.Lang "UserName"}}" name="userName" value = "{{.UserName}}" required>
			    	</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="passwd" class="col-sm-3 control-label">{{i18n $.Lang "Passwd"}}</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="passwd" placeholder="{{i18n $.Lang "Passwd"}}" name="passwd" value = "{{.Passwd}}" required>
			    	</div>
			  	</div>
			  	<div class="form-group">
			    	<label for="host" class="col-sm-3 control-label">{{i18n $.Lang "Host"}}</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="host" placeholder="{{i18n $.Lang "Host"}}" name="host" value = "{{.Host}}" required>
			    	</div>
			  	</div>
				<div class="form-group">
				    <label for="dataCenterName" class="col-sm-3 control-label">{{i18n $.Lang "DataCenterName"}}</label>
				    <div class="col-sm-7">
				      	<input type="text" class="form-control" id="dataCenterName" placeholder="{{i18n $.Lang "DataCenterName"}}" name="dataCenterName" value = "{{.DataCenterName}}" required>
				    </div>
				</div>
			  	<div class="form-group">
			    	<label for="vmFolder" class="col-sm-3 control-label">{{i18n $.Lang "VmFolder"}}</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="vmFolder" placeholder="{{i18n $.Lang "VmFolder"}}" name="vmFolder" value = "{{.VmFolder}}" required>
			    	</div>
			  	</div>
				<div class="form-group">
			    	<label for="templateFolder" class="col-sm-3 control-label">{{i18n $.Lang "TemplateFolder"}}</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="templateFolder" placeholder="{{i18n $.Lang "TemplateFolder"}}" name="templateFolder" value = "{{.TemplateFolder}}" required>
			    	</div>
			  	</div>
				<div class="form-group">
			    	<label for="diskPath" class="col-sm-3 control-label">{{i18n $.Lang "DiskPath"}}</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="diskPath" placeholder="{{i18n $.Lang "DiskPath"}}" name="diskPath" value = "{{.DiskPath}}" required>
			    	</div>
			  	</div>
				<div class="form-group">
			    	<label for="dataStore" class="col-sm-3 control-label">{{i18n $.Lang "DataStore"}}</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="dataStore" placeholder="{{i18n $.Lang "DataStore"}}" name="dataStore" value = "{{.DataStore}}" required>
			    	</div>
			  	</div>
				<div class="form-group">
			    	<label for="clustersName" class="col-sm-3 control-label">{{i18n $.Lang "ClustersName"}}</label>
			    	<div class="col-sm-7">
			      		<input type="text" class="form-control" id="clustersName" placeholder="{{i18n $.Lang "ClustersName"}}" name="clustersName" value = "{{.ClustersName}}" required>
			    	</div>
			  	</div>
				{{end}}	
				{{end}}
			</form>
  		</div>
	</div>
</div>
