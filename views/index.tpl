{{template "public/header.tpl"}}
  <link href="/static/css/signin.css" rel="stylesheet">
{{template "public/bodyheader.tpl"}}
<!--nav bar-->
{{template "public/index-nav.tpl"}}
<!--body start-->
{{.LayoutContent}}
<!--body end-->
{{template "public/index-footer.tpl"}}