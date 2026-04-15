import { WebSocketServer } from "ws";
import * as cp from "child_process";
import * as path from "path";
import * as rpc from "vscode-ws-jsonrpc";
import * as server from "vscode-ws-jsonrpc/server";

const wss = new WebSocketServer({ port: 3000 });

wss.on("connection", (webSocket) => {
  const socket: rpc.IWebSocket = {
    send: (content) =>
      webSocket.send(content, (err) => {
        if (err) throw err;
      }),
    onMessage: (cb) => webSocket.on("message", cb),
    onError: (cb) => webSocket.on("error", cb),
    onClose: (cb) =>
      webSocket.on("close", (code, reason) => cb(code, reason.toString())),
    dispose: () => webSocket.close(),
  };

  // creating child process (the extension)
  const lspServerPath = path.resolve(__dirname, "../../server/out/server.js");
  const serverProcess = cp.spawn("node", [lspServerPath, "--stdio"]);

  const connection = server.createWebSocketConnection(socket);
  const serverConnection = server.createProcessStreamConnection(serverProcess);

  // the connection
  server.forward(connection, serverConnection!, (msg) => msg);

  webSocket.on("close", () => serverProcess.kill());
});

console.log("WebSocket Server running on ws://localhost:3000");
