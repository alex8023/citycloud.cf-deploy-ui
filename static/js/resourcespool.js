$(document).ready(function(){
	var index = parseInt($("#poollength").val());
	var panel = "deletePanel"
	if ($("#add_resources_pool")){
		$("#add_resources_pool").on("click",function(){
			var divs = $("form.form-horizontal").find("div.form-horizontal:first").html();
			var panelId = panel + index;
			$("<div class= \"panel panel-default\" id = \""+panelId +"\"><div class=\"panel-heading\" ><h2 class=\"panel-title\">ResourcesPools<button name=\"delete\" class=\"btn btn-default\" type=\"button\" onclick=\"deletePanel('"+panelId+"')\">Delete</button></h2></div><div class=\"form-horizontal\">"+divs + "</div></div>").appendTo($("form.form-horizontal"));
			index = index + 1;
		})
	}
	
	if ($("#resourceform")){
		//校验name是否重复
		$("#resourceform").submit(function(){
			var resourceName = $("[name=name]");
			var nameArr = new Array();
			for (var i=0;i< resourceName.length;i++){
				nameArr[i] = resourceName[i].value
			}
			var copyArr = nameArr.sort();
			for (var i=0;i<resourceName.length;i++){
				if (nameArr[i]==copyArr[i+1]){
					alert("资源池名称不能重复");
					return false;
				}
			}
			return true;
		});
	}
	
	
});
function deletePanel(panelId){
	var res = $("form.form-horizontal").find("div.form-horizontal")
	if (res.length==1){
		alert("必须保留至少一个资源池！");
		return;
	}
	$("#"+panelId).remove();
}