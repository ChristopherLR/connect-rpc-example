# Connect-RPC Example

This project demonstrates a Connect-RPC application with a Go backend and a React/TypeScript frontend.

## Prerequisites

- [Go](https://go.dev/dl/) (1.20+)
- [Node.js](https://nodejs.org/) (18+)
- [Buf CLI](https://buf.build/docs/installation) (for Protobuf generation)

## Project Structure

- `proto/`: Protobuf service definitions.
- `backend/`: Go backend server using `connect-go`.
- `frontend/`: React/Vite frontend using `connect-web` and `connect-es`.

## Generating Protobuf Types

If you make changes to the `.proto` files in the `proto/` directory, you need to regenerate the code for both the backend and frontend.

Run the following command from the root of the repository:

```bash
buf generate
```

This uses `buf.gen.yaml` to generate:
- Go code in `backend/`
- TypeScript code in `frontend/src/gen/`

## Running the Application

### 1. Start the Backend

Open a terminal and run the Go server:

```bash
cd backend
go run ./cmd/server/main.go
```

The server will start on `http://localhost:8080`.

### 2. Start the Frontend

Open a **new** terminal and run the Vite development server:

```bash
cd frontend
npm install  # only needed the first time
npm run dev
```

The application will be available at [`http://localhost:5173`](http://localhost:5173).

## Usage

1. Enter a name in the input field.
2. Click **Unary Greet** to send a simple request-response.
3. Click **Stream Greetings** to receive a stream of responses from the server.
