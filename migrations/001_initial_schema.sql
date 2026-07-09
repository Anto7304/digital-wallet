BEGIN;
CREATE TABLE IF NOT EXISTS users(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(100) NOT NULL,  -- Changed from fullname to full_name
    phone VARCHAR(20),
    is_active BOOLEAN DEFAULT TRUE,
    is_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    last_login TIMESTAMP WITH TIME ZONE,

    CONSTRAINT valid_email CHECK (email ~ '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}')
);

COMMENT ON TABLE users IS 'Registered users of the wallet system';
COMMENT ON COLUMN users.id IS 'Unique id using UUId for security';
COMMENT ON COLUMN users.email IS 'Email addres used for loggin';
COMMENT on COLUMN users.password_hash IS 'bcrypt hassh. - never store plain passwor text';


--wallet table
CREATE TABLE IF NOT EXISTS wallets(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    balance DECIMAL(15,2) DEFAULT 0.00,
    currency VARCHAR(3) DEFAULT 'KES',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_wallet_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT unique_user_wallet  UNIQUE (user_id),
    CONSTRAINT positive_balance CHECK (balance >=0)
);

COMMENT ON TABLE wallets IS 'User wallets storing Balance';
COMMENT ON COLUMN wallets.balance IS 'Decimal for exact monetary without floating errors';


CREATE TABLE IF NOT EXISTS transactions(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    reference VARCHAR(50) UNIQUE NOT NULL,
    sender_id UUID,
    receiver_id UUID,
    amount DECIMAL(15,2) NOT NULL,
    type VARCHAR(20) NOT NULL,
    category VARCHAR(50),
    description TEXT,
    status VARCHAR(20) DEFAULT 'pending',
    metadata JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME zone,


    CONSTRAINT fk_transaction_sender_id FOREIGN KEY(sender_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_transaction_receiver_id FOREIGN KEY (receiver_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT valid_transaction_type CHECK (type IN('transfer','deposit','withdrawal')),
    CONSTRAINT valid_transaction_status CHECK(status IN ('pending','completed','failed'))
);


COMMENT ON TABLE transactions IS 'All financial transaction';
COMMENT ON COLUMN transactions.reference IS 'idempotency key';



-- INDEXES
-- for email lookups
CREATE INDEX IF NOT EXISTS idx_user_email ON users(email);

--wallet look ups
CREATE INDEX IF NOT EXISTS idx_wallet_user_id ON wallets(user_id);

--transaction lookups
CREATE INDEX IF NOT EXISTS idx_transaction_receiver ON transactions(receiver_id);
CREATE INDEX IF NOT EXISTS idx_transaction_reference ON transactions(reference);
CREATE INDEX IF NOT EXISTS idx_transaction_created_at ON transactions(created_at);
CREATE INDEX IF NOT EXISTS idx_transaction_status ON transactions(status);

--composite indexes for common queerries
CREATE INDEX IF NOT EXISTS idx_transaction_receiver_status ON transactions(receiver_id,status);
CREATE INDEX IF NOT EXISTS idx_transaction_sender_status ON transactions(sender_id,status);
CREATE INDEX IF NOT EXISTS idx_transaction_sender_created ON transactions(sender_id,created_at DESC);


COMMIT;


--verification queries
SELECT table_name FROM information_schema.tables 
WHERE table_schema = 'public' 
ORDER BY table_name;

SELECT indexname FROM pg_indexes 
WHERE schemaname = 'public' 
ORDER BY indexname;



--to run migration = psql -U walletuser -d walletdb -h localhost -f migrations/001_initial_schema.sql