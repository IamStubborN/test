-- +migrate Up
-- +migrate StatementBegin
insert into transaction_types(type_id, name)
values(1,'Win'), (2,'Bet')
-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin
-- +migrate StatementEnd