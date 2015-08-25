    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">Jobs-{{.JobName}}</h2>
		</div>
		<div class="form-horizontal">
	  		<div class="panel-body">
				<div class="form-group">
					<label class="col-sm-3 control-label">Name</label>
					<div class="col-sm-5">
						<p class="form-control-static">{{.Name}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-3 control-label">ResourcesPool</label>
					<div class="col-sm-5">
						<p class="form-control-static">{{.ResourcesPool}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-3 control-label">Instances</label>
					<div class="col-sm-5">
						<p class="form-control-static">{{.Instances}}</p>
					</div>
				</div>
				<div class="form-group">
					<label class="col-sm-3 control-label">StaticIp</label>
					<div class="col-sm-5">
						<p class="form-control-static">{{.StaticIp}}</p>
					</div>
				</div>
			</div>
		</div>
	</div>