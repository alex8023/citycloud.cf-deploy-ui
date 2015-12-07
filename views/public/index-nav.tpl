<nav class="navbar navbar-default navbar-fixed-top">
	<div class="container">
	    <div class="navbar-header">
			<a class="navbar-brand" href="#">PaaS Ops</a>
		</div>
		<div id="navbar">
			<ul class="nav navbar-nav" id="navbarul">
				<li id = "home"><a href="index?action=home">Home</a></li>
				<li id = "microbosh"><a href="microbosh">Deploy MicroBOSH</a></li>
				<li id = "cloudfoundry"><a href="cloudfoundry">Deploy PaaS</a></li>
				<li id = "templates"><a href="templates">Deploy Custom</a></li>
				<li id = "ops"><a href="index?action=ops">PaaS Ops</a></li>
				<li class="dropdown">
					<a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Dropdown <span class="caret"></span></a>
					<ul class="dropdown-menu">
						<li><a href="#">Action</a></li>
						<li><a href="#">Another action</a></li>
						<li><a href="#">Something else here</a></li>
						<li role="separator" class="divider"></li>
						<li class="dropdown-header">Nav header</li>
						<li><a href="#">Separated link</a></li>
						<li><a href="#">One more separated link</a></li>
					</ul>
				</li>
			</ul>
			<ul class="nav navbar-nav navbar-right">
			  	<li><a href="logout"><span class="glyphicon glyphicon-log-out"></span> LogOut</a></li>
			</ul>
		</div>
	</div>
</nav>
<input type="hidden" id="navfocus" value = "{{.NavBarFocus}}">
<div style="height:20px"></div>