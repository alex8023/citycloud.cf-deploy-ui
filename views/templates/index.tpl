  <div class="container">
	<div class="alert alert-warning alert-dismissible hide" role="alert" id = "warning-block">
		<button type="button" class="close" data-dismiss="alert" aria-label="Close">x</button>
		<strong>Warning!</strong> {{.MessageErr}}
	</div>
	<input type="hidden" id="hidden-message" value="{{.MessageErr}}">
	<script type="text/javascript">
		if ($('#hidden-message').val()!=""){
			$('#warning-block').attr("class","alert alert-warning alert-dismissible")
		}
	</script>
	<div class="alert alert-warning alert-dismissible hide" role="alert" id = "warning-block-service">
		<button type="button" class="close" data-dismiss="alert" aria-label="Close">x</button>
		<strong>Oh snap!</strong> Change a few things up and try submitting again.
	</div>
	<div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">{{i18n $.Lang "Custom Services"}}</h2>
		</div>
		<div class="panel-body">
			<div class="col-sm-offset-2 col-sm-10">
				<button type="button" class="btn btn-default" data-toggle="modal" data-target="#config-custom-services" data-whatever="add"><span class="glyphicon glyphicon-plus"></span> {{i18n $.Lang "Add-Service"}}</button>
			</div>
		    <table class="table table-responsive ">
				<thead>
					<tr>
					<th class ="col-sm-2">{{i18n $.Lang "SoftName"}}</th>
					<th class ="col-sm-4">{{i18n $.Lang "Description"}}</th>
					<th class ="col-sm-2">{{i18n $.Lang "Platform"}}</th>
					<th class ="col-sm-4">{{i18n $.Lang "Action"}}</th>
					</tr>
				</thead>
				<tbody>
				{{range $index,$element := .Service}}
					<tr>
					<td>{{$element.Name}}</td>
					<td>{{$element.Description}}</td>
					<td>{{$element.Where}}</td>
					<td>
					<button type="button" class="btn btn-default" data-toggle="modal" data-target="#config-custom-services" data-whatever="{{$element.Id}}" title = "{{i18n $.Lang "Edit"}}"><span class="glyphicon glyphicon-edit"></span> </button>
					<a class="btn btn-default"  href="templatesdetail?serviceId={{$element.Id}}" title = "{{i18n $.Lang "Config"}}" role="button"><span class="glyphicon glyphicon-edit" ></span> </a>
					<button type="button" class="btn btn-default" data-toggle="modal" data-target="#delete-custom-services" data-whatever="{{$element.Id}}" title = "{{i18n $.Lang "Delete"}}"><span class="glyphicon glyphicon-remove" ></span> </button>
					<a class="btn btn-default"  href="servicedeploy?serviceId={{$element.Id}}" title = "{{i18n $.Lang "Deploy"}}" role="button"><span class="glyphicon glyphicon-cog" ></span> </a>
					</td>
					</tr>
				{{end}}
				</tbody>
			</table>
		</div>
	</div>
  </div>
