use actix_web::{web, App, HttpResponse, HttpServer, Responder};

fn main() {
    if let Err(e) = run() {
        tracing::error!("{:?}", e);
    }
}

#[actix_web::main]
async fn run() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .service(morphological_analysis::server::get_aggregate_nouns)
            .route("/systems/ping", web::get().to(pong))
    })
    .bind(("127.0.0.1", 3008))?
    .run()
    .await
}

async fn pong() -> impl Responder {
    HttpResponse::Ok().body("pong")
}
