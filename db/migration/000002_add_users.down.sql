DROP INDEX IF EXISTS "account_owner_currency_idx";
ALTER TABLE "account" DROP CONSTRAINT IF EXISTS "account_owner_fkey";
DROP TABLE IF EXISTS "user";
