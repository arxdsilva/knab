CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE OR REPLACE FUNCTION UUID() RETURNS uuid AS $$
	BEGIN
		RETURN uuid_generate_v4();
	END;
$$ LANGUAGE plpgsql;

CREATE TABLE public.accounts (
	id serial NOT NULL,
	uuid uuid NOT NULL DEFAULT uuid(),
	document_number text NOT NULL,
	active boolean NOT NULL DEFAULT true,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	last_update timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT accounts_pkey PRIMARY KEY (id)
);

CREATE INDEX if not exists accounts_uuid_idx ON public.accounts USING btree (uuid);
