use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
pub struct EmbedVideo {
    pub url: Option<String>,
    pub proxy_url: Option<String>,
    pub height: Option<u64>,
    pub width: Option<u64>,
}
