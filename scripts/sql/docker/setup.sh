retry_count=0
max_retries=10
retry_delay=5

until psql -h pg -U cook -d cook -c "SELECT 1" > /dev/null 2>&1; do
    if [ $retry_count -eq $max_retries ]; then
        echo "Unable to connect to PostgreSQL container after $max_retries retries. Exiting..."
        exit 1
    fi

    echo "Waiting for PostgreSQL container to be ready..."
    sleep $retry_delay
    ((retry_count++))
done

echo "Setup completed. Starting the application..."
