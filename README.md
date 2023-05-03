# Concept
 - Log in (to create calendar entries, comments, show and tell session)
 - Pages:
   - Landing page:
     - Log in (from app bar)
      - Log in using Keycloak: https://levelup.gitconnected.com/building-micro-services-in-go-using-keycloak-for-authorisation-e00a29b80a43
      - Or https://mikebolshakov.medium.com/keycloak-with-go-web-services-why-not-f806c0bc820a
     - View schedule
     - Select item from schedule
   - Schedules:
     - View schedule
     - Look at planned agenda
     - Add question/comment:
       - provide name and store in localstorage
     - Action items?
 - Features:
   - Time each agenda speaker (start, stop and automatic. Alert when overshot)
   - Seeder: seed db with testable data: https://ieftimov.com/posts/simple-golang-database-seeding-abstraction-gorm/
   - Sockets for live updates of time, comments: https://github.com/gofiber/websocket
 - Database:
   - Models:
     - DONE User
     - DONE Schedule
     - DONE Speaker (nullable fk to User)
     - DONE Comment (nullable fk to User)

# Technologies
 - Golang
   - Fiber
   - Gorm:
     - go-gormigrate/gormigrate
   - gofiber/swagger
 - sqlite
 - VueJS
 - Flowbite
 - Docker
 - Keycloak (auth)

# TODO
 - Implement caching at route level
 - For a given schedule entry it should allow for:
   - Internal/Private (to EVS) or External
   - Location (i.e. chat room, physical room)
 - Add route UI transitions
 - Support reloading page - add fiber route for all pages (catch all)
 - Add light/dark mode selection support (save in localstorage)
 - For client side ui inject 'VITE_' before selected variables and copy to web/static/src
 - Setup docker commands for dev and for production:
   - dev: npm run dev; air
 - Graceful shutdown: https://github.com/gofiber/recipes/tree/master/graceful-shutdown
