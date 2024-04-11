CREATE TABLE accounts (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  email text NOT NULL,
  password_hash text NOT NULL,
  is_active boolean DEFAULT TRUE,
  is_verified boolean DEFAULT FALSE,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE accounts_sessions (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  account_id UUID REFERENCES accounts(id) ON DELETE CASCADE,
  refresh_token text NOT NULL,
  user_agent text,
  ip_address text,
  issued_at timestamp,
  expires_at timestamp,
  created_at timestamp NOT NULL DEFAULT NOW()
);
