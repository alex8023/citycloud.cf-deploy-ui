mysql -uroot -p$1 < createDb.sql
mysql -uroot -p$1 cf_deploy_ui < mysql.sql
mysql -uroot -p$1 cf_deploy_ui < job_properties.sql