{{template "header.tpl"}}
  <link href="/static/css/signin.css" rel="stylesheet">
{{template "bodyheader.tpl"}}
    <h2 style="text-align: right;">你好，{{.UserName}}！</h2><a href="/citycloud.cf-deploy-ui/logout">LogOut</a>
	<div class = "container"></div>
{{template "footer.tpl"}}