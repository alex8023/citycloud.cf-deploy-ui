<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">PaaS {{i18n $.Lang "JobsProperties"}}</h2>
		</div>
  		<div class="panel-body">
			<form class="form-horizontal" method="post" action="vspherecloudfoundry">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				     	<button type="submit" class="btn btn-default " data-loading-text="Saving..."><span class="glyphicon glyphicon-floppy-save"></span> {{i18n $.Lang "Save"}}</button>
						<a class="btn btn-default " href="vspherecloudfoundry"><span class="glyphicon glyphicon-step-backward"></span> {{i18n $.Lang "Back"}}</a>
				    </div>
			  	</div>
				{{with .Properties}}
				{{range .JobProperties}}
			  	<div class="form-group">
			    	<label for="{{.Name}}" class="col-sm-5 control-label">{{.Name}}</label>
				    <div class="col-sm-4">
				      <input type="text" class="form-control" id="{{.Name}}" placeholder="{{.Name}}" name="{{.Name}}" value = "{{.Value}}" required>
				    </div>
			  	</div>
				{{end}}
				{{end}}
				<input type = "hidden" name="model" value = "{{.Model}}">
			</form>
  		</div>
	</div>
</div>
