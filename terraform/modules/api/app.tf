resource "heroku_app" "default" {
  name   = "golang-api-apf"
  region = "us"

  config_vars = {
    GOVERSION="go1.17"
    JWT_SECRET = "${var.jwt_secret}"
  }

  buildpacks = [
    "heroku/go"
  ]
}

# Build code & release to the app
resource "heroku_build" "example" {
    app_id = heroku_app.default.id

    source {
        url = "https://github.com/albertopformoso/golang-gin-postgresql-gorm-jwt/archive/refs/tags/v1.0.1.tar.gz"
        version = "v1.0.0"
    }
}

