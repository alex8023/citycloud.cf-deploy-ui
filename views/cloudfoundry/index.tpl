<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">PaaS Deployment</h2>
		</div>
		<div class="form-horizontal">
	  		<div class="panel-body">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      	<button class="btn btn-default " id = "config-deploy">Deploy</button>
				    </div>
					<script type="text/javascript">
						$('#config-deploy').on('click', function(){
				    		window.location.href = "cloudfoundry?action=deploy";
				  		})
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
