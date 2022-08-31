use actix_web::{get, web, HttpResponse, Responder};

use crate::usecase::text_toknizer;

#[derive(serde::Deserialize)]
pub struct GetQueryParams {
    query: String,
}

#[get("/nouns")]
pub async fn get_aggregate_nouns(params: web::Query<GetQueryParams>) -> impl Responder {
    match text_toknizer::aggregate_group_by_noun(params.query.clone()) {
        Ok(nouns) => HttpResponse::Ok().json(nouns),
        Err(err) => {
            tracing::error!("Failed to: {:?}", err);
            HttpResponse::InternalServerError().json(err.to_string())
        }
    }
}
