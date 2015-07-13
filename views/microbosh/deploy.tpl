<input type="hidden" id="navfocus" value = "{{.NavBarFocus}}">
<div style="height:20px"></div>
<div class="container" >
    <div class="panel panel-default">
		<div class="panel-heading" >
			<h2 class="panel-title">MicroBOSH Deployment</h2>
		</div>
	  		<div class="panel-body">
				    <div class="col-sm-offset-2 col-sm-10">
				      	<button class="btn btn-default " id = "config-microbosh">Config</button>
					  	<button class="btn btn-default " id = "deploy-microbosh">Deploy</button>
				    </div>
					<div id = "websocketmessage">
					hello world!
					<div>
			</div>
	</div>
</div>
<script type="text/javascript">
	$(function(){
		var conn;
	    var log = $("#websocketmessage");
	    if (window["WebSocket"]) {
	        conn = new WebSocket("ws://{{.HOST}}{{.AppName}}/websocket?action=mi");
	        conn.onclose = function(evt) {
	            $("<div><b>Connection closed.</b></div>").appendTo(log);
	        }
	        conn.onmessage = function(evt) {
	            $("<div/>").text(evt.data).appendTo(log);
	        }
	    } else {
	        $("<div><b>Your browser does not support WebSockets.</b></div>").appendTo(log);
	    }
	});
</script>