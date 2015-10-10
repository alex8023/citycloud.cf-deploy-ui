  <div class="container">
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
			<h2 class="panel-title">Custom Services</h2>
		</div>
		<div class="panel-body">
		    <table class="table table-hover">
				<thead>
					<tr>
					<th>#</th>
					<th>Name</th>
					<th>Description</th>
					<th>Action</th>
					</tr>
				</thead>
				<tbody>
				{{range $index,$element := .Service}}
					<tr>
					<td>{{$index}}</td>
					<td>{{$element.Name}}</td>
					<td>{{$element.Description}}</td>
					<td><button type="button" class="btn btn-default"><span class="glyphicon glyphicon-edit"></span> Config</button>	<button type="button" class="btn btn-default"><span class="glyphicon glyphicon-remove"></span> Delete</button>
					</td>
					</tr>
				{{end}}
				</tbody>
			</table>
		</div>
	</div>
  </div>