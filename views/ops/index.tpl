	<div class="container" >
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
		<div class="alert alert-warning alert-dismissible hide" role="alert" id = "warning-block-agenterr">
			<button type="button" class="close" data-dismiss="alert" aria-label="Close">x</button>
			<strong>Warning!</strong> {{.AgentError}}
		</div>
		<input type="hidden" id="hidden-message-agent" value="{{.AgentError}}">
		<script type="text/javascript">
			if ($('#hidden-message-agent').val()!=""){
				$('#warning-block-agenterr').attr("class","alert alert-warning alert-dismissible")
			}
		</script>
		<div class="panel panel-default">
			<div class="panel-heading" >
				<h2 class="panel-title">PaaS {{i18n $.Lang "Ops"}}</h2>
			</div>
			<div class="panel-body">
				<div class="col-sm-offset-2 col-sm-10">
					<button type="button" class="btn btn-default" data-toggle="modal" data-target="#config-agent" data-whatever="add"><span class="glyphicon glyphicon-plus"></span> {{i18n $.Lang "Add Agent"}}</button>
				</div>
			 	<table class="table table-responsive ">
					<thead>
						<tr>
						<th class ="col-sm-3">{{i18n $.Lang "JobName"}}</th>
						<th class ="col-sm-3">{{i18n $.Lang "Alias"}}</th>
						<th class ="col-sm-2">{{i18n $.Lang "Index"}}</th>
						<th class ="col-sm-2">{{i18n $.Lang "Stats"}}</th>
						<th class ="col-sm-5">{{i18n $.Lang "Action"}}</th>
						</tr>
					</thead>
					<tbody>
					{{range $jobName,$jobMonitorStruct := .JobMonitorStruct}}
						{{if $jobMonitorStruct.Monitors}}
							{{range $index,$monitor := $jobMonitorStruct.Monitors}}
								<tr>
								<td>{{$jobName}}</td>
								<td>{{$jobMonitorStruct.CloudFoundryJobs.Name}}</td>
								<td>{{$monitor.Index}}</td>
								<td>{{$monitor.JobState}}</td>
								<td>
								<button type="button" class="btn btn-default" data-toggle="modal" data-target="#config-agent" data-whatever="{{$monitor.AgentId}}" title = "{{i18n $.Lang "Edit"}}"><span class="glyphicon glyphicon-edit"></span> </button>
								<a class="btn btn-default"  href="opsmonitor?agent_id={{$monitor.AgentId}}" title = "{{i18n $.Lang "Stats"}}" role="button"><span class="glyphicon glyphicon-stats" ></span> </a>
								<button type="button" class="btn btn-default" data-toggle="modal" data-target="#delete-agent" data-whatever="{{$monitor.AgentId}}" title = "{{i18n $.Lang "Delete"}}"><span class="glyphicon glyphicon-remove" ></span> </button>
								</td>
								</tr>
							{{end}}
						{{else}}
							<tr>
							<td>{{$jobName}}</td>
							<td>{{$jobMonitorStruct.CloudFoundryJobs.Name}}</td>
							<td>0</td>
							<td>Unknow</td>
							<td>
							<button type="button" class="btn btn-default" data-toggle="modal" data-target="#config-agent" data-whatever="add" title = "{{i18n $.Lang "Edit"}}"><span class="glyphicon glyphicon-edit"></span> </button>
							</td>
							</tr>
						{{end}}

					{{end}}
					</tbody>
				</table>
			</div>
	</div>
<!--config-agent-->
<div class="modal fade" id="config-agent" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
 	<div class="modal-dialog" role="document">
    		<div class="modal-content">
      		<div class="modal-header">
        			<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        			<h4 class="modal-title" id="myModalLabel" title-add = "{{i18n $.Lang "Add Agent"}}" title-update = "{{i18n $.Lang "Update Agent"}}" >{{i18n $.Lang "Add Agent"}}</h4>
      		</div>
			<form class="form-horizontal" method="post" action="ops" enctype="multipart/form-data">
    			<div class="modal-body">
					<div class="form-horizontal">
			  			<div class="panel-body">
				    		<label for="ops_agent_id" class="col-sm-3 control-label">{{i18n $.Lang "Agent"}}</label>
					      	<div class="col-sm-9">
								<input type="text" class="form-control" id="ops_agent_id" placeholder="{{i18n $.Lang "Agent"}}" name="agent_id"  >
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
<!--config-agent-->
<!--delete-agent-->
<div class="modal fade" id="delete-agent" tabindex="-1" role="dialog" aria-labelledby="myModalLabel1">
 	<div class="modal-dialog" role="document">
    		<div class="modal-content">
      		<div class="modal-header">
        			<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        			<h4 class="modal-title" id="myModalLabel1">{{i18n $.Lang "Delete Agent"}}</h4>
      		</div>
			<form class="form-horizontal" method="post" action="ops" enctype="multipart/form-data">
    				<div class="modal-body">
					<div class="form-horizontal">
			  			<div class="panel-body">
				  			<input type="hidden" name = "agent_id" id = "deleteAgentId">
							<input type="hidden" name = "action" value="delete">
							{{i18n $.Lang "Delete this Agent"}} ?
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
<!--delete-agent-->

<script src="/static/js/ops.js"></script>