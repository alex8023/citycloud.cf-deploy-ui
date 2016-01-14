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
		<strong>Oh snap!</strong> {{i18n $.Lang "Change a few things up and try submitting again"}}.
	</div>
	<div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">{{i18n $.Lang "Custom Services Config"}}- <b>{{.Service.Name}}</b></h2>
		</div>
		<div class="panel-body">
			<div class="col-sm-offset-2 col-sm-10">
				<button type="button" class="btn btn-default" data-toggle="modal" data-target="#config-custom-template" data-whatever="add"><span class="glyphicon glyphicon-plus"></span> {{i18n $.Lang "Add-File"}}</button>
				<a class="btn btn-default"  href="templates" title = "back" role="button"><span class="glyphicon glyphicon-backward" ></span> {{i18n $.Lang "Back"}}</a>
			</div>
		    <table class="table table-responsive ">
				<thead>
					<tr>
					<th class ="col-sm-3">{{i18n $.Lang "Name"}}</th>
					<th class ="col-sm-5">{{i18n $.Lang "Description"}}</th>
					<th class ="col-sm-2">{{i18n $.Lang "FileType"}}</th>
					<th class ="col-sm-2">{{i18n $.Lang "Action"}}</th>
					</tr>
				</thead>
				<tbody>
				{{range $index,$element := .Template}}
					<tr>
					<td>{{$element.Name}}</td>
					<td>{{$element.Description}}</td>
					<td>{{$element.FileType}}</td>
					<td>
					<button type="button" class="btn btn-default" data-toggle="modal" data-target="#config-custom-template" data-whatever="{{$element.Id}}" title = "{{i18n $.Lang "Edit"}}"><span class="glyphicon glyphicon-edit"></span> </button>
					<button type="button" class="btn btn-default" data-toggle="modal" data-target="#delete-custom-template" data-whatever="{{$element.Id}}" title = "{{i18n $.Lang "Delete"}}"><span class="glyphicon glyphicon-remove" ></span> </button>
					</td>
					</tr>
				{{end}}
				</tbody>
			</table>
		</div>
	</div>
	
	<div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">{{i18n $.Lang "Custom Service Environments"}}- <b>{{.Service.Name}}</b></h2>
		</div>
		<div class="panel-body">
			<div class="col-sm-offset-2 col-sm-10">
				<button type="button" class="btn btn-default" data-toggle="modal" data-target="#config-custom-component" data-whatever="add"><span class="glyphicon glyphicon-plus"></span> {{i18n $.Lang "Add-Environments"}}</button>
			</div>
		    <table class="table table-responsive ">
				<thead>
					<tr>
					<th class ="col-sm-3">{{i18n $.Lang "Name"}}</th>
					<th class ="col-sm-7">{{i18n $.Lang "Value"}}</th>
					<th class ="col-sm-2">{{i18n $.Lang "Action"}}</th>
					</tr>
				</thead>
				<tbody>
				{{range $index,$element := .Component}}
					<tr>
					<td>{{$element.Name}}</td>
					<td>{{$element.Value}}</td>
					<td>
					<button type="button" class="btn btn-default" data-toggle="modal" data-target="#config-custom-component" data-whatever="{{$element.Id}}" title = "{{i18n $.Lang "Edit"}}"><span class="glyphicon glyphicon-edit"></span> </button>
					<button type="button" class="btn btn-default" data-toggle="modal" data-target="#delete-custom-component" data-whatever="{{$element.Id}}" title = "{{i18n $.Lang "Delete"}}"><span class="glyphicon glyphicon-remove" ></span> </button>
					</td>
					</tr>
				{{end}}
				</tbody>
			</table>
		</div>
	</div>
	{{with .Operation}}
	<div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">{{i18n $.Lang "Custom Services Operation"}}- <b>{{$.Service.Name}}</b></h2>
		</div>
		<div class="form-horizontal">
			<div class="panel-body">
				<div class="col-sm-offset-2 col-sm-10">
					<button type="button" class="btn btn-default" data-toggle="modal" data-target="#config-custom-operation" data-whatever="{{$.Service.Id}}"><span class="glyphicon glyphicon-edit"></span> {{i18n $.Lang "Config"}}</button>
				</div>
				<div class="form-group">
					<div class="form-group">
						<label class="col-sm-3 control-label">{{i18n $.Lang "StartCommand"}}</label>
						<div class="col-sm-7">
							<p class="form-control-static">{{.Start}}</p>
						</div>
					</div>
					<div class="form-group">
						<label class="col-sm-3 control-label">{{i18n $.Lang "RestartCommand"}}</label>
						<div class="col-sm-7">
							<p class="form-control-static">{{.Restart}}</p>
						</div>
					</div>
					<div class="form-group">
						<label class="col-sm-3 control-label">{{i18n $.Lang "StopCommand"}}</label>
						<div class="col-sm-7">
							<p class="form-control-static">{{.Stop}}</p>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
	{{end}}
	
