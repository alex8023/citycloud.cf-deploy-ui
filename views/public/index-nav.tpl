<nav class="navbar navbar-default navbar-fixed-top">
	<div class="container">
	    <div class="navbar-header">
			<a class="navbar-brand" href="#">CCI-PaaS</a>
		</div>
		<div id="navbar">
			<ul class="nav navbar-nav" id="navbarul">
				<li id = "home"><a href="index?action=home">{{i18n $.Lang "Home"}}</a></li>
				{{if ne $.IaaSVersion "Vsphere"}}
				<li id = "microbosh"><a href="microbosh">MicroBOSH {{i18n $.Lang "Deploy"}}</a></li>
				<li id = "cloudfoundry"><a href="cloudfoundry">PaaS {{i18n $.Lang "Deploy"}}</a></li>
				{{end}}
				{{if eq $.IaaSVersion "Vsphere"}}
				<li id = "microbosh"><a href="vspheremicrobosh">MicroBOSH {{i18n $.Lang "Deploy"}}</a></li>
				<li id = "cloudfoundry"><a href="vspherecloudfoundry">PaaS {{i18n $.Lang "Deploy"}}</a></li>
				{{end}}
				<li id = "templates"><a href="templates">{{i18n $.Lang "CustomService"}}</a></li>
				<li id = "ops"><a href="ops">PaaS {{i18n $.Lang "Operations"}}</a></li>
			</ul>
			<ul class="nav navbar-nav navbar-right">
			  	<li><a href="logout"><span class="glyphicon glyphicon-log-out"></span> {{i18n $.Lang "LogOut"}}</a></li>
			</ul>
			<ul class="nav navbar-nav navbar-right">
				<li class="dropdown">
					<a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">{{if eq $.Lang "zh-CN"}}简体中文{{end}}{{if eq $.Lang "en-US"}}English{{end}} <span class="caret"></span></a>
					<ul class="dropdown-menu pull-right">
					{{if eq $.Lang "zh-CN"}}
						<li><a href="#" data-lang="en-US" class="lang-changed">English </a></li>
					{{end}}
					{{if eq $.Lang "en-US"}}
						<li><a href="#" data-lang="zh-CN" class="lang-changed">简体中文 </a></li>
					{{end}}
					</ul>
				</li>
			</ul>
		</div>
	</div>
</nav>
<input type="hidden" id="navfocus" value = "{{.NavBarFocus}}">
<div style="height:20px"></div>