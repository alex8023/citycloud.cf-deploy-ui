$(document).ready(function(){
    var conn;
	var deploylock = false;
	var host = $("#host").val();
	var url = "ws://" + host + "/microboshwebsocket";
	console.log(host);
	var log = $("#websocketmessage");
    if (window["WebSocket"]) {
        conn = new WebSocket(url);
        conn.onclose = function(evt) {
            $("<div><b>Connection closed.</b></div>").appendTo(log);
			deploylock = false;
        }
        conn.onmessage = function(evt) {
            $("<div/>").html($("<b/>").text(evt.data)).appendTo(log);
        }
		conn.onopen = function(){
			conn.send("hello client");
		}
    } else {
        $("<div><b>Your browser does not support WebSockets.</b></div>").appendTo(log);
    }
	
	$('#config-microbosh').on('click', function(){
		window.location.href = "microbosh?action=config";
	})
	
	$('#deploy-microbosh').on('click',function(){
	
		if (!conn){
			console.log("conn has close")
		}else{
			conn.send("ping")
		}
		console.log("send message");
	})
});