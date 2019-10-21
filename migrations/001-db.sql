-- +migrate Up
-- +migrate StatementBegin
create table if not exists users
(
	id bigint not null
		constraint users_pk
			primary key,
	balance numeric(10,2)
);

create table if not exists transaction_types
(
	type_id smallserial not null
		constraint transaction_types_pk
			primary key,
	name varchar(50) not null
);

create table if not exists deposits
(
	id bigint not null
		constraint deposits_pk
			primary key,
	user_id bigint
		constraint users___fk
			references users (id),
	amount numeric(10,2) not null,
	balance_before numeric(10,2) not null,
	balance_after numeric(10,2) not null,
	date timestamp not null
);

create table if not exists transactions
(
	id bigint not null
		constraint transaction_pk
			primary key,
	user_id bigint not null
		constraint users___fk
			references users (id),
	transaction_type_id smallint not null
		constraint transaction_types___fk
			references transaction_types,
	amount numeric(10,2) not null,
	balance_before numeric(10,2) not null,
	balance_after numeric(10,2) not null,
	date timestamp not null
);


-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin
drop table IF EXISTS deposits cascade;
drop table IF EXISTS transactions cascade;
drop table IF EXISTS transaction_types cascade;
drop table IF EXISTS users cascade;
-- +migrate StatementEnd