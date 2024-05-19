BEGIN;

-- Create Rates Table
CREATE TABLE IF NOT EXISTS rates (
    id SERIAL PRIMARY KEY,
    currency_code CHAR(3) NOT NULL,
    rate NUMERIC(10, 4) NOT NULL,
    exchange_date DATE NOT NULL,
    CONSTRAINT unique_rate_per_day UNIQUE (currency_code, exchange_date)
);

-- Create Subscribers Table
CREATE TABLE IF NOT EXISTS subscribers (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE
    );

-- Create Email Notifications Table
CREATE TABLE IF NOT EXISTS email_notifications (
    id SERIAL PRIMARY KEY,
    subscriber_id INTEGER NOT NULL,
    rate_id INTEGER NOT NULL,
    sent_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (subscriber_id) REFERENCES subscribers (id) ON DELETE CASCADE,
    FOREIGN KEY (rate_id) REFERENCES rates (id) ON DELETE CASCADE
    );

COMMIT;
