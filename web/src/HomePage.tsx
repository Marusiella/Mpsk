import React, { useEffect, useMemo, useState } from "react";
import {
  AppShell,
  Navbar,
  Header,
  Footer,
  Aside,
  Text,
  MediaQuery,
  Burger,
  useMantineTheme,
  Center,
  Box,
  MantineTheme,
} from "@mantine/core";
import { address } from "./main";
import useSWR from "swr";
import { Group, User } from "./models";

function BoxTheme(theme: MantineTheme) {
  return {
    backgroundColor:
      theme.colorScheme === "dark"
        ? theme.colors.dark[6]
        : theme.colors.gray[0],
    textAlign: "left",
    padding: theme.spacing.lg,
    borderRadius: theme.radius.md,
    cursor: "pointer",

    "&:hover": {
      backgroundColor:
        theme.colorScheme === "dark"
          ? theme.colors.dark[5]
          : theme.colors.gray[1],
    },
  };
}

export default function HomePage() {
  const theme = useMantineTheme();
  const [opened, setOpened] = useState(false);
  const [root, setRoot] = useState<Group[]>([]);
  const [users, setUsers] = useState<User[]>([]);
  function Fetch() {
    fetch(address + "/api/v1/gettasks", {
      credentials: "include",
    })
      .then((res) => res.json())
      .then((data: Group[]) => {
        console.log([data]);
        try {
          data.map(() => {});
          setRoot(data);
        } catch (error) {
          // @ts-ignore
          // TODO: it's work, but it's not good, need to fix it
          setRoot([data]);
        }
      });
    // http://127.0.0.1:3000/api/v1/getusers
    fetch(address + "/api/v1/getusers", {
      credentials: "include",
    })
      .then((res) => res.json())
      .then((data: User[]) => {
        console.log(data);
        setUsers(data);
      });
  }
  useEffect(() => {
    Fetch();
  }, []);

  return (
    <AppShell
      styles={{
        main: {
          background:
            theme.colorScheme === "dark"
              ? theme.colors.dark[8]
              : theme.colors.gray[0],
        },
      }}
      navbarOffsetBreakpoint="sm"
      asideOffsetBreakpoint="sm"
      fixed
      navbar={
        <Navbar
          p="md"
          hiddenBreakpoint="sm"
          hidden={!opened}
          width={{ sm: 200, lg: 300 }}
        >
          <Navbar.Section>{/* Header with logo */}</Navbar.Section>
          <Navbar.Section grow mt="md">
            {/* Links sections */}
          </Navbar.Section>
          <Navbar.Section>{/* Footer with user */}</Navbar.Section>
        </Navbar>
      }
      aside={
        <MediaQuery smallerThan="sm" styles={{ display: "none" }}>
          <Aside p="md" hiddenBreakpoint="sm" width={{ sm: 200, lg: 300 }}>
            {users?.map((user) => (
              <Box
                key={user.ID}
                onClick={() => {
                  console.log(user);
                }}
                sx={{
                  backgroundColor:
                    theme.colorScheme === "dark"
                      ? theme.colors.dark[6]
                      : theme.colors.gray[0],
                  textAlign: "left",
                  padding: theme.spacing.xs,
                  marginBottom: theme.spacing.xs,
                  borderRadius: theme.radius.md,
                  cursor: "pointer",

                  "&:hover": {
                    backgroundColor:
                      theme.colorScheme === "dark"
                        ? theme.colors.dark[5]
                        : theme.colors.gray[1],
                  },
                }}
              >
                <Text>
                  {user.Name} {user.Surname}
                </Text>
                <Text
                  sx={{
                    fontSize: 12,
                    textOverflow: "ellipsis",
                    overflow: "hidden",
                    maxWidth: "100%",
                    whiteSpace: "nowrap",
                  }}
                >
                  {user.Email}
                </Text>
              </Box>
            ))}
          </Aside>
        </MediaQuery>
      }
      header={
        <Header height={70} p="md">
          <div
            style={{ display: "flex", alignItems: "center", height: "100%" }}
          >
            <MediaQuery largerThan="sm" styles={{ display: "none" }}>
              <Burger
                opened={opened}
                onClick={() => setOpened((o) => !o)}
                size="sm"
                color={theme.colors.gray[6]}
                mr="xl"
              />
            </MediaQuery>

            <Text weight={"bold"} style={{ fontSize: "40px" }}>
              Mpsk
            </Text>
          </div>
        </Header>
      }
    >
      {root?.map((item) => (
        <>
          {/* @ts-ignore sx={(theme) => BoxTheme(theme)} */}
          <Box
            id={String(item.ID)}
            sx={(theme) => ({ padding: theme.spacing.lg })}
          >
            <Text
              sx={{
                fontSize: "3em",
                fontWeight: "bold",
                borderBottom: "2px dashed",
              }}
            >
              {item.Name}
            </Text>
          </Box>
          {/* @ts-ignore */}
          <Box sx={(theme) => BoxTheme(theme)} id={String(item.ID)}>
            {item.Tasks.map((task) => (
              <div key={task.ID}>
                <Text sx={{ fontSize: 25 }}>{task.Name}</Text>
                <Text
                  sx={{
                    fontSize: 15,
                    textOverflow: "ellipsis",
                    overflow: "hidden",
                    maxWidth: "30vw",
                    whiteSpace: "nowrap",
                  }}
                >
                  {task.Description}
                </Text>
              </div>
            ))}
          </Box>
        </>
      ))}
    </AppShell>
  );
}
