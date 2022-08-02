buildCmd="docker-compose up"
for f in build/docker/*.yaml; do
  buildCmd="{$buildCmd} -f {$f}"
done
bash "{$buildCmd} --build"