const express = require('express')
const passport = require('passport')
const session = require('express-session')
const redis = require('redis')
const cookieParser = require('cookie-parser')
const bodyParser = require('body-parser')
const app = require('./next')
const config = require('../config')
const routes = require('./routes')

const port = config.port

const RedisStore = require('connect-redis')(session)
const redisClient = redis.createClient({ url: config.session.url })

const nextjsPath = config.kore.showPrototypes ? '*' : /^((?!\/prototype).)*$/

if (config.disableAnimations) {
  console.log('Animations disabled - unset KORE_UI_DISABLE_ANIMATIONS to restore.')
}

app.prepare().then(() => {
  const server = express()

  server.use(cookieParser())
  server.use(bodyParser.urlencoded({ extended: true }))
  server.use(bodyParser.json())
  server.use(session({
    store: new RedisStore({
      client: redisClient,
      url: config.session.url,
      ttl: config.session.ttlInSeconds
    }),
    secret: config.session.sessionSecret,
    resave: false,
    saveUninitialized: true
  }))

  server.use(passport.initialize())
  server.use(passport.session())

  server.use(routes)

  server.all(nextjsPath, (req, res) => {
    const handle = app.getRequestHandler()
    return handle(req, res)
  })

  server.listen(port, err => {
    if (err) {
      throw err
    }
    console.log(`> Ready on http://localhost:${port}`)
  })
})
