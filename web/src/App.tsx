import React from "react";
import {
  Box,
  Button,
  Center,
  Checkbox,
  Group,
  PasswordInput,
  Text,
  TextInput,
} from "@mantine/core";
import { useForm } from "@mantine/form";
import { useNavigate } from "react-router-dom";
import { adress } from "./main";

interface FormData {
  email: string;
  password: string;
}
interface LoginResponse {
  HaveToCreateNewUser: boolean;
  result: string;
}

export default function App() {
  let navigate = useNavigate();
  async function HandleLogin(params: FormData) {
    var a = await fetch(adress + "/api/v1/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify(params),
    });
    var b: LoginResponse = await a.json();
    console.log(b);
    if (b.HaveToCreateNewUser) {
      navigate("/changeAdmin");
    }
  }

  const form = useForm({
    initialValues: {
      email: "",
      password: "",
    },

    validate: {
      email: (value) => (/^\S+@\S+$/.test(value) ? null : "Invalid email"),
    },
  });
  return (
    <Center>
      <Box px={"30vw"}>
        <Text weight={"bold"} style={{ fontSize: 100 }}>
          Mpsk
        </Text>
        <Text weight={"inherit"} style={{ fontSize: 40 }} align={"center"}>
          Login into
        </Text>
        <form onSubmit={form.onSubmit((values) => HandleLogin(values))}>
          <TextInput
            required
            label="Email"
            type="email"
            placeholder="your@email.com"
            {...form.getInputProps("email")}
            size={"md"}
          />

          <PasswordInput
            required
            label="Password"
            {...form.getInputProps("password")}
            size={"md"}
          />
          <Group position="right" mt="md">
            <Button type="submit">Submit</Button>
          </Group>
        </form>
      </Box>
    </Center>
  );
}
