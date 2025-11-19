-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS guitars (
    id                  BIGSERIAL PRIMARY KEY,
    
    -- Basic info
    brand               VARCHAR(100) NOT NULL,
    model               VARCHAR(100) NOT NULL,
    name                VARCHAR(200),
    year                SMALLINT CHECK (year > 1900 AND year <= EXTRACT(YEAR FROM CURRENT_DATE) + 1),
    country_of_origin   VARCHAR(100),
    
    -- Type & construction
    body_type           VARCHAR(50)     CHECK (body_type IN ('solid', 'semi-hollow', 'hollow', 'acoustic')),
    body_wood           VARCHAR(100),
    top_wood            VARCHAR(100),
    neck_wood           VARCHAR(100),
    fretboard_wood      VARCHAR(100),
    scale_length        DECIMAL(5,2),
    fretboard_radius    DECIMAL(5,2),
    
    -- Hardware
    pickups_config      VARCHAR(20)     CHECK (pickups_config IN ('SSS', 'HHH', 'HH', 'HSH', 'HSS', 'SS', 'P90', 'single_coil', 'humbucker', 'piezo', 'other')),
    pickup_brand        VARCHAR(100),
    bridge_type         VARCHAR(100),
    tuners              VARCHAR(100),
    
    -- Aesthetics
    color               VARCHAR(100),
    finish              VARCHAR(50)     CHECK (finish IN ('nitro', 'poly', 'satin', 'gloss', 'oil', 'relic')),
    pickguard           VARCHAR(100),
    
    -- Condition & pricing
    condition           VARCHAR(50)     CHECK (condition IN ('mint', 'near_mint', 'excellent', 'very_good', 'good', 'fair', 'poor')),
    purchase_price      DECIMAL(12,2),
    purchase_date       DATE,
    selling_price       DECIMAL(12,2),
    is_for_sale         BOOLEAN DEFAULT FALSE,
    is_sold             BOOLEAN DEFAULT FALSE,
    sold_date           DATE,
    sold_to             VARCHAR(200),
    
    -- Extra details
    weight_kg           DECIMAL(5,3),
    serial_number       VARCHAR(100) UNIQUE,
    case_type           VARCHAR(100),
    modifications       TEXT,
    notes               TEXT,
    
    -- Timestamps
    created_at          TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    updated_at          TIMESTAMPTZ DEFAULT NOW() NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS guitars;
-- +goose StatementEnd
