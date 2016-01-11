<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">PaaS NetWorks</h2>
		</div>
  		<div class="panel-body">
			<form class="form-horizontal" method="post" action="vspherecloudfoundry">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      <button type="submit" class="btn btn-default " data-loading-text="Saving..."><span class="glyphicon glyphicon-floppy-save"></span> Save</button>
					<a class="btn btn-default " href="vspherecloudfoundry"><span class="glyphicon glyphicon-step-backward"></span> Back</a>
				    </div>
			  	</div>
				{{template "cloudfoundry/vsphere/config_networksv2.tpl" .}}
			</form>
  		</div>
	</div>
</div>
