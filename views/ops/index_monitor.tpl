  	<script src="/static/js/highcharts.js"></script>
	{{if $.AgentId}}
	<script src="/static/js/monitor.js"></script>
	{{end}}
	<div class="container" >
		<div class="alert alert-warning alert-dismissible hide" role="alert" id = "warning-block">
			<button type="button" class="close" data-dismiss="alert" aria-label="Close">x</button>
			<strong>Warning!</strong> Request Error
		</div>
		<div class="alert alert-warning alert-dismissible hide" role="alert" id = "warning-block-request">
			<button type="button" class="close" data-dismiss="alert" aria-label="Close">x</button>
			<strong>Warning!</strong><span id = "warning-block-request-message"></span>
		</div>
		<input type="hidden" id="hidden-message" value="{{.MessageErr}}">
		<div class="panel panel-default">
			<div class="panel-heading" >
				<h2 class="panel-title">PaaS {{i18n $.Lang "Ops"}}</h2>
			</div>

			<div class="panel-body">
			{{if $.AgentId}}
			 	<div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">PaaS {{i18n $.Lang "CPU"}}</h2>
					</div>
					<div class="panel-body">
						<div id="cpu-monitor" style="min-width: 310px; height: 400px; margin: 0 auto"></div>
					</div>
				</div>
			 	<div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">PaaS {{i18n $.Lang "MEM"}}</h2>
					</div>
					<div class="panel-body">
						<div id="mem-monitor" style="min-width: 310px; height: 400px; margin: 0 auto"></div>
					</div>
				</div>
			 	<div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">PaaS {{i18n $.Lang "Disk"}}</h2>
					</div>
					<div class="panel-body">
						<div id="disk-monitor" style="min-width: 310px; height: 400px; margin: 0 auto"></div>
					</div>
				</div>
				<div class="panel panel-default">
					<div class="panel-heading" >
						<h2 class="panel-title">PaaS {{i18n $.Lang "Load Average"}}</h2>
					</div>
					<div class="panel-body">
						<div id="load-monitor" style="min-width: 310px; height: 400px; margin: 0 auto"></div>
					</div>
				</div>
			{{end}}
			</div>
			<input type = "hidden" id ="agent_id" name = "agentId" value = "{{$.AgentId}}">
	</div>