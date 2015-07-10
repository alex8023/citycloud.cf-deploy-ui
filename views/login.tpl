{{template "public/header.tpl"}}
    <link href="/static/css/signin.css" rel="stylesheet">
{{template "public/bodyheader.tpl"}}
	<script src="/static/js/login.js"></script>
    <h2></h2>
    <div class="container">
      <form class="form-signin" action = "{{.LoginAction}}" method = "post">
        <h2 class="form-signin-heading">Please sign in</h2>
        <label for="inputEmail" class="sr-only">Email address</label>
        <input type="text" id="inputEmail" name="username" class="form-control" placeholder="Email address" required autofocus>
        <label for="inputPassword" class="sr-only">Password</label>
        <input type="password" id="inputPassword" name= "password" class="form-control" placeholder="Password" required>
        <div class="checkbox">
          <label>
            <input type="checkbox" value="remember-me"> Remember me
          </label>
        </div>
        <button class="btn btn-lg btn-primary btn-block" type="submit">Sign in</button>
      </form>
    </div>
	<input id = "message" type="hidden" value="{{.Message}}">
{{template "public/footer.tpl"}}