use author::EmbedAuthor;
use field::EmbedField;
use footer::EmbedFooter;
use image::EmbedImage;
use kind::EmbedKind;
use provider::EmbedProvider;
use serde::{Deserialize, Serialize};
use thumbnail::EmbedThumbnail;
use video::EmbedVideo;

pub mod author;
pub mod field;
pub mod footer;
pub mod image;
pub mod kind;
pub mod provider;
pub mod thumbnail;
pub mod video;

#[derive(Debug, Serialize, Deserialize)]
pub struct Embed {
    pub title: Option<String>,
    #[serde(rename = "type")]
    pub kind: EmbedKind,
    pub description: Option<String>,
    pub url: Option<String>,
    pub timestamp: Option<String>,
    pub color: Option<u64>,
    pub footer: EmbedFooter,
    pub image: EmbedImage,
    pub thumbnail: EmbedThumbnail,
    pub video: EmbedVideo,
    pub provider: EmbedProvider,
    pub author: EmbedAuthor,
    pub fields: Vec<EmbedField>,
}
