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
	<div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">Custom Services</h2>
		</div>
		<div class="panel-body">
			<div class="col-sm-offset-2 col-sm-10">
				<button type="button" class="btn btn-default" data-toggle="modal" data-target="#myModal"><span class="glyphicon glyphicon-plus"></span> Add Custom-Service</button>
			</div>
		    <table class="table table-responsive ">
				<thead>
					<tr>
					<th class ="col-sm-2">Name</th>
					<th class ="col-sm-8">Description</th>
					<th class ="col-sm-2">Action</th>
					</tr>
				</thead>
				<tbody>
				{{range $index,$element := .Service}}
					<tr>
					<td>{{$element.Name}}</td>
					<td>{{$element.Description}}</td>
					<td><button type="button" class="btn btn-default"><span class="glyphicon glyphicon-edit"></span></button><button type="button" class="btn btn-default"><span class="glyphicon glyphicon-remove"></span></button>
					</td>
					</tr>
				{{end}}
				</tbody>
			</table>
		</div>
	</div>
  </div>
<div class="modal fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
 	<div class="modal-dialog" role="document">
    		<div class="modal-content">
      		<div class="modal-header">
        			<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        			<h4 class="modal-title" id="myModalLabel">Add Custom Service</h4>
      		</div>
			<form class="form-horizontal" method="post" action="templates" enctype="multipart/form-data">
    				<div class="modal-body">
					<div class="form-horizontal">
			  			<div class="panel-body">
				  			<div class="form-group">
				    				<label for="customServiceName" class="col-sm-3 control-label">Name</label>
					   			<div class="col-sm-7">
					      			<input type="text" class="form-control" id="customServiceName" placeholder="Name" name="name"  required>
					    			</div>
				  			</div>
							<div class="form-group">
			    					<label for="customServiceDescription" class="col-sm-3 control-label">Description</label>
				    				<div class="col-sm-7">
									<textarea class="form-control" rows="15" id="customServiceDescription" placeholder="Description" name="description"></textarea>
				    				</div>
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
<script src="/static/js/templates.js"></script>
