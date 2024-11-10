use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
pub enum EmbedKind {
    Rich,
    Image,
    Video,
    Gifv,
    Article,
    Link,
    PollResults,
}
