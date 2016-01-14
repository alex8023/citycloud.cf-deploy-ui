<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">{{i18n $.Lang "Deploying Custom Service"}} - {{.Service.Name}}</h2>
		</div>
	  		<div class="panel-body">
				<!--   <div class="col-sm-offset-2 col-sm-10">-->
					<div class="btn-group btn-group" role="group">
						<button  class="btn btn-default " id = "deploy-custom-service" data-toggle="tooltip" data-placement="right" title="Deploy The Service" data-service="{{.ServiceId}}"><span class="glyphicon glyphicon-cog"></span> {{i18n $.Lang "Deploy"}}</button>
					</div>
					<div class="btn-group btn-group" role="group">
						<button  class="btn btn-default " id = "start-custom-service" data-toggle="tooltip" data-placement="right" title="Start The Service" data-service="{{.ServiceId}}"><span class="glyphicon glyphicon-refresh"></span> {{i18n $.Lang "Start"}}</button>
						<button  class="btn btn-default " id = "restart-custom-service" data-toggle="tooltip" data-placement="right" title="Restart The Service" data-service="{{.ServiceId}}"><span class="glyphicon glyphicon-refresh"></span> {{i18n $.Lang "Restart"}}</button>
						<button  class="btn btn-default " id = "stop-custom-service" data-toggle="tooltip" data-placement="right" title="Stop The Service" data-service="{{.ServiceId}}"><span class="glyphicon glyphicon-off"></span> {{i18n $.Lang "Stop"}}</button>
					</div>
					<div class="btn-group btn-group" role="group">
					 	<a class="btn btn-default"  href="templates" title = "back" role="button"><span class="glyphicon glyphicon-backward" ></span> {{i18n $.Lang "Back"}}</a>
					</div>
					<div class="btn-group btn-group" role="group">
						<button class="btn btn-default " id = "clean-console" data-toggle="tooltip" data-placement="right" title="clean console"><span class="glyphicon glyphicon-remove"></span> {{i18n $.Lang "Clean-Console"}}</button>
					</div>
					<div id = "websocketmessage">
					<div><b>{{.Message}}</b></div>
					<div>
			</div>
	</div>
</div>
<input type="hidden" id = "host" value="{{.HOST}}{{.AppName}}{{.WebSocket}}">
<script src="/static/js/servicedeploy.js"></script>