<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">{{if eq .IaaSVersion "Vsphere"}}Vsphere{{end}}MicroBOSH Deployment</h2>
		</div>
	  		<div class="panel-body">
				<!--   <div class="col-sm-offset-2 col-sm-10">-->
				<div class="btn-group btn-group" role="group">
				      	<button class="btn btn-default " id = "config-microbosh" data-toggle="tooltip" data-placement="right" title="配置部署文件"><span class="glyphicon glyphicon-edit"></span> Config</button>
						<script type="text/javascript">
							$('#config-microbosh').on('click', function(){
								{{if ne .IaaSVersion "Vsphere"}}
								window.location.href = "microbosh?action=config";
								{{end}}
								{{if eq .IaaSVersion "Vsphere"}}
								window.location.href = "vspheremicrobosh?action=config";
								console.log(window.location.href)
								{{end}}
							})
						</script>
					  	<button class="btn btn-default " id = "deploy-all" data-toggle="tooltip" data-placement="right" title="一键部署MicroBOSH"><span class="glyphicon glyphicon-cog"></span> All-In-One</button>
				    </div>
					<div class="btn-group btn-group" role="group">
						<button class="btn btn-default " id = "set-microbosh" data-toggle="tooltip" data-placement="right" title="设置部署文件"><span class="glyphicon glyphicon-edit"></span> SetDeployment</button>
						<button class="btn btn-default " id = "deploy-microbosh" data-toggle="tooltip" data-placement="right" title="部署MicroBOSH"><span class="glyphicon glyphicon-cog"></span> Deploy</button>
						<button class="btn btn-default " id = "redeploy-microbosh" data-toggle="tooltip" data-placement="right" title="重部署MicroBOSH"><span class="glyphicon glyphicon-cog"></span> Re-Deploy</button>
						<button class="btn btn-default " id = "login-microbosh" data-toggle="tooltip" data-placement="right" title="登录MicroBOSH"><span class="glyphicon glyphicon-log-in"></span> Login</button>
						<button class="btn btn-default " id = "status-microbosh" data-toggle="tooltip" data-placement="right" title="查看MicroBOSH"><span class="glyphicon glyphicon-stats"></span> Status</button>
					</div>
					<div class="btn-group btn-group" role="group">
						<button class="btn btn-default " id = "clean-console" data-toggle="tooltip" data-placement="right" title="clean console"><span class="glyphicon glyphicon-remove"></span> Clean-Console</button>
					</div>
					<div id = "websocketmessage">
					<div><b>{{.Message}}</b></div>
					<div>
			</div>
	</div>
</div>
<input type="hidden" id = "host" value="{{.HOST}}{{.AppName}}{{.WebSocket}}">
<script src="/static/js/microbosh.js"></script>