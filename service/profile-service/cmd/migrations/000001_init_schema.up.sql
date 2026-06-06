CREATE TABLE IF NOT EXISTS profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    profile_email VARCHAR(255) UNIQUE NOT NULL,
    profile_phone VARCHAR(50),
    profile_address TEXT,
    profile_title VARCHAR(100),
    profile_subtitle VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO profiles (id, profile_email, profile_phone, profile_title, profile_subtitle)
VALUES ('123e4567-e89b-12d3-a456-426614174000', 'contact@hairstudio.com', '+123456789', 'Welcome to Redmond Hair Studio', 'Premium Hair Styling')
    ON CONFLICT (id) DO NOTHING;