mod domain;
mod server;
mod usecase;
use axum::Router;
use server::handler;

pub fn routes() -> Router {
    Router::new().nest("/nouns", handler::routes())
}
