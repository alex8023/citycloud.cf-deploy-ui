$(document).ready(function(){
    var conn;
	var host = $("#host").val();
	var url = "ws://" + host;
	var cango = "yes";
	var log = $("#websocketmessage");
    if (window["WebSocket"]) {
        conn = new WebSocket(url);
        conn.onclose = function(evt) {
            $("<div><b>Connection closed.</b></div>").appendTo(log);
        }
        conn.onmessage = function(evt) {
			var message = evt.data;
			var replaceMessage = "da39a3ee5e6b4b0d3255bfef95601890afd80709";
			if (message == replaceMessage){
				cango = "yes";
			} else {
				$("<div/>").html($("<b/>").text(evt.data)).appendTo(log);
			}
        }
		conn.onopen = function(){
			$("<div/>").html($("<b/>").text("Connection connected")).appendTo(log);
		}
	$('#deploy-all').on('click',function(){
		SendMessage("AllStep");
	})
	
	$('#set-deployment').on('click',function(){
		SendMessage("SetDeploy");
	})
	
	$('#upload-release').on('click',function(){
		SendMessage("UpRelease");
	})
	
	$('#upload-stemcell').on('click',function(){
		SendMessage("UpStemcell");
	})
	
	$('#deploy-cloudfoundry').on('click',function(){
		SendMessage("Deploy");
	})
	
	$('#login-microbosh').on('click',function(){
		SendMessage("Login");
	})
		
    } else {
        $("<div><b>Your browser does not support WebSockets.</b></div>").appendTo(log);
    }
	
	$('#config-cloudfoundry').on('click', function(){
		window.location.href = "cloudfoundry";
	})
	
	function SendMessage(msg){
		if (cango == "yes") {
			conn.send(msg);
		}else{
			$("<div/>").html($("<b/>").text("Job is still running!")).appendTo(log);
			alert("Job is still running!")
		}
		cango = "no";
	}

});