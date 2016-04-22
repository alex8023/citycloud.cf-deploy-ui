mysql -uroot -pc1oudc0w < sql/createDb.sql
mysql -uroot -pc1oudc0w cf_deploy_ui < sql/mysql.sql
mysql -uroot -pc1oudc0w cf_deploy_ui < sql/job_properties.sql
mysql -uroot -pc1oudc0w cf_deploy_ui < sql/operation.sql
mysql -uroot -pc1oudc0w cf_deploy_ui < sql/service.sql
mysql -uroot -pc1oudc0w cf_deploy_ui < sql/template.sql
mysql -uroot -pc1oudc0w cf_deploy_ui < sql/on_custom.sql
mysql -uroot -pc1oudc0w cf_deploy_ui < sql/on_paas.sql
mysql -uroot -pc1oudc0w cf_deploy_ui < sql/component.sql