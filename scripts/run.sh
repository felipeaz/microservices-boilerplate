buildCmd="docker-compose"
for f in build/docker/*.yaml; do
  buildCmd="$buildCmd -f $f"
done
bash "$buildCmd up"