<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">BOSH Deployment</h2>
		</div>
		<div class="form-horizontal">
	  		<div class="panel-body">
				<div class="form-group">
				    <div class="col-sm-offset-2 col-sm-10">
				      	<button class="btn btn-default " id = "config-bosh">Config</button>
					  	<button class="btn btn-default " id = "deploy-bosh">Deploy</button>
					<!--<button class="btn btn-primary ">primary</button>
						<button class="btn btn-warning ">warning</button>
						<button class="btn btn-danger ">danger</button>
						<button class="btn btn-link ">link</button>-->
				    </div>
					<script type="text/javascript">
						$('#config-bosh').on('click', function(){
				    		window.location.href = "bosh?action=config";
				  		})
						$('#deploy-bosh').on('click',function(){
							window.location.href = "bosh?action=deploy";
						})
					</script>
				</div>
				{{with .BOSH}}
				<div class="form-group">
					<label class="col-sm-2 control-label">Deployment Name</label>
					<div class="col-sm-10">
						<p class="form-control-static">{{.Name}}</p>
					</div>
				</div>
				{{end}}
			</div>
		</div>
	</div>
</div>
