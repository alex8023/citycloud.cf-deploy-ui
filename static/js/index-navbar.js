$(document).ready(function(){
	//	$("#navbarul").children("li").removeClass("active");
	var value = $("#navfocus").val();
	$("#navbarul").children("li#"+value).addClass("active");
});