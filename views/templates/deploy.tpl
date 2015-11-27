<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">Deploying Custom Service</h2>
		</div>
	  		<div class="panel-body">
				<!--   <div class="col-sm-offset-2 col-sm-10">-->
					<div class="btn-group btn-group" role="group">
						<button  class="btn btn-default " id = "deploy-custom-service" data-toggle="tooltip" data-placement="right" title="Deploy The Service" data-service="{{.ServiceId}}"><span class="glyphicon glyphicon-cog"></span> Deploy</button>
						<button  class="btn btn-default " id = "restart-custom-service" data-toggle="tooltip" data-placement="right" title="Restart The Service" data-service="{{.ServiceId}}"><span class="glyphicon glyphicon-refresh"></span> Restart</button>
						<button  class="btn btn-default " id = "stop-custom-service" data-toggle="tooltip" data-placement="right" title="Stop The Service" data-service="{{.ServiceId}}"><span class="glyphicon glyphicon-off"></span> Stop-Service</button>
					</div>
					<div id = "websocketmessage">
					<div><b>{{.Message}}</b></div>
					<div>
			</div>
	</div>
</div>
<input type="hidden" id = "host" value="{{.HOST}}{{.AppName}}{{.WebSocket}}">
<script src="/static/js/servicedeploy.js"></script>