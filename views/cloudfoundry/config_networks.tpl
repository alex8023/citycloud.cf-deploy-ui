<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">PaaS {{i18n $.Lang "NetWorks"}} </h2>
		</div>
  		<div class="panel-body">
			<form class="form-horizontal" method="post" action="cloudfoundry">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      <button type="submit" class="btn btn-default " data-loading-text="Saving..."><span class="glyphicon glyphicon-floppy-save"></span> {{i18n $.Lang "Save"}}</button>
					<a class="btn btn-default " href="cloudfoundry"><span class="glyphicon glyphicon-step-backward"></span> {{i18n $.Lang "Back"}}</a>
				    </div>
			  	</div>
			{{if eq .IaaSVersion .DefaultVersion}}
				{{template "cloudfoundry/config_networksv3.tpl" .}}
			{{else}}
				{{template "cloudfoundry/config_networksv2.tpl" .}}
			{{end}}
			</form>
  		</div>
	</div>
</div>
