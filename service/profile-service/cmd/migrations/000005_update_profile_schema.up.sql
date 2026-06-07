-- 1. Create the new tables
CREATE TABLE IF NOT EXISTS home (
                                    id UUID PRIMARY KEY,
                                    title TEXT,
                                    subtitle TEXT,
                                    aboutStudioTitle TEXT,
                                    aboutStudioSubtitle TEXT,
                                    category VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
                             );

CREATE TABLE IF NOT EXISTS contact (
                                       id UUID PRIMARY KEY,
                                       email VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    title TEXT,
    subtitle TEXT,
    work_hours JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
                             );

CREATE TABLE IF NOT EXISTS about (
                                     id UUID PRIMARY KEY,
                                     title TEXT,
                                     subtitle TEXT,
                                     ourStoryTitle TEXT,
                                     ourStorySubtitle TEXT,
                                     aboutFounderTitle TEXT,
                                     aboutFounderSubtitle TEXT,
                                     aboutFounderReplica TEXT,
                                     category VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
                             );

-- 2. Initialize with Mock Data
-- Using the same ID ensures consistency if you are replacing the old record
INSERT INTO home (id, title, subtitle, aboutStudioTitle, aboutStudioSubtitle, category)
VALUES ('123e4567-e89b-12d3-a456-426614174000', 'Welcome to Redmond Hair Studio', 'Premium Hair Styling', 'About Us', 'Expert styling since 2026', 'main');

INSERT INTO contact (id, email, phone, address, title, subtitle, work_hours)
VALUES ('123e4567-e89b-12d3-a456-426614174000', 'contact@hairstudio.com', '+123456789', '123 Redmond Way', 'Contact Us', 'We are here to help', '{"monday": {"open": 900, "close": 1800}}');

INSERT INTO about (id, title, subtitle, ourStoryTitle, ourStorySubtitle, aboutFounderTitle, aboutFounderSubtitle, aboutFounderReplica, category)
VALUES ('123e4567-e89b-12d3-a456-426614174000', 'About Our Studio', 'The story behind the scissors', 'Our Story', 'Started from a small chair', 'Founder Name', 'Master Stylist', 'Best quality in Redmond', 'main');