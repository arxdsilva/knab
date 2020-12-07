ALTER TABLE accounts
    ADD COLUMN available_credit_limit decimal DEFAULT 0.0;
ALTER TABLE accounts
    ADD COLUMN total_credit_limit decimal DEFAULT 0.0;