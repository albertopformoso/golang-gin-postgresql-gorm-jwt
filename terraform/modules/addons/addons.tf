resource "heroku_addon" "database" {
  app_id = var.app_id
  plan = "heroku-postgresql:hobby-dev"
}