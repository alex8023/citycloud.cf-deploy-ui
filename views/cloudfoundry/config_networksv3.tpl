	{{with .CloudFoundry}}
		{{with .NetWorks}}
			{{with .private}}
			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">NetWork-Private</h2>
					</div>
					<div class="form-horizontal">
					  	<div class="panel-body">
						    	<div class="form-group">
						    	<label for="name" class="col-sm-2 control-label">Name</label>
							    <div class="col-sm-10">
							      <input type="text" class="form-control" id="name" placeholder="Name" name="private-name" value = "{{.Name}}" required>
							    <input type="hidden" name = "private-netWorkName" value = "{{.NetWorkName}}">
								</div>
						  	</div>
						  	<div class="form-group">
						    	<label for="netType" class="col-sm-2 control-label">NetType</label>
							    <div class="col-sm-10">
							      <input type="text" class="form-control" id="netType" placeholder="NetType" name="private-netType" value = "{{.NetType}}" readonly required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="netId" class="col-sm-2 control-label">NetId</label>
							    <div class="col-sm-10">
							      <input type="text" class="form-control" id="netId" placeholder="NetId" name="private-netId" value = "{{.NetId}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="cidr" class="col-sm-2 control-label">Cidr</label>
							    <div class="col-sm-10">
							      <input type="text" class="form-control" id="cidr" placeholder="Cidr" name="private-cidr" value = "{{.Cidr}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="dns" class="col-sm-2 control-label">Dns</label>
							    <div class="col-sm-10">
							      <input type="text" class="form-control" id="dns" placeholder="Dns" name="private-dns" value = "{{.Dns}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="reservedIp" class="col-sm-2 control-label">ReservedIp</label>
							    <div class="col-sm-10">
							      <input type="text" class="form-control" id="reservedIp" placeholder="ReservedIp" name="private-reservedIp" value = "{{.ReservedIp}}" required>
							    </div>
						  	</div>
							<div class="form-group">
						    	<label for="staticIp" class="col-sm-2 control-label">StaticIp</label>
							    <div class="col-sm-10">
							      <input type="text" class="form-control" id="staticIp" placeholder="StaticIp" name="private-staticIp" value = "{{.StaticIp}}" required>
							    </div>
						  	</div>
						 </div>
					</div>
				</div>
				{{end}}
				{{with .public}}
			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">NetWork-Public</h2>
					</div>
					<div class="form-horizontal">
					  	<div class="panel-body">
						  	<div class="form-group">
						    	<label for="name" class="col-sm-2 control-label">Name</label>
							    <div class="col-sm-10">
							      	<input type="text" class="form-control" id="name" placeholder="Name" name="public-name" value = "{{.Name}}" required>
									<input type="hidden" name = "public-netWorkName" value = "{{.NetWorkName}}">
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="netType" class="col-sm-2 control-label">NetType</label>
							    <div class="col-sm-10">
							      <input type="text" class="form-control" id="netType" placeholder="NetType" name="public-netType" value = "{{.NetType}}" readonly required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="netId" class="col-sm-2 control-label">NetId</label>
							    <div class="col-sm-10">
							      <input type="text" class="form-control" id="netId" placeholder="NetId" name="public-netId" value = "{{.NetId}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="cidr" class="col-sm-2 control-label">Cidr</label>
							    <div class="col-sm-10">
							      <input type="text" class="form-control" id="cidr" placeholder="Cidr" name="public-cidr" value = "{{.Cidr}}" required>
							    </div>
						  	</div>
							<div class="form-group">
						    	<label for="staticIp" class="col-sm-2 control-label">StaticIp</label>
							    <div class="col-sm-10">
							      <input type="text" class="form-control" id="staticIp" placeholder="StaticIp" name="public-staticIp" value = "{{.StaticIp}}" required>
							    </div>
						  	</div>
						 </div>
					</div>
				</div>
				{{end}}
			{{end}}
		{{end}}
		<input type = "hidden" name="model" value = "{{.Model}}">
		<input type = "hidden" name="iaasVsersion" value = "{{.IaaSVersion}}">

