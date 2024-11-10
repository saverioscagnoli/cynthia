use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
pub struct AvatarDecorationData {
    pub asset: String,
    pub sku_id: String,
}
