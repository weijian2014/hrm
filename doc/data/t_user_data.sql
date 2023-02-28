INSERT INTO t_user (user_id, user_name, password, user_type, create_date, modify_date)
VALUES
	(nextval('t_user_user_id_seq'), 'admin', 'wzn123456789', 0, now(), now()),
	(nextval('t_user_user_id_seq'), 'user', 'wzn123456', 1, now(), now());