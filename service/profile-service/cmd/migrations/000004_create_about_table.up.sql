CREATE TABLE IF NOT EXISTS about (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
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