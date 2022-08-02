echo "Configuring Kong API Gateway"
for f in setup/kong/*.sh; do
  bash "$f"
done