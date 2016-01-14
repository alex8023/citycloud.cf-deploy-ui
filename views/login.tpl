{{template "public/header.tpl"}}
    <link href="/static/css/signin.css" rel="stylesheet">
{{template "public/bodyheader.tpl"}}
	<script src="/static/js/login.js"></script>
    <h2></h2>
    <div class="container">
      <form class="form-signin" action = "{{.LoginAction}}" method = "post">
        <h2 class="form-signin-heading">{{i18n .Lang "Please sign in"}}</h2>
        <label for="inputEmail" class="sr-only">{{i18n .Lang "Email address"}}</label>
        <input type="text" id="inputEmail" name="username" class="form-control" placeholder="{{i18n .Lang "Email address"}}" required autofocus value="admin">
        <label for="inputPassword" class="sr-only">{{i18n .Lang "Password"}}</label>
        <input type="password" id="inputPassword" name= "password" class="form-control" placeholder="{{i18n .Lang "Password"}}" required value="c1oudc0w">
   
        <button class="btn btn-lg btn-primary btn-block" type="submit">{{i18n .Lang "Sign in"}}</button>
      </form>
    </div>
	<input id = "message" type="hidden" value="{{.Message}}">
{{template "public/footer.tpl" .}}