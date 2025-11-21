-- +goose Up
-- +goose StatementBegin
ALTER TABLE guitars
    DROP COLUMN IF EXISTS name,
    DROP COLUMN IF EXISTS country_of_origin,

    DROP COLUMN IF EXISTS body_type,
    DROP COLUMN IF EXISTS body_wood,
    DROP COLUMN IF EXISTS top_wood,
    DROP COLUMN IF EXISTS neck_wood,
    DROP COLUMN IF EXISTS fretboard_wood,
    DROP COLUMN IF EXISTS scale_length,
    DROP COLUMN IF EXISTS fretboard_radius,

    DROP COLUMN IF EXISTS pickups_config,
    DROP COLUMN IF EXISTS pickup_brand,
    DROP COLUMN IF EXISTS bridge_type,
    DROP COLUMN IF EXISTS tuners,

    DROP COLUMN IF EXISTS color,
    DROP COLUMN IF EXISTS finish,
    DROP COLUMN IF EXISTS pickguard,

    DROP COLUMN IF EXISTS condition,
    DROP COLUMN IF EXISTS purchase_price,
    DROP COLUMN IF EXISTS purchase_date,
    DROP COLUMN IF EXISTS selling_price,
    DROP COLUMN IF EXISTS is_for_sale,
    DROP COLUMN IF EXISTS is_sold,
    DROP COLUMN IF EXISTS sold_date,
    DROP COLUMN IF EXISTS sold_to,

    DROP COLUMN IF EXISTS weight_kg,
    DROP COLUMN IF EXISTS serial_number,
    DROP COLUMN IF EXISTS case_type,
    DROP COLUMN IF EXISTS modifications;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE guitars
    ADD COLUMN IF NOT EXISTS name VARCHAR(200),
    ADD COLUMN IF NOT EXISTS country_of_origin VARCHAR(100),

    ADD COLUMN IF NOT EXISTS body_type VARCHAR(50) CHECK (body_type IN ('solid', 'semi-hollow', 'hollow', 'acoustic')),
    ADD COLUMN IF NOT EXISTS body_wood VARCHAR(100),
    ADD COLUMN IF NOT EXISTS top_wood VARCHAR(100),
    ADD COLUMN IF NOT EXISTS neck_wood VARCHAR(100),
    ADD COLUMN IF NOT EXISTS fretboard_wood VARCHAR(100),
    ADD COLUMN IF NOT EXISTS scale_length DECIMAL(5,2),
    ADD COLUMN IF NOT EXISTS fretboard_radius DECIMAL(5,2),

    ADD COLUMN IF NOT EXISTS pickups_config VARCHAR(20) CHECK (pickups_config IN ('SSS', 'HHH', 'HH', 'HSH', 'HSS', 'SS', 'P90', 'single_coil', 'humbucker', 'piezo', 'other')),
    ADD COLUMN IF NOT EXISTS pickup_brand VARCHAR(100),
    ADD COLUMN IF NOT EXISTS bridge_type VARCHAR(100),
    ADD COLUMN IF NOT EXISTS tuners VARCHAR(100),

    ADD COLUMN IF NOT EXISTS color VARCHAR(100),
    ADD COLUMN IF NOT EXISTS finish VARCHAR(50) CHECK (finish IN ('nitro', 'poly', 'satin', 'gloss', 'oil', 'relic')),
    ADD COLUMN IF NOT EXISTS pickguard VARCHAR(100),

    ADD COLUMN IF NOT EXISTS condition VARCHAR(50) CHECK (condition IN ('mint', 'near_mint', 'excellent', 'very_good', 'good', 'fair', 'poor')),
    ADD COLUMN IF NOT EXISTS purchase_price DECIMAL(12,2),
    ADD COLUMN IF NOT EXISTS purchase_date DATE,
    ADD COLUMN IF NOT EXISTS selling_price DECIMAL(12,2),
    ADD COLUMN IF NOT EXISTS is_for_sale BOOLEAN DEFAULT FALSE,
    ADD COLUMN IF NOT EXISTS is_sold BOOLEAN DEFAULT FALSE,
    ADD COLUMN IF NOT EXISTS sold_date DATE,
    ADD COLUMN IF NOT EXISTS sold_to VARCHAR(200),

    ADD COLUMN IF NOT EXISTS weight_kg DECIMAL(5,3),
    ADD COLUMN IF NOT EXISTS serial_number VARCHAR(100) UNIQUE,
    ADD COLUMN IF NOT EXISTS case_type VARCHAR(100),
    ADD COLUMN IF NOT EXISTS modifications TEXT;
-- +goose StatementEnd
