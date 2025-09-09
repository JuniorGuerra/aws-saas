-- Create user SMTP configurations table
CREATE TABLE IF NOT EXISTS user_smtp_configs (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID NOT NULL, -- references auth.users(id) from Supabase Auth
    config_name VARCHAR(255) NOT NULL,
    api_token VARCHAR(255) UNIQUE NOT NULL,
    smtp_host VARCHAR(255) NOT NULL,
    smtp_port INTEGER NOT NULL DEFAULT 465,
    smtp_username VARCHAR(255) NOT NULL,
    smtp_password TEXT NOT NULL, -- encrypted
    smtp_use_ssl BOOLEAN DEFAULT true,
    smtp_timeout INTEGER DEFAULT 5,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    -- Constraints
    UNIQUE(user_id, config_name),
);

-- Create index for faster API token lookups
CREATE INDEX IF NOT EXISTS idx_user_smtp_configs_api_token ON user_smtp_configs(api_token);

-- Create index for user_id lookups
CREATE INDEX IF NOT EXISTS idx_user_smtp_configs_user_id ON user_smtp_configs(user_id);

-- Create index for active configs
CREATE INDEX IF NOT EXISTS idx_user_smtp_configs_active ON user_smtp_configs(user_id, is_active);