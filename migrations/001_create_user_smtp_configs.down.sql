-- Drop policies
DROP POLICY IF EXISTS "Users can delete own smtp configs" ON user_smtp_configs;
DROP POLICY IF EXISTS "Users can update own smtp configs" ON user_smtp_configs;
DROP POLICY IF EXISTS "Users can insert own smtp configs" ON user_smtp_configs;
DROP POLICY IF EXISTS "Users can view own smtp configs" ON user_smtp_configs;

-- Drop indexes
DROP INDEX IF EXISTS idx_user_smtp_configs_active;
DROP INDEX IF EXISTS idx_user_smtp_configs_user_id;
DROP INDEX IF EXISTS idx_user_smtp_configs_api_token;

-- Drop table
DROP TABLE IF EXISTS user_smtp_configs;