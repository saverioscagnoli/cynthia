use serde::{Deserialize, Serialize};

use super::kind::ChannelKind;

#[derive(Debug, Serialize, Deserialize)]
pub struct ChannelMention {
    pub id: String,
    pub guild_id: String,
    #[serde(rename = "type")]
    pub kind: ChannelKind,
    pub name: String,
}
