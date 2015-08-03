<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">PaaS NetWorks</h2>
		</div>
  		<div class="panel-body">
			<form class="form-horizontal" method="post" action="cloudfoundry">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      <button type="submit" class="btn btn-default " data-loading-text="Saving...">Save</button>
				    </div>
			  	</div>
				{{with .CloudFoundry}}
				{{with .NetWorks}}{{with .private}}
			  	<div class="form-group">
			    	<label for="name" class="col-sm-2 control-label">Name</label>
				    <div class="col-sm-10">
				      <input type="text" class="form-control" id="name" placeholder="Name" name="name" value = "{{.Name}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="netType" class="col-sm-2 control-label">NetType</label>
				    <div class="col-sm-10">
				      <input type="text" class="form-control" id="netType" placeholder="NetType" name="netType" value = "{{.NetType}}" readonly required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="netId" class="col-sm-2 control-label">NetId</label>
				    <div class="col-sm-10">
				      <input type="text" class="form-control" id="netId" placeholder="NetId" name="netId" value = "{{.NetId}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="cidr" class="col-sm-2 control-label">Cidr</label>
				    <div class="col-sm-10">
				      <input type="text" class="form-control" id="cidr" placeholder="Cidr" name="cidr" value = "{{.Cidr}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="dns" class="col-sm-2 control-label">Dns</label>
				    <div class="col-sm-10">
				      <input type="text" class="form-control" id="dns" placeholder="Dns" name="dns" value = "{{.Dns}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="reservedIp" class="col-sm-2 control-label">ReservedIp</label>
				    <div class="col-sm-10">
				      <input type="text" class="form-control" id="reservedIp" placeholder="ReservedIp" name="reservedIp" value = "{{.ReservedIp}}" required>
				    </div>
			  	</div>
				<div class="form-group">
			    	<label for="staticIp" class="col-sm-2 control-label">StaticIp</label>
				    <div class="col-sm-10">
				      <input type="text" class="form-control" id="staticIp" placeholder="StaticIp" name="staticIp" value = "{{.StaticIp}}" required>
				    </div>
			  	</div>
				{{end}}
				{{end}}{{end}}
				<input type = "hidden" name="model" value = "{{.Model}}">
			</form>
  		</div>
	</div>
</div>
