use super::{channel::mention::ChannelMention, user::User};
use crate::client::{http, token};
use attachment::Attachment;
use embed::Embed;
use serde::{Deserialize, Serialize};
use serde_json::json;

pub mod attachment;
pub mod embed;

#[derive(Debug, Serialize, Deserialize)]
pub struct Message {
    pub id: String,
    pub channel_id: String,
    pub author: User,
    pub content: String,
    pub timestamp: String,
    pub edited_timestamp: Option<String>,
    pub tts: bool,
    pub mention_everyone: bool,
    pub mentions: Vec<User>,
    pub mention_roles: Vec<String>,
    pub mention_channels: Option<Vec<ChannelMention>>,
    pub attachments: Vec<Attachment>,
    pub embeds: Vec<Embed>,
}

impl Message {
    pub async fn reply(&self, content: &str) {
        let channel_id = self.channel_id.clone();
        let content = content.to_string();

        http()
            .post(
                format!(
                    "https://discord.com/api/v10/channels/{}/messages",
                    channel_id
                )
                .as_str(),
            )
            .header("Authorization", format!("Bot {}", token()))
            .header("Content-Type", "application/json")
            .body(json!({ "content": content.to_string(), "message_reference": { "message_id": self.id } }).to_string())
            .send()
            .await
            .unwrap();
    }

    pub async fn send(&self, content: &str) {
        http()
            .post(
                format!(
                    "https://discord.com/api/v10/channels/{}/messages",
                    self.channel_id
                )
                .as_str(),
            )
            .header("Authorization", format!("Bot {}", token()))
            .header("Content-Type", "application/json")
            .body(json!({ "content": content.to_string() }).to_string())
            .send()
            .await
            .unwrap();
    }
}
