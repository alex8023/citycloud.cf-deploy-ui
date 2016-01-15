$(document).ready(function(){
	$("#config-vsphere-resource").on('show.bs.modal', function (event) {
  		var button = $(event.relatedTarget);
  		var resourceId = button.data('whatever');
		var modal = $(this)
		if (resourceId == "add") {
			modal.find('.modal-title').text(modal.find('.modal-title').attr('title-add'));
			modal.find('.modal-body #vsphereresourceId').val(0);
			modal.find('.modal-body #ram').val('4096');
			modal.find('.modal-body #cpu').val('2');
			modal.find('.modal-body #disk').val('4096');
			modal.find('.modal-body #persistentDisk').val('0');
			return 
		}
  		modal.find('.modal-title').text(modal.find('.modal-title').attr('title-update'))
		modal.find('.modal-body #vsphereresourceId').val(resourceId);
		$.ajax({
			url:'vsphereresource?action=detail',
			dataType:'json',
			data:{id:resourceId},
			success:function(data){
				modal.find('.modal-body #ram').val(data.Ram);
				modal.find('.modal-body #cpu').val(data.Cpu);
				modal.find('.modal-body #disk').val(data.Disk);
				modal.find('.modal-body #persistentDisk').val(data.PersistentDisk);
			},
			error:function(data){
				$("#warning-block-service").attr("class","alert alert-danger alert-dismissible")
			}
		});
	})
	
	$("#delete-vsphere-resource").on('show.bs.modal', function (event) {
  		var button = $(event.relatedTarget);
  		var resourceId = button.data('whatever');
		var modal = $(this)
		modal.find('.modal-body #vsphereresourceId').val(resourceId);
	})
			
});