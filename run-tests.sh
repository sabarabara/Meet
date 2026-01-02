set -euo pipefail

echo "Starting services..."
docker compose up -d db server-client ws

echo "Waiting for Postgres to be ready..."
until docker compose exec db pg_isready -U "$DB_USER" -d "$DB_NAME"; do
  sleep 1
done

echo "Running migrations..."
docker compose run --rm migrate

echo "Seeding database..."
docker compose run --rm seed

echo "Running tests..."
docker compose run --rm test

echo "Tearing down..."
docker compose down -v

echo "Done!"
