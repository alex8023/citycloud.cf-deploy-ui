package entity

const (
	CloudFoundryUpdateTempalte string = `{{define "update"}}
update:
  canaries: 1
  canary_watch_time: 30000 - 90000
  update_watch_time: 30000 - 90000
  max_in_flight: 4
  max_errors: 4
{{end}}`
)