</div>


<!--edit template-->
<div class="modal fade" id="config-custom-template" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
 	<div class="modal-dialog" role="document">
    		<div class="modal-content">
      		<div class="modal-header">
        			<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        			<h4 class="modal-title" id="myModalLabel" title-add = "{{i18n $.Lang "Add Custom Template"}}" title-update ="Update Custom Template"></h4>
      		</div>
			<form class="form-horizontal" method="post" action="templatesdetail" enctype="multipart/form-data">
    				<div class="modal-body">
					<div class="form-horizontal">
			  			<div class="panel-body">
				  			<div class="form-group">
				    				<label for="customTemplateName" class="control-label">{{i18n $.Lang "Name"}}</label>
								<input type="hidden" name = "id" id = "customTemplateId">
								<input type="hidden" name = "sid" value = "{{.ServiceId}}">
					      		<input type="text" class="form-control" id="customTemplateName" placeholder="{{i18n $.Lang "Name"}}" name="name"  required>
				  			</div>
				  			<div class="form-group">
				    				<label for="customTemplateFile" class="control-label">{{i18n $.Lang "TemplateFile"}}</label>
					      		<input type="text" class="form-control" id="customTemplateFile" placeholder="{{i18n $.Lang "TemplateFile"}}" name="templatefile"  required>
				  			</div>
							<div class="form-group hidden" id="templatedetail_targetfile">
				    				<label for="customTargetFile" class="control-label">{{i18n $.Lang "TargetFile"}}</label>
					      		<input type="text" class="form-control" id="customTargetFile" placeholder="{{i18n $.Lang "TargetFile"}}" name="targetfile" >
				  			</div>

							<div class="radio">
								  <label>
								    <input type="radio" name="fileType" id="fileType1" value="War" checked>
								    {{i18n $.Lang "Software Package"}}
								  </label>
							</div>
							<div class="radio">
								  <label>
								    <input type="radio" name="fileType" id="fileType2" value="Template">
								    {{i18n $.Lang "TemplateFileType"}}
								  </label>
							</div>

							<div class="form-group">
			    					<label for="customTemplateDescription" class="control-label">{{i18n $.Lang "Description"}}</label>
				    				<textarea class="form-control" rows="8" id="customTemplateDescription" placeholder="{{i18n $.Lang "Description"}}" name="description"></textarea>
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
<!--delete template-->
<div class="modal fade" id="delete-custom-template" tabindex="-1" role="dialog" aria-labelledby="myModalLabel1">
 	<div class="modal-dialog" role="document">
    		<div class="modal-content">
      		<div class="modal-header">
        			<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        			<h4 class="modal-title" id="myModalLabel1">{{i18n $.Lang "Delete Custom Template"}}</h4>
      		</div>
			<form class="form-horizontal" method="post" action="templatesdetail" enctype="multipart/form-data">
    				<div class="modal-body">
					<div class="form-horizontal">
			  			<div class="panel-body">
				  			<input type="hidden" name = "id" id = "deleteCustomTemplateId">
							<input type="hidden" name = "action" value="delete">
							<input type="hidden" name = "sid" value = "{{.ServiceId}}">
							{{i18n $.Lang "Delete this Custom-Template"}} ?
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
<!--edit component-->
<div class="modal fade" id="config-custom-component" tabindex="-1" role="dialog" aria-labelledby="myModalLabel2">
 	<div class="modal-dialog" role="document">
    		<div class="modal-content">
      		<div class="modal-header">
        			<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        			<h4 class="modal-title" id="myModalLabel1" title-add = "{{i18n $.Lang "Add Custom Service Environment"}}" title-update = "{{i18n $.Lang "Update Custom Service Environment"}}">{{i18n $.Lang "Add Custom Service Environment"}}</h4>
      		</div>
			<form class="form-horizontal" method="post" action="servicecomponent" enctype="multipart/form-data">
    				<div class="modal-body">
					<div class="form-horizontal">
			  			<div class="panel-body">
				  			<div class="form-group">
				    				<label for="customcomponentName" class="control-label">{{i18n $.Lang "Name"}}</label>
								<input type="hidden" name = "id" id = "customComponentId">
								<input type="hidden" name = "sid" value = "{{.ServiceId}}">
					      		<input type="text" class="form-control" id="customComponentName" placeholder="{{i18n $.Lang "Name"}}" name="name"  required>
				  			</div>
				  			<div class="form-group">
				    			<label for="customcomponentValue" class="control-label">{{i18n $.Lang "Value"}}</label>
					      		<input type="text" class="form-control" id="customComponentValue" placeholder="{{i18n $.Lang "Value"}}" name="value"  required>
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

