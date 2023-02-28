INSERT INTO t_parts(
	parts_id, vehicle_name, parts_name, parts_no, parts_count, parts_price)
	VALUES
 	(nextval('t_parts_parts_id_seq'), '包装辅料', '4#周转箱500*400*380', '011611', 3, 7.12),
	(nextval('t_parts_parts_id_seq'), 'SDH1P52QMl-5', '活塞销', '13111-KVJ-910', 10, 2.80),
 	(nextval('t_parts_parts_id_seq'), '(B11)SDH1P52QM', '链条涨紧器', '14510-KCW-850', 5, 9.60),
	(nextval('t_parts_parts_id_seq'), 'SDH152FMI-4', '链条导板', '14610-KCW-880', 3, 4.20);