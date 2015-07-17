<input type="hidden" id="navfocus" value = "{{.NavBarFocus}}">
<div style="height:20px"></div>
<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">MicroBOSH Deployment</h2>
		</div>
	  		<div class="panel-body">
				    <div class="col-sm-offset-2 col-sm-10">
				      	<button class="btn btn-default " id = "config-microbosh">Config</button>
					  	<button class="btn btn-default " id = "deploy-all">DeployMicroBOSH</button>
						<button class="btn btn-default " id = "set-microbosh">SetDeployment</button>
						<button class="btn btn-default " id = "deploy-microbosh">Deploy</button>
						<button class="btn btn-default " id = "login-microbosh">Login</button>
				    </div>
					<div id = "websocketmessage">
					<div><b>{{.Message}}</b></div>
					<div>
			</div>
	</div>
</div>
<input type="hidden" id = "host" value="{{.HOST}}{{.AppName}}">
<script src="/static/js/microbosh.js"></script>