<div class="modal fade" id="config-custom-services" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
 	<div class="modal-dialog" role="document">
    		<div class="modal-content">
      		<div class="modal-header">
        			<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        			<h4 class="modal-title" id="myModalLabel" title-add = "{{i18n $.Lang "Add Custom Service"}}" title-update = "{{i18n $.Lang "Update Custom Service"}}">{{i18n $.Lang "Add Custom Service"}}</h4>
      		</div>
			<form class="form-horizontal" method="post" action="templates" enctype="multipart/form-data">
    				<div class="modal-body">
					<div class="form-horizontal">
			  			<div class="panel-body">
				  			<div class="form-group">
				    				<label for="customServiceName" class="col-sm-2 control-label">{{i18n $.Lang "SoftName"}}</label>
								<input type="hidden" name = "id" id = "customServiceId">
								 <div class="col-sm-10">
					      			<input type="text" class="form-control" id="customServiceName" placeholder="{{i18n $.Lang "SoftName"}}" name="name"  required>
				  				</div>
							</div>
							<div class="form-group">
			    					<label for="customServiceDescription" class="col-sm-2 control-label">{{i18n $.Lang "Description"}}</label>
				    				<div class="col-sm-10">
									<textarea class="form-control" rows="5" id="customServiceDescription" placeholder="{{i18n $.Lang "Description"}}" name="description"></textarea>
			  					</div>
							</div>
							<div class="radio">
								  <label>
								    <input type="radio" name="where" id="customServiceWhere1" value="PaaS" checked>
								    {{i18n $.Lang "Services will be deployed to CCI-PaaS"}}
								  </label>
							</div>
							<div class="radio">
								  <label>
								    <input type="radio" name="where" id="customServiceWhere2" value="Vms">
								    {{i18n $.Lang "Services will be deployed to a virtual machine"}}
								  </label>
							</div>
						</div>
						<div class="panel-body" id="deploy-onpaas">
							<div class="form-group">
				    				<label for="customServiceApi" class="col-sm-2 control-label">{{i18n $.Lang "Api"}}</label>
								<div class="col-sm-10">
									<input type="hidden" name = "paasid" id = "customServicePaaSId">
									<input type="text" class="form-control" id="customServiceApi" placeholder="{{i18n $.Lang "Api"}}" name="Api"  required>
				  				</div>
							</div>
							<div class="form-group">
				    				<label for="customServiceUser" class="col-sm-2 control-label">{{i18n $.Lang "UserName"}}</label>
					      		<div class="col-sm-10">
									<input type="text" class="form-control" id="customServiceUser" placeholder="{{i18n $.Lang "UserName"}}" name="UserName"  required>
				  				</div>
							</div>
							<div class="form-group">
				    				<label for="customServicePasswd" class="col-sm-2 control-label">{{i18n $.Lang "Passwd"}}</label>
					      		<div class="col-sm-10">
									<input type="password" class="form-control" id="customServicePasswd" placeholder="{{i18n $.Lang "Passwd"}}" name="Passwd"  required>
				  				</div>
							</div>
							<div class="form-group">
				    				<label for="customServiceOrg" class="col-sm-2 control-label">{{i18n $.Lang "Org"}}</label>
					      		<div class="col-sm-10">
									<input type="text" class="form-control" id="customServiceOrg" placeholder="{{i18n $.Lang "Org"}}" name="Org"  required>
				  				</div>
							</div>	
							<div class="form-group">
				    				<label for="customServiceSpace" class="col-sm-2 control-label">{{i18n $.Lang "Space"}}</label>
					      		<div class="col-sm-10">
									<input type="text" class="form-control" id="customServiceSpace" placeholder="{{i18n $.Lang "Space"}}" name="Space"  required>
				  				</div>
							</div>														
						</div>
						<div class="panel-body hidden" id="deploy-onvms">
							<div class="form-group">
				    				<label for="customServiceIp" class="col-sm-2 control-label">{{i18n $.Lang "HostIp"}}</label>
								<div class="col-sm-10">
									<input type="hidden" name = "vmid" id = "customServiceVmId">
									<input type="text" class="form-control" id="customServiceIp" placeholder="{{i18n $.Lang "HostIp"}}" name="ip"  >
				  				</div>
							</div>
							<div class="form-group">
				    				<label for="customServiceVmUser" class="col-sm-2 control-label">{{i18n $.Lang "UserName"}}</label>
					      		<div class="col-sm-10">
									<input type="text" class="form-control" id="customServiceVmUser" placeholder="{{i18n $.Lang "UserName"}}" name="VmUserName"  >
				  				</div>
							</div>
							<div class="form-group">
				    				<label for="customServiceVmPasswd" class="col-sm-2 control-label">{{i18n $.Lang "Passwd"}}</label>
					      		<div class="col-sm-10">
									<input type="password" class="form-control" id="customServiceVmPasswd" placeholder="{{i18n $.Lang "Passwd"}}" name="VmPasswd"  >
				  				</div>
							</div>
						</div>
					</div>
      			</div>
		      	<div class="modal-footer">
		        		<button type="button" class="btn btn-default" data-dismiss="modal"><span class="glyphicon glyphicon-remove"></span> {{i18n $.Lang "Close"}}</button>
		        		<button type="submit" class="btn btn-primary"><span class="glyphicon glyphicon-save"></span>  {{i18n $.Lang "Save"}}</button>
		      	</div>
			</form>
    		</div>
  	</div>
</div>
<div class="modal fade" id="delete-custom-services" tabindex="-1" role="dialog" aria-labelledby="myModalLabel1">
 	<div class="modal-dialog" role="document">
    		<div class="modal-content">
      		<div class="modal-header">
        			<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        			<h4 class="modal-title" id="myModalLabel1" title ="{{i18n $.Lang "Delete Custom Service"}}" >{{i18n $.Lang "Delete Custom Service"}}</h4>
      		</div>
			<form class="form-horizontal" method="post" action="templates" enctype="multipart/form-data">
    				<div class="modal-body">
					<div class="form-horizontal">
			  			<div class="panel-body">
				  			<input type="hidden" name = "id" id = "deleteCustomServiceId">
							<input type="hidden" name = "action" value="delete">
							{{i18n $.Lang "Delete this Custom Service"}} ?
						</div>
					</div>
      			</div>
		      	<div class="modal-footer">
		        		<button type="button" class="btn btn-default" data-dismiss="modal"><span class="glyphicon glyphicon-remove"></span> {{i18n $.Lang "Close"}}</button>
		        		<button type="submit" class="btn btn-primary"><span class="glyphicon glyphicon-remove"></span>  {{i18n $.Lang "Delete"}}</button>
		      	</div>
			</form>
    		</div>
  	</div>
</div>
<script src="/static/js/templates.js"></script>
