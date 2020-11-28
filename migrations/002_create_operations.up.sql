CREATE TABLE public.operations (
	id serial NOT NULL,
	description text NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	last_update timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT operations_pkey PRIMARY KEY (id)
);

INSERT INTO public.operations (description) VALUES ('COMPRA A VISTA');
INSERT INTO public.operations (description) VALUES ('COMPRA PARCELADA');
INSERT INTO public.operations (description) VALUES ('SAQUE');
INSERT INTO public.operations (description) VALUES ('PAGAMENTO');
