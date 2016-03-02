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
		<div class="panel panel-default">
			<div class="panel-heading" >
				<h2 class="panel-title">PaaS {{i18n $.Lang "Ops"}}</h2>
			</div>
			<div class="panel-body">
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
								<button type="button" class="btn btn-default" data-toggle="modal" data-target="#config-custom-template" data-whatever="{{$monitor.AgentId}}" title = "{{i18n $.Lang "Edit"}}"><span class="glyphicon glyphicon-edit"></span> </button>
								<a class="btn btn-default"  href="opsmonitor?agent_id={{$monitor.AgentId}}" title = "{{i18n $.Lang "Config"}}" role="button"><span class="glyphicon glyphicon-edit" ></span> </a>
								<button type="button" class="btn btn-default" data-toggle="modal" data-target="#delete-custom-template" data-whatever="{{$monitor.AgentId}}" title = "{{i18n $.Lang "Delete"}}"><span class="glyphicon glyphicon-remove" ></span> </button>
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
							<button type="button" class="btn btn-default" data-toggle="modal" data-target="#config-custom-template" data-whatever="{{$jobName}}" title = "{{i18n $.Lang "Edit"}}"><span class="glyphicon glyphicon-edit"></span> </button>
							<button type="button" class="btn btn-default" data-toggle="modal" data-target="#delete-custom-template" data-whatever="{{$jobName}}" title = "{{i18n $.Lang "Delete"}}"><span class="glyphicon glyphicon-remove" ></span> </button>
							</td>
							</tr>
						{{end}}

					{{end}}
					</tbody>
				</table>
			</div>
	</div>