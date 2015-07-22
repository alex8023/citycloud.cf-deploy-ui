$(document).ready(function(){
	var index = 1;
	var panel = "deletePanel"
	if ($("#add_resources_pool")){
		$("#add_resources_pool").on("click",function(){
			index = index +1;
			var divs = $("form.form-horizontal").find("div.form-horizontal:first").html();
			var panelId = panel + index;
			$("<div class= \"panel panel-default\" id = \""+panelId +"\"><div class=\"panel-heading\" ><h2 class=\"panel-title\">ResourcesPools<button name=\"delete\" class=\"btn btn-default\" type=\"button\" onclick=\"deletePanel('"+panelId+"')\">Delete</button></h2></div><div class=\"form-horizontal\">"+divs + "</div></div>").appendTo($("form.form-horizontal"));
		})
	}
	
	
	
});
function deletePanel(panelId){
	console.log(panelId)
	$("#"+panelId).remove();
	console.log($("form.form-horizontal"))
}