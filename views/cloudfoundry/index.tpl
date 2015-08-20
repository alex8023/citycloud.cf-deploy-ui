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
			<h2 class="panel-title">PaaS Deployment</h2>
		</div>
		<div class="form-horizontal">
	  		<div class="panel-body">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      	<button class="btn btn-default " id = "config-deploy">Deploy</button>
						<button class="btn btn-default " id = "config-more">MoreConf</button>
				    </div>
					<script type="text/javascript">
						if ($('#config-deploy')){
							$('#config-deploy').on('click', function(){
				    				window.location.href = "cloudfoundry?action=deploy";
				  			})
						}
						if ($('#config-more')){
							$('#config-more').on('click', function(){
				    				window.location.href = "cloudfoundry?action=more";
				  			})
						}						
					</script>
				</div>
				{{template "cloudfoundry/index_properties.tpl" .}}
				{{template "cloudfoundry/index_networks.tpl" .}}
				{{template "cloudfoundry/index_compilation.tpl" .}}
				{{template "cloudfoundry/index_resourcespools.tpl" .}}
				{{template "cloudfoundry/index_jobs.tpl" .}}
			</div>
		</div>
	</div>
</div>
