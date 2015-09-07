<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">MicroBOSH Deployment</h2>
		</div>
	  		<div class="panel-body">
				<!--   <div class="col-sm-offset-2 col-sm-10">-->
				<div class="btn-group btn-group" role="group">
				      	<button class="btn btn-default " id = "config-microbosh" data-toggle="tooltip" data-placement="right" title="配置部署文件">Config</button>
					  	<button class="btn btn-default " id = "deploy-all" data-toggle="tooltip" data-placement="right" title="一键部署MicroBOSH">OneKey</button>
				    </div>
					<div class="btn-group btn-group" role="group">
						<button class="btn btn-default " id = "set-microbosh" data-toggle="tooltip" data-placement="right" title="设置部署文件">SetDeployment</button>
						<button class="btn btn-default " id = "deploy-microbosh" data-toggle="tooltip" data-placement="right" title="部署MicroBOSH">Deploy</button>
						<button class="btn btn-default " id = "redeploy-microbosh" data-toggle="tooltip" data-placement="right" title="重部署MicroBOSH">Re-Deploy</button>
						<button class="btn btn-default " id = "login-microbosh" data-toggle="tooltip" data-placement="right" title="登录MicroBOSH">Login</button>
						<button class="btn btn-default " id = "status-microbosh" data-toggle="tooltip" data-placement="right" title="查看MicroBOSH">Status</button>
					</div>
					<div id = "websocketmessage">
					<div><b>{{.Message}}</b></div>
					<div>
			</div>
	</div>
</div>
<input type="hidden" id = "host" value="{{.HOST}}{{.AppName}}{{.WebSocket}}">
<script src="/static/js/microbosh.js"></script>