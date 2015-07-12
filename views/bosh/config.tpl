<input type="hidden" id="navfocus" value = "{{.NavBarFocus}}">
<div style="height:20px"></div>
<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">BOSH Deployment</h2>
		</div>
  		<div class="panel-body">
			<form class="form-horizontal" method="post" action="bosh">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      <button type="submit" class="btn btn-default " data-loading-text="Saving...">Save</button>
				    </div>
			  	</div>
				{{with .BOSH}}
			  	<div class="form-group">
			    	<label for="name" class="col-sm-2 control-label">Deployment Name</label>
				    <div class="col-sm-10">
				      <input type="text" class="form-control" id="name" placeholder="Deployment Name" name="name" value = "{{.Name}}">
				    </div>
			  	</div>
				{{end}}
			</form>
  		</div>
	</div>
</div>
