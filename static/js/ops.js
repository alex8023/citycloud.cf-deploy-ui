$(document).ready(function(){
	$("#config-agent").on('show.bs.modal', function (event) {
  		var button = $(event.relatedTarget);
  		var agent_id = button.data('whatever');
		var modal = $(this)
		if (agent_id == "add") {
			modal.find('.modal-title').text(modal.find('.modal-title').attr('title-add'));
			modal.find('.modal-body #ops_agent_id').val('');
			return 
		}
  		modal.find('.modal-title').text(modal.find('.modal-title').attr('title-update'))
		modal.find('.modal-body #ops_agent_id').val(agent_id);
	})
	
	$("#delete-agent").on('show.bs.modal', function (event) {
  		var button = $(event.relatedTarget);
  		var agentId = button.data('whatever');
		var modal = $(this)
		modal.find('.modal-body #deleteAgentId').val(agentId);
	})
});