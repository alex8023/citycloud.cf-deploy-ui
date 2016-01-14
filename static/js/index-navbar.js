$(document).ready(function(){
	//	$("#navbarul").children("li").removeClass("active");
	var value = $("#navfocus").val();
	$("#navbarul").children("li#"+value).addClass("active");
	
	$(document).on('click', '.lang-changed', function(){
		var $e = $(this);
		var lang = $e.data('lang');
		$.cookie('lang', lang, {path: '/', expires: 365});
		window.location.reload();
	});
});