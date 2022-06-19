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

interface FormData {
  email: string;
  password: string;
}
function HandleLogin(params: FormData) {
  fetch("http://localhost:3000/api/v1/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
    body: JSON.stringify(params),
  })
    .then((res) => res.json())
    .then((res) => {
      console.log(res);
    });
}

export default function App() {
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
        <Text weight={"bold"} style={{ fontSize: 100 }} align={"center"}>
          Mpsk
        </Text>
        <Text weight={"inherit"} style={{ fontSize: 40 }} align={"center"}>
          You have to create username for admin user
        </Text>
        <form onSubmit={form.onSubmit((values) => HandleLogin(values))}>
          <TextInput
            required
            label="Name"
            type="text"
            placeholder="Adam"
            {...form.getInputProps("email")}
            size={"md"}
          />
          <TextInput
            required
            label="Surname"
            type="text"
            placeholder="Nowak"
            {...form.getInputProps("email")}
            size={"md"}
          />

          <TextInput
            required
            label="Email"
            type="email"
            placeholder="admin@admin.pl"
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
