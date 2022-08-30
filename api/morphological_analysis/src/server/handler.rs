use axum::{extract, response::IntoResponse, routing::get, Router};

use crate::server::response;
use crate::usecase::text_toknizer;

pub fn routes() -> Router {
    Router::new().route("/", get(get_aggregate_nouns))
}

#[derive(serde::Deserialize)]
pub struct GetQueryParams {
    query: String,
}

pub async fn get_aggregate_nouns(
    extract::Query(query): extract::Query<GetQueryParams>,
) -> impl IntoResponse {
    match text_toknizer::aggregate_group_by_noun(query.query) {
        Ok(nouns) => {
            let body_str = serde_json::to_string(&nouns).unwrap();
            response::ok_with_body_str(body_str)
        }
        Err(err) => {
            tracing::error!("Failed to: {:?}", err);
            response::internal_server_error()
        }
    }
}
