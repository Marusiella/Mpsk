import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import { Global, MantineProvider } from "@mantine/core";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import ChangeAdmin from "./ChangeAdmin";

export var adress = import.meta.env.PROD ? "" : "http://localhost:3000";

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
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<App />} />
          <Route path="/changeAdmin" element={<ChangeAdmin />} />
        </Routes>
      </BrowserRouter>
    </MantineProvider>
  </React.StrictMode>
);
