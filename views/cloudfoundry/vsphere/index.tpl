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
			<h2 class="panel-title">PaaS {{i18n $.Lang "Deployment"}}</h2>
		</div>
		<div class="form-horizontal">
	  		<div class="panel-body">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      	<button class="btn btn-default " id = "config-deploy"><span class="glyphicon glyphicon-cog"></span> {{i18n $.Lang "Deploy"}}</button>
						<button class="btn btn-default " id = "config-more"><span class="glyphicon glyphicon-edit"></span> {{i18n $.Lang "MoreConf"}}</button>
				    </div>
					<script type="text/javascript">
						if ($('#config-deploy')){
							$('#config-deploy').on('click', function(){
				    				window.location.href = "vspherecloudfoundry?action=deploy";
				  			})
						}
						if ($('#config-more')){
							$('#config-more').on('click', function(){
				    				window.location.href = "vspherecloudfoundry?action=moreconf";
				  			})
						}						
					</script>
				</div>
				{{template "cloudfoundry/vsphere/index_properties.tpl" .}}
				{{template "cloudfoundry/vsphere/index_networks.tpl" .}}
				{{template "cloudfoundry/vsphere/index_compilation.tpl" .}}
				{{template "cloudfoundry/vsphere/index_resourcespools.tpl" .}}
				{{template "cloudfoundry/vsphere/index_jobs.tpl" .}}
			</div>
		</div>
	</div>
</div>
