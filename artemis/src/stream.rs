use std::ops::Deref;

use anyhow::anyhow;
use futures::{
    stream::{SplitSink, SplitStream},
    SinkExt,
};
use serde::Serialize;
use tokio::net::TcpStream;
use tokio_tungstenite::{tungstenite::Message, MaybeTlsStream, WebSocketStream};

pub type Reader = SplitStream<WebSocketStream<MaybeTlsStream<TcpStream>>>;
pub struct Writer(SplitSink<WebSocketStream<MaybeTlsStream<TcpStream>>, Message>);

impl Deref for Writer {
    type Target = SplitSink<WebSocketStream<MaybeTlsStream<TcpStream>>, Message>;

    fn deref(&self) -> &Self::Target {
        &self.0
    }
}

impl Writer {
    pub(crate) fn new(
        writer: SplitSink<WebSocketStream<MaybeTlsStream<TcpStream>>, Message>,
    ) -> Self {
        Writer(writer)
    }

    pub(crate) async fn send<P: Serialize>(&mut self, payload: P) -> anyhow::Result<()> {
        let msg = Message::Text(serde_json::to_string(&payload)?);
        self.0.send(msg).await.map_err(|e| anyhow!(e))
    }
}

unsafe impl Send for Writer {}
unsafe impl Sync for Writer {}
