use serde::{Deserialize, Serialize};
use serde_json::json;

use crate::{client::token, opcode::Opcode};

#[derive(Debug, Serialize, Deserialize)]
pub struct Payload {
    pub op: u8,
    pub d: serde_json::Value,
    pub s: Option<u64>,
    pub t: Option<String>,
}

impl Payload {
    pub fn identify() -> Self {
        Self {
            op: Opcode::Identify.into(),
            d: json!({
                "token": token().to_string(),
                "intents": 53608447,
                "properties": {
                    "$os": "linux",
                    "$browser": "artemis",
                    "$device": "artemis",
                },
            }),
            s: None,
            t: None,
        }
    }

    pub fn heartbeat() -> Self {
        Self {
            op: Opcode::Heartbeat.into(),
            d: serde_json::Value::Null,
            s: None,
            t: None,
        }
    }
}
