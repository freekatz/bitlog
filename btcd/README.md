# Bootstrap a btcd peer

## clone this project and prepare

```shell
git clone https://github.com/1uvu/bitlog.git
cd bitlog/btcd
mkdir docker-compose/tmpl/
```

## modify `.env.tmpl` with the help of comments

```shell
cp docker-compose/example/.env.example docker-compose/tmpl/.env.tmpl
nano docker-compose/tmpl/.env.tmpl
```

## modify `btcd.conf.tmpl` with the help of comments

```shell
cp docker-compose/example/btcd.config.example docker-compose/tmpl/btcd.config.tmpl
nano docker-compose/tmpl/btcd.config.tmpl
```

## run bootstrap.sh

```shell
bash bootstrap.sh
```
