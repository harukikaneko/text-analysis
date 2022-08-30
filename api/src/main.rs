use std::net::SocketAddr;

use axum::{routing::get, Router, response::IntoResponse};
use hyper::{Body, Response, StatusCode};
use tracing_subscriber::{fmt::time::OffsetTime, EnvFilter};

fn main() {
    let style = std::env::var("RUST_LOG_STYLE").unwrap_or_else(|_| "auto".into());

    let builder = tracing_subscriber::fmt()
        .with_env_filter(EnvFilter::from_default_env())
        .with_timer(OffsetTime::local_rfc_3339().expect("could not get local time offset"))
        .with_ansi(style.to_lowercase() != "never")
        .with_file(true)
        .with_line_number(true);

    builder.json().init();

    if let Err(e) = run() {
        tracing::error!("{:?}", e);
    }
}

#[tokio::main]
async fn run() -> hyper::Result<()> {
    let app = Router::new()
        .route("/systems/ping", get(pong))
        .merge(morphological_analysis::routes());

    let addr = SocketAddr::from(([0, 0, 0, 0], 3008));
    tracing::debug!("listening on {}", addr);

    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
}

async fn pong() -> impl IntoResponse {
    Response::builder()
        .status(StatusCode::OK)
        .body(Body::from("pong"))
        .unwrap()
}