<!--delete component-->
<div class="modal fade" id="delete-custom-component" tabindex="-1" role="dialog" aria-labelledby="myModalLabel3">
 	<div class="modal-dialog" role="document">
    		<div class="modal-content">
      		<div class="modal-header">
        			<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        			<h4 class="modal-title" id="myModalLabel1">{{i18n $.Lang "Delete Custom Service Environment"}}</h4>
      		</div>
			<form class="form-horizontal" method="post" action="servicecomponent" enctype="multipart/form-data">
    				<div class="modal-body">
					<div class="form-horizontal">
			  			<div class="panel-body">
				  			<input type="hidden" name = "id" id = "deleteservicecomponentId">
							<input type="hidden" name = "action" value="delete">
							<input type="hidden" name = "sid" value = "{{.ServiceId}}">
							{{i18n $.Lang "Delete this Custom Service Environment"}} ?
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


<!--edit Operation-->
{{with .Operation}}
<div class="modal fade" id="config-custom-operation" tabindex="-1" role="dialog" aria-labelledby="myModalLabel4">
 	<div class="modal-dialog" role="document">
    		<div class="modal-content">
      		<div class="modal-header">
        			<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        			<h4 class="modal-title" id="myModalLabel1">{{i18n $.Lang "Update Custom Service Operation"}}</h4>
      		</div>
			<form class="form-horizontal" method="post" action="serviceoperation" enctype="multipart/form-data">
    				<div class="modal-body">
					<div class="form-horizontal">
			  			<div class="panel-body">
				  			<div class="form-group">
				    			<label for="customOperationStart" class="control-label">{{i18n $.Lang "StartCommand"}}</label>
								<input type="hidden" name = "id" id = "customOperationId" value="{{.Id}}">
								<input type="hidden" name = "sid" value = "{{$.ServiceId}}">
					      		<input type="text" class="form-control" id="customOperationStart" placeholder="{{i18n $.Lang "StartCommand"}}" name="start" value="{{.Start}}"  required>
				  			</div>
				  			<div class="form-group">
				    			<label for="customOperationRestart" class="control-label">{{i18n $.Lang "RestartCommand"}}</label>
					      		<input type="text" class="form-control" id="customOperationRestart" placeholder="{{i18n $.Lang "RestartCommand"}}" name="restart"  value="{{.Restart}}" required>
				  			</div>
				  			<div class="form-group">
				    			<label for="customOperationStop" class="control-label">{{i18n $.Lang "StopCommand"}}</label>
					      		<input type="text" class="form-control" id="customOperationStop" placeholder="{{i18n $.Lang "StopCommand"}}" name="stop"  value="{{.Stop}}" required>
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
{{end}}
<script src="/static/js/templates.js"></script>
