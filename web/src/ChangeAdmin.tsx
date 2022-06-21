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
import { address } from "./main";
import { useNavigate } from "react-router-dom";

interface FormData {
  name: string;
  surname: string;
  email: string;
  password: string;
}
interface ChangeData {
  result: string;
}

export default function App() {
  let navigate = useNavigate();
  function HandleLogin(params: FormData) {
    fetch(address + "/api/v1/adduser", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify({
        Name: params.name,
        Surname: params.surname,
        Email: params.email,
        Password: params.password,
        Role: 0,
      }),
    })
      .then((res) => res.json())
      .then((res) => {
        console.log(res);
      });
    fetch(address + "/api/v1/changeadmin", { credentials: "include" })
      .then((res) => res.json())
      .then((res: ChangeData) =>
        res.result == "success" ? navigate("/") : console.log("error")
      );
  }
  const form = useForm({
    initialValues: {
      name: "",
      surname: "",
      email: "",
      password: "",
    },

    validate: {
      email: (value) =>
        /^\S+@\S+$/.test(value) || value == "admin@admin.pl"
          ? null
          : "Invalid email",
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
            {...form.getInputProps("name")}
            size={"md"}
          />
          <TextInput
            required
            label="Surname"
            type="text"
            placeholder="Nowak"
            {...form.getInputProps("surname")}
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
