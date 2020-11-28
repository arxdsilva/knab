CREATE TABLE public.transactions (
	id serial NOT NULL,
	uuid uuid NOT NULL DEFAULT uuid(),
	account_id integer NOT NULL,
	operation_id integer NOT NULL,
    amount decimal NOT NULL DEFAULT 0.0,
	event_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	last_update timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT transactions_pkey PRIMARY KEY (id),
    FOREIGN KEY (account_id)
        REFERENCES accounts (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    FOREIGN KEY (operation_id)
        REFERENCES operations (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

CREATE INDEX if not exists transactions_uuid_idx ON public.transactions USING btree (uuid);
