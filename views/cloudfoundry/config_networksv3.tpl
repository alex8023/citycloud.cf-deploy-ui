	{{with .CloudFoundry}}
		{{with .NetWorks}}
			{{with .private}}
			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">{{i18n $.Lang "NetWorks-Private"}}</h2>
					</div>
					<div class="form-horizontal">
					  	<div class="panel-body">
						    	<div class="form-group">
						    	<label for="name" class="col-sm-3 control-label">{{i18n $.Lang "Name"}}</label>
							    <div class="col-sm-7">
							      <input type="text" class="form-control" id="name" placeholder="{{i18n $.Lang "Name"}}" name="private-name" value = "{{.Name}}" required>
							    <input type="hidden" name = "private-netWorkName" value = "{{.NetWorkName}}">
								</div>
						  	</div>
						  	<div class="form-group">
						    	<label for="netType" class="col-sm-3 control-label">{{i18n $.Lang "NetType"}}</label>
							    <div class="col-sm-7">
							      <input type="text" class="form-control" id="netType" placeholder="{{i18n $.Lang "NetType"}}" name="private-netType" value = "{{.NetType}}" readonly required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="netId" class="col-sm-3 control-label">{{i18n $.Lang "NetId"}}</label>
							    <div class="col-sm-7">
							      <input type="text" class="form-control" id="netId" placeholder="{{i18n $.Lang "NetId"}}" name="private-netId" value = "{{.NetId}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="cidr" class="col-sm-3 control-label">{{i18n $.Lang "Cidr"}}</label>
							    <div class="col-sm-7">
							      <input type="text" class="form-control" id="cidr" placeholder="{{i18n $.Lang "Cidr"}}" name="private-cidr" value = "{{.Cidr}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="dns" class="col-sm-3 control-label">{{i18n $.Lang "Dns"}}</label>
							    <div class="col-sm-7">
							      <input type="text" class="form-control" id="dns" placeholder="Dns" name="private-dns" value = "{{.Dns}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="reservedIp" class="col-sm-3 control-label">{{i18n $.Lang "ReservedIps"}}</label>
							    <div class="col-sm-7">
							      <input type="text" class="form-control" id="reservedIp" placeholder="{{i18n $.Lang "ReservedIps"}}" name="private-reservedIp" value = "{{.ReservedIp}}" required>
							    </div>
						  	</div>
							<div class="form-group">
						    	<label for="staticIp" class="col-sm-3 control-label">{{i18n $.Lang "StaticIps"}}</label>
							    <div class="col-sm-7">
							      <input type="text" class="form-control" id="staticIp" placeholder="{{i18n $.Lang "StaticIps"}}" name="private-staticIp" value = "{{.StaticIp}}" required>
							    </div>
						  	</div>
						 </div>
					</div>
				</div>
				{{end}}
				{{with .public}}
			    <div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">{{i18n $.Lang "NetWorks-Public"}}</h2>
					</div>
					<div class="form-horizontal">
					  	<div class="panel-body">
						  	<div class="form-group">
						    	<label for="name" class="col-sm-3 control-label">{{i18n $.Lang "Name"}}</label>
							    <div class="col-sm-7">
							      	<input type="text" class="form-control" id="name" placeholder="{{i18n $.Lang "Name"}}" name="public-name" value = "{{.Name}}" required>
									<input type="hidden" name = "public-netWorkName" value = "{{.NetWorkName}}">
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="netType" class="col-sm-3 control-label">{{i18n $.Lang "NetType"}}</label>
							    <div class="col-sm-7">
							      <input type="text" class="form-control" id="netType" placeholder="{{i18n $.Lang "NetType"}}" name="public-netType" value = "{{.NetType}}" readonly required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="netId" class="col-sm-3 control-label">{{i18n $.Lang "NetId"}}</label>
							    <div class="col-sm-7">
							      <input type="text" class="form-control" id="netId" placeholder="{{i18n $.Lang "NetId"}}" name="public-netId" value = "{{.NetId}}" required>
							    </div>
						  	</div>
						  	<div class="form-group">
						    	<label for="cidr" class="col-sm-3 control-label">{{i18n $.Lang "Cidr"}}</label>
							    <div class="col-sm-7">
							      <input type="text" class="form-control" id="cidr" placeholder="{{i18n $.Lang "Cidr"}}" name="public-cidr" value = "{{.Cidr}}" required>
							    </div>
						  	</div>
							<div class="form-group">
						    	<label for="staticIp" class="col-sm-3 control-label">{{i18n $.Lang "StaticIps"}}</label>
							    <div class="col-sm-7">
							      <input type="text" class="form-control" id="staticIp" placeholder="{{i18n $.Lang "StaticIps"}}" name="public-staticIp" value = "{{.StaticIp}}" required>
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

