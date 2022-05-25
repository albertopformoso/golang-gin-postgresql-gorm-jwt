output "api_url" {
    description = "Application URL"
    value = "https://${heroku_app.default.name}.herokuapp.com"
}

output "heroku_app_id" {
    description = "Heroku app id"
  value = heroku_app.default.id
}