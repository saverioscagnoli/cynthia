use super::user::User;
use crate::client::{http, token};
use serde::{Deserialize, Serialize};
use serde_json::json;

#[derive(Debug, Serialize, Deserialize)]
pub struct Message {
    pub id: String,
    pub channel_id: String,
    pub content: String,
    pub author: User,
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
