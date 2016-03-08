<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">PaaS {{i18n $.Lang "Deployment"}}</h2>
		</div>
	  		<div class="panel-body">
				<!--   <div class="col-sm-offset-2 col-sm-10">-->
					<div class="btn-group btn-group" role="group">
				      	<button class="btn btn-default " id = "config-cloudfoundry" data-toggle="tooltip" data-placement="right"><span class="glyphicon glyphicon-edit"></span> {{i18n $.Lang "Config"}}</button>
					  	<script type="text/javascript">	
							$('#config-cloudfoundry').on('click', function(){
								{{if ne .IaaSVersion "Vsphere"}}
								window.location.href = "cloudfoundry";
								{{end}}
								{{if eq .IaaSVersion "Vsphere"}}
								window.location.href = "vspherecloudfoundry";
								{{end}}
							})
						</script>
						<button class="btn btn-default " id = "deploy-all" data-toggle="tooltip" data-placement="right" ><span class="glyphicon glyphicon-cog"></span> {{i18n $.Lang "All-In-One"}}</button>
				    </div>
					<div class="btn-group btn-group" role="group">
						<button class="btn btn-default " id = "set-deployment" data-toggle="tooltip" data-placement="right" ><span class="glyphicon glyphicon-edit"></span> {{i18n $.Lang "SetDeployment"}}</button>
						<button class="btn btn-default " id = "upload-release" data-toggle="tooltip" data-placement="right" ><span class="glyphicon glyphicon-upload"></span> {{i18n $.Lang "UpLoad-Release"}}</button>
						<button class="btn btn-default " id = "upload-stemcell" data-toggle="tooltip" data-placement="right"><span class="glyphicon glyphicon-upload"></span> {{i18n $.Lang "UpLoad-Stemcell"}}</button>
						<button class="btn btn-default " id = "deploy-cloudfoundry" data-toggle="tooltip" data-placement="right" ><span class="glyphicon glyphicon-cog"></span> {{i18n $.Lang "Deploy"}}</button>
						<button class="btn btn-default " id = "stats-cloudfoundry" data-toggle="tooltip" data-placement="right" ><span class="glyphicon glyphicon-stats"></span> {{i18n $.Lang "Status"}}</button>
						<!--<button class="btn btn-default " id = "login-cloudfoundry" data-toggle="tooltip" data-placement="right" ><span class="glyphicon glyphicon-log-in"></span> {{i18n $.Lang "LoginPaaS"}}</button>-->
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
<script src="/static/js/cloudfoundry.js"></script>