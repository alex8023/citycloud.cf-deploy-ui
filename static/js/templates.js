$(document).ready(function(){
	$("#config-custom-services").on('show.bs.modal', function (event) {
  		var button = $(event.relatedTarget);
  		var serviceId = button.data('whatever');
		var modal = $(this)
		if (serviceId == "add") {
			modal.find('.modal-title').text('Add Custom Service');
			modal.find('.modal-body #customServiceId').val(0);
			modal.find('.modal-body #customServiceName').val('');
			modal.find('.modal-body #customServiceDescription').val('');
			
			modal.find('.modal-body input[name="where"][value="PaaS"]').prop("checked",true);
			modal.find('.modal-body #customServicePaaSId').val(0);
			modal.find('.modal-body #customServiceApi').val('');
			modal.find('.modal-body #customServiceUser').val('');
			modal.find('.modal-body #customServicePasswd').val('');
			modal.find('.modal-body #customServiceOrg').val('');
			modal.find('.modal-body #customServiceSpace').val('');
			
			modal.find('.modal-body #customServiceVmId').val(0);
			modal.find('.modal-body #customServiceIp').val('');
			modal.find('.modal-body #customServiceVmUser').val('');
			modal.find('.modal-body #customServiceVmPasswd').val('');
			return 
		}
  		modal.find('.modal-title').text('Update Custom Service')
		modal.find('.modal-body #customServiceId').val(serviceId);
		$.ajax({
			url:'templates?action=detail',
			dataType:'json',
			data:{id:serviceId},
			success:function(data){
				console.log(data)
				modal.find('.modal-body #customServiceName').val(data.Service.Name);
				modal.find('.modal-body #customServiceDescription').val(data.Service.Description);
				if (data.Service.Where == "PaaS"){
					$("#deploy-onpaas").attr("class","panel-body");
					requiredablePaaS(true);
					$("#deploy-onvms").attr("class","panel-body hidden");
					requiredableVms(false);
					modal.find('.modal-body input[name="where"][value="PaaS"]').prop("checked",true);
					modal.find('.modal-body #customServicePaaSId').val(data.OnPaaS.Id);
					modal.find('.modal-body #customServiceApi').val(data.OnPaaS.Api);
					modal.find('.modal-body #customServiceUser').val(data.OnPaaS.User);
					modal.find('.modal-body #customServicePasswd').val(data.OnPaaS.Passwd);
					modal.find('.modal-body #customServiceOrg').val(data.OnPaaS.Org);
					modal.find('.modal-body #customServiceSpace').val(data.OnPaaS.Space);
				}else if (data.Service.Where == "Vms"){
					$("#deploy-onpaas").attr("class","panel-body hidden");
					requiredablePaaS(false);
					$("#deploy-onvms").attr("class","panel-body");
					requiredableVms(true);
					modal.find('.modal-body input[name="where"][value="Vms"]').prop("checked",true);
				
					modal.find('.modal-body #customServiceVmId').val(data.OnCustom.Id);
					modal.find('.modal-body #customServiceIp').val(data.OnCustom.Ip);
					modal.find('.modal-body #customServiceVmUser').val(data.OnCustom.User);
					modal.find('.modal-body #customServiceVmPasswd').val(data.OnCustom.Passwd);
				}
				


			},
			error:function(data){
				$("#warning-block-service").attr("class","alert alert-danger alert-dismissible")
			}
		});
	})
	
	$("#delete-custom-services").on('show.bs.modal', function (event) {
  		var button = $(event.relatedTarget);
  		var serviceId = button.data('whatever');
		var modal = $(this)
		modal.find('.modal-body #deleteCustomServiceId').val(serviceId);
	})
	
	//template
	$("#config-custom-template").on('show.bs.modal', function (event) {
  		var button = $(event.relatedTarget);
  		var templateId = button.data('whatever');
		var modal = $(this)
		if (templateId == "add") {
			modal.find('.modal-title').text('Add Custom Template');
			modal.find('.modal-body #customTemplateId').val(0);
			modal.find('.modal-body #customTemplateName').val('');
			modal.find('.modal-body #customTemplateFile').val('');
			modal.find('.modal-body #customTargetFile').val('');
			modal.find('.modal-body #customTemplateDescription').val('');
			return 
		}
  		modal.find('.modal-title').text('Update Custom Template')
		modal.find('.modal-body #customTemplateId').val(templateId);
		$.ajax({
			url:'templatesdetail?action=detail',
			dataType:'json',
			data:{id:templateId},
			success:function(data){
				modal.find('.modal-body #customTemplateName').val(data.Name);
				modal.find('.modal-body #customTemplateFile').val(data.TemplateFile);
				modal.find('.modal-body #customTargetFile').val(data.TargetFile);
				modal.find('.modal-body #customTemplateDescription').val(data.Description);
				if (data.FileType == "War"){
					modal.find('.modal-body input[name="fileType"][value="War"]').prop("checked",true);
				}
				if (data.FileType == "Template"){
					modal.find('.modal-body input[name="fileType"][value="Template"]').prop("checked",true);
				}
				
			},
			error:function(data){
				$("#warning-block-service").attr("class","alert alert-danger alert-dismissible")
			}
		});
	})

	$("#delete-custom-template").on('show.bs.modal', function (event) {
  		var button = $(event.relatedTarget);
  		var templateId = button.data('whatever');
		var modal = $(this)
		modal.find('.modal-body #deleteCustomTemplateId').val(templateId);
	})
	
	if ($("#customServiceWhere1")){
		$("#customServiceWhere1").on('click',function(){
			$("#deploy-onpaas").attr("class","panel-body");
			requiredablePaaS(true);
			$("#deploy-onvms").attr("class","panel-body hidden");
			requiredableVms(false);

		})
	}
	if ($("#customServiceWhere2")){
		$("#customServiceWhere2").on('click',function(){
			$("#deploy-onpaas").attr("class","panel-body hidden");
			requiredablePaaS(false);
			$("#deploy-onvms").attr("class","panel-body");
			requiredableVms(true);
		})
	}
	
	function requiredablePaaS(bool){
		$("#customServiceApi").attr("required",bool);
		$("#customServiceUser").attr("required",bool);
		$("#customServicePasswd").attr("required",bool);
		$("#customServiceOrg").attr("required",bool);
		$("#customServiceSpace").attr("required",bool);
	
		$("#customServiceApi").attr("disabled",!bool);
		$("#customServiceUser").attr("disabled",!bool);
		$("#customServicePasswd").attr("disabled",!bool);
		$("#customServiceOrg").attr("disabled",!bool);
		$("#customServiceSpace").attr("disabled",!bool);
		
	}
	
	function requiredableVms(bool){
		$("#customServiceIp").attr("required",bool);
		$("#customServiceVmUser").attr("required",bool);
		$("#customServiceVmPasswd").attr("required",bool);
		
		$("#customServiceIp").attr("disabled",!bool);
		$("#customServiceVmUser").attr("disabled",!bool);
		$("#customServiceVmPasswd").attr("disabled",!bool);
	}
	
	$("#config-custom-component").on('show.bs.modal', function (event) {
		var button = $(event.relatedTarget);
  		var componentId = button.data('whatever');
		var modal = $(this)
		if (componentId == "add") {
			modal.find('.modal-title').text('Add Custom Service Component');
			modal.find('.modal-body #customComponentsId').val(0);
			modal.find('.modal-body #customComponentName').val('');
			modal.find('.modal-body #customComponentValue').val('');
			return 
		}
		modal.find('.modal-title').text('Update Custom Service Component')
		modal.find('.modal-body #customComponentId').val(componentId);
		$.ajax({
			url:'servicecomponent',
			dataType:'json',
			data:{id:componentId},
			success:function(data){
				modal.find('.modal-body #customComponentName').val(data.Name);
				modal.find('.modal-body #customComponentValue').val(data.Value);
			},
			error:function(data){
				$("#warning-block-service").attr("class","alert alert-danger alert-dismissible")
			}
		});
	});
	$("#delete-custom-component").on('show.bs.modal', function (event) {
  		var button = $(event.relatedTarget);
  		var componentId = button.data('whatever');
		var modal = $(this)
		modal.find('.modal-body #deleteservicecomponentId').val(componentId);
	});
});