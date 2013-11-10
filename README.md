# readme
free to use blog written in go. it's pretty basic. the root displays a list of
your posts, going to `/create/session` lets you log in, and you can manage your
posts at `/show/session` (don't ask me about the dumb path naming). once you're
logged in you can add new posts at `/new`.

## setup
easiest way to get it up is using heroku. create your heroku site, then set the buildpack on it:
`$ heroku config:add BUILDPACK_URL=https://github.com/robfig/heroku-buildpack-go-revel.git`

push like you normally would to heroku

## logging in
you have to set the environment variables that will be used for authentication.
they are BLOGTOKEN, HASHED_PASSWORD, and HASHED_USER. Use the env var setting
executable, `setup_env_vars`.

## postgres
there is a bit of setup to get your database working

first, you need to add postgres to your herokuapp:
`$ heroku addons:add heroku-postgresql:dev`

then you need to change the datatype for the body column in posts:

```
$ heroku pg:psql
psql=> alter table posts alter column body type text;
psql=> \q;
```
