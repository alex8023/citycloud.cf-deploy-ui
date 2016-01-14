		{{with .CloudFoundry}}
			{{with .NetWorks}}
				{{with .private}}
			  	<div class="form-group">
			    	<label for="name" class="col-sm-3 control-label">{{i18n $.Lang "Name"}}</label>
				    <div class="col-sm-7">
				      	<input type="text" class="form-control" id="name" placeholder="{{i18n $.Lang "Name"}}" name="private-name" value = "{{.Name}}" required>
						<input type="hidden" name = "private-netWorkName" value = "{{.NetWorkName}}">
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="netId" class="col-sm-3 control-label">{{i18n $.Lang "VlanName"}}</label>
				    <div class="col-sm-7">
				      <input type="text" class="form-control" id="netId" placeholder="{{i18n $.Lang "VlanName"}}" name="private-netId" value = "{{.NetId}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="cidr" class="col-sm-3 control-label">{{i18n $.Lang "Cidr"}}</label>
				    <div class="col-sm-7">
				      <input type="text" class="form-control" id="cidr" placeholder="{{i18n $.Lang "Cidr"}}" name="private-cidr" value = "{{.Cidr}}" required>
				    </div>
			  	</div>
				<div class="form-group">
			    	<label for="netType" class="col-sm-3 control-label">{{i18n $.Lang "GateWay"}}</label>
				    <div class="col-sm-7">
				      <input type="text" class="form-control" id="netType" placeholder="{{i18n $.Lang "GateWay"}}" name="private-netType" value = "{{.NetType}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="dns" class="col-sm-3 control-label">{{i18n $.Lang "Dns"}}</label>
				    <div class="col-sm-7">
				      <input type="text" class="form-control" id="dns" placeholder="{{i18n $.Lang "Dns"}}" name="private-dns" value = "{{.Dns}}" required>
				    </div>
			  	</div>
			  	<div class="form-group">
			    	<label for="reservedIp" class="col-sm-3 control-label">{{i18n $.Lang "ReservedIps"}}</label>
				    <div class="col-sm-7">
				      <input type="text" class="form-control" id="reservedIp" placeholder="{{i18n $.Lang "ReservedIps"}}" name="private-reservedIp" value = "{{.ReservedIp}}" required>
				    </div>
			  	</div>
				<div class="form-group">
			    	<label for="staticIp" class="col-sm-3 control-label">{{i18n $.Lang "StaticIps"}}StaticIp</label>
				    <div class="col-sm-7">
				      <input type="text" class="form-control" id="staticIp" placeholder="{{i18n $.Lang "StaticIps"}}" name="private-staticIp" value = "{{.StaticIp}}" required>
				    </div>
			  	</div>
				{{end}}
			{{end}}
		{{end}}
<input type = "hidden" name="model" value = "{{.Model}}">
<input type = "hidden" name="iaasVsersion" value = "{{.IaaSVersion}}">