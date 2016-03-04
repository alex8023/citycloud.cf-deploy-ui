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
			<h2 class="panel-title">{{i18n $.Lang "vSphere Flavor Config"}}</b></h2>
		</div>
		<div class="panel-body">
			<div class="col-sm-offset-2 col-sm-10">
				<button type="button" class="btn btn-default" data-toggle="modal" data-target="#config-vsphere-resource" data-whatever="add"><span class="glyphicon glyphicon-plus"></span> {{i18n $.Lang "Add Flavor"}}</button>
				<a class="btn btn-default"  href="vspherecloudfoundry" title = "back" role="button"><span class="glyphicon glyphicon-backward" ></span> {{i18n $.Lang "Back"}}</a>
			</div>
		    <table class="table table-responsive ">
				<thead>
					<tr>
					<th class ="col-sm-2">#</th>
					<th class ="col-sm-2">{{i18n $.Lang "Ram"}}(M)</th>
					<th class ="col-sm-2">{{i18n $.Lang "Cpu"}}</th>
					<th class ="col-sm-2">{{i18n $.Lang "Disk"}}(M)</th>
					<th class ="col-sm-2">{{i18n $.Lang "Action"}}</th>
					</tr>
				</thead>
				<tbody>
				{{range $index,$element := .VsphereResource}}
					<tr>
					<td>{{$index}}</td>
					<td>{{$element.Ram}}</td>
					<td>{{$element.Cpu}}</td>
					<td>{{$element.Disk}}</td>
					<td>
					<button type="button" class="btn btn-default" data-toggle="modal" data-target="#config-vsphere-resource" data-whatever="{{$element.Id}}" title = "{{i18n $.Lang "Edit"}}"><span class="glyphicon glyphicon-edit"></span> </button>
					<button type="button" class="btn btn-default" data-toggle="modal" data-target="#delete-vsphere-resource" data-whatever="{{$element.Id}}" title = "{{i18n $.Lang "Delete"}}"><span class="glyphicon glyphicon-remove" ></span> </button>
					</td>
					</tr>
				{{end}}
				</tbody>
			</table>
		</div>
	</div>
</div>

<!--edit template-->
<div class="modal fade" id="config-vsphere-resource" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
 	<div class="modal-dialog" role="document">
    		<div class="modal-content">
      		<div class="modal-header">
        			<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        			<h4 class="modal-title" id="myModalLabel" title-add = "{{i18n $.Lang "Add vSphere Flavor"}}" title-update ="{{i18n $.Lang "Update vSphere Flavor"}}"></h4>
      		</div>
			<form class="form-horizontal" method="post" action="vsphereresource" enctype="multipart/form-data">
    				<div class="modal-body">
					<div class="form-horizontal">
			  			<div class="panel-body">
				  			<div class="form-group">
				    				<label for="ram" class="control-label">{{i18n $.Lang "Ram"}}</label>
								<input type="hidden" name = "id" id = "vsphereresourceId">
								<input type="hidden" name = "action" value="save">
								<input type="hidden" name = "persistentDisk" id = "persistentDisk" >
					      		<input type="text" class="form-control" id="ram" placeholder="{{i18n $.Lang "Ram"}}" name="ram"  required>
				  			</div>
				  			<div class="form-group">
				    			<label for="cpu" class="control-label">{{i18n $.Lang "Cpu"}}</label>
					      		<input type="text" class="form-control" id="cpu" placeholder="{{i18n $.Lang "Cpu"}}" name="cpu"  required>
				  			</div>
				  			<div class="form-group">
				    			<label for="disk" class="control-label">{{i18n $.Lang "Disk"}}</label>
					      		<input type="text" class="form-control" id="disk" placeholder="{{i18n $.Lang "Disk"}}" name="disk"  required>
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
<div class="modal fade" id="delete-vsphere-resource" tabindex="-1" role="dialog" aria-labelledby="myModalLabel1">
 	<div class="modal-dialog" role="document">
    		<div class="modal-content">
      		<div class="modal-header">
        			<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        			<h4 class="modal-title" id="myModalLabel1">{{i18n $.Lang "Delete vSphere Flavor"}}</h4>
      		</div>
			<form class="form-horizontal" method="post" action="vsphereresource" enctype="multipart/form-data">
    				<div class="modal-body">
					<div class="form-horizontal">
			  			<div class="panel-body">
				  			<input type="hidden" name = "id" id = "vsphereresourceId">
							<input type="hidden" name = "action" value="delete">
							{{i18n $.Lang "Delete this vSphere Flavor"}} ?
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
<script src="/static/js/config-vsphereresource.js"></script>
