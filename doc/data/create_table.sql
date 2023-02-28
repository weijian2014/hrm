/*
PostgreSQL Backup
Database: db_hi/public
Backup Time: 2018-03-28 23:26:26
*/

SET lc_monetary='zh_CN.UTF-8';
DROP SEQUENCE IF EXISTS "public"."t_parts_parts_id_seq";
DROP SEQUENCE IF EXISTS "public"."t_user_user_id_seq";
DROP TABLE IF EXISTS "public"."t_parts";
DROP TABLE IF EXISTS "public"."t_user";

CREATE SEQUENCE "public"."t_parts_parts_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
CREATE SEQUENCE "public"."t_user_user_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- Table: public.t_parts
CREATE TABLE public.t_parts
(
    parts_id integer NOT NULL DEFAULT nextval('t_parts_parts_id_seq'::regclass),
    vehicle_name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    parts_name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    parts_no character varying(255) COLLATE pg_catalog."default" NOT NULL,
    parts_count bigint NOT NULL,
    parts_price money,
    parts_selling_price money,
    CONSTRAINT t_parts_pkey PRIMARY KEY (parts_id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.t_parts
    OWNER to postgres;

COMMENT ON COLUMN public.t_parts.parts_id
    IS '自动增长';

COMMENT ON COLUMN public.t_parts.vehicle_name
    IS '车型名称';

COMMENT ON COLUMN public.t_parts.parts_name
    IS '配件名称';

COMMENT ON COLUMN public.t_parts.parts_no
    IS '配件编码';

COMMENT ON COLUMN public.t_parts.parts_count
    IS '配件个数';

COMMENT ON COLUMN public.t_parts.parts_price
    IS '配件单价';

COMMENT ON COLUMN public.t_parts.parts_selling_price
    IS '配件销售价格';


-- Table: public.t_user
CREATE TABLE public.t_user
(
    user_id integer NOT NULL DEFAULT nextval('t_user_user_id_seq'::regclass),
    user_name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    password character varying(255) COLLATE pg_catalog."default" NOT NULL,
    user_type integer NOT NULL,
    create_date date NOT NULL,
    modify_date date,
    CONSTRAINT t_user_pkey PRIMARY KEY (user_id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.t_user
    OWNER to postgres;

COMMENT ON COLUMN public.t_user.user_id
    IS '自封增长';

COMMENT ON COLUMN public.t_user.user_name
    IS '登录用户名';

COMMENT ON COLUMN public.t_user.password
    IS '登录密码';

COMMENT ON COLUMN public.t_user.user_type
    IS '用户类型 0:管理员 1普通用户';

COMMENT ON COLUMN public.t_user.create_date
    IS '创建日期';

COMMENT ON COLUMN public.t_user.modify_date
    IS '修改日期';


INSERT INTO t_user (user_id, user_name, password, user_type, create_date, modify_date)
VALUES
	(nextval('t_user_user_id_seq'), 'admin', 'wzn123456789', 0, now(), now()),
	(nextval('t_user_user_id_seq'), 'user', 'wzn123456', 1, now(), now());