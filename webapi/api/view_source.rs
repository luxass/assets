use vercel_runtime::{Body, Error, Request, Response, StatusCode};

pub async fn handler(_req: Request) -> Result<Response<Body>, Error> {
    // make a redirect to the source code
    let response = Response::builder()
        .status(StatusCode::FOUND)
        .header(
            "Location",
            "https://github.com/luxass/assets/tree/main/webapi",
        )
        .body(Body::Empty)?;

    Ok(response)
}
