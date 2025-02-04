create table customers
(
    id         bigserial
        primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name       text
        constraint uni_customers_name
            unique,
    surname    text,
    address    text,
    phone      text,
    email      text
);

alter table customers
    owner to admin_panel_user;

create index idx_customers_deleted_at
    on customers (deleted_at);

create table invoices
(
    id                    bigserial
        primary key,
    created_at            timestamp with time zone,
    updated_at            timestamp with time zone,
    deleted_at            timestamp with time zone,
    customer_id           bigint
        constraint fk_invoices_customer
            references customers,
    year                  bigint,
    number                text,
    payment_method        text,
    amount                bigint,
    paid_amount           bigint,
    date                  timestamp with time zone,
    payment_date          timestamp with time zone,
    expected_payment_date timestamp with time zone,
    note                  text
);

alter table invoices
    owner to admin_panel_user;

create index idx_invoices_deleted_at
    on invoices (deleted_at);

create table invoice_rows
(
    id         bigserial
        primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    invoice_id bigint
        constraint fk_invoices_rows
            references invoices,
    number     bigint,
    amount     bigint
);

alter table invoice_rows
    owner to admin_panel_user;

create index idx_invoice_rows_deleted_at
    on invoice_rows (deleted_at);

