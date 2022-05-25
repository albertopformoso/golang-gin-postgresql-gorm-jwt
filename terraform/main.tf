module "api" {
  source     = "./modules/api"
  jwt_secret = var.jwt_secret
}

module "addons" {
  source = "./modules/addons"
  app_id = module.api.heroku_app_id
}