# ql-rename
批量重命名文件

## 事件定义
Event Name 在 frontend/src/events.gen.ts 中定义
然后执行 go generate ./backend 生成前端用的事件定义events.gen.ts

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.
### 构建 Windows AMD64 版本
wails build -clean -platform windows/amd64
### 构建 Windows ARM64 版本
wails build -clean -platform windows/arm64
