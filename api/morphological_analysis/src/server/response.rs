use hyper::{Body, Response, StatusCode};

pub fn internal_server_error() -> Response<Body> {
    Response::builder()
        .status(StatusCode::INTERNAL_SERVER_ERROR)
        .body(Body::empty())
        .unwrap()
}

pub fn ok_with_body_str(body_str: String) -> Response<Body> {
    Response::builder()
        .status(StatusCode::OK)
        .body(Body::from(body_str))
        .unwrap()
}
