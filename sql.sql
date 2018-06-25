create table if not exists model
(
  id uuid not null
    constraint model_pkey
    primary key,
  rec_time timestamp default timezone('utc'::text, now()) not null
)
;

create table if not exists model_2018_06
(
  constraint model_2018_06_pk
  primary key (id),
  constraint model_2018_06_rec_time_check
  check ((rec_time >= '2018-06-01 00:00:00'::timestamp without time zone) AND (rec_time < '2018-07-21 00:00:00'::timestamp without time zone))
)
  inherits (model)
;

create or replace function get_month_partition_name(t timestamp without time zone) returns character varying
immutable
language sql
as $$
select to_char(t, 'yyyy_mm');
$$
;

create or replace function month_partition_insert() returns trigger
language plpgsql
as $$
declare
  table_name      varchar;
begin
  table_name := TG_TABLE_SCHEMA || '.' || TG_RELNAME || '_' || get_month_partition_name(new.rec_time);

  execute 'insert into ' || table_name || ' values ($1.*)' using new;
  return null;
end;
$$
;

create trigger model_partition_insert
  before insert
  on model
  for each row
execute procedure month_partition_insert()
;

