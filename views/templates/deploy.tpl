<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">Deploying Custom Service</h2>
		</div>
	  		<div class="panel-body">
				<!--   <div class="col-sm-offset-2 col-sm-10">-->
				<div class="btn-group btn-group" role="group">
				      	<button class="btn btn-default " id = "config-cloudfoundry" data-toggle="tooltip" data-placement="right" title="配置部署文件"><span class="glyphicon glyphicon-edit"></span> Config</button>
					  	<button class="btn btn-default " id = "deploy-all" data-toggle="tooltip" data-placement="right" title="一键部署PaaS"><span class="glyphicon glyphicon-cog"></span> OneKey</button>
				    </div>
					<div class="btn-group btn-group" role="group">
						<button class="btn btn-default " id = "set-deployment" data-toggle="tooltip" data-placement="right" title="设置部署文件"><span class="glyphicon glyphicon-edit"></span> SetDeployment</button>
						<button class="btn btn-default " id = "upload-release" data-toggle="tooltip" data-placement="right" title="上传软件包"><span class="glyphicon glyphicon-upload"></span> UpLoad-Release</button>
						<button class="btn btn-default " id = "upload-stemcell" data-toggle="tooltip" data-placement="right" title="上传系统镜像"><span class="glyphicon glyphicon-upload"></span> UpLoad-Stemcell</button>
						<button class="btn btn-default " id = "deploy-cloudfoundry" data-toggle="tooltip" data-placement="right" title="部署PaaS"><span class="glyphicon glyphicon-cog"></span> Deploy</button>
						<button class="btn btn-default " id = "login-cloudfoundry" data-toggle="tooltip" data-placement="right" title="登录PaaS"><span class="glyphicon glyphicon-log-in"></span> Login</button>
					</div>
					<div id = "websocketmessage">
					<div><b>{{.Message}}</b></div>
					<div>
			</div>
	</div>
</div>
<input type="hidden" id = "host" value="{{.HOST}}{{.AppName}}{{.WebSocket}}">
<script src="/static/js/servicedeploy.js"></script>