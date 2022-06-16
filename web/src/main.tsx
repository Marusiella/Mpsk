import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import { Global, MantineProvider } from "@mantine/core";
import { css } from "@linaria/core";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <MantineProvider
      theme={{
        fontFamily: "Roboto, sans serif",
       
        spacing: { xs: 15, sm: 20, md: 25, lg: 30, xl: 40 },
      }}
    >
      <Global
        styles={(theme) => ({
          "*": {
            boxSizing: "border-box",
            margin: 0,
            padding: 0,
          },
          body: {
            backgroundColor: theme.colors.background,
            color: theme.colors.text,
            fontFamily: theme.fontFamily,
          },
        })}
      />
      <App />
    </MantineProvider>
  </React.StrictMode>
);
