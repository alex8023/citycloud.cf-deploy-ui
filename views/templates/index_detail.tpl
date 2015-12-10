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
			<h2 class="panel-title">Custom Services - <b>{{.Service.Name}}</b></h2>
		</div>
		<div class="panel-body">
			<div class="col-sm-offset-2 col-sm-10">
				<button type="button" class="btn btn-default" data-toggle="modal" data-target="#config-custom-template" data-whatever="add"><span class="glyphicon glyphicon-plus"></span> Add-File</button>
				<a class="btn btn-default"  href="templates" title = "back" role="button"><span class="glyphicon glyphicon-backward" ></span> Back</a>
			</div>
		    <table class="table table-responsive ">
				<thead>
					<tr>
					<th class ="col-sm-3">Name</th>
					<th class ="col-sm-5">Description</th>
					<th class ="col-sm-2">FileType</th>
					<th class ="col-sm-2">Action</th>
					</tr>
				</thead>
				<tbody>
				{{range $index,$element := .Template}}
					<tr>
					<td>{{$element.Name}}</td>
					<td>{{$element.Description}}</td>
					<td>{{$element.FileType}}</td>
					<td>
					<button type="button" class="btn btn-default" data-toggle="modal" data-target="#config-custom-template" data-whatever="{{$element.Id}}" title = "edit"><span class="glyphicon glyphicon-edit"></span> </button>
					<button type="button" class="btn btn-default" data-toggle="modal" data-target="#delete-custom-template" data-whatever="{{$element.Id}}" title = "delete"><span class="glyphicon glyphicon-remove" ></span> </button>
					</td>
					</tr>
				{{end}}
				</tbody>
			</table>
		</div>
	</div>
	
	<div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">Custom Services Component- <b>{{.Service.Name}}</b></h2>
		</div>
		<div class="panel-body">
			<div class="col-sm-offset-2 col-sm-10">
				<button type="button" class="btn btn-default" data-toggle="modal" data-target="#config-custom-component" data-whatever="add"><span class="glyphicon glyphicon-plus"></span> Add-Config</button>
			</div>
		    <table class="table table-responsive ">
				<thead>
					<tr>
					<th class ="col-sm-3">Name</th>
					<th class ="col-sm-7">Value</th>
					<th class ="col-sm-2">Action</th>
					</tr>
				</thead>
				<tbody>
				{{range $index,$element := .Component}}
					<tr>
					<td>{{$element.Name}}</td>
					<td>{{$element.Value}}</td>
					<td>
					<button type="button" class="btn btn-default" data-toggle="modal" data-target="#config-custom-component" data-whatever="{{$element.Id}}" title = "edit"><span class="glyphicon glyphicon-edit"></span> </button>
					<button type="button" class="btn btn-default" data-toggle="modal" data-target="#delete-custom-component" data-whatever="{{$element.Id}}" title = "delete"><span class="glyphicon glyphicon-remove" ></span> </button>
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
			<h2 class="panel-title">Custom Services Operation- <b>{{$.Service.Name}}</b></h2>
		</div>
		<div class="form-horizontal">
			<div class="panel-body">
				<div class="col-sm-offset-2 col-sm-10">
					<button type="button" class="btn btn-default" data-toggle="modal" data-target="#config-custom-operation" data-whatever="{{$.Service.Id}}"><span class="glyphicon glyphicon-edit"></span> Config</button>
				</div>
				<div class="form-group">
					<div class="form-group">
						<label class="col-sm-3 control-label">StartCommand</label>
						<div class="col-sm-7">
							<p class="form-control-static">{{.Start}}</p>
						</div>
					</div>
					<div class="form-group">
						<label class="col-sm-3 control-label">RestartCommand</label>
						<div class="col-sm-7">
							<p class="form-control-static">{{.Restart}}</p>
						</div>
					</div>
					<div class="form-group">
						<label class="col-sm-3 control-label">StopCommand</label>
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
        			<h4 class="modal-title" id="myModalLabel">Add Custom Template</h4>
      		</div>
			<form class="form-horizontal" method="post" action="templatesdetail" enctype="multipart/form-data">
    				<div class="modal-body">
					<div class="form-horizontal">
			  			<div class="panel-body">
				  			<div class="form-group">
				    				<label for="customTemplateName" class="control-label">Name</label>
								<input type="hidden" name = "id" id = "customTemplateId">
								<input type="hidden" name = "sid" value = "{{.ServiceId}}">
					      		<input type="text" class="form-control" id="customTemplateName" placeholder="Name" name="name"  required>
				  			</div>
				  			<div class="form-group">
				    				<label for="customTemplateFile" class="control-label">TemplateFile</label>
					      		<input type="text" class="form-control" id="customTemplateFile" placeholder="TemplateFile" name="templatefile"  required>
				  			</div>
							<div class="form-group hidden" id="templatedetail_targetfile">
				    				<label for="customTargetFile" class="control-label">TargetFile</label>
					      		<input type="text" class="form-control" id="customTargetFile" placeholder="TargetFile" name="targetfile" >
				  			</div>

							<div class="radio">
								  <label>
								    <input type="radio" name="fileType" id="fileType1" value="War" checked>
								    Software package (e.g cfWeb.war)
								  </label>
							</div>
							<div class="radio">
								  <label>
								    <input type="radio" name="fileType" id="fileType2" value="Template">
								    TemplateFile (e.g spring-datasource.xml)
								  </label>
							</div>

							<div class="form-group">
			    					<label for="customTemplateDescription" class="control-label">Description</label>
				    				<textarea class="form-control" rows="8" id="customTemplateDescription" placeholder="Description" name="description"></textarea>
			  				</div>
						</div>
					</div>
      			</div>
		      	<div class="modal-footer">
		        		<button type="button" class="btn btn-default" data-dismiss="modal"><span class="glyphicon glyphicon-remove"></span> Close</button>
		        		<button type="submit" class="btn btn-primary"><span class="glyphicon glyphicon-save"></span>  Save</button>
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
        			<h4 class="modal-title" id="myModalLabel1">Delete Custom Service</h4>
      		</div>
			<form class="form-horizontal" method="post" action="templatesdetail" enctype="multipart/form-data">
    				<div class="modal-body">
					<div class="form-horizontal">
			  			<div class="panel-body">
				  			<input type="hidden" name = "id" id = "deleteCustomTemplateId">
							<input type="hidden" name = "action" value="delete">
							<input type="hidden" name = "sid" value = "{{.ServiceId}}">
							Delete this Custom-Template ?
						</div>
					</div>
      			</div>
		      	<div class="modal-footer">
		        		<button type="button" class="btn btn-default" data-dismiss="modal"><span class="glyphicon glyphicon-remove"></span> Close</button>
		        		<button type="submit" class="btn btn-primary"><span class="glyphicon glyphicon-remove"></span>  Delete</button>
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
        			<h4 class="modal-title" id="myModalLabel1">Add Custom Service Component</h4>
      		</div>
			<form class="form-horizontal" method="post" action="servicecomponent" enctype="multipart/form-data">
    				<div class="modal-body">
					<div class="form-horizontal">
			  			<div class="panel-body">
				  			<div class="form-group">
				    				<label for="customcomponentName" class="control-label">Name</label>
								<input type="hidden" name = "id" id = "customComponentId">
								<input type="hidden" name = "sid" value = "{{.ServiceId}}">
					      		<input type="text" class="form-control" id="customComponentName" placeholder="Name" name="name"  required>
				  			</div>
				  			<div class="form-group">
				    			<label for="customcomponentValue" class="control-label">Value</label>
					      		<input type="text" class="form-control" id="customComponentValue" placeholder="Value" name="value"  required>
				  			</div>
						</div>
					</div>
      			</div>
		      	<div class="modal-footer">
		        		<button type="button" class="btn btn-default" data-dismiss="modal"><span class="glyphicon glyphicon-remove"></span> Close</button>
		        		<button type="submit" class="btn btn-primary"><span class="glyphicon glyphicon-save"></span>  Save</button>
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
        			<h4 class="modal-title" id="myModalLabel1">Delete Custom Service Component</h4>
      		</div>
			<form class="form-horizontal" method="post" action="servicecomponent" enctype="multipart/form-data">
    				<div class="modal-body">
					<div class="form-horizontal">
			  			<div class="panel-body">
				  			<input type="hidden" name = "id" id = "deleteservicecomponentId">
							<input type="hidden" name = "action" value="delete">
							<input type="hidden" name = "sid" value = "{{.ServiceId}}">
							Delete this Custom Service Component ?
						</div>
					</div>
      			</div>
		      	<div class="modal-footer">
		        		<button type="button" class="btn btn-default" data-dismiss="modal"><span class="glyphicon glyphicon-remove"></span> Close</button>
		        		<button type="submit" class="btn btn-primary"><span class="glyphicon glyphicon-remove"></span>  Delete</button>
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
        			<h4 class="modal-title" id="myModalLabel1">Update Custom Service Operation</h4>
      		</div>
			<form class="form-horizontal" method="post" action="serviceoperation" enctype="multipart/form-data">
    				<div class="modal-body">
					<div class="form-horizontal">
			  			<div class="panel-body">
				  			<div class="form-group">
				    			<label for="customOperationStart" class="control-label">StartCommand</label>
								<input type="hidden" name = "id" id = "customOperationId" value="{{.Id}}">
								<input type="hidden" name = "sid" value = "{{$.ServiceId}}">
					      		<input type="text" class="form-control" id="customOperationStart" placeholder="StartCommand" name="start" value="{{.Start}}"  required>
				  			</div>
				  			<div class="form-group">
				    			<label for="customOperationRestart" class="control-label">RestartCommand</label>
					      		<input type="text" class="form-control" id="customOperationRestart" placeholder="RestartCommand" name="restart"  value="{{.Restart}}" required>
				  			</div>
				  			<div class="form-group">
				    			<label for="customOperationStop" class="control-label">StopCommand</label>
					      		<input type="text" class="form-control" id="customOperationStop" placeholder="StopCommand" name="stop"  value="{{.Stop}}" required>
				  			</div>
						</div>
					</div>
      			</div>
		      	<div class="modal-footer">
		        		<button type="button" class="btn btn-default" data-dismiss="modal"><span class="glyphicon glyphicon-remove"></span> Close</button>
		        		<button type="submit" class="btn btn-primary"><span class="glyphicon glyphicon-save"></span>  Save</button>
		      	</div>
			</form>
    		</div>
  	</div>
</div>
{{end}}
<script src="/static/js/templates.js"></script>
