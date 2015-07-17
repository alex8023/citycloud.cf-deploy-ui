$(document).ready(function(){
    var conn;
	var host = $("#host").val();
	var url = "ws://" + host + "/microboshwebsocket";
	console.log(url);
	var log = $("#websocketmessage");
    if (window["WebSocket"]) {
        conn = new WebSocket(url);
        conn.onclose = function(evt) {
            $("<div><b>Connection closed.</b></div>").appendTo(log);
        }
        conn.onmessage = function(evt) {
            $("<div/>").html($("<b/>").text(evt.data)).appendTo(log);
        }
		conn.onopen = function(){
			$("<div/>").html($("<b/>").text("Connection connected")).appendTo(log);
		}
	
	$('#deploy-all').on('click',function(){
		SendMessage("AllStep");
	})
	
	$('#set-microbosh').on('click',function(){
		SendMessage("SetDeploy");
	})
	
	$('#deploy-microbosh').on('click',function(){
		SendMessage("Deploy");
	})
	
	$('#login-microbosh').on('click',function(){
		SendMessage("Login");
	})
		
    } else {
        $("<div><b>Your browser does not support WebSockets.</b></div>").appendTo(log);
    }
	
	$('#config-microbosh').on('click', function(){
		window.location.href = "microbosh?action=config";
	})
	
	function SendMessage(msg){
		console.log(conn.readyState);
		if (conn.readyState!= conn.OPEN) {
			console.log("reconnect")
			conn.close();
			conn = new WebSocket(url);
			setTimeout(conn.send(msg),2000 );
		}else{
			conn.send(msg);
		}
	}

